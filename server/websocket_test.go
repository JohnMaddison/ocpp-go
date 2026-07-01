package server

import (
	"net/http"
	"net/http/httptest"
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
		// Inject path value for testing since r.PathValue doesn't work outside chi/router
		r.SetPathValue("cpid", "CP123")
		server.wshandler(w, r)
	})

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	// Convert HTTP URL to WebSocket URL
	wsURL := "ws" + testServer.URL[4:] + "/ws/CP123"
	header := http.Header{}
	header.Add("Sec-WebSocket-Protocol", "ocpp1.6")

	conn, resp, err := websocket.DefaultDialer.Dial(wsURL, header)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode)
	defer conn.Close()

	// Send a BootNotification message
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

	assert.Equal(t, float64(3), response[0]) // MessageTypeCallResult
	assert.Equal(t, "abc123", response[1])   // MessageID
	payload, ok := response[2].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "Accepted", payload["status"])

	// ✅ Assert that the callbacks were called
	assert.True(t, calledConnected, "Connected callback should be called")
	assert.Equal(t, "CP123", connectedChargePointID)
	assert.True(t, calledBootNotification, "BootNotification callback should be called")
}
