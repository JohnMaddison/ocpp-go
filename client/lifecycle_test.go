package client

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/johnmaddison/ocpp-go"
)

func TestAppendChargePointIDPreservesURLAndEscapesSegment(t *testing.T) {
	got, err := appendChargePointID("wss://example.test/ocpp/?tenant=one", "CP / 1")
	if err != nil {
		t.Fatalf("append charge point id: %v", err)
	}
	if got != "wss://example.test/ocpp/CP%20%2F%201?tenant=one" {
		t.Fatalf("URL = %q", got)
	}
}

func TestDuplicateConnectDisconnectAndReconnect(t *testing.T) {
	server, connections := newLifecycleServer(t)
	defer server.Close()

	var connected, disconnected atomic.Int32
	c := New16("CP_1", websocketAddress(server.URL)).
		WithConnectedHandler(func(info ocpp.ConnectionInfo) { connected.Add(1) }).
		WithDisconnectHandler(func(info ocpp.ConnectionInfo) { disconnected.Add(1) })

	if err := c.Connect(); err != nil {
		t.Fatalf("first connect: %v", err)
	}
	waitForCount(t, connections, 1)
	if err := c.Connect(); !errors.Is(err, ErrAlreadyConnected) {
		t.Fatalf("duplicate connect error = %v, want ErrAlreadyConnected", err)
	}
	if err := c.Disconnect(); err != nil {
		t.Fatalf("first disconnect: %v", err)
	}
	if err := c.Disconnect(); err != nil {
		t.Fatalf("idempotent disconnect: %v", err)
	}
	waitAtomic(t, &disconnected, 1)

	if err := c.Connect(); err != nil {
		t.Fatalf("reconnect: %v", err)
	}
	waitForCount(t, connections, 2)
	if err := c.Disconnect(); err != nil {
		t.Fatalf("second disconnect: %v", err)
	}
	waitAtomic(t, &disconnected, 2)
	if connected.Load() != 2 {
		t.Fatalf("connected callbacks = %d, want 2", connected.Load())
	}
}

func TestDisconnectCancelsInProgressDial(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	defer listener.Close()
	accepted := make(chan net.Conn, 1)
	go func() {
		conn, acceptErr := listener.Accept()
		if acceptErr == nil {
			accepted <- conn
		}
	}()

	c := New16("CP_1", "ws://"+listener.Addr().String()+"/ocpp")
	connectResult := make(chan error, 1)
	go func() { connectResult <- c.ConnectContext(context.Background()) }()
	conn := <-accepted
	defer conn.Close()

	disconnectResult := make(chan error, 1)
	go func() { disconnectResult <- c.Disconnect() }()
	select {
	case err := <-disconnectResult:
		if err != nil {
			t.Fatalf("disconnect: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("disconnect did not cancel dial")
	}
	select {
	case err := <-connectResult:
		if err == nil {
			t.Fatal("connect unexpectedly succeeded")
		}
	case <-time.After(time.Second):
		t.Fatal("ConnectContext did not return after disconnect")
	}
}

func TestConnectContextCancellationClearsAttempt(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	defer listener.Close()
	go func() {
		conn, acceptErr := listener.Accept()
		if acceptErr == nil {
			defer conn.Close()
			<-time.After(time.Second)
		}
	}()

	c := New16("CP_1", "ws://"+listener.Addr().String()+"/ocpp")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	if err := c.ConnectContext(ctx); err == nil {
		t.Fatal("ConnectContext unexpectedly succeeded")
	}
	if err := c.Disconnect(); err != nil {
		t.Fatalf("disconnect after canceled context: %v", err)
	}
}

func newLifecycleServer(t *testing.T) (*httptest.Server, <-chan int) {
	t.Helper()
	connections := make(chan int, 4)
	var count atomic.Int32
	upgrader := websocket.Upgrader{
		Subprotocols: []string{"ocpp1.6"},
		CheckOrigin:  func(r *http.Request) bool { return true },
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		connections <- int(count.Add(1))
		defer conn.Close()
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}))
	return server, connections
}

func websocketAddress(httpAddress string) string {
	return "ws" + strings.TrimPrefix(httpAddress, "http") + "/ocpp"
}

func waitForCount(t *testing.T, counts <-chan int, wanted int) {
	t.Helper()
	select {
	case got := <-counts:
		if got != wanted {
			t.Fatalf("connection count = %d, want %d", got, wanted)
		}
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for connection %d", wanted)
	}
}

func waitAtomic(t *testing.T, value *atomic.Int32, wanted int32) {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for value.Load() != wanted && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if value.Load() != wanted {
		t.Fatalf("callback count = %d, want %d", value.Load(), wanted)
	}
}
