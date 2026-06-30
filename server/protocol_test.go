package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/JohnMaddison/ocpp-go/ocpp16"
	"github.com/JohnMaddison/ocpp-go/ocpp21"
	"github.com/gorilla/websocket"
)

func TestServerProtocols(t *testing.T) {
	tests := []struct {
		name string
		srv  *Server
		want []string
	}{
		{name: "default", srv: NewServer(":0"), want: []string{"ocpp1.6"}},
		{name: "ocpp16", srv: NewServer(":0", WithOCPP16Callbacks(ocpp16.OCPPCallbacks{})), want: []string{"ocpp1.6"}},
		{name: "ocpp21", srv: NewServer(":0", WithOCPP21Callbacks(ocpp21.OCPPCallbacks{})), want: []string{"ocpp2.1"}},
		{name: "mixed", srv: NewServer(":0", WithOCPP16Callbacks(ocpp16.OCPPCallbacks{}), WithOCPP21Callbacks(ocpp21.OCPPCallbacks{})), want: []string{"ocpp1.6", "ocpp2.1"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.srv.protocols()
			if len(got) != len(tt.want) {
				t.Fatalf("protocol count = %d, want %d: %#v", len(got), len(tt.want), got)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Fatalf("protocols = %#v, want %#v", got, tt.want)
				}
			}
		})
	}
}

func TestWshandler_RoutesBySubprotocol(t *testing.T) {
	var called16 bool
	var called21 bool
	srv := NewServer(":0",
		WithOCPP16Callbacks(ocpp16.OCPPCallbacks{
			BootNotification: func(ctx *ocpp16.OCPPContext, request ocpp16.BootNotificationRequest) (*ocpp16.BootNotificationResponse, *ocpp16.OCPPError) {
				called16 = true
				return &ocpp16.BootNotificationResponse{CurrentTime: time.Now(), Interval: 30, Status: ocpp16.RegistrationStatusAccepted}, nil
			},
		}),
		WithOCPP21Callbacks(ocpp21.OCPPCallbacks{
			BootNotification: func(ctx *ocpp21.OCPPContext, request ocpp21.BootNotificationRequest) (*ocpp21.BootNotificationResponse, *ocpp21.OCPPError) {
				called21 = true
				return &ocpp21.BootNotificationResponse{CurrentTime: time.Now(), Interval: 30, Status: ocpp21.RegistrationStatusEnumAccepted}, nil
			},
		}),
	)
	testServer := newWebsocketTestServer(srv)
	defer testServer.Close()

	conn16 := dialProtocol(t, testServer.URL, "ocpp1.6")
	defer conn16.Close()
	if err := conn16.WriteJSON([]any{2, "msg-16", "BootNotification", map[string]any{"chargePointModel": "M", "chargePointVendor": "V"}}); err != nil {
		t.Fatalf("write ocpp1.6 message: %v", err)
	}
	var response16 []any
	if err := conn16.ReadJSON(&response16); err != nil {
		t.Fatalf("read ocpp1.6 response: %v", err)
	}

	conn21 := dialProtocol(t, testServer.URL, "ocpp2.1")
	defer conn21.Close()
	if err := conn21.WriteJSON([]any{2, "msg-21", "BootNotification", map[string]any{}}); err != nil {
		t.Fatalf("write ocpp2.1 message: %v", err)
	}
	var response21 []any
	if err := conn21.ReadJSON(&response21); err != nil {
		t.Fatalf("read ocpp2.1 response: %v", err)
	}

	if !called16 || !called21 {
		t.Fatalf("callbacks called: ocpp16=%t ocpp21=%t", called16, called21)
	}
}

func TestWshandler_RejectsMissingOrUnsupportedSubprotocol(t *testing.T) {
	srv := NewServer(":0", WithOCPP21Callbacks(ocpp21.OCPPCallbacks{}))
	testServer := newWebsocketTestServer(srv)
	defer testServer.Close()
	wsURL := "ws" + testServer.URL[4:] + "/ws/CP123"

	if _, _, err := websocket.DefaultDialer.Dial(wsURL, nil); err == nil {
		t.Fatal("expected missing subprotocol to be rejected")
	}

	dialer := websocket.Dialer{Subprotocols: []string{"ocpp1.6"}}
	if _, _, err := dialer.Dial(wsURL, nil); err == nil {
		t.Fatal("expected unsupported subprotocol to be rejected")
	}
}

func newWebsocketTestServer(srv *Server) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws/{cpid}", func(w http.ResponseWriter, r *http.Request) {
		r.SetPathValue("cpid", "CP123")
		srv.wshandler(w, r)
	})
	return httptest.NewServer(mux)
}

func dialProtocol(t *testing.T, serverURL, protocol string) *websocket.Conn {
	t.Helper()
	wsURL := "ws" + serverURL[4:] + "/ws/CP123"
	dialer := websocket.Dialer{Subprotocols: []string{protocol}}
	conn, resp, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("dial %s failed: %v", protocol, err)
	}
	if resp.StatusCode != http.StatusSwitchingProtocols {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusSwitchingProtocols)
	}
	if conn.Subprotocol() != protocol {
		t.Fatalf("negotiated protocol = %q, want %q", conn.Subprotocol(), protocol)
	}
	return conn
}
