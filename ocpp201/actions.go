// Package ocpp201 contains the data structures for OCPP 2.0.1 messages.
package ocpp201

// Action defines the type for all OCPP 2.0.1 message actions.
type Action string

// Defines the complete list of supported actions for OCPP 2.0.1.
const (
	// Security
	ActionSecurityEventNotification Action = "SecurityEventNotification"
	ActionSignCertificate           Action = "SignCertificate"
	ActionCertificateSigned         Action = "CertificateSigned"

	// Provisioning
	ActionBootNotification        Action = "BootNotification"
	ActionSetVariables            Action = "SetVariables"
	ActionGetVariables            Action = "GetVariables"
	ActionGetBaseReport           Action = "GetBaseReport"
	ActionSetNetworkProfile       Action = "SetNetworkProfile"
	ActionReset                   Action = "Reset"
	ActionTriggerMessage          Action = "TriggerMessage"
	ActionNotifyReport            Action = "NotifyReport"
	ActionSetMonitoringBase       Action = "SetMonitoringBase"
	ActionSetMonitoringLevel      Action = "SetMonitoringLevel"
	ActionSetVariableMonitoring   Action = "SetVariableMonitoring"
	ActionClearVariableMonitoring Action = "ClearVariableMonitoring"
	ActionGetMonitoringReport     Action = "GetMonitoringReport"
	ActionNotifyMonitoringReport  Action = "NotifyMonitoringReport"
	ActionNotifyEvent             Action = "NotifyEvent"
	ActionDataTransfer            Action = "DataTransfer"

	// Authorization
	ActionAuthorize                 Action = "Authorize"
	ActionNotifyCustomerInformation Action = "NotifyCustomerInformation"
	ActionGetLocalListVersion       Action = "GetLocalListVersion"
	ActionSendLocalList             Action = "SendLocalList"
	ActionClearCache                Action = "ClearCache"

	// Transactions
	ActionTransactionEvent        Action = "TransactionEvent"
	ActionGetTransactionStatus    Action = "GetTransactionStatus"
	ActionRequestStartTransaction Action = "RequestStartTransaction"
	ActionRequestStopTransaction  Action = "RequestStopTransaction"
	ActionMeterValues             Action = "MeterValues"

	// Availability
	ActionStatusNotification Action = "StatusNotification"
	ActionHeartbeat          Action = "Heartbeat"
	ActionChangeAvailability Action = "ChangeAvailability"

	// Reservation
	ActionReserveNow              Action = "ReserveNow"
	ActionCancelReservation       Action = "CancelReservation"
	ActionReservationStatusUpdate Action = "ReservationStatusUpdate"

	// Smart Charging
	ActionSetChargingProfile       Action = "SetChargingProfile"
	ActionGetChargingProfiles      Action = "GetChargingProfiles"
	ActionClearChargingProfile     Action = "ClearChargingProfile"
	ActionGetCompositeSchedule     Action = "GetCompositeSchedule"
	ActionNotifyChargingLimit      Action = "NotifyChargingLimit"
	ActionClearedChargingLimit     Action = "ClearedChargingLimit"
	ActionNotifyEVChargingNeeds    Action = "NotifyEVChargingNeeds"
	ActionNotifyEVChargingSchedule Action = "NotifyEVChargingSchedule"
	ActionReportChargingProfiles   Action = "ReportChargingProfiles"
	ActionNotifyDisplayMessages    Action = "NotifyDisplayMessages"
	ActionGetDisplayMessages       Action = "GetDisplayMessages"
	ActionClearDisplayMessage      Action = "ClearDisplayMessage"
	ActionCostUpdated              Action = "CostUpdated"
	ActionCustomerInformation      Action = "CustomerInformation"

	// Remote Control
	ActionUnlockConnector Action = "UnlockConnector"

	// Firmware Management
	ActionUpdateFirmware                    Action = "UpdateFirmware"
	ActionPublishFirmware                   Action = "PublishFirmware"
	ActionUnpublishFirmware                 Action = "UnpublishFirmware"
	ActionGetFirmware                       Action = "GetFirmware"
	ActionNotifyFirmwareStatusNotification  Action = "NotifyFirmwareStatusNotification"
	ActionFirmwareStatusNotification        Action = "FirmwareStatusNotification"
	ActionPublishFirmwareStatusNotification Action = "PublishFirmwareStatusNotification"

	// Certificate Management
	ActionInstallCertificate         Action = "InstallCertificate"
	ActionDeleteCertificate          Action = "DeleteCertificate"
	ActionGetInstalledCertificateIds Action = "GetInstalledCertificateIds"
	ActionGetCertificateStatus       Action = "GetCertificateStatus"

	// Diagnostics
	ActionGetLog                      Action = "GetLog"
	ActionLogStatusNotification       Action = "LogStatusNotification"
	ActionGetDiagnostics              Action = "GetDiagnostics"
	ActionNotifyDiagnosticsStatus     Action = "NotifyDiagnosticsStatus"
	ActionUploadLogStatusNotification Action = "UploadLogStatusNotification"

	// ISO 15118 Certificate Management
	ActionGet15118EVCertificate Action = "Get15118EVCertificate"

	// Clock
	ActionSetClock Action = "SetClock"
)
