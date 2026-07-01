package client

import "github.com/johnmaddison/ocpp-go/ocpp21"

func (c *Client) WithOCPP21AFRRSignalHandler(callback ocpp21.AFRRSignalCallback) *Client {
	c.ocpp21Callbacks.AFRRSignal = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21AdjustPeriodicEventStreamHandler(callback ocpp21.AdjustPeriodicEventStreamCallback) *Client {
	c.ocpp21Callbacks.AdjustPeriodicEventStream = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21AuthorizeHandler(callback ocpp21.AuthorizeCallback) *Client {
	c.ocpp21Callbacks.Authorize = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21BatterySwapHandler(callback ocpp21.BatterySwapCallback) *Client {
	c.ocpp21Callbacks.BatterySwap = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21BootNotificationHandler(callback ocpp21.BootNotificationCallback) *Client {
	c.ocpp21Callbacks.BootNotification = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21CancelReservationHandler(callback ocpp21.CancelReservationCallback) *Client {
	c.ocpp21Callbacks.CancelReservation = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21CertificateSignedHandler(callback ocpp21.CertificateSignedCallback) *Client {
	c.ocpp21Callbacks.CertificateSigned = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ChangeAvailabilityHandler(callback ocpp21.ChangeAvailabilityCallback) *Client {
	c.ocpp21Callbacks.ChangeAvailability = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ChangeTransactionTariffHandler(callback ocpp21.ChangeTransactionTariffCallback) *Client {
	c.ocpp21Callbacks.ChangeTransactionTariff = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ClearCacheHandler(callback ocpp21.ClearCacheCallback) *Client {
	c.ocpp21Callbacks.ClearCache = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ClearChargingProfileHandler(callback ocpp21.ClearChargingProfileCallback) *Client {
	c.ocpp21Callbacks.ClearChargingProfile = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ClearDERControlHandler(callback ocpp21.ClearDERControlCallback) *Client {
	c.ocpp21Callbacks.ClearDERControl = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ClearDisplayMessageHandler(callback ocpp21.ClearDisplayMessageCallback) *Client {
	c.ocpp21Callbacks.ClearDisplayMessage = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ClearTariffsHandler(callback ocpp21.ClearTariffsCallback) *Client {
	c.ocpp21Callbacks.ClearTariffs = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ClearVariableMonitoringHandler(callback ocpp21.ClearVariableMonitoringCallback) *Client {
	c.ocpp21Callbacks.ClearVariableMonitoring = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ClearedChargingLimitHandler(callback ocpp21.ClearedChargingLimitCallback) *Client {
	c.ocpp21Callbacks.ClearedChargingLimit = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ClosePeriodicEventStreamHandler(callback ocpp21.ClosePeriodicEventStreamCallback) *Client {
	c.ocpp21Callbacks.ClosePeriodicEventStream = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21CostUpdatedHandler(callback ocpp21.CostUpdatedCallback) *Client {
	c.ocpp21Callbacks.CostUpdated = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21CustomerInformationHandler(callback ocpp21.CustomerInformationCallback) *Client {
	c.ocpp21Callbacks.CustomerInformation = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21DataTransferHandler(callback ocpp21.DataTransferCallback) *Client {
	c.ocpp21Callbacks.DataTransfer = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21DeleteCertificateHandler(callback ocpp21.DeleteCertificateCallback) *Client {
	c.ocpp21Callbacks.DeleteCertificate = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21FirmwareStatusNotificationHandler(callback ocpp21.FirmwareStatusNotificationCallback) *Client {
	c.ocpp21Callbacks.FirmwareStatusNotification = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21Get15118EVCertificateHandler(callback ocpp21.Get15118EVCertificateCallback) *Client {
	c.ocpp21Callbacks.Get15118EVCertificate = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetBaseReportHandler(callback ocpp21.GetBaseReportCallback) *Client {
	c.ocpp21Callbacks.GetBaseReport = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetCertificateChainStatusHandler(callback ocpp21.GetCertificateChainStatusCallback) *Client {
	c.ocpp21Callbacks.GetCertificateChainStatus = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetCertificateStatusHandler(callback ocpp21.GetCertificateStatusCallback) *Client {
	c.ocpp21Callbacks.GetCertificateStatus = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetChargingProfilesHandler(callback ocpp21.GetChargingProfilesCallback) *Client {
	c.ocpp21Callbacks.GetChargingProfiles = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetCompositeScheduleHandler(callback ocpp21.GetCompositeScheduleCallback) *Client {
	c.ocpp21Callbacks.GetCompositeSchedule = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetDERControlHandler(callback ocpp21.GetDERControlCallback) *Client {
	c.ocpp21Callbacks.GetDERControl = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetDisplayMessagesHandler(callback ocpp21.GetDisplayMessagesCallback) *Client {
	c.ocpp21Callbacks.GetDisplayMessages = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetInstalledCertificateIdsHandler(callback ocpp21.GetInstalledCertificateIdsCallback) *Client {
	c.ocpp21Callbacks.GetInstalledCertificateIds = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetLocalListVersionHandler(callback ocpp21.GetLocalListVersionCallback) *Client {
	c.ocpp21Callbacks.GetLocalListVersion = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetLogHandler(callback ocpp21.GetLogCallback) *Client {
	c.ocpp21Callbacks.GetLog = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetMonitoringReportHandler(callback ocpp21.GetMonitoringReportCallback) *Client {
	c.ocpp21Callbacks.GetMonitoringReport = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetPeriodicEventStreamHandler(callback ocpp21.GetPeriodicEventStreamCallback) *Client {
	c.ocpp21Callbacks.GetPeriodicEventStream = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetReportHandler(callback ocpp21.GetReportCallback) *Client {
	c.ocpp21Callbacks.GetReport = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetTariffsHandler(callback ocpp21.GetTariffsCallback) *Client {
	c.ocpp21Callbacks.GetTariffs = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetTransactionStatusHandler(callback ocpp21.GetTransactionStatusCallback) *Client {
	c.ocpp21Callbacks.GetTransactionStatus = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21GetVariablesHandler(callback ocpp21.GetVariablesCallback) *Client {
	c.ocpp21Callbacks.GetVariables = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21HeartbeatHandler(callback ocpp21.HeartbeatCallback) *Client {
	c.ocpp21Callbacks.Heartbeat = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21InstallCertificateHandler(callback ocpp21.InstallCertificateCallback) *Client {
	c.ocpp21Callbacks.InstallCertificate = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21LogStatusNotificationHandler(callback ocpp21.LogStatusNotificationCallback) *Client {
	c.ocpp21Callbacks.LogStatusNotification = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21MeterValuesHandler(callback ocpp21.MeterValuesCallback) *Client {
	c.ocpp21Callbacks.MeterValues = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyAllowedEnergyTransferHandler(callback ocpp21.NotifyAllowedEnergyTransferCallback) *Client {
	c.ocpp21Callbacks.NotifyAllowedEnergyTransfer = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyChargingLimitHandler(callback ocpp21.NotifyChargingLimitCallback) *Client {
	c.ocpp21Callbacks.NotifyChargingLimit = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyCustomerInformationHandler(callback ocpp21.NotifyCustomerInformationCallback) *Client {
	c.ocpp21Callbacks.NotifyCustomerInformation = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyDERAlarmHandler(callback ocpp21.NotifyDERAlarmCallback) *Client {
	c.ocpp21Callbacks.NotifyDERAlarm = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyDERStartStopHandler(callback ocpp21.NotifyDERStartStopCallback) *Client {
	c.ocpp21Callbacks.NotifyDERStartStop = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyDisplayMessagesHandler(callback ocpp21.NotifyDisplayMessagesCallback) *Client {
	c.ocpp21Callbacks.NotifyDisplayMessages = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyEVChargingNeedsHandler(callback ocpp21.NotifyEVChargingNeedsCallback) *Client {
	c.ocpp21Callbacks.NotifyEVChargingNeeds = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyEVChargingScheduleHandler(callback ocpp21.NotifyEVChargingScheduleCallback) *Client {
	c.ocpp21Callbacks.NotifyEVChargingSchedule = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyEventHandler(callback ocpp21.NotifyEventCallback) *Client {
	c.ocpp21Callbacks.NotifyEvent = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyMonitoringReportHandler(callback ocpp21.NotifyMonitoringReportCallback) *Client {
	c.ocpp21Callbacks.NotifyMonitoringReport = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyPeriodicEventStreamHandler(callback ocpp21.NotifyPeriodicEventStreamCallback) *Client {
	c.ocpp21Callbacks.NotifyPeriodicEventStream = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyPriorityChargingHandler(callback ocpp21.NotifyPriorityChargingCallback) *Client {
	c.ocpp21Callbacks.NotifyPriorityCharging = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyReportHandler(callback ocpp21.NotifyReportCallback) *Client {
	c.ocpp21Callbacks.NotifyReport = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifySettlementHandler(callback ocpp21.NotifySettlementCallback) *Client {
	c.ocpp21Callbacks.NotifySettlement = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21NotifyWebPaymentStartedHandler(callback ocpp21.NotifyWebPaymentStartedCallback) *Client {
	c.ocpp21Callbacks.NotifyWebPaymentStarted = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21OpenPeriodicEventStreamHandler(callback ocpp21.OpenPeriodicEventStreamCallback) *Client {
	c.ocpp21Callbacks.OpenPeriodicEventStream = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21PublishFirmwareHandler(callback ocpp21.PublishFirmwareCallback) *Client {
	c.ocpp21Callbacks.PublishFirmware = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21PublishFirmwareStatusNotificationHandler(callback ocpp21.PublishFirmwareStatusNotificationCallback) *Client {
	c.ocpp21Callbacks.PublishFirmwareStatusNotification = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21PullDynamicScheduleUpdateHandler(callback ocpp21.PullDynamicScheduleUpdateCallback) *Client {
	c.ocpp21Callbacks.PullDynamicScheduleUpdate = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ReportChargingProfilesHandler(callback ocpp21.ReportChargingProfilesCallback) *Client {
	c.ocpp21Callbacks.ReportChargingProfiles = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ReportDERControlHandler(callback ocpp21.ReportDERControlCallback) *Client {
	c.ocpp21Callbacks.ReportDERControl = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21RequestBatterySwapHandler(callback ocpp21.RequestBatterySwapCallback) *Client {
	c.ocpp21Callbacks.RequestBatterySwap = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21RequestStartTransactionHandler(callback ocpp21.RequestStartTransactionCallback) *Client {
	c.ocpp21Callbacks.RequestStartTransaction = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21RequestStopTransactionHandler(callback ocpp21.RequestStopTransactionCallback) *Client {
	c.ocpp21Callbacks.RequestStopTransaction = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ReservationStatusUpdateHandler(callback ocpp21.ReservationStatusUpdateCallback) *Client {
	c.ocpp21Callbacks.ReservationStatusUpdate = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ReserveNowHandler(callback ocpp21.ReserveNowCallback) *Client {
	c.ocpp21Callbacks.ReserveNow = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21ResetHandler(callback ocpp21.ResetCallback) *Client {
	c.ocpp21Callbacks.Reset = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SecurityEventNotificationHandler(callback ocpp21.SecurityEventNotificationCallback) *Client {
	c.ocpp21Callbacks.SecurityEventNotification = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SendLocalListHandler(callback ocpp21.SendLocalListCallback) *Client {
	c.ocpp21Callbacks.SendLocalList = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SetChargingProfileHandler(callback ocpp21.SetChargingProfileCallback) *Client {
	c.ocpp21Callbacks.SetChargingProfile = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SetDERControlHandler(callback ocpp21.SetDERControlCallback) *Client {
	c.ocpp21Callbacks.SetDERControl = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SetDefaultTariffHandler(callback ocpp21.SetDefaultTariffCallback) *Client {
	c.ocpp21Callbacks.SetDefaultTariff = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SetDisplayMessageHandler(callback ocpp21.SetDisplayMessageCallback) *Client {
	c.ocpp21Callbacks.SetDisplayMessage = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SetMonitoringBaseHandler(callback ocpp21.SetMonitoringBaseCallback) *Client {
	c.ocpp21Callbacks.SetMonitoringBase = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SetMonitoringLevelHandler(callback ocpp21.SetMonitoringLevelCallback) *Client {
	c.ocpp21Callbacks.SetMonitoringLevel = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SetNetworkProfileHandler(callback ocpp21.SetNetworkProfileCallback) *Client {
	c.ocpp21Callbacks.SetNetworkProfile = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SetVariableMonitoringHandler(callback ocpp21.SetVariableMonitoringCallback) *Client {
	c.ocpp21Callbacks.SetVariableMonitoring = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SetVariablesHandler(callback ocpp21.SetVariablesCallback) *Client {
	c.ocpp21Callbacks.SetVariables = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21SignCertificateHandler(callback ocpp21.SignCertificateCallback) *Client {
	c.ocpp21Callbacks.SignCertificate = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21StatusNotificationHandler(callback ocpp21.StatusNotificationCallback) *Client {
	c.ocpp21Callbacks.StatusNotification = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21TransactionEventHandler(callback ocpp21.TransactionEventCallback) *Client {
	c.ocpp21Callbacks.TransactionEvent = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21TriggerMessageHandler(callback ocpp21.TriggerMessageCallback) *Client {
	c.ocpp21Callbacks.TriggerMessage = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21UnlockConnectorHandler(callback ocpp21.UnlockConnectorCallback) *Client {
	c.ocpp21Callbacks.UnlockConnector = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21UnpublishFirmwareHandler(callback ocpp21.UnpublishFirmwareCallback) *Client {
	c.ocpp21Callbacks.UnpublishFirmware = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21UpdateDynamicScheduleHandler(callback ocpp21.UpdateDynamicScheduleCallback) *Client {
	c.ocpp21Callbacks.UpdateDynamicSchedule = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21UpdateFirmwareHandler(callback ocpp21.UpdateFirmwareCallback) *Client {
	c.ocpp21Callbacks.UpdateFirmware = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21UsePriorityChargingHandler(callback ocpp21.UsePriorityChargingCallback) *Client {
	c.ocpp21Callbacks.UsePriorityCharging = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}

func (c *Client) WithOCPP21VatNumberValidationHandler(callback ocpp21.VatNumberValidationCallback) *Client {
	c.ocpp21Callbacks.VatNumberValidation = callback
	c.ocpp21Callbacks.InitHandlers()
	c.subprotocol = "ocpp2.1"
	return c
}
