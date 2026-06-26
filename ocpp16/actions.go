package ocpp16

type Action string

const (
	// Core Profile
	ActionAuthorize              Action = "Authorize"
	ActionBootNotification       Action = "BootNotification"
	ActionChangeAvailability     Action = "ChangeAvailability"
	ActionChangeConfiguration    Action = "ChangeConfiguration"
	ActionClearCache             Action = "ClearCache"
	ActionDataTransfer           Action = "DataTransfer"
	ActionGetConfiguration       Action = "GetConfiguration"
	ActionHeartbeat              Action = "Heartbeat"
	ActionMeterValues            Action = "MeterValues"
	ActionRemoteStartTransaction Action = "RemoteStartTransaction"
	ActionRemoteStopTransaction  Action = "RemoteStopTransaction"
	ActionReset                  Action = "Reset"
	ActionStartTransaction       Action = "StartTransaction"
	ActionStopTransaction        Action = "StopTransaction"
	ActionStatusNotification     Action = "StatusNotification"
	ActionUnlockConnector        Action = "UnlockConnector"

	// Local Auth List Management Profile
	ActionGetLocalListVersion Action = "GetLocalListVersion"
	ActionSendLocalList       Action = "SendLocalList"

	// Reservation Profile
	ActionCancelReservation Action = "CancelReservation"
	ActionReserveNow        Action = "ReserveNow"

	// Smart Charging Profile
	ActionSetChargingProfile   Action = "SetChargingProfile"
	ActionClearChargingProfile Action = "ClearChargingProfile"
	ActionGetCompositeSchedule Action = "GetCompositeSchedule"

	// Firmware Management Profile
	ActionGetDiagnostics                Action = "GetDiagnostics"
	ActionDiagnosticsStatusNotification Action = "DiagnosticsStatusNotification"
	ActionFirmwareStatusNotification    Action = "FirmwareStatusNotification"
	ActionUpdateFirmware                Action = "UpdateFirmware"

	// Remote Trigger Profile
	ActionTriggerMessage Action = "TriggerMessage"

	// Security Extension
	ActionCertificateSigned                Action = "CertificateSigned"
	ActionDeleteCertificate                Action = "DeleteCertificate"
	ActionGetInstalledCertificateIds       Action = "GetInstalledCertificateIds"
	ActionInstallCertificate               Action = "InstallCertificate"
	ActionLogStatusNotification            Action = "LogStatusNotification"
	ActionSecurityEventNotification        Action = "SecurityEventNotification"
	ActionSignCertificate                  Action = "SignCertificate"
	ActionSignedUpdateFirmware             Action = "SignedUpdateFirmware"
	ActionSignedFirmwareStatusNotification Action = "SignedFirmwareStatusNotification"
	ActionGetLog                           Action = "GetLog"
	ActionExtendedTriggerMessage           Action = "ExtendedTriggerMessage"
)
