package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestServerShutdownStopsServe(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	address := listener.Addr().String()
	if err := listener.Close(); err != nil {
		t.Fatalf("close listener: %v", err)
	}

	srv := NewServer(address)
	serveErr := make(chan error, 1)
	go func() {
		serveErr <- srv.Serve()
	}()

	waitForServer(t, address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		t.Fatalf("shutdown: %v", err)
	}

	select {
	case err := <-serveErr:
		if !errors.Is(err, http.ErrServerClosed) {
			t.Fatalf("Serve() error = %v, want %v", err, http.ErrServerClosed)
		}
	case <-time.After(time.Second):
		t.Fatal("Serve() did not return after Shutdown")
	}
}

func TestServerShutdownClosesActiveWebsockets(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	address := listener.Addr().String()
	if err := listener.Close(); err != nil {
		t.Fatalf("close listener: %v", err)
	}

	srv := NewServer(address)
	serveErr := make(chan error, 1)
	go func() {
		serveErr <- srv.Serve()
	}()

	waitForServer(t, address)

	dialer := websocket.Dialer{Subprotocols: []string{"ocpp1.6"}}
	conn, resp, err := dialer.Dial("ws://"+address+"/ocpp/CP123", nil)
	if err != nil {
		t.Fatalf("dial websocket: %v", err)
	}
	if resp.StatusCode != http.StatusSwitchingProtocols {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusSwitchingProtocols)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		t.Fatalf("shutdown: %v", err)
	}

	if got := srv.activeWebsocketCount(); got != 0 {
		t.Fatalf("active websockets = %d, want 0", got)
	}

	if err := conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond)); err != nil {
		t.Fatalf("set read deadline: %v", err)
	}
	if _, _, err := conn.ReadMessage(); err == nil {
		t.Fatal("expected websocket to be closed")
	}

	select {
	case err := <-serveErr:
		if !errors.Is(err, http.ErrServerClosed) {
			t.Fatalf("Serve() error = %v, want %v", err, http.ErrServerClosed)
		}
	case <-time.After(time.Second):
		t.Fatal("Serve() did not return after Shutdown")
	}
}

func waitForServer(t *testing.T, address string) {
	t.Helper()

	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		resp, err := http.Get("http://" + address + "/ocpp/CP123")
		if err == nil {
			_ = resp.Body.Close()
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("server did not start")
}
