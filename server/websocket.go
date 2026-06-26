package server

import (
	"log"
	"net/http"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/internal/ws"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{Subprotocols: []string{"ocpp1.6"}, ReadBufferSize: 2048,
	WriteBufferSize: 2048, EnableCompression: true}

func (o *Server) wshandler(w http.ResponseWriter, r *http.Request) {
	cpid := r.PathValue("cpid")

	// Optional Basic Auth check
	if o.basicAuthEnabled {
		user, pass, ok := r.BasicAuth()
		if !ok || user != o.basicUser || pass != o.basicPass {
			w.Header().Set("WWW-Authenticate", "Basic realm=\"restricted\"")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}

	// Allow/deny via user-defined callback if provided
	if o.socketcallbacks.ConnectRequest != nil && !o.socketcallbacks.ConnectRequest(ocpp.ConnectRequest{R: r, W: &w}) {
		return
	}

	log.Printf("Received connection from %s", cpid)

	// Apply server-configured subprotocols if set
	if len(o.subprotocols) > 0 {
		upgrader.Subprotocols = o.subprotocols
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %s", err)
		return
	}
	defer c.Close()

	// Create a context
	context := ocpp16.NewOcppContext(cpid)

	// Choose parser: prefer custom parser when provided
	parser := o.parser
	if parser == nil {
		parser = o.ocppcallbacks.ParseMessage
	}
	// Use common runner with optional traffic logging and keepalive settings
	ws.Run(c, parser, context, o.socketcallbacks, &ws.Options{
		LogSent:      o.logTraffic,
		LogKeepalive: o.logKeepalive,
		PingInterval: o.pingInterval,
		PongTimeout:  o.pongTimeout,
	})
}
