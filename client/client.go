// Package client provides an OCPP charge point websocket client.
package client

import (
	"time"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
	"github.com/JohnMaddison/ocpp-go/ocpp21"
	"github.com/google/uuid"
)

type Client struct {
	chargePointID   string
	address         string
	ocppCallbacks   ocpp16.OCPPCallbacks
	ocpp21Callbacks ocpp21.OCPPCallbacks
	socketCallbacks ocpp.SocketCallbacks
	OCPPContext     *ocpp16.OCPPContext
	OCPP21Context   *ocpp21.OCPPContext
	subprotocol     string
	logTraffic      bool
	logKeepalive    bool
	username        string
	password        string
	pingInterval    time.Duration
	pongTimeout     time.Duration
}

func NewOCPP16Client(chargePointID, address string) *Client {
	client := &Client{
		chargePointID: chargePointID,
		address:       address,
		OCPPContext:   nil,
		subprotocol:   "ocpp1.6",
	}
	client.ocppCallbacks.InitHandlers()
	return client
}

func NewOCPP21Client(chargePointID, address string) *Client {
	client := &Client{
		chargePointID: chargePointID,
		address:       address,
		OCPP21Context: nil,
		subprotocol:   "ocpp2.1",
	}
	client.ocpp21Callbacks.InitHandlers()
	return client
}

// EnableTrafficLogging toggles websocket sent-traffic logging.
func (c *Client) WithTrafficLogging() *Client {
	c.logTraffic = true
	return c
}

// EnableKeepaliveLogging toggles logging of websocket ping/pong frames.
func (c *Client) WithKeepaliveLogging() *Client {
	c.logKeepalive = true
	return c
}

func (c *Client) WithAddress(address string) *Client {
	c.address = address
	return c
}

func (c *Client) WithOCPP16Callbacks(callbacks ocpp16.OCPPCallbacks) *Client {
	callbacks.InitHandlers()
	c.ocppCallbacks = callbacks
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP21Callbacks(callbacks ocpp21.OCPPCallbacks) *Client {
	callbacks.InitHandlers()
	c.ocpp21Callbacks = callbacks
	c.subprotocol = "ocpp2.1"
	return c
}

// WithBasicAuth configures HTTP Basic authentication for the websocket handshake.
func (c *Client) WithBasicAuth(username, password string) *Client {
	c.username = username
	c.password = password
	return c
}

func (c *Client) WithConnectedHandler(callback ocpp.ConnectedCallback) *Client {
	c.socketCallbacks.Connected = callback
	return c
}

func (c *Client) WithDisconnectHandler(callback ocpp.DisconnectCallback) *Client {
	c.socketCallbacks.Disconnect = callback
	return c
}

func (c *Client) WithOfflineHandler(callback ocpp.DisconnectCallback) *Client {
	return c.WithDisconnectHandler(callback)
}

// WithWebsocketKeepalive configures websocket ping/pong behaviour.
// A non-positive interval disables keepalive.
// When pongTimeout is zero or negative, a default of twice the interval is used.
func (c *Client) WithWebsocketKeepalive(interval, pongTimeout time.Duration) *Client {
	c.pingInterval = interval
	c.pongTimeout = pongTimeout
	return c
}

func (c *Client) SendOCPP16Call(action ocpp16.Action, payload any) (*ocpp16.ResultOrError, error) {
	return c.OCPPContext.Send(ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: uuid.New().String(), Action: string(action), Payload: payload})
}

func (c *Client) SendOCPP21Call(action ocpp21.Action, payload any) (*ocpp21.ResultOrError, error) {
	return c.OCPP21Context.Send(ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: uuid.New().String(), Action: string(action), Payload: payload})
}
