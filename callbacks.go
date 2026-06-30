package ocpp

import (
	"net"
	"net/http"
)

// CONNECTION CALLBACKS
type ConnectRequestCallback func(ctx ConnectRequest) bool
type ConnectedCallback func(info ConnectionInfo)
type DisconnectCallback func(info ConnectionInfo)

type ConnectRequest struct {
	R *http.Request
	W *http.ResponseWriter
}

type ConnectionInfo struct {
	ChargePointID string
	RemoteAddr    net.Addr
	LocalAddr     net.Addr
}

type SocketCallbacks struct {
	ConnectRequest ConnectRequestCallback
	Connected      ConnectedCallback
	Disconnect     DisconnectCallback
}
