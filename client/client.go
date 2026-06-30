// Package client provides an OCPP charge point websocket client.
package client

import (
	"time"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
	"github.com/JohnMaddison/ocpp-go/ocpp21"
	"github.com/google/uuid"
)

type Client struct {
	chargePointID   string
	address         string
	ocppCallbacks   ocpp16.OCPPCallbacks
	ocpp21Callbacks ocpp21.OCPPCallbacks
	socketCallbacks ocpp.SocketCallbacks
	OCPPContext     *ocpp16.OCPPContext
	OCPP21Context   *ocpp21.OCPPContext
	subprotocol     string
	logTraffic      bool
	logKeepalive    bool
	username        string
	password        string
	pingInterval    time.Duration
	pongTimeout     time.Duration
}

func NewOCPP16Client(chargePointID, address string) *Client {
	client := &Client{
		chargePointID: chargePointID,
		address:       address,
		OCPPContext:   nil,
		subprotocol:   "ocpp1.6",
	}
	client.ocppCallbacks.InitHandlers()
	return client
}

func NewOCPP21Client(chargePointID, address string) *Client {
	client := &Client{
		chargePointID: chargePointID,
		address:       address,
		OCPP21Context: nil,
		subprotocol:   "ocpp2.1",
	}
	client.ocpp21Callbacks.InitHandlers()
	return client
}

// EnableTrafficLogging toggles websocket sent-traffic logging.
func (c *Client) WithTrafficLogging() *Client {
	c.logTraffic = true
	return c
}

// EnableKeepaliveLogging toggles logging of websocket ping/pong frames.
func (c *Client) WithKeepaliveLogging() *Client {
	c.logKeepalive = true
	return c
}

func (c *Client) WithAddress(address string) *Client {
	c.address = address
	return c
}

func (c *Client) WithOCPP21Callbacks(callbacks ocpp21.OCPPCallbacks) *Client {
	callbacks.InitHandlers()
	c.ocpp21Callbacks = callbacks
	c.subprotocol = "ocpp2.1"
	return c
}

// WithBasicAuth configures HTTP Basic authentication for the websocket handshake.
func (c *Client) WithBasicAuth(username, password string) *Client {
	c.username = username
	c.password = password
	return c
}

func (c *Client) WithConnectedHandler(callback ocpp.ConnectedCallback) *Client {
	c.socketCallbacks.Connected = callback
	return c
}

func (c *Client) WithDisconnectHandler(callback ocpp.DisconnectCallback) *Client {
	c.socketCallbacks.Disconnect = callback
	return c
}

func (c *Client) WithOfflineHandler(callback ocpp.DisconnectCallback) *Client {
	return c.WithDisconnectHandler(callback)
}

// WithWebsocketKeepalive configures websocket ping/pong behaviour.
// A non-positive interval disables keepalive.
// When pongTimeout is zero or negative, a default of twice the interval is used.
func (c *Client) WithWebsocketKeepalive(interval, pongTimeout time.Duration) *Client {
	c.pingInterval = interval
	c.pongTimeout = pongTimeout
	return c
}

func (c *Client) WithAuthorizeHandler(callback ocpp16.AuthorizeCallback) *Client {
	c.ocppCallbacks.Authorize = callback
	return c
}

func (c *Client) WithBootNotificationHandler(callback ocpp16.BootNotificationCallback) *Client {
	c.ocppCallbacks.BootNotification = callback
	return c
}

func (c *Client) WithCancelReservationHandler(callback ocpp16.CancelReservationCallback) *Client {
	c.ocppCallbacks.CancelReservation = callback
	return c
}

func (c *Client) WithCertificateSignedHandler(callback ocpp16.CertificateSignedCallback) *Client {
	c.ocppCallbacks.CertificateSigned = callback
	return c
}

func (c *Client) WithChangeAvailabilityHandler(callback ocpp16.ChangeAvailabilityCallback) *Client {
	c.ocppCallbacks.ChangeAvailability = callback
	return c
}

func (c *Client) WithChangeConfigurationHandler(callback ocpp16.ChangeConfigurationCallback) *Client {
	c.ocppCallbacks.ChangeConfiguration = callback
	return c
}

func (c *Client) WithClearCacheHandler(callback ocpp16.ClearCacheCallback) *Client {
	c.ocppCallbacks.ClearCache = callback
	return c
}

func (c *Client) WithClearChargingProfileHandler(callback ocpp16.ClearChargingProfileCallback) *Client {
	c.ocppCallbacks.ClearChargingProfile = callback
	return c
}

func (c *Client) WithClearReservationHandler(callback ocpp16.ClearReservationCallback) *Client {
	c.ocppCallbacks.ClearReservation = callback
	return c
}

func (c *Client) WithDataTransferHandler(callback ocpp16.DataTransferCallback) *Client {
	c.ocppCallbacks.DataTransfer = callback
	return c
}

func (c *Client) WithDeleteCertificateHandler(callback ocpp16.DeleteCertificateCallback) *Client {
	c.ocppCallbacks.DeleteCertificate = callback
	return c
}

func (c *Client) WithDiagnosticsStatusNotificationHandler(callback ocpp16.DiagnosticsStatusNotificationCallback) *Client {
	c.ocppCallbacks.DiagnosticsStatusNotification = callback
	return c
}

