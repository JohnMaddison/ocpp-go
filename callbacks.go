package ocpp

import (
	"net"
	"net/http"
)

// ConnectRequestCallback decides whether an incoming websocket upgrade may continue.
type ConnectRequestCallback func(ctx ConnectRequest) bool

// ConnectedCallback is called after a websocket connection has been established.
type ConnectedCallback func(info ConnectionInfo)

// DisconnectCallback is called after a websocket connection has closed.
type DisconnectCallback func(info ConnectionInfo)

// MessageDirection identifies whether an OCPP websocket text frame was sent or received.
type MessageDirection string

const (
	// MessageDirectionReceived identifies an inbound OCPP websocket text frame.
	MessageDirectionReceived MessageDirection = "received"
	// MessageDirectionSent identifies an outbound OCPP websocket text frame.
	MessageDirectionSent MessageDirection = "sent"
)

// MessageInfo describes an observed OCPP websocket text frame.
type MessageInfo struct {
	ConnectionInfo ConnectionInfo
	Protocol       string
	Direction      MessageDirection
	Message        []byte
}

// MessageCallback observes an OCPP websocket text frame.
type MessageCallback func(info MessageInfo)

// ConnectRequest contains the HTTP upgrade request and response writer.
type ConnectRequest struct {
	R *http.Request
	W *http.ResponseWriter
}

// ConnectionInfo describes an established OCPP websocket connection.
type ConnectionInfo struct {
	ChargePointID string
	RemoteAddr    net.Addr
	LocalAddr     net.Addr
}

// SocketCallbacks groups websocket lifecycle callbacks.
type SocketCallbacks struct {
	ConnectRequest  ConnectRequestCallback
	Connected       ConnectedCallback
	Disconnect      DisconnectCallback
	MessageReceived MessageCallback
	MessageSent     MessageCallback
}
