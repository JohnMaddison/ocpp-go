package client

import "github.com/JohnMaddison/ocpp-go/ocpp16"

func (c *Client) WithOCPP16AuthorizeHandler(callback ocpp16.AuthorizeCallback) *Client {
	c.ocppCallbacks.Authorize = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16BootNotificationHandler(callback ocpp16.BootNotificationCallback) *Client {
	c.ocppCallbacks.BootNotification = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16CancelReservationHandler(callback ocpp16.CancelReservationCallback) *Client {
	c.ocppCallbacks.CancelReservation = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16CertificateSignedHandler(callback ocpp16.CertificateSignedCallback) *Client {
	c.ocppCallbacks.CertificateSigned = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16ChangeAvailabilityHandler(callback ocpp16.ChangeAvailabilityCallback) *Client {
	c.ocppCallbacks.ChangeAvailability = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16ChangeConfigurationHandler(callback ocpp16.ChangeConfigurationCallback) *Client {
	c.ocppCallbacks.ChangeConfiguration = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16ClearCacheHandler(callback ocpp16.ClearCacheCallback) *Client {
	c.ocppCallbacks.ClearCache = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16ClearChargingProfileHandler(callback ocpp16.ClearChargingProfileCallback) *Client {
	c.ocppCallbacks.ClearChargingProfile = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16ClearReservationHandler(callback ocpp16.ClearReservationCallback) *Client {
	c.ocppCallbacks.ClearReservation = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16DataTransferHandler(callback ocpp16.DataTransferCallback) *Client {
	c.ocppCallbacks.DataTransfer = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16DeleteCertificateHandler(callback ocpp16.DeleteCertificateCallback) *Client {
	c.ocppCallbacks.DeleteCertificate = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16DiagnosticsStatusNotificationHandler(callback ocpp16.DiagnosticsStatusNotificationCallback) *Client {
	c.ocppCallbacks.DiagnosticsStatusNotification = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16ExtendedTriggerMessageHandler(callback ocpp16.ExtendedTriggerMessageCallback) *Client {
	c.ocppCallbacks.ExtendedTriggerMessage = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16GetCompositeScheduleHandler(callback ocpp16.GetCompositeScheduleCallback) *Client {
	c.ocppCallbacks.GetCompositeSchedule = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16GetConfigurationHandler(callback ocpp16.GetConfigurationCallback) *Client {
	c.ocppCallbacks.GetConfiguration = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16GetDiagnosticsHandler(callback ocpp16.GetDiagnosticsCallback) *Client {
	c.ocppCallbacks.GetDiagnostics = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16GetInstalledCertificatesHandler(callback ocpp16.GetInstalledCertificatesCallback) *Client {
	c.ocppCallbacks.GetInstalledCertificates = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16GetLogHandler(callback ocpp16.GetLogCallback) *Client {
	c.ocppCallbacks.GetLog = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16HeartbeatHandler(callback ocpp16.HeartbeatCallback) *Client {
	c.ocppCallbacks.Heartbeat = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16InstallCertificateHandler(callback ocpp16.InstallCertificateCallback) *Client {
	c.ocppCallbacks.InstallCertificate = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16LogStatusNotificationHandler(callback ocpp16.LogStatusNotificationCallback) *Client {
	c.ocppCallbacks.LogStatusNotification = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16MeterValuesHandler(callback ocpp16.MeterValuesCallback) *Client {
	c.ocppCallbacks.MeterValues = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16RemoteStartTransactionHandler(callback ocpp16.RemoteStartTransactionCallback) *Client {
	c.ocppCallbacks.RemoteStartTransaction = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16RemoteStopTransactionHandler(callback ocpp16.RemoteStopTransactionCallback) *Client {
	c.ocppCallbacks.RemoteStopTransaction = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16ReserveNowHandler(callback ocpp16.ReserveNowCallback) *Client {
	c.ocppCallbacks.ReserveNow = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16ResetHandler(callback ocpp16.ResetCallback) *Client {
	c.ocppCallbacks.Reset = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16SendLocalListHandler(callback ocpp16.SendLocalListCallback) *Client {
	c.ocppCallbacks.SendLocalList = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16SetChargingProfileHandler(callback ocpp16.SetChargingProfileCallback) *Client {
	c.ocppCallbacks.SetChargingProfile = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16SignCertificateHandler(callback ocpp16.SignCertificateCallback) *Client {
	c.ocppCallbacks.SignCertificate = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16SignedFirmwareStatusNotificationHandler(callback ocpp16.SignedFirmwareStatusNotificationCallback) *Client {
	c.ocppCallbacks.SignedFirmwareStatusNotification = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16SignedUpdateFirmwareHandler(callback ocpp16.SignedUpdateFirmwareCallback) *Client {
	c.ocppCallbacks.SignedUpdateFirmware = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16StartTransactionHandler(callback ocpp16.StartTransactionCallback) *Client {
	c.ocppCallbacks.StartTransaction = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16StatusNotificationHandler(callback ocpp16.StatusNotificationCallback) *Client {
	c.ocppCallbacks.StatusNotification = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16StopTransactionHandler(callback ocpp16.StopTransactionCallback) *Client {
	c.ocppCallbacks.StopTransaction = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16TriggerMessageHandler(callback ocpp16.TriggerMessageCallback) *Client {
	c.ocppCallbacks.TriggerMessage = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16UnlockConnectorHandler(callback ocpp16.UnlockConnectorCallback) *Client {
	c.ocppCallbacks.UnlockConnector = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}

func (c *Client) WithOCPP16UpdateFirmwareHandler(callback ocpp16.UpdateFirmwareCallback) *Client {
	c.ocppCallbacks.UpdateFirmware = callback
	c.ocppCallbacks.InitHandlers()
	c.subprotocol = "ocpp1.6"
	return c
}
