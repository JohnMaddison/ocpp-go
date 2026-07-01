package server

import "github.com/johnmaddison/ocpp-go/ocpp21"

func With21AFRRSignalHandler(callback ocpp21.AFRRSignalCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.AFRRSignal = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21AdjustPeriodicEventStreamHandler(callback ocpp21.AdjustPeriodicEventStreamCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.AdjustPeriodicEventStream = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21AuthorizeHandler(callback ocpp21.AuthorizeCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.Authorize = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21BatterySwapHandler(callback ocpp21.BatterySwapCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.BatterySwap = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21BootNotificationHandler(callback ocpp21.BootNotificationCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.BootNotification = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21CancelReservationHandler(callback ocpp21.CancelReservationCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.CancelReservation = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21CertificateSignedHandler(callback ocpp21.CertificateSignedCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.CertificateSigned = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ChangeAvailabilityHandler(callback ocpp21.ChangeAvailabilityCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ChangeAvailability = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ChangeTransactionTariffHandler(callback ocpp21.ChangeTransactionTariffCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ChangeTransactionTariff = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ClearCacheHandler(callback ocpp21.ClearCacheCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ClearCache = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ClearChargingProfileHandler(callback ocpp21.ClearChargingProfileCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ClearChargingProfile = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ClearDERControlHandler(callback ocpp21.ClearDERControlCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ClearDERControl = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ClearDisplayMessageHandler(callback ocpp21.ClearDisplayMessageCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ClearDisplayMessage = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ClearTariffsHandler(callback ocpp21.ClearTariffsCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ClearTariffs = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ClearVariableMonitoringHandler(callback ocpp21.ClearVariableMonitoringCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ClearVariableMonitoring = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ClearedChargingLimitHandler(callback ocpp21.ClearedChargingLimitCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ClearedChargingLimit = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ClosePeriodicEventStreamHandler(callback ocpp21.ClosePeriodicEventStreamCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ClosePeriodicEventStream = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21CostUpdatedHandler(callback ocpp21.CostUpdatedCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.CostUpdated = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21CustomerInformationHandler(callback ocpp21.CustomerInformationCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.CustomerInformation = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21DataTransferHandler(callback ocpp21.DataTransferCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.DataTransfer = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21DeleteCertificateHandler(callback ocpp21.DeleteCertificateCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.DeleteCertificate = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21FirmwareStatusNotificationHandler(callback ocpp21.FirmwareStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.FirmwareStatusNotification = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21Get15118EVCertificateHandler(callback ocpp21.Get15118EVCertificateCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.Get15118EVCertificate = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetBaseReportHandler(callback ocpp21.GetBaseReportCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetBaseReport = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetCertificateChainStatusHandler(callback ocpp21.GetCertificateChainStatusCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetCertificateChainStatus = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetCertificateStatusHandler(callback ocpp21.GetCertificateStatusCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetCertificateStatus = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetChargingProfilesHandler(callback ocpp21.GetChargingProfilesCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetChargingProfiles = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetCompositeScheduleHandler(callback ocpp21.GetCompositeScheduleCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetCompositeSchedule = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetDERControlHandler(callback ocpp21.GetDERControlCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetDERControl = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetDisplayMessagesHandler(callback ocpp21.GetDisplayMessagesCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetDisplayMessages = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetInstalledCertificateIdsHandler(callback ocpp21.GetInstalledCertificateIdsCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetInstalledCertificateIds = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetLocalListVersionHandler(callback ocpp21.GetLocalListVersionCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetLocalListVersion = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetLogHandler(callback ocpp21.GetLogCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetLog = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetMonitoringReportHandler(callback ocpp21.GetMonitoringReportCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetMonitoringReport = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetPeriodicEventStreamHandler(callback ocpp21.GetPeriodicEventStreamCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetPeriodicEventStream = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetReportHandler(callback ocpp21.GetReportCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetReport = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetTariffsHandler(callback ocpp21.GetTariffsCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetTariffs = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetTransactionStatusHandler(callback ocpp21.GetTransactionStatusCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetTransactionStatus = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21GetVariablesHandler(callback ocpp21.GetVariablesCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.GetVariables = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21HeartbeatHandler(callback ocpp21.HeartbeatCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.Heartbeat = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21InstallCertificateHandler(callback ocpp21.InstallCertificateCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.InstallCertificate = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21LogStatusNotificationHandler(callback ocpp21.LogStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.LogStatusNotification = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21MeterValuesHandler(callback ocpp21.MeterValuesCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.MeterValues = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyAllowedEnergyTransferHandler(callback ocpp21.NotifyAllowedEnergyTransferCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyAllowedEnergyTransfer = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyChargingLimitHandler(callback ocpp21.NotifyChargingLimitCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyChargingLimit = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyCustomerInformationHandler(callback ocpp21.NotifyCustomerInformationCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyCustomerInformation = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyDERAlarmHandler(callback ocpp21.NotifyDERAlarmCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyDERAlarm = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyDERStartStopHandler(callback ocpp21.NotifyDERStartStopCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyDERStartStop = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyDisplayMessagesHandler(callback ocpp21.NotifyDisplayMessagesCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyDisplayMessages = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyEVChargingNeedsHandler(callback ocpp21.NotifyEVChargingNeedsCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyEVChargingNeeds = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyEVChargingScheduleHandler(callback ocpp21.NotifyEVChargingScheduleCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyEVChargingSchedule = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyEventHandler(callback ocpp21.NotifyEventCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyEvent = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyMonitoringReportHandler(callback ocpp21.NotifyMonitoringReportCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyMonitoringReport = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyPeriodicEventStreamHandler(callback ocpp21.NotifyPeriodicEventStreamCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyPeriodicEventStream = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyPriorityChargingHandler(callback ocpp21.NotifyPriorityChargingCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyPriorityCharging = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyReportHandler(callback ocpp21.NotifyReportCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyReport = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifySettlementHandler(callback ocpp21.NotifySettlementCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifySettlement = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21NotifyWebPaymentStartedHandler(callback ocpp21.NotifyWebPaymentStartedCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.NotifyWebPaymentStarted = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21OpenPeriodicEventStreamHandler(callback ocpp21.OpenPeriodicEventStreamCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.OpenPeriodicEventStream = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21PublishFirmwareHandler(callback ocpp21.PublishFirmwareCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.PublishFirmware = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21PublishFirmwareStatusNotificationHandler(callback ocpp21.PublishFirmwareStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.PublishFirmwareStatusNotification = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21PullDynamicScheduleUpdateHandler(callback ocpp21.PullDynamicScheduleUpdateCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.PullDynamicScheduleUpdate = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ReportChargingProfilesHandler(callback ocpp21.ReportChargingProfilesCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ReportChargingProfiles = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ReportDERControlHandler(callback ocpp21.ReportDERControlCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ReportDERControl = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21RequestBatterySwapHandler(callback ocpp21.RequestBatterySwapCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.RequestBatterySwap = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21RequestStartTransactionHandler(callback ocpp21.RequestStartTransactionCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.RequestStartTransaction = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21RequestStopTransactionHandler(callback ocpp21.RequestStopTransactionCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.RequestStopTransaction = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ReservationStatusUpdateHandler(callback ocpp21.ReservationStatusUpdateCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ReservationStatusUpdate = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ReserveNowHandler(callback ocpp21.ReserveNowCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.ReserveNow = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21ResetHandler(callback ocpp21.ResetCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.Reset = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SecurityEventNotificationHandler(callback ocpp21.SecurityEventNotificationCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SecurityEventNotification = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SendLocalListHandler(callback ocpp21.SendLocalListCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SendLocalList = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SetChargingProfileHandler(callback ocpp21.SetChargingProfileCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SetChargingProfile = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SetDERControlHandler(callback ocpp21.SetDERControlCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SetDERControl = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SetDefaultTariffHandler(callback ocpp21.SetDefaultTariffCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SetDefaultTariff = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SetDisplayMessageHandler(callback ocpp21.SetDisplayMessageCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SetDisplayMessage = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SetMonitoringBaseHandler(callback ocpp21.SetMonitoringBaseCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SetMonitoringBase = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SetMonitoringLevelHandler(callback ocpp21.SetMonitoringLevelCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SetMonitoringLevel = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SetNetworkProfileHandler(callback ocpp21.SetNetworkProfileCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SetNetworkProfile = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SetVariableMonitoringHandler(callback ocpp21.SetVariableMonitoringCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SetVariableMonitoring = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SetVariablesHandler(callback ocpp21.SetVariablesCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SetVariables = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21SignCertificateHandler(callback ocpp21.SignCertificateCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.SignCertificate = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21StatusNotificationHandler(callback ocpp21.StatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.StatusNotification = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21TransactionEventHandler(callback ocpp21.TransactionEventCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.TransactionEvent = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21TriggerMessageHandler(callback ocpp21.TriggerMessageCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.TriggerMessage = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21UnlockConnectorHandler(callback ocpp21.UnlockConnectorCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.UnlockConnector = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21UnpublishFirmwareHandler(callback ocpp21.UnpublishFirmwareCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.UnpublishFirmware = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21UpdateDynamicScheduleHandler(callback ocpp21.UpdateDynamicScheduleCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.UpdateDynamicSchedule = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21UpdateFirmwareHandler(callback ocpp21.UpdateFirmwareCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.UpdateFirmware = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21UsePriorityChargingHandler(callback ocpp21.UsePriorityChargingCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.UsePriorityCharging = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}

func With21VatNumberValidationHandler(callback ocpp21.VatNumberValidationCallback) Option {
	return func(s *Server) {
		s.ocpp21Callbacks.VatNumberValidation = callback
		s.ocpp21Callbacks.InitHandlers()
		s.ocpp21Enabled = true
	}
}
