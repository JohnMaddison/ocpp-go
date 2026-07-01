package server

import "github.com/johnmaddison/ocpp-go/ocpp16"

func WithOCPP16AuthorizeHandler(callback ocpp16.AuthorizeCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.Authorize = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16BootNotificationHandler(callback ocpp16.BootNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.BootNotification = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16CancelReservationHandler(callback ocpp16.CancelReservationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.CancelReservation = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16CertificateSignedHandler(callback ocpp16.CertificateSignedCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.CertificateSigned = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16ChangeAvailabilityHandler(callback ocpp16.ChangeAvailabilityCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ChangeAvailability = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16ChangeConfigurationHandler(callback ocpp16.ChangeConfigurationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ChangeConfiguration = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16ClearCacheHandler(callback ocpp16.ClearCacheCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ClearCache = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16ClearChargingProfileHandler(callback ocpp16.ClearChargingProfileCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ClearChargingProfile = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16ClearReservationHandler(callback ocpp16.ClearReservationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ClearReservation = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16DataTransferHandler(callback ocpp16.DataTransferCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.DataTransfer = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16DeleteCertificateHandler(callback ocpp16.DeleteCertificateCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.DeleteCertificate = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16DiagnosticsStatusNotificationHandler(callback ocpp16.DiagnosticsStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.DiagnosticsStatusNotification = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16ExtendedTriggerMessageHandler(callback ocpp16.ExtendedTriggerMessageCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ExtendedTriggerMessage = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16GetCompositeScheduleHandler(callback ocpp16.GetCompositeScheduleCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetCompositeSchedule = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16GetConfigurationHandler(callback ocpp16.GetConfigurationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetConfiguration = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16GetDiagnosticsHandler(callback ocpp16.GetDiagnosticsCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetDiagnostics = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16GetInstalledCertificatesHandler(callback ocpp16.GetInstalledCertificatesCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetInstalledCertificates = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16GetLogHandler(callback ocpp16.GetLogCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetLog = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16HeartbeatHandler(callback ocpp16.HeartbeatCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.Heartbeat = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16InstallCertificateHandler(callback ocpp16.InstallCertificateCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.InstallCertificate = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16LogStatusNotificationHandler(callback ocpp16.LogStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.LogStatusNotification = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16MeterValuesHandler(callback ocpp16.MeterValuesCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.MeterValues = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16RemoteStartTransactionHandler(callback ocpp16.RemoteStartTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.RemoteStartTransaction = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16RemoteStopTransactionHandler(callback ocpp16.RemoteStopTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.RemoteStopTransaction = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16ReserveNowHandler(callback ocpp16.ReserveNowCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ReserveNow = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16ResetHandler(callback ocpp16.ResetCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.Reset = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16SendLocalListHandler(callback ocpp16.SendLocalListCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SendLocalList = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16SetChargingProfileHandler(callback ocpp16.SetChargingProfileCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SetChargingProfile = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16SignCertificateHandler(callback ocpp16.SignCertificateCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SignCertificate = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16SignedFirmwareStatusNotificationHandler(callback ocpp16.SignedFirmwareStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SignedFirmwareStatusNotification = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16SignedUpdateFirmwareHandler(callback ocpp16.SignedUpdateFirmwareCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SignedUpdateFirmware = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16StartTransactionHandler(callback ocpp16.StartTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.StartTransaction = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16StatusNotificationHandler(callback ocpp16.StatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.StatusNotification = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16StopTransactionHandler(callback ocpp16.StopTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.StopTransaction = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16TriggerMessageHandler(callback ocpp16.TriggerMessageCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.TriggerMessage = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16UnlockConnectorHandler(callback ocpp16.UnlockConnectorCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.UnlockConnector = callback
		ensureOCPP16Parser(s)
	}
}

func WithOCPP16UpdateFirmwareHandler(callback ocpp16.UpdateFirmwareCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.UpdateFirmware = callback
		ensureOCPP16Parser(s)
	}
}
