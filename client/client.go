package client

import (
	"time"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
	"github.com/google/uuid"
)

type Client struct {
	chargepointId   string
	address         string
	ocppcallbacks   ocpp16.OcppCallbacks
	socketcallbacks ocpp.SocketCallbacks
	Ocppcontext     *ocpp16.OcppContext
	logTraffic      bool
	logKeepalive    bool
	username        string
	password        string
	pingInterval    time.Duration
	pongTimeout     time.Duration
}

func NewExtensibleOcppClient(chargepointId string, address string, ocppcallbacks ocpp16.OcppCallbacks, socketcallbacks ocpp.SocketCallbacks) *Client {

	client := &Client{
		chargepointId:   chargepointId,
		address:         address,
		ocppcallbacks:   ocppcallbacks,
		socketcallbacks: socketcallbacks,
		Ocppcontext:     nil,
	}
	client.ocppcallbacks.InitHandlers()
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

// WithBasicAuth configures HTTP Basic authentication for the websocket handshake.
func (c *Client) WithBasicAuth(username, password string) *Client {
	c.username = username
	c.password = password
	return c
}

func (c *Client) WithOfflineHandler(callback ocpp.DisconnectCallback) *Client {
	c.socketcallbacks.Disconnect = callback
	return c
}

// WithWebsocketKeepalive configures websocket ping/pong behaviour.
// A non-positive interval disables keepalive.
// When pongTimeout is zero or negative, a default of twice the interval is used.
func (c *Client) WithWebsocketKeepalive(interval, pongTimeout time.Duration) *Client {
	c.pingInterval = interval
	c.pongTimeout = pongTimeout
	return c
}

func (c *Client) WithDataTransferHandler(callback ocpp16.DataTransferCallback) *Client {
	c.ocppcallbacks.DataTransfer = callback
	return c
}

func (c *Client) WithConfigurationHandler(callback ocpp16.GetConfigurationCallback) *Client {
	c.ocppcallbacks.GetConfiguration = callback
	return c
}

func (c *Client) WithChangeConfigurationHandler(callback ocpp16.ChangeConfigurationCallback) *Client {
	c.ocppcallbacks.ChangeConfiguration = callback
	return c
}

func (c *Client) WithClearChargingProfileHandler(callback ocpp16.ClearChargingProfileCallback) *Client {
	c.ocppcallbacks.ClearChargingProfile = callback
	return c
}

// Convenience method for sending a call
func (c *Client) SendCall(action ocpp16.Action, payload any) (*ocpp16.ResultOrError, error) {
	return c.Ocppcontext.Send(ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: uuid.New().String(), Action: string(action), Payload: payload})
}
