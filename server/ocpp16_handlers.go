package server

import "github.com/johnmaddison/ocpp-go/ocpp16"

func With16AuthorizeHandler(callback ocpp16.AuthorizeCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.Authorize = callback
		ensure16Parser(s)
	}
}

func With16BootNotificationHandler(callback ocpp16.BootNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.BootNotification = callback
		ensure16Parser(s)
	}
}

func With16CancelReservationHandler(callback ocpp16.CancelReservationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.CancelReservation = callback
		ensure16Parser(s)
	}
}

func With16CertificateSignedHandler(callback ocpp16.CertificateSignedCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.CertificateSigned = callback
		ensure16Parser(s)
	}
}

func With16ChangeAvailabilityHandler(callback ocpp16.ChangeAvailabilityCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ChangeAvailability = callback
		ensure16Parser(s)
	}
}

func With16ChangeConfigurationHandler(callback ocpp16.ChangeConfigurationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ChangeConfiguration = callback
		ensure16Parser(s)
	}
}

func With16ClearCacheHandler(callback ocpp16.ClearCacheCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ClearCache = callback
		ensure16Parser(s)
	}
}

func With16ClearChargingProfileHandler(callback ocpp16.ClearChargingProfileCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ClearChargingProfile = callback
		ensure16Parser(s)
	}
}

func With16ClearReservationHandler(callback ocpp16.ClearReservationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ClearReservation = callback
		ensure16Parser(s)
	}
}

func With16DataTransferHandler(callback ocpp16.DataTransferCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.DataTransfer = callback
		ensure16Parser(s)
	}
}

func With16DeleteCertificateHandler(callback ocpp16.DeleteCertificateCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.DeleteCertificate = callback
		ensure16Parser(s)
	}
}

func With16DiagnosticsStatusNotificationHandler(callback ocpp16.DiagnosticsStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.DiagnosticsStatusNotification = callback
		ensure16Parser(s)
	}
}

func With16ExtendedTriggerMessageHandler(callback ocpp16.ExtendedTriggerMessageCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ExtendedTriggerMessage = callback
		ensure16Parser(s)
	}
}

func With16GetCompositeScheduleHandler(callback ocpp16.GetCompositeScheduleCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetCompositeSchedule = callback
		ensure16Parser(s)
	}
}

func With16GetConfigurationHandler(callback ocpp16.GetConfigurationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetConfiguration = callback
		ensure16Parser(s)
	}
}

func With16GetDiagnosticsHandler(callback ocpp16.GetDiagnosticsCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetDiagnostics = callback
		ensure16Parser(s)
	}
}

func With16GetInstalledCertificatesHandler(callback ocpp16.GetInstalledCertificatesCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetInstalledCertificates = callback
		ensure16Parser(s)
	}
}

func With16GetLogHandler(callback ocpp16.GetLogCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.GetLog = callback
		ensure16Parser(s)
	}
}

func With16HeartbeatHandler(callback ocpp16.HeartbeatCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.Heartbeat = callback
		ensure16Parser(s)
	}
}

func With16InstallCertificateHandler(callback ocpp16.InstallCertificateCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.InstallCertificate = callback
		ensure16Parser(s)
	}
}

func With16LogStatusNotificationHandler(callback ocpp16.LogStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.LogStatusNotification = callback
		ensure16Parser(s)
	}
}

func With16MeterValuesHandler(callback ocpp16.MeterValuesCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.MeterValues = callback
		ensure16Parser(s)
	}
}

func With16RemoteStartTransactionHandler(callback ocpp16.RemoteStartTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.RemoteStartTransaction = callback
		ensure16Parser(s)
	}
}

func With16RemoteStopTransactionHandler(callback ocpp16.RemoteStopTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.RemoteStopTransaction = callback
		ensure16Parser(s)
	}
}

func With16ReserveNowHandler(callback ocpp16.ReserveNowCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.ReserveNow = callback
		ensure16Parser(s)
	}
}

func With16ResetHandler(callback ocpp16.ResetCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.Reset = callback
		ensure16Parser(s)
	}
}

func With16SendLocalListHandler(callback ocpp16.SendLocalListCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SendLocalList = callback
		ensure16Parser(s)
	}
}

func With16SetChargingProfileHandler(callback ocpp16.SetChargingProfileCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SetChargingProfile = callback
		ensure16Parser(s)
	}
}

func With16SignCertificateHandler(callback ocpp16.SignCertificateCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SignCertificate = callback
		ensure16Parser(s)
	}
}

func With16SignedFirmwareStatusNotificationHandler(callback ocpp16.SignedFirmwareStatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SignedFirmwareStatusNotification = callback
		ensure16Parser(s)
	}
}

func With16SignedUpdateFirmwareHandler(callback ocpp16.SignedUpdateFirmwareCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.SignedUpdateFirmware = callback
		ensure16Parser(s)
	}
}

func With16StartTransactionHandler(callback ocpp16.StartTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.StartTransaction = callback
		ensure16Parser(s)
	}
}

func With16StatusNotificationHandler(callback ocpp16.StatusNotificationCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.StatusNotification = callback
		ensure16Parser(s)
	}
}

func With16StopTransactionHandler(callback ocpp16.StopTransactionCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.StopTransaction = callback
		ensure16Parser(s)
	}
}

func With16TriggerMessageHandler(callback ocpp16.TriggerMessageCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.TriggerMessage = callback
		ensure16Parser(s)
	}
}

func With16UnlockConnectorHandler(callback ocpp16.UnlockConnectorCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.UnlockConnector = callback
		ensure16Parser(s)
	}
}

func With16UpdateFirmwareHandler(callback ocpp16.UpdateFirmwareCallback) Option {
	return func(s *Server) {
		s.ocppCallbacks.UpdateFirmware = callback
		ensure16Parser(s)
	}
}
