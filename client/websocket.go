package client

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/johnmaddison/ocpp-go/internal/ws"
	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/johnmaddison/ocpp-go/ocpp21"
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
		Subprotocols: []string{c.subprotocol}, EnableCompression: true, ReadBufferSize: 2048, WriteBufferSize: 2048, HandshakeTimeout: 10 * time.Second,
	}

	// Connect
	conn, resp, err := dialer.Dial(url, headers)
	if err != nil {
		if resp != nil {
			return fmt.Errorf("dial failed: %v, HTTP status: %v", err, resp.Status)
		}
		return fmt.Errorf("dial failed: %v", err)
	}
	runtime, err := c.runtime()
	if err != nil {
		conn.Close()
		return err
	}

	// Use common runner with optional traffic logging and keepalive settings
	go ws.Run(conn, runtime, c.socketCallbacks, &ws.Options{
		LogSent:      c.logTraffic,
		LogKeepalive: c.logKeepalive,
		PingInterval: c.pingInterval,
		PongTimeout:  c.pongTimeout,
	})
	return nil
}

func (c *Client) runtime() (ws.Runtime, error) {
	switch c.subprotocol {
	case "ocpp1.6":
		c.Context16 = ocpp16.NewContextWithMessageIDGenerator(c.chargePointID, c.messageIDGenerator)
		return ws.Runtime{
			ChargePointID: c.Context16.ChargePointID,
			OutgoingCalls: requestChannel[ocpp16.Request](c.Context16.Queue),
			Parse: func(message []byte) ([]byte, error) {
				return c.ocppCallbacks.ParseMessage(message, c.Context16)
			},
			Serialize: func(call any) ([]byte, error) {
				request := call.(ocpp16.Request)
				return request.Call.SerializeOCPP()
			},
		}, nil
	case "ocpp2.1":
		c.Context21 = ocpp21.NewContextWithMessageIDGenerator(c.chargePointID, c.messageIDGenerator)
		return ws.Runtime{
			ChargePointID: c.Context21.ChargePointID,
			OutgoingCalls: requestChannel[ocpp21.Request](c.Context21.Queue),
			Parse: func(message []byte) ([]byte, error) {
				return c.ocpp21Callbacks.ParseMessage(message, c.Context21)
			},
			Serialize: func(call any) ([]byte, error) {
				request := call.(ocpp21.Request)
				return request.Call.SerializeOCPP()
			},
		}, nil
	default:
		return ws.Runtime{}, fmt.Errorf("unsupported OCPP subprotocol: %s", c.subprotocol)
	}
}

func requestChannel[T any](in <-chan T) func(done <-chan struct{}) <-chan any {
	return func(done <-chan struct{}) <-chan any {
		out := make(chan any)
		go func() {
			defer close(out)
			for {
				select {
				case item, ok := <-in:
					if !ok {
						return
					}
					select {
					case out <- item:
					case <-done:
						return
					}
				case <-done:
					return
				}
			}
		}()
		return out
	}
}
