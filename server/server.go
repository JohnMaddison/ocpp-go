// Package server provides an OCPP central system websocket server.
package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/johnmaddison/ocpp-go"
	"github.com/johnmaddison/ocpp-go/internal/uuidgenerator"
	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/johnmaddison/ocpp-go/ocpp21"
)

type Server struct {
	mu               sync.Mutex
	address          string
	path             string
	httpServer       *http.Server
	activeWebsockets map[*websocket.Conn]struct{}
	sessions         map[string]*Session
	ocppCallbacks    ocpp16.Callbacks
	ocpp21Callbacks  ocpp21.Callbacks
	socketCallbacks  ocpp.SocketCallbacks
	logTraffic       bool
	logKeepalive     bool
	parser           func(message []byte, ctx *ocpp16.Context) ([]byte, error)
	subprotocols     []string
	ocpp16Enabled    bool
	ocpp21Enabled    bool
	// Basic auth (optional)
	basicAuthEnabled bool
	basicUser        string
	basicPass        string
	// TLS (optional)
	tlsEnabled bool
	tlsCert    string
	tlsKey     string
	// Websocket keepalive options
	pingInterval       time.Duration
	pongTimeout        time.Duration
	messageIDGenerator uuidgenerator.MessageIDGeneratorMethod
}

// Option configures a Server.
type Option func(*Server)

// NewServer creates a new configurable OCPP server using options.
func NewServer(address string, opts ...Option) *Server {
	s := &Server{
		address:            address,
		path:               "ocpp",
		activeWebsockets:   make(map[*websocket.Conn]struct{}),
		sessions:           make(map[string]*Session),
		messageIDGenerator: uuidgenerator.DefaultUUIDGenerator,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func ensure16Parser(s *Server) {
	s.ocppCallbacks.InitHandlers()
	s.ocpp16Enabled = true
	if s.parser == nil {
		s.parser = s.ocppCallbacks.ParseMessage
	}
}

// With16Callbacks sets OCPP 1.6 callbacks and uses its ParseMessage.
func With16Callbacks(cb ocpp16.Callbacks) Option {
	return func(s *Server) {
		// Ensure callback handlers are initialized so ParseMessage can route actions
		cb.InitHandlers()
		s.ocppCallbacks = cb
		s.parser = s.ocppCallbacks.ParseMessage
		s.ocpp16Enabled = true
	}
}

// With21Callbacks sets OCPP 2.1 callbacks and uses its ParseMessage for ocpp2.1 connections.
func With21Callbacks(cb ocpp21.Callbacks) Option {
	return func(s *Server) {
		cb.InitHandlers()
		s.ocpp21Callbacks = cb
		s.ocpp21Enabled = true
	}
}

// WithParser sets a custom message parse function.
func WithParser(p func(message []byte, ctx *ocpp16.Context) ([]byte, error)) Option {
	return func(s *Server) {
		s.parser = p
		s.ocpp16Enabled = true
	}
}

// WithSocketCallbacks sets socket lifecycle callbacks.
func WithSocketCallbacks(cb ocpp.SocketCallbacks) Option {
	return func(s *Server) { s.socketCallbacks = cb }
}

func WithConnectRequestHandler(callback ocpp.ConnectRequestCallback) Option {
	return func(s *Server) { s.socketCallbacks.ConnectRequest = callback }
}

func WithConnectedHandler(callback ocpp.ConnectedCallback) Option {
	return func(s *Server) { s.socketCallbacks.Connected = callback }
}

func WithDisconnectHandler(callback ocpp.DisconnectCallback) Option {
	return func(s *Server) { s.socketCallbacks.Disconnect = callback }
}

// WithPath sets the URL base path (default: "ocpp").
func WithPath(path string) Option { return func(s *Server) { s.path = path } }

// WithTrafficLogging enables/disables websocket sent-traffic logging.
func WithTrafficLogging() Option { return func(s *Server) { s.logTraffic = true } }

// WithSubprotocols sets allowed WebSocket subprotocols (default: ["ocpp1.6"]).
func WithSubprotocols(protocols ...string) Option {
	return func(s *Server) { s.subprotocols = protocols }
}

func (s *Server) protocols() []string {
	if len(s.subprotocols) > 0 {
		return s.subprotocols
	}
	if s.ocpp16Enabled && s.ocpp21Enabled {
		return []string{"ocpp1.6", "ocpp2.1"}
	}
	if s.ocpp21Enabled {
		return []string{"ocpp2.1"}
	}
	return []string{"ocpp1.6"}
}

// WithWebsocketKeepalive configures periodic websocket pings and pong timeout.
// A non-positive interval disables the keepalive.
// When pongTimeout is zero or negative, a default of twice the interval is used.
func WithWebsocketKeepalive(interval, pongTimeout time.Duration) Option {
	return func(s *Server) {
		s.pingInterval = interval
		s.pongTimeout = pongTimeout
	}
}

// WithKeepaliveLogging enables logging of websocket ping/pong frames.
func WithKeepaliveLogging() Option { return func(s *Server) { s.logKeepalive = true } }

func WithMessageIDGenerator(f func() string) Option {
	return func(s *Server) {
		if f != nil {
			s.messageIDGenerator = f
		}
	}
}

// WithBasicAuth enables HTTP Basic Auth using the given username and password.
func WithBasicAuth(username, password string) Option {
	return func(s *Server) {
		s.basicAuthEnabled = true
		s.basicUser = username
		s.basicPass = password
	}
}

// WithTLS enables TLS using the provided certificate and key files.
func WithTLS(certFile, keyFile string) Option {
	return func(s *Server) {
		s.tlsEnabled = true
		s.tlsCert = certFile
		s.tlsKey = keyFile
	}
}

// Serve starts the HTTP server using the configured path.
func (s *Server) Serve() error {
	// Ensure a parser is configured. If none provided, initialize OCPP 1.6 callbacks parser.
	if !s.ocpp21Enabled {
		s.ocpp16Enabled = true
	}
	s.ocppCallbacks.InitHandlers()
	if s.parser == nil {

		s.parser = s.ocppCallbacks.ParseMessage
	}

	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("/%s/{cpid}", s.path), func(w http.ResponseWriter, r *http.Request) {
		s.wshandler(w, r)
	})

	httpServer := &http.Server{
		Addr:    s.address,
		Handler: mux,
	}
	s.mu.Lock()
	s.httpServer = httpServer
	s.mu.Unlock()

	if s.tlsEnabled {
		return httpServer.ListenAndServeTLS(s.tlsCert, s.tlsKey)
	}
	return httpServer.ListenAndServe()
}

// ListenAndServe starts the server after setting the URL base path.
func (s *Server) ListenAndServe(path string) error { s.path = path; return s.Serve() }

// Shutdown gracefully shuts down the underlying HTTP server.
//
// Shutdown stops accepting new connections and waits for active HTTP handlers to
// return until ctx expires, matching net/http.Server.Shutdown semantics.
func (s *Server) Shutdown(ctx context.Context) error {
	s.mu.Lock()
	httpServer := s.httpServer
	s.mu.Unlock()
	if httpServer == nil {
		return http.ErrServerClosed
	}

	err := httpServer.Shutdown(ctx)
	s.closeActiveWebsockets()

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	for {
		if s.activeWebsocketCount() == 0 {
			return err
		}
		select {
		case <-ticker.C:
		case <-ctx.Done():
			if err != nil {
				return err
			}
			return ctx.Err()
		}
	}
}
