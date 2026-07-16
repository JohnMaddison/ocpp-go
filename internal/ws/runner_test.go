package ws

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/johnmaddison/ocpp-go"
)

func TestRunMessageReceivedCallback(t *testing.T) {
	received := make(chan ocpp.MessageInfo, 1)
	parsed := make(chan []byte, 1)
	payload := []byte(`[2,"msg-1","BootNotification",{}]`)
	runtime := Runtime{
		ChargePointID: "CP_1",
		Protocol:      "ocpp1.6",
		OutgoingCalls: closedOutgoingCalls,
		Parse: func(message []byte) ([]byte, error) {
			parsed <- append([]byte(nil), message...)
			return nil, nil
		},
	}
	callbacks := ocpp.SocketCallbacks{
		MessageReceived: func(info ocpp.MessageInfo) {
			info.Message[0] = 'x'
			received <- info
		},
	}
	conn, cleanup := newRunnerTestPeer(t, runtime, callbacks)
	defer cleanup()

	if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
		t.Fatalf("write message: %v", err)
	}

	info := readMessageInfo(t, received)
	if info.Direction != ocpp.MessageDirectionReceived {
		t.Fatalf("direction = %q, want %q", info.Direction, ocpp.MessageDirectionReceived)
	}
	if info.ConnectionInfo.ChargePointID != "CP_1" {
		t.Fatalf("charge point ID = %q, want CP_1", info.ConnectionInfo.ChargePointID)
	}
	if info.ConnectionInfo.RemoteAddr == nil {
		t.Fatal("expected remote address")
	}
	if info.ConnectionInfo.LocalAddr == nil {
		t.Fatal("expected local address")
	}
	if info.Protocol != "ocpp1.6" {
		t.Fatalf("protocol = %q, want ocpp1.6", info.Protocol)
	}
	if string(info.Message) != `x2,"msg-1","BootNotification",{}]` {
		t.Fatalf("callback message = %q", string(info.Message))
	}

	select {
	case got := <-parsed:
		if string(got) != string(payload) {
			t.Fatalf("parser message = %q, want %q", string(got), string(payload))
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for parser")
	}
}

func TestRunMessageSentCallbackForQueuedCalls(t *testing.T) {
	sent := make(chan ocpp.MessageInfo, 1)
	outgoing := make(chan any, 1)
	payload := []byte(`[2,"msg-1","Reset",{}]`)
	runtime := Runtime{
		ChargePointID: "CP_1",
		Protocol:      "ocpp1.6",
		OutgoingCalls: func(done <-chan struct{}) <-chan any { return outgoing },
		Parse:         func(message []byte) ([]byte, error) { return nil, nil },
		Serialize: func(call any) ([]byte, error) {
			return append([]byte(nil), call.([]byte)...), nil
		},
	}
	conn, cleanup := newRunnerTestPeer(t, runtime, ocpp.SocketCallbacks{
		MessageSent: func(info ocpp.MessageInfo) { sent <- info },
	})
	defer cleanup()

	outgoing <- payload

	mt, got, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("read message: %v", err)
	}
	if mt != websocket.TextMessage {
		t.Fatalf("message type = %d, want text", mt)
	}
	if string(got) != string(payload) {
		t.Fatalf("sent websocket payload = %q, want %q", string(got), string(payload))
	}
	info := readMessageInfo(t, sent)
	if info.Direction != ocpp.MessageDirectionSent {
		t.Fatalf("direction = %q, want %q", info.Direction, ocpp.MessageDirectionSent)
	}
	if string(info.Message) != string(payload) {
		t.Fatalf("callback message = %q, want %q", string(info.Message), string(payload))
	}
}

func TestRunMessageSentCallbackForParserResponses(t *testing.T) {
	sent := make(chan ocpp.MessageInfo, 1)
	response := []byte(`[3,"msg-1",{}]`)
	runtime := Runtime{
		ChargePointID: "CP_1",
		Protocol:      "ocpp2.1",
		OutgoingCalls: closedOutgoingCalls,
		Parse: func(message []byte) ([]byte, error) {
			return append([]byte(nil), response...), nil
		},
	}
	conn, cleanup := newRunnerTestPeer(t, runtime, ocpp.SocketCallbacks{
		MessageSent: func(info ocpp.MessageInfo) { sent <- info },
	})
	defer cleanup()

	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[2,"msg-1","BootNotification",{}]`)); err != nil {
		t.Fatalf("write message: %v", err)
	}
	_, got, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("read response: %v", err)
	}
	if string(got) != string(response) {
		t.Fatalf("response = %q, want %q", string(got), string(response))
	}
	info := readMessageInfo(t, sent)
	if info.Protocol != "ocpp2.1" {
		t.Fatalf("protocol = %q, want ocpp2.1", info.Protocol)
	}
	if string(info.Message) != string(response) {
		t.Fatalf("callback message = %q, want %q", string(info.Message), string(response))
	}
}

func TestRunNilMessageCallbacksAreSafe(t *testing.T) {
	parsed := make(chan struct{}, 1)
	runtime := Runtime{
		ChargePointID: "CP_1",
		Protocol:      "ocpp1.6",
		OutgoingCalls: closedOutgoingCalls,
		Parse: func(message []byte) ([]byte, error) {
			parsed <- struct{}{}
			return nil, nil
		},
	}
	conn, cleanup := newRunnerTestPeer(t, runtime, ocpp.SocketCallbacks{})
	defer cleanup()

	if err := conn.WriteMessage(websocket.TextMessage, []byte(`[3,"msg-1",{}]`)); err != nil {
		t.Fatalf("write message: %v", err)
	}
	select {
	case <-parsed:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for parser")
	}
}

func newRunnerTestPeer(t *testing.T, runtime Runtime, callbacks ocpp.SocketCallbacks) (*websocket.Conn, func()) {
	t.Helper()
	if runtime.OutgoingCalls == nil {
		runtime.OutgoingCalls = closedOutgoingCalls
	}
	if runtime.Parse == nil {
		runtime.Parse = func(message []byte) ([]byte, error) { return nil, nil }
	}
	upgrader := websocket.Upgrader{}
	runnerDone := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("upgrade: %v", err)
			return
		}
		defer close(runnerDone)
		Run(conn, runtime, callbacks, &Options{PingInterval: 0})
	}))

	conn, _, err := websocket.DefaultDialer.Dial("ws"+server.URL[4:], nil)
	if err != nil {
		server.Close()
		t.Fatalf("dial: %v", err)
	}
	cleanup := func() {
		_ = conn.Close()
		select {
		case <-runnerDone:
		case <-time.After(time.Second):
			t.Fatal("timed out waiting for runner shutdown")
		}
		server.Close()
	}
	return conn, cleanup
}

func closedOutgoingCalls(done <-chan struct{}) <-chan any {
	ch := make(chan any)
	go func() {
		<-done
		close(ch)
	}()
	return ch
}

func readMessageInfo(t *testing.T, ch <-chan ocpp.MessageInfo) ocpp.MessageInfo {
	t.Helper()
	select {
	case info := <-ch:
		return info
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for message callback")
		return ocpp.MessageInfo{}
	}
}
