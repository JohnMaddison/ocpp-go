// Package server provides an OCPP central system websocket server.
package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
)

type Server struct {
	address         string
	path            string
	ocppCallbacks   ocpp16.OCPPCallbacks
	socketCallbacks ocpp.SocketCallbacks
	logTraffic      bool
	logKeepalive    bool
	parser          func(message []byte, ctx *ocpp16.OCPPContext) ([]byte, error)
	subprotocols    []string
	// Basic auth (optional)
	basicAuthEnabled bool
	basicUser        string
	basicPass        string
	// TLS (optional)
	tlsEnabled bool
	tlsCert    string
	tlsKey     string
	// Websocket keepalive options
	pingInterval time.Duration
	pongTimeout  time.Duration
}

func NewServerWithCallbacks(address string, ocppCallbacks ocpp16.OCPPCallbacks, socketCallbacks ocpp.SocketCallbacks) *Server {
	return &Server{
		address:         address,
		path:            "ocpp",
		ocppCallbacks:   ocppCallbacks,
		socketCallbacks: socketCallbacks,
		parser:          ocppCallbacks.ParseMessage,
	}
}

// Option configures a Server.
type Option func(*Server)

// NewServer creates a new configurable OCPP server using options.
func NewServer(address string, opts ...Option) *Server {
	s := &Server{
		address: address,
		path:    "ocpp",
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// WithOCPP16Callbacks sets OCPP 1.6 callbacks and uses its ParseMessage.
func WithOCPP16Callbacks(cb ocpp16.OCPPCallbacks) Option {
	return func(s *Server) {
		// Ensure callback handlers are initialized so ParseMessage can route actions
		cb.InitHandlers()
		s.ocppCallbacks = cb
		s.parser = cb.ParseMessage
	}
}

// WithParser sets a custom message parse function.
func WithParser(p func(message []byte, ctx *ocpp16.OCPPContext) ([]byte, error)) Option {
	return func(s *Server) { s.parser = p }
}

// WithSocketCallbacks sets socket lifecycle callbacks.
func WithSocketCallbacks(cb ocpp.SocketCallbacks) Option {
	return func(s *Server) { s.socketCallbacks = cb }
}

// WithPath sets the URL base path (default: "ocpp").
func WithPath(path string) Option { return func(s *Server) { s.path = path } }

// WithTrafficLogging enables/disables websocket sent-traffic logging.
func WithTrafficLogging() Option { return func(s *Server) { s.logTraffic = true } }

// WithSubprotocols sets allowed WebSocket subprotocols (default: ["ocpp1.6"]).
func WithSubprotocols(protocols ...string) Option {
	return func(s *Server) { s.subprotocols = protocols }
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
	s.ocppCallbacks.InitHandlers()
	if s.parser == nil {

		s.parser = s.ocppCallbacks.ParseMessage
	}
	http.HandleFunc(fmt.Sprintf("/%s/{cpid}", s.path), func(w http.ResponseWriter, r *http.Request) {
		s.wshandler(w, r)
	})
	if s.tlsEnabled {
		return http.ListenAndServeTLS(s.address, s.tlsCert, s.tlsKey, nil)
	}
	return http.ListenAndServe(s.address, nil)
}

// ListenAndServe starts the server after setting the URL base path.
func (s *Server) ListenAndServe(path string) error { s.path = path; return s.Serve() }
