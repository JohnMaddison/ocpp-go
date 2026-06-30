package ws

import (
	"log"
	"sync"
	"time"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
	"github.com/gorilla/websocket"
)

// Options controls behavior of the websocket runner.
type Options struct {
	// LogSent enables logging of outgoing websocket traffic (responses and calls).
	LogSent bool
	// Logf receives log lines; defaults to log.Printf when nil.
	Logf func(format string, args ...any)
	// PingInterval enables websocket pings when greater than zero.
	PingInterval time.Duration
	// PongTimeout controls how long to wait for the corresponding pong response.
	// When zero, a default of twice the PingInterval is used.
	PongTimeout time.Duration
	// LogKeepalive records ping/pong traffic when enabled.
	LogKeepalive bool
}

// Run starts read/write pumps for a websocket OCPP connection.
// It invokes socket callbacks for connect/disconnect lifecycle.
func Run(conn *websocket.Conn, parse func(message []byte, ctx *ocpp16.OCPPContext) ([]byte, error), ctx *ocpp16.OCPPContext, socketCallbacks ocpp.SocketCallbacks, opts ...*Options) {
	var opt *Options
	if len(opts) > 0 {
		opt = opts[0]
	}
	logf := log.Printf
	if opt != nil && opt.Logf != nil {
		logf = opt.Logf
	}
	var pingInterval time.Duration
	var pongTimeout time.Duration
	if opt != nil {
		pingInterval = opt.PingInterval
		pongTimeout = opt.PongTimeout
		if pingInterval > 0 && pongTimeout <= 0 {
			pongTimeout = pingInterval * 2
		}
	}

	if pingInterval > 0 && pongTimeout <= 0 {
		pongTimeout = pingInterval * 2
	}

	if pingInterval > 0 {
		if err := conn.SetReadDeadline(time.Now().Add(pongTimeout)); err != nil {
			logf("failed to set initial read deadline: %v", err)
		}
		conn.SetPongHandler(func(appData string) error {
			if opt != nil && opt.LogKeepalive {
				logf("recv pong")
			}
			return conn.SetReadDeadline(time.Now().Add(pongTimeout))
		})
	}
	// Connected callback
	connectionInfo := ocpp.ConnectionInfo{
		ChargePointID: ctx.ChargePointID,
		RemoteAddr:    conn.RemoteAddr(),
		LocalAddr:     conn.LocalAddr(),
	}
	if socketCallbacks.Connected != nil {
		socketCallbacks.Connected(connectionInfo)
	}

	outgoingResponses := make(chan []byte, 10)
	done := make(chan struct{})
	var closeDone sync.Once
	closeDoneFunc := func() { closeDone.Do(func() { close(done) }) }

	var wg sync.WaitGroup
	wg.Add(2)

	// Read goroutine
	go func() {
		defer wg.Done()
		defer closeDoneFunc()

		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if mt != websocket.TextMessage {
				log.Printf("Unsupported messages type: %d", mt)
				continue
			}
			if opt != nil && opt.LogSent {
				logf("recv: %s", string(message))
			}
			if pingInterval > 0 {
				if err := conn.SetReadDeadline(time.Now().Add(pongTimeout)); err != nil {
					logf("failed to extend read deadline: %v", err)
				}
			}
			response, err := parse(message, ctx)
			if err != nil {
				log.Printf("parseMessage error: %s", err)
				conn.Close()
				break
			}
			if response == nil {
				continue
			}

			select {
			case outgoingResponses <- response:
			case <-done:
				return
			}
		}
	}()

	// Write goroutine
	go func() {
		defer wg.Done()
		defer close(outgoingResponses)

		var pingTicker *time.Ticker
		var pingC <-chan time.Time
		if pingInterval > 0 {
			pingTicker = time.NewTicker(pingInterval)
			pingC = pingTicker.C
			defer pingTicker.Stop()
		}

		for {
			select {
			case response, ok := <-outgoingResponses:
				if !ok {
					return
				}

				if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
					log.Printf("WriteMessage error: %s", err)
					closeDoneFunc()
					return
				}
				if opt != nil && opt.LogSent {
					logf("sent: %s", string(response))
				}
			case message, ok := <-ctx.Queue:
				if !ok {
					return
				}

				payload, serializeErr := message.Call.SerializeOCPP()
				if serializeErr != nil {
					log.Printf("serialize error: %s", serializeErr)
					return
				}
				if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
					log.Printf("WriteMessage error: %s", err)
					closeDoneFunc()
					return
				}
				if opt != nil && opt.LogSent {
					logf("sent: %s", string(payload))
				}
			case <-pingC:
				deadline := time.Now().Add(pongTimeout)
				if err := conn.SetWriteDeadline(deadline); err != nil {
					log.Printf("SetWriteDeadline error: %s", err)
					closeDoneFunc()
					return
				}
				if err := conn.WriteControl(websocket.PingMessage, nil, deadline); err != nil {
					log.Printf("WriteControl ping error: %s", err)
					closeDoneFunc()
					return
				}
				// Reset deadline so subsequent writes aren't constrained by the ping timeout.
				_ = conn.SetWriteDeadline(time.Time{})
				if opt != nil && opt.LogKeepalive {
					logf("sent ping")
				}
			case <-done:
				return
			}
		}
	}()

	// Wait until both pumps stop
	wg.Wait()

	// Disconnected callback
	if socketCallbacks.Disconnect != nil {
		socketCallbacks.Disconnect(connectionInfo)
	}
}
