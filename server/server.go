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
	ocppcallbacks   ocpp16.OcppCallbacks
	socketcallbacks ocpp.SocketCallbacks
	logTraffic      bool
	logKeepalive    bool
	parser          func(message []byte, ctx *ocpp16.OcppContext) ([]byte, error)
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

func NewExtensibleOcppServer(address string, ocppcallbacks ocpp16.OcppCallbacks, socketcallbacks ocpp.SocketCallbacks) *Server {
	return &Server{
		address:         address,
		path:            "ocpp",
		ocppcallbacks:   ocppcallbacks,
		socketcallbacks: socketcallbacks,
		parser:          ocppcallbacks.ParseMessage,
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

// WithOcpp16Callbacks sets OCPP 1.6 callbacks and uses its ParseMessage.
func WithOcpp16Callbacks(cb ocpp16.OcppCallbacks) Option {
	return func(s *Server) {
		// Ensure callback handlers are initialized so ParseMessage can route actions
		cb.InitHandlers()
		s.ocppcallbacks = cb
		s.parser = cb.ParseMessage
	}
}

// WithParser sets a custom message parse function.
func WithParser(p func(message []byte, ctx *ocpp16.OcppContext) ([]byte, error)) Option {
	return func(s *Server) { s.parser = p }
}

// WithSocketCallbacks sets socket lifecycle callbacks.
func WithSocketCallbacks(cb ocpp.SocketCallbacks) Option {
	return func(s *Server) { s.socketcallbacks = cb }
}

// WithPath sets the URL base path (default: "ocpp").
func WithPath(path string) Option { return func(s *Server) { s.path = path } }

// WithTrafficLogging enables/disables websocket sent-traffic logging.
func WithTrafficLogging(enable bool) Option { return func(s *Server) { s.logTraffic = enable } }

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
func WithKeepaliveLogging(enable bool) Option { return func(s *Server) { s.logKeepalive = enable } }

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

// EnableTrafficLogging toggles websocket sent-traffic logging.
func (s *Server) EnableTrafficLogging(enable bool) *Server {
	s.logTraffic = enable
	return s
}

// EnableKeepaliveLogging toggles websocket ping/pong logging.
func (s *Server) EnableKeepaliveLogging(enable bool) *Server {
	s.logKeepalive = enable
	return s
}

// Serve starts the HTTP server using the configured path.
func (s *Server) Serve() error {
	// Ensure a parser is configured. If none provided, initialize OCPP 1.6 callbacks parser.
	s.ocppcallbacks.InitHandlers()
	if s.parser == nil {

		s.parser = s.ocppcallbacks.ParseMessage
	}
	http.HandleFunc(fmt.Sprintf("/%s/{cpid}", s.path), func(w http.ResponseWriter, r *http.Request) {
		s.wshandler(w, r)
	})
	if s.tlsEnabled {
		return http.ListenAndServeTLS(s.address, s.tlsCert, s.tlsKey, nil)
	}
	return http.ListenAndServe(s.address, nil)
}

// Backward-compatible entrypoint that accepts the path explicitly.
func (s *Server) ListenAndServe(path string) error { s.path = path; return s.Serve() }
