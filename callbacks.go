package ocpp

import (
	"net/http"
)

// CONNECTION CALLBACKS
type ConnectRequestCallback func(ctx ConnectRequest) bool
type ConnectedCallback func()
type DisconnectCallback func()

type ConnectRequest struct {
	R *http.Request
	W *http.ResponseWriter
}

type SocketCallbacks struct {
	ConnectRequest ConnectRequestCallback
	Connected      ConnectedCallback
	Disconnect     DisconnectCallback
}
