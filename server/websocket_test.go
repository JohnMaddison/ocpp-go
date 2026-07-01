package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/johnmaddison/ocpp-go"
	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/stretchr/testify/assert"
)

func TestWshandler_BootNotification(t *testing.T) {

	var calledConnected bool
	var connectedChargePointID string
	var calledBootNotification bool

	server := &Server{
		socketCallbacks: ocpp.SocketCallbacks{
			Connected: func(info ocpp.ConnectionInfo) {
				t.Log("Connected callback triggered")
				calledConnected = true
				connectedChargePointID = info.ChargePointID
			},
		},
		ocppCallbacks: ocpp16.Callbacks{
			BootNotification: func(ctx *ocpp16.Context, request ocpp16.BootNotificationRequest) (*ocpp16.BootNotificationResponse, *ocpp16.Error) {
				t.Log("BootNotification callback triggered")
				calledBootNotification = true
				return &ocpp16.BootNotificationResponse{
					CurrentTime: time.Now(),
					Interval:    30,
					Status:      "Accepted",
				}, nil
			},
		},
	}
	server.ocppCallbacks.InitHandlers()

	mux := http.NewServeMux()
	mux.HandleFunc("/ws/{cpid}", func(w http.ResponseWriter, r *http.Request) {
		r.SetPathValue("cpid", "CP123")
		server.wshandler(w, r)
	})

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	wsURL := "ws" + testServer.URL[4:] + "/ws/CP123"
	header := http.Header{}
	header.Add("Sec-WebSocket-Protocol", "ocpp1.6")

	conn, resp, err := websocket.DefaultDialer.Dial(wsURL, header)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode)
	defer conn.Close()

	msg := []any{
		2,
		"abc123",
		"BootNotification",
		map[string]any{
			"chargePointModel":  "TestModel",
			"chargePointVendor": "TestVendor",
		},
	}
	err = conn.WriteJSON(msg)
	assert.NoError(t, err)

	var response []any
	err = conn.ReadJSON(&response)
	assert.NoError(t, err)

	assert.Equal(t, float64(3), response[0])
	assert.Equal(t, "abc123", response[1])
	payload, ok := response[2].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "Accepted", payload["status"])

	assert.True(t, calledConnected, "Connected callback should be called")
	assert.Equal(t, "CP123", connectedChargePointID)
	assert.True(t, calledBootNotification, "BootNotification callback should be called")
}

func TestWshandler_AppliesWebsocketReadLimit(t *testing.T) {
	server := NewServer(":0",
		WithWebsocketReadLimit(32),
		WithWebsocketKeepalive(0, 0),
	)
	testServer := newWebsocketTestServer(server)
	defer testServer.Close()

	conn := dialProtocol(t, testServer.URL, "ocpp1.6")
	defer conn.Close()

	if err := conn.WriteMessage(websocket.TextMessage, []byte(strings.Repeat("x", 128))); err != nil {
		t.Fatalf("write oversized message: %v", err)
	}
	if err := conn.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
		t.Fatalf("set read deadline: %v", err)
	}
	if _, _, err := conn.ReadMessage(); err == nil {
		t.Fatal("expected connection read to fail after oversized message")
	}
}

func TestWshandler_WebsocketCompressionCanBeDisabled(t *testing.T) {
	server := NewServer(":0", WithWebsocketCompression(false))
	testServer := newWebsocketTestServer(server)
	defer testServer.Close()

	wsURL := "ws" + testServer.URL[4:] + "/ws/CP123"
	dialer := websocket.Dialer{
		Subprotocols:      []string{"ocpp1.6"},
		EnableCompression: true,
	}
	conn, resp, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("dial failed: %v", err)
	}
	defer conn.Close()

	if got := resp.Header.Get("Sec-WebSocket-Extensions"); got != "" {
		t.Fatalf("Sec-WebSocket-Extensions = %q, want empty", got)
	}
}