func (c *Client) WithExtendedTriggerMessageHandler(callback ocpp16.ExtendedTriggerMessageCallback) *Client {
	c.ocppCallbacks.ExtendedTriggerMessage = callback
	return c
}

func (c *Client) WithGetCompositeScheduleHandler(callback ocpp16.GetCompositeScheduleCallback) *Client {
	c.ocppCallbacks.GetCompositeSchedule = callback
	return c
}

func (c *Client) WithGetConfigurationHandler(callback ocpp16.GetConfigurationCallback) *Client {
	c.ocppCallbacks.GetConfiguration = callback
	return c
}

func (c *Client) WithConfigurationHandler(callback ocpp16.GetConfigurationCallback) *Client {
	return c.WithGetConfigurationHandler(callback)
}

func (c *Client) WithGetDiagnosticsHandler(callback ocpp16.GetDiagnosticsCallback) *Client {
	c.ocppCallbacks.GetDiagnostics = callback
	return c
}

func (c *Client) WithGetInstalledCertificatesHandler(callback ocpp16.GetInstalledCertificatesCallback) *Client {
	c.ocppCallbacks.GetInstalledCertificates = callback
	return c
}

func (c *Client) WithGetLogHandler(callback ocpp16.GetLogCallback) *Client {
	c.ocppCallbacks.GetLog = callback
	return c
}

func (c *Client) WithHeartbeatHandler(callback ocpp16.HeartbeatCallback) *Client {
	c.ocppCallbacks.Heartbeat = callback
	return c
}

func (c *Client) WithInstallCertificateHandler(callback ocpp16.InstallCertificateCallback) *Client {
	c.ocppCallbacks.InstallCertificate = callback
	return c
}

func (c *Client) WithLogStatusNotificationHandler(callback ocpp16.LogStatusNotificationCallback) *Client {
	c.ocppCallbacks.LogStatusNotification = callback
	return c
}

func (c *Client) WithMeterValuesHandler(callback ocpp16.MeterValuesCallback) *Client {
	c.ocppCallbacks.MeterValues = callback
	return c
}

func (c *Client) WithRemoteStartTransactionHandler(callback ocpp16.RemoteStartTransactionCallback) *Client {
	c.ocppCallbacks.RemoteStartTransaction = callback
	return c
}

func (c *Client) WithRemoteStopTransactionHandler(callback ocpp16.RemoteStopTransactionCallback) *Client {
	c.ocppCallbacks.RemoteStopTransaction = callback
	return c
}

func (c *Client) WithReserveNowHandler(callback ocpp16.ReserveNowCallback) *Client {
	c.ocppCallbacks.ReserveNow = callback
	return c
}

func (c *Client) WithResetHandler(callback ocpp16.ResetCallback) *Client {
	c.ocppCallbacks.Reset = callback
	return c
}

func (c *Client) WithSendLocalListHandler(callback ocpp16.SendLocalListCallback) *Client {
	c.ocppCallbacks.SendLocalList = callback
	return c
}

func (c *Client) WithSetChargingProfileHandler(callback ocpp16.SetChargingProfileCallback) *Client {
	c.ocppCallbacks.SetChargingProfile = callback
	return c
}

func (c *Client) WithSignCertificateHandler(callback ocpp16.SignCertificateCallback) *Client {
	c.ocppCallbacks.SignCertificate = callback
	return c
}

func (c *Client) WithSignedFirmwareStatusNotificationHandler(callback ocpp16.SignedFirmwareStatusNotificationCallback) *Client {
	c.ocppCallbacks.SignedFirmwareStatusNotification = callback
	return c
}

func (c *Client) WithSignedUpdateFirmwareHandler(callback ocpp16.SignedUpdateFirmwareCallback) *Client {
	c.ocppCallbacks.SignedUpdateFirmware = callback
	return c
}

func (c *Client) WithStartTransactionHandler(callback ocpp16.StartTransactionCallback) *Client {
	c.ocppCallbacks.StartTransaction = callback
	return c
}

func (c *Client) WithStatusNotificationHandler(callback ocpp16.StatusNotificationCallback) *Client {
	c.ocppCallbacks.StatusNotification = callback
	return c
}

func (c *Client) WithStopTransactionHandler(callback ocpp16.StopTransactionCallback) *Client {
	c.ocppCallbacks.StopTransaction = callback
	return c
}

func (c *Client) WithTriggerMessageHandler(callback ocpp16.TriggerMessageCallback) *Client {
	c.ocppCallbacks.TriggerMessage = callback
	return c
}

func (c *Client) WithUnlockConnectorHandler(callback ocpp16.UnlockConnectorCallback) *Client {
	c.ocppCallbacks.UnlockConnector = callback
	return c
}

func (c *Client) WithUpdateFirmwareHandler(callback ocpp16.UpdateFirmwareCallback) *Client {
	c.ocppCallbacks.UpdateFirmware = callback
	return c
}

func (c *Client) SendOCPP16Call(action ocpp16.Action, payload any) (*ocpp16.ResultOrError, error) {
	return c.OCPPContext.Send(ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: uuid.New().String(), Action: string(action), Payload: payload})
}

func (c *Client) SendOCPP21Call(action ocpp21.Action, payload any) (*ocpp21.ResultOrError, error) {
	return c.OCPP21Context.Send(ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: uuid.New().String(), Action: string(action), Payload: payload})
}
