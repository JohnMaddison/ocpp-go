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

func (c *Client) WithMessageIDGenerator(f func() string) *Client {
	if f != nil {
		c.messageIDGenerator = f
	}
	return c
}

func (c *Client) WithAddress(address string) *Client {
	c.address = address
	return c
}

func (c *Client) With16Callbacks(callbacks ocpp16.Callbacks) *Client {
	callbacks.InitHandlers()
	c.ocppCallbacks = callbacks
	c.subprotocol = "ocpp1.6"
	return c
}

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

func (c *Client) Send16Call(action ocpp16.Action, payload any) (*ocpp16.ResultOrError, error) {
	return c.Context16.SendCall(action, payload)
}

func (c *Client) Send21Call(action ocpp21.Action, payload any) (*ocpp21.ResultOrError, error) {
	return c.Context21.SendCall(action, payload)
}
