// Package client provides an OCPP charge point websocket client.
package client

import (
	"time"

	"github.com/johnmaddison/ocpp-go"
	"github.com/johnmaddison/ocpp-go/internal/uuidgenerator"
	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/johnmaddison/ocpp-go/ocpp21"
)

type Client struct {
	chargePointID      string
	address            string
	ocppCallbacks      ocpp16.Callbacks
	ocpp21Callbacks    ocpp21.Callbacks
	socketCallbacks    ocpp.SocketCallbacks
	Context16          *ocpp16.Context
	Context21          *ocpp21.Context
	subprotocol        string
	logTraffic         bool
	logKeepalive       bool
	username           string
	password           string
	pingInterval       time.Duration
	pongTimeout        time.Duration
	messageIDGenerator uuidgenerator.MessageIDGeneratorMethod
}

// New16 creates a client configured for the OCPP 1.6 subprotocol.
func New16(chargePointID, address string) *Client {
	client := &Client{
		chargePointID:      chargePointID,
		address:            address,
		Context16:          nil,
		subprotocol:        "ocpp1.6",
		messageIDGenerator: uuidgenerator.DefaultUUIDGenerator,
	}
	client.ocppCallbacks.InitHandlers()
	return client
}

// New21 creates a client configured for the OCPP 2.1 subprotocol.
func New21(chargePointID, address string) *Client {
	client := &Client{
		chargePointID:      chargePointID,
		address:            address,
		Context21:          nil,
		subprotocol:        "ocpp2.1",
		messageIDGenerator: uuidgenerator.DefaultUUIDGenerator,
	}
	client.ocpp21Callbacks.InitHandlers()
	return client
}

// WithTrafficLogging enables websocket sent-traffic logging.
func (c *Client) WithTrafficLogging() *Client {
	c.logTraffic = true
	return c
}

// WithKeepaliveLogging enables logging of websocket ping/pong frames.
func (c *Client) WithKeepaliveLogging() *Client {
	c.logKeepalive = true
	return c
}

// WithMessageIDGenerator sets the generator used for outbound CALL message IDs.
func (c *Client) WithMessageIDGenerator(f func() string) *Client {
	if f != nil {
		c.messageIDGenerator = f
	}
	return c
}

// WithAddress updates the websocket server base URL.
func (c *Client) WithAddress(address string) *Client {
	c.address = address
	return c
}

// With16Callbacks sets OCPP 1.6 callbacks and selects the ocpp1.6 subprotocol.
func (c *Client) With16Callbacks(callbacks ocpp16.Callbacks) *Client {
	callbacks.InitHandlers()
	c.ocppCallbacks = callbacks
	c.subprotocol = "ocpp1.6"
	return c
}

// With21Callbacks sets OCPP 2.1 callbacks and selects the ocpp2.1 subprotocol.
func (c *Client) With21Callbacks(callbacks ocpp21.Callbacks) *Client {
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

// WithConnectedHandler sets a callback for established websocket connections.
func (c *Client) WithConnectedHandler(callback ocpp.ConnectedCallback) *Client {
	c.socketCallbacks.Connected = callback
	return c
}

// WithDisconnectHandler sets a callback for closed websocket connections.
func (c *Client) WithDisconnectHandler(callback ocpp.DisconnectCallback) *Client {
	c.socketCallbacks.Disconnect = callback
	return c
}

// WithOfflineHandler is an alias for WithDisconnectHandler.
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

// Send16Call sends a typed OCPP 1.6 CALL through the active connection.
func (c *Client) Send16Call(action ocpp16.Action, payload any) (*ocpp16.ResultOrError, error) {
	return c.Context16.SendCall(action, payload)
}

// Send21Call sends a typed OCPP 2.1 CALL through the active connection.
func (c *Client) Send21Call(action ocpp21.Action, payload any) (*ocpp21.ResultOrError, error) {
	return c.Context21.SendCall(action, payload)
}
