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

func ensureOCPP16Parser(s *Server) {
	s.ocppCallbacks.InitHandlers()
	if s.parser == nil {
		s.parser = s.ocppCallbacks.ParseMessage
	}
}

// WithOCPP16Callbacks sets OCPP 1.6 callbacks and uses its ParseMessage.
func WithOCPP16Callbacks(cb ocpp16.OCPPCallbacks) Option {
	return func(s *Server) {
		// Ensure callback handlers are initialized so ParseMessage can route actions
		cb.InitHandlers()
		s.ocppCallbacks = cb
		s.parser = s.ocppCallbacks.ParseMessage
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

func WithConnectRequestHandler(callback ocpp.ConnectRequestCallback) Option {
	return func(s *Server) { s.socketCallbacks.ConnectRequest = callback }
}

func WithConnectedHandler(callback ocpp.ConnectedCallback) Option {
	return func(s *Server) { s.socketCallbacks.Connected = callback }
}

func WithDisconnectHandler(callback ocpp.DisconnectCallback) Option {
	return func(s *Server) { s.socketCallbacks.Disconnect = callback }
}

func WithAuthorizeHandler(callback ocpp16.AuthorizeCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.Authorize = callback
		ensureOCPP16Parser(s)
	}
}

func WithBootNotificationHandler(callback ocpp16.BootNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.BootNotification = callback
		ensureOCPP16Parser(s)
	}
}

func WithCancelReservationHandler(callback ocpp16.CancelReservationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.CancelReservation = callback
		ensureOCPP16Parser(s)
	}
}

func WithCertificateSignedHandler(callback ocpp16.CertificateSignedCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.CertificateSigned = callback
		ensureOCPP16Parser(s)
	}
}

func WithChangeAvailabilityHandler(callback ocpp16.ChangeAvailabilityCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ChangeAvailability = callback
		ensureOCPP16Parser(s)
	}
}

func WithChangeConfigurationHandler(callback ocpp16.ChangeConfigurationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ChangeConfiguration = callback
		ensureOCPP16Parser(s)
	}
}

func WithClearCacheHandler(callback ocpp16.ClearCacheCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ClearCache = callback
		ensureOCPP16Parser(s)
	}
}

func WithClearChargingProfileHandler(callback ocpp16.ClearChargingProfileCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ClearChargingProfile = callback
		ensureOCPP16Parser(s)
	}
}

func WithClearReservationHandler(callback ocpp16.ClearReservationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ClearReservation = callback
		ensureOCPP16Parser(s)
	}
}

func WithDataTransferHandler(callback ocpp16.DataTransferCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.DataTransfer = callback
		ensureOCPP16Parser(s)
	}
}

func WithDeleteCertificateHandler(callback ocpp16.DeleteCertificateCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.DeleteCertificate = callback
		ensureOCPP16Parser(s)
	}
}

func WithDiagnosticsStatusNotificationHandler(callback ocpp16.DiagnosticsStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.DiagnosticsStatusNotification = callback
		ensureOCPP16Parser(s)
	}
}

func WithExtendedTriggerMessageHandler(callback ocpp16.ExtendedTriggerMessageCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ExtendedTriggerMessage = callback
		ensureOCPP16Parser(s)
	}
}

func WithGetCompositeScheduleHandler(callback ocpp16.GetCompositeScheduleCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetCompositeSchedule = callback
		ensureOCPP16Parser(s)
	}
}

func WithGetConfigurationHandler(callback ocpp16.GetConfigurationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetConfiguration = callback
		ensureOCPP16Parser(s)
	}
}

func WithConfigurationHandler(callback ocpp16.GetConfigurationCallback) Option {
	return WithGetConfigurationHandler(callback)
}

func WithGetDiagnosticsHandler(callback ocpp16.GetDiagnosticsCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetDiagnostics = callback
		ensureOCPP16Parser(s)
	}
}

func WithGetInstalledCertificatesHandler(callback ocpp16.GetInstalledCertificatesCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetInstalledCertificates = callback
		ensureOCPP16Parser(s)
	}
}

func WithGetLogHandler(callback ocpp16.GetLogCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetLog = callback
		ensureOCPP16Parser(s)
	}
}

func WithHeartbeatHandler(callback ocpp16.HeartbeatCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.Heartbeat = callback
		ensureOCPP16Parser(s)
	}
}

func WithInstallCertificateHandler(callback ocpp16.InstallCertificateCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.InstallCertificate = callback
		ensureOCPP16Parser(s)
	}
}

func WithLogStatusNotificationHandler(callback ocpp16.LogStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.LogStatusNotification = callback
		ensureOCPP16Parser(s)
	}
}

func WithMeterValuesHandler(callback ocpp16.MeterValuesCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.MeterValues = callback
		ensureOCPP16Parser(s)
	}
}

func WithRemoteStartTransactionHandler(callback ocpp16.RemoteStartTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.RemoteStartTransaction = callback
		ensureOCPP16Parser(s)
	}
}

func WithRemoteStopTransactionHandler(callback ocpp16.RemoteStopTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.RemoteStopTransaction = callback
		ensureOCPP16Parser(s)
	}
}

func WithReserveNowHandler(callback ocpp16.ReserveNowCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ReserveNow = callback
		ensureOCPP16Parser(s)
	}
}

func WithResetHandler(callback ocpp16.ResetCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.Reset = callback
		ensureOCPP16Parser(s)
	}
}

func WithSendLocalListHandler(callback ocpp16.SendLocalListCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SendLocalList = callback
		ensureOCPP16Parser(s)
	}
}

func WithSetChargingProfileHandler(callback ocpp16.SetChargingProfileCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SetChargingProfile = callback
		ensureOCPP16Parser(s)
	}
}

func WithSignCertificateHandler(callback ocpp16.SignCertificateCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SignCertificate = callback
		ensureOCPP16Parser(s)
	}
}

func WithSignedFirmwareStatusNotificationHandler(callback ocpp16.SignedFirmwareStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SignedFirmwareStatusNotification = callback
		ensureOCPP16Parser(s)
	}
}

func WithSignedUpdateFirmwareHandler(callback ocpp16.SignedUpdateFirmwareCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SignedUpdateFirmware = callback
		ensureOCPP16Parser(s)
	}
}

func WithStartTransactionHandler(callback ocpp16.StartTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.StartTransaction = callback
		ensureOCPP16Parser(s)
	}
}

func WithStatusNotificationHandler(callback ocpp16.StatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.StatusNotification = callback
		ensureOCPP16Parser(s)
	}
}

func WithStopTransactionHandler(callback ocpp16.StopTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.StopTransaction = callback
		ensureOCPP16Parser(s)
	}
}

func WithTriggerMessageHandler(callback ocpp16.TriggerMessageCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.TriggerMessage = callback
		ensureOCPP16Parser(s)
	}
}

func WithUnlockConnectorHandler(callback ocpp16.UnlockConnectorCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.UnlockConnector = callback
		ensureOCPP16Parser(s)
	}
}

func WithUpdateFirmwareHandler(callback ocpp16.UpdateFirmwareCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.UpdateFirmware = callback
		ensureOCPP16Parser(s)
	}
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
