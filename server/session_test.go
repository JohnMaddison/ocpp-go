package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/johnmaddison/ocpp-go"
	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/johnmaddison/ocpp-go/ocpp21"
)

func TestServerSessionMissing(t *testing.T) {
	srv := NewServer(":0")
	if _, ok := srv.Session("missing"); ok {
		t.Fatal("Session returned true for missing charge point")
	}
}

func TestServerSessionSend16(t *testing.T) {
	srv := NewServer(":0",
		With16Callbacks(ocpp16.Callbacks{}),
		WithMessageIDGenerator(func() string { return "server-16" }),
	)
	testServer := newSessionTestServer(srv)
	defer testServer.Close()

	conn := dialSessionProtocol(t, testServer.URL, "CP123", "ocpp1.6")
	defer conn.Close()

	session := waitForSession(t, srv, "CP123")
	if session.ChargePointID() != "CP123" {
		t.Fatalf("ChargePointID = %q, want CP123", session.ChargePointID())
	}
	if session.Protocol() != "ocpp1.6" {
		t.Fatalf("Protocol = %q, want ocpp1.6", session.Protocol())
	}
	if session.RemoteAddr() == nil || session.LocalAddr() == nil {
		t.Fatal("expected local and remote addresses to be set")
	}
	if _, ok := session.Context16(); !ok {
		t.Fatal("expected Context16")
	}
	if _, ok := session.Context21(); ok {
		t.Fatal("did not expect Context21")
	}

	resultCh := make(chan send16Result, 1)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		result, err := session.Send16CallWithContext(ctx, ocpp16.ActionDataTransfer, ocpp16.DataTransferRequest{VendorID: "vendor"})
		resultCh <- send16Result{result: result, err: err}
	}()

	var call []any
	if err := conn.ReadJSON(&call); err != nil {
		t.Fatalf("read server CALL: %v", err)
	}
	assertCall(t, call, "server-16", "DataTransfer")
	if err := conn.WriteJSON([]any{ocpp.MessageTypeCallResult, "server-16", map[string]any{"status": "Accepted"}}); err != nil {
		t.Fatalf("write CALLRESULT: %v", err)
	}

	got := <-resultCh
	if got.err != nil {
		t.Fatalf("Send16WithContext returned error: %v", got.err)
	}
	payload, ok := got.result.CallResult.Payload.(ocpp16.DataTransferResponse)
	if !ok {
		t.Fatalf("payload type = %T, want ocpp16.DataTransferResponse", got.result.CallResult.Payload)
	}
	if payload.Status != ocpp16.DataTransferStatusAccepted {
		t.Fatalf("status = %q, want Accepted", payload.Status)
	}
}

func TestServerSessionSend21(t *testing.T) {
	srv := NewServer(":0",
		With21Callbacks(ocpp21.Callbacks{}),
		WithMessageIDGenerator(func() string { return "server-21" }),
	)
	testServer := newSessionTestServer(srv)
	defer testServer.Close()

	conn := dialSessionProtocol(t, testServer.URL, "CP123", "ocpp2.1")
	defer conn.Close()

	session := waitForSession(t, srv, "CP123")
	if session.Protocol() != "ocpp2.1" {
		t.Fatalf("Protocol = %q, want ocpp2.1", session.Protocol())
	}
	if _, ok := session.Context21(); !ok {
		t.Fatal("expected Context21")
	}
	if _, ok := session.Context16(); ok {
		t.Fatal("did not expect Context16")
	}

	resultCh := make(chan send21Result, 1)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		result, err := session.Send21CallWithContext(ctx, ocpp21.ActionDataTransfer, ocpp21.DataTransferRequest{VendorID: "vendor"})
		resultCh <- send21Result{result: result, err: err}
	}()

	var call []any
	if err := conn.ReadJSON(&call); err != nil {
		t.Fatalf("read server CALL: %v", err)
	}
	assertCall(t, call, "server-21", "DataTransfer")
	if err := conn.WriteJSON([]any{ocpp.MessageTypeCallResult, "server-21", map[string]any{"status": "Accepted"}}); err != nil {
		t.Fatalf("write CALLRESULT: %v", err)
	}

	got := <-resultCh
	if got.err != nil {
		t.Fatalf("Send21WithContext returned error: %v", got.err)
	}
	payload, ok := got.result.CallResult.Payload.(ocpp21.DataTransferResponse)
	if !ok {
		t.Fatalf("payload type = %T, want ocpp21.DataTransferResponse", got.result.CallResult.Payload)
	}
	if payload.Status != ocpp21.DataTransferStatusEnumAccepted {
		t.Fatalf("status = %q, want Accepted", payload.Status)
	}
}

