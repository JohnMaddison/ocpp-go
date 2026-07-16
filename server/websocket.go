package server

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/johnmaddison/ocpp-go"
	"github.com/johnmaddison/ocpp-go/internal/ws"
	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/johnmaddison/ocpp-go/ocpp21"
)

func (o *Server) wshandler(w http.ResponseWriter, r *http.Request) {
	cpid := r.PathValue("cpid")

	if o.basicAuthEnabled {
		user, pass, ok := r.BasicAuth()
		if !ok || user != o.basicUser || pass != o.basicPass {
			w.Header().Set("WWW-Authenticate", "Basic realm=\"restricted\"")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}

	if o.socketCallbacks.ConnectRequest != nil && !o.socketCallbacks.ConnectRequest(ocpp.ConnectRequest{R: r, W: &w}) {
		return
	}

	protocols := o.protocols()
	if !hasNegotiableSubprotocol(r, protocols) {
		http.Error(w, "Unsupported WebSocket subprotocol", http.StatusBadRequest)
		return
	}

	localUpgrader := websocket.Upgrader{
		Subprotocols:      protocols,
		ReadBufferSize:    2048,
		WriteBufferSize:   2048,
		EnableCompression: o.websocketCompression,
	}

	c, err := localUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %s", err)
		return
	}
	defer c.Close()
	o.registerWebsocket(c)
	defer o.unregisterWebsocket(c)

	runtime, session, ok := o.runtime(cpid, c.Subprotocol(), c.RemoteAddr(), c.LocalAddr())
	if !ok {
		log.Printf("unsupported negotiated subprotocol: %s", c.Subprotocol())
		return
	}
	o.registerSession(session)
	defer o.unregisterSession(session)

	ws.Run(c, runtime, o.socketCallbacks, &ws.Options{
		LogSent:      o.logTraffic,
		LogKeepalive: o.logKeepalive,
		PingInterval: o.pingInterval,
		PongTimeout:  o.pongTimeout,
		ReadLimit:    o.websocketReadLimit,
	})
}

func (o *Server) registerWebsocket(conn *websocket.Conn) {
	o.mu.Lock()
	if o.activeWebsockets == nil {
		o.activeWebsockets = make(map[*websocket.Conn]struct{})
	}
	o.activeWebsockets[conn] = struct{}{}
	o.mu.Unlock()
}

func (o *Server) unregisterWebsocket(conn *websocket.Conn) {
	o.mu.Lock()
	delete(o.activeWebsockets, conn)
	o.mu.Unlock()
}

func (o *Server) closeActiveWebsockets() {
	o.mu.Lock()
	conns := make([]*websocket.Conn, 0, len(o.activeWebsockets))
	for conn := range o.activeWebsockets {
		conns = append(conns, conn)
	}
	o.mu.Unlock()

	deadline := time.Now().Add(time.Second)
	for _, conn := range conns {
		_ = conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server shutting down"), deadline)
		_ = conn.Close()
	}
}

func (o *Server) activeWebsocketCount() int {
	o.mu.Lock()
	defer o.mu.Unlock()
	return len(o.activeWebsockets)
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

func (o *Server) runtime(cpid, protocol string, remoteAddr, localAddr net.Addr) (ws.Runtime, *Session, bool) {
	switch protocol {
	case "ocpp1.6":
		context := ocpp16.NewContextWithMessageIDGenerator(cpid, o.messageIDGenerator)
		session := &Session{
			chargePointID: cpid,
			protocol:      protocol,
			remoteAddr:    remoteAddr,
			localAddr:     localAddr,
			ocpp16Context: context,
		}
		parser := o.parser
		if parser == nil {
			parser = o.ocppCallbacks.ParseMessage
		}
		return ws.Runtime{
			ChargePointID: context.ChargePointID,
			Protocol:      protocol,
			OutgoingCalls: requestChannel[ocpp16.Request](context.Queue),
			Parse: func(message []byte) ([]byte, error) {
				return parser(message, context)
			},
			Serialize: func(call any) ([]byte, error) {
				request := call.(ocpp16.Request)
				return request.Call.SerializeOCPP()
			},
		}, session, true
	case "ocpp2.1":
		context := ocpp21.NewContextWithMessageIDGenerator(cpid, o.messageIDGenerator)
		session := &Session{
			chargePointID: cpid,
			protocol:      protocol,
			remoteAddr:    remoteAddr,
			localAddr:     localAddr,
			ocpp21Context: context,
		}
		return ws.Runtime{
			ChargePointID: context.ChargePointID,
			Protocol:      protocol,
			OutgoingCalls: requestChannel[ocpp21.Request](context.Queue),
			Parse: func(message []byte) ([]byte, error) {
				return o.ocpp21Callbacks.ParseMessage(message, context)
			},
			Serialize: func(call any) ([]byte, error) {
				request := call.(ocpp21.Request)
				return request.Call.SerializeOCPP()
			},
		}, session, true
	default:
		return ws.Runtime{}, nil, false
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
