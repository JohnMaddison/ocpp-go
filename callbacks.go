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
	ConnectRequest ConnectRequestCallback
	Connected      ConnectedCallback
	Disconnect     DisconnectCallback
}
