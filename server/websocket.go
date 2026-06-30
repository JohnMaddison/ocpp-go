package server

import (
	"log"
	"net/http"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/internal/ws"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
	"github.com/JohnMaddison/ocpp-go/ocpp21"
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
	if o.socketCallbacks.ConnectRequest != nil && !o.socketCallbacks.ConnectRequest(ocpp.ConnectRequest{R: r, W: &w}) {
		return
	}

	protocols := o.protocols()
	if !hasNegotiableSubprotocol(r, protocols) {
		http.Error(w, "Unsupported WebSocket subprotocol", http.StatusBadRequest)
		return
	}

	localUpgrader := upgrader
	localUpgrader.Subprotocols = protocols

	c, err := localUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %s", err)
		return
	}
	defer c.Close()

	runtime, ok := o.runtime(cpid, c.Subprotocol())
	if !ok {
		log.Printf("unsupported negotiated subprotocol: %s", c.Subprotocol())
		return
	}

	// Use common runner with optional traffic logging and keepalive settings
	ws.Run(c, runtime, o.socketCallbacks, &ws.Options{
		LogSent:      o.logTraffic,
		LogKeepalive: o.logKeepalive,
		PingInterval: o.pingInterval,
		PongTimeout:  o.pongTimeout,
	})
}

func hasNegotiableSubprotocol(r *http.Request, supported []string) bool {
	offered := websocket.Subprotocols(r)
	for _, offer := range offered {
		for _, protocol := range supported {
			if offer == protocol {
				return true
			}
		}
	}
	return false
}

func (o *Server) runtime(cpid, protocol string) (ws.Runtime, bool) {
	switch protocol {
	case "ocpp1.6":
		context := ocpp16.NewOCPPContext(cpid)
		parser := o.parser
		if parser == nil {
			parser = o.ocppCallbacks.ParseMessage
		}
		return ws.Runtime{
			ChargePointID: context.ChargePointID,
			OutgoingCalls: requestChannel[ocpp16.Request](context.Queue),
			Parse: func(message []byte) ([]byte, error) {
				return parser(message, context)
			},
			Serialize: func(call any) ([]byte, error) {
				request := call.(ocpp16.Request)
				return request.Call.SerializeOCPP()
			},
		}, true
	case "ocpp2.1":
		context := ocpp21.NewOCPPContext(cpid)
		return ws.Runtime{
			ChargePointID: context.ChargePointID,
			OutgoingCalls: requestChannel[ocpp21.Request](context.Queue),
			Parse: func(message []byte) ([]byte, error) {
				return o.ocpp21Callbacks.ParseMessage(message, context)
			},
			Serialize: func(call any) ([]byte, error) {
				request := call.(ocpp21.Request)
				return request.Call.SerializeOCPP()
			},
		}, true
	default:
		return ws.Runtime{}, false
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
