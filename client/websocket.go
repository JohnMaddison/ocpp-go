package client

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/JohnMaddison/ocpp-go/internal/ws"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
	"github.com/gorilla/websocket"
)

func (c *Client) Connect() error {
	// WebSocket server URL (from client configuration)
	url := c.address + "/" + c.chargePointID

	// Optional Basic Auth
	headers := http.Header{}
	if c.username != "" || c.password != "" {
		auth := c.username + ":" + c.password
		authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
		headers.Add("Authorization", authHeader)
	}

	// Create a dialer with supported subprotocols
	dialer := websocket.Dialer{
		Subprotocols: []string{"ocpp1.6"}, EnableCompression: true, ReadBufferSize: 2048, WriteBufferSize: 2048, HandshakeTimeout: 10 * time.Second,
	}

	// Connect
	conn, resp, err := dialer.Dial(url, headers)
	if err != nil {
		if resp != nil {
			return fmt.Errorf("dial failed: %v, HTTP status: %v", err, resp.Status)
		}
		return fmt.Errorf("dial failed: %v", err)
	}
	// Create a context
	c.OCPPContext = ocpp16.NewOCPPContext(c.chargePointID)

	// Use common runner with optional traffic logging and keepalive settings
	go ws.Run(conn, c.ocppCallbacks.ParseMessage, c.OCPPContext, c.socketCallbacks, &ws.Options{
		LogSent:      c.logTraffic,
		LogKeepalive: c.logKeepalive,
		PingInterval: c.pingInterval,
		PongTimeout:  c.pongTimeout,
	})
	return nil
}