func TestSessionWrongProtocolErrors(t *testing.T) {
	session16 := &Session{
		chargePointID: "CP16",
		protocol:      "ocpp1.6",
		ocpp16Context: ocpp16.NewContext("CP16"),
	}
	if _, err := session16.Send21(ocpp.Call{}); err == nil || !strings.Contains(err.Error(), "session CP16 uses ocpp1.6, not ocpp2.1") {
		t.Fatalf("Send21 error = %v", err)
	}
	if _, err := session16.Send21Call(ocpp21.ActionDataTransfer, ocpp21.DataTransferRequest{}); err == nil || !strings.Contains(err.Error(), "session CP16 uses ocpp1.6, not ocpp2.1") {
		t.Fatalf("Send21Call error = %v", err)
	}
	if _, err := session16.Send21CallWithContext(context.Background(), ocpp21.ActionDataTransfer, ocpp21.DataTransferRequest{}); err == nil || !strings.Contains(err.Error(), "session CP16 uses ocpp1.6, not ocpp2.1") {
		t.Fatalf("Send21CallWithContext error = %v", err)
	}

	session21 := &Session{
		chargePointID: "CP21",
		protocol:      "ocpp2.1",
		ocpp21Context: ocpp21.NewContext("CP21"),
	}
	if _, err := session21.Send16(ocpp.Call{}); err == nil || !strings.Contains(err.Error(), "session CP21 uses ocpp2.1, not ocpp1.6") {
		t.Fatalf("Send16 error = %v", err)
	}
	if _, err := session21.Send16Call(ocpp16.ActionDataTransfer, ocpp16.DataTransferRequest{}); err == nil || !strings.Contains(err.Error(), "session CP21 uses ocpp2.1, not ocpp1.6") {
		t.Fatalf("Send16Call error = %v", err)
	}
	if _, err := session21.Send16CallWithContext(context.Background(), ocpp16.ActionDataTransfer, ocpp16.DataTransferRequest{}); err == nil || !strings.Contains(err.Error(), "session CP21 uses ocpp2.1, not ocpp1.6") {
		t.Fatalf("Send16CallWithContext error = %v", err)
	}
}

func TestServerSessionDuplicateIDLatestWins(t *testing.T) {
	disconnected := make(chan string, 2)
	srv := NewServer(":0",
		With16Callbacks(ocpp16.Callbacks{}),
		WithDisconnectHandler(func(info ocpp.ConnectionInfo) {
			disconnected <- info.ChargePointID
		}),
	)
	testServer := newSessionTestServer(srv)
	defer testServer.Close()

	conn1 := dialSessionProtocol(t, testServer.URL, "CP123", "ocpp1.6")
	session1 := waitForSession(t, srv, "CP123")

	conn2 := dialSessionProtocol(t, testServer.URL, "CP123", "ocpp1.6")
	defer conn2.Close()
	session2 := waitForNewSession(t, srv, "CP123", session1)
	if session2 == session1 {
		t.Fatal("expected second connection to replace addressable session")
	}

	if err := conn1.Close(); err != nil {
		t.Fatalf("close first connection: %v", err)
	}
	select {
	case cpid := <-disconnected:
		if cpid != "CP123" {
			t.Fatalf("disconnected charge point = %q, want CP123", cpid)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for first disconnect")
	}

	current, ok := srv.Session("CP123")
	if !ok || current != session2 {
		t.Fatalf("latest session was unregistered after older disconnect: ok=%t current=%p session2=%p", ok, current, session2)
	}
}

type send16Result struct {
	result *ocpp16.ResultOrError
	err    error
}

type send21Result struct {
	result *ocpp21.ResultOrError
	err    error
}

func newSessionTestServer(srv *Server) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws/{cpid}", func(w http.ResponseWriter, r *http.Request) {
		srv.wshandler(w, r)
	})
	return httptest.NewServer(mux)
}

func dialSessionProtocol(t *testing.T, serverURL, cpid, protocol string) *websocket.Conn {
	t.Helper()
	wsURL := fmt.Sprintf("ws%s/ws/%s", serverURL[4:], cpid)
	dialer := websocket.Dialer{Subprotocols: []string{protocol}}
	conn, resp, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("dial %s failed: %v", protocol, err)
	}
	if resp.StatusCode != http.StatusSwitchingProtocols {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusSwitchingProtocols)
	}
	return conn
}

func waitForSession(t *testing.T, srv *Server, cpid string) *Session {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for {
		session, ok := srv.Session(cpid)
		if ok {
			return session
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for session %s", cpid)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func waitForNewSession(t *testing.T, srv *Server, cpid string, old *Session) *Session {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for {
		session, ok := srv.Session(cpid)
		if ok && session != old {
			return session
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for replacement session %s", cpid)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func assertCall(t *testing.T, call []any, messageID, action string) {
	t.Helper()
	if len(call) != 4 {
		t.Fatalf("CALL length = %d, want 4: %#v", len(call), call)
	}
	if call[0] != float64(ocpp.MessageTypeCall) {
		t.Fatalf("message type = %#v, want CALL", call[0])
	}
	if call[1] != messageID {
		t.Fatalf("message ID = %#v, want %s", call[1], messageID)
	}
	if call[2] != action {
		t.Fatalf("action = %#v, want %s", call[2], action)
	}
}
