// Package ocpp16 contains the data structures for OCPP 1.6 messages.
package ocpp16

// Defines all the enumerated types used in the OCPP 1.6 specification.

type AuthorizationStatus string

const (
	AuthorizationStatusAccepted     AuthorizationStatus = "Accepted"
	AuthorizationStatusBlocked      AuthorizationStatus = "Blocked"
	AuthorizationStatusExpired      AuthorizationStatus = "Expired"
	AuthorizationStatusInvalid      AuthorizationStatus = "Invalid"
	AuthorizationStatusConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

type AvailabilityStatus string

const (
	AvailabilityStatusAccepted  AvailabilityStatus = "Accepted"
	AvailabilityStatusRejected  AvailabilityStatus = "Rejected"
	AvailabilityStatusScheduled AvailabilityStatus = "Scheduled"
)

type AvailabilityType string

const (
	AvailabilityTypeInoperative AvailabilityType = "Inoperative"
	AvailabilityTypeOperative   AvailabilityType = "Operative"
)

type CancelReservationStatus string

const (
	CancelReservationStatusAccepted CancelReservationStatus = "Accepted"
	CancelReservationStatusRejected CancelReservationStatus = "Rejected"
)

type CertificateSignedStatus string

const (
	CertificateSignedStatusAccepted CertificateSignedStatus = "Accepted"
	CertificateSignedStatusRejected CertificateSignedStatus = "Rejected"
)

type CertificateUse string

const (
	CertificateUseCentralSystemRootCertificate CertificateUse = "CentralSystemRootCertificate"
	CertificateUseManufacturerRootCertificate  CertificateUse = "ManufacturerRootCertificate"
)

type ChangeConfigurationStatus string

const (
	ChangeConfigurationStatusAccepted       ChangeConfigurationStatus = "Accepted"
	ChangeConfigurationStatusRejected       ChangeConfigurationStatus = "Rejected"
	ChangeConfigurationStatusRebootRequired ChangeConfigurationStatus = "RebootRequired"
	ChangeConfigurationStatusNotSupported   ChangeConfigurationStatus = "NotSupported"
)

type ChargingProfileKindType string

const (
	ChargingProfileKindAbsolute  ChargingProfileKindType = "Absolute"
	ChargingProfileKindRecurring ChargingProfileKindType = "Recurring"
	ChargingProfileKindRelative  ChargingProfileKindType = "Relative"
)

type ChargingProfilePurposeType string

const (
	ChargingProfilePurposeChargePointMaxProfile ChargingProfilePurposeType = "ChargePointMaxProfile"
	ChargingProfilePurposeTxDefaultProfile      ChargingProfilePurposeType = "TxDefaultProfile"
	ChargingProfilePurposeTxProfile             ChargingProfilePurposeType = "TxProfile"
)

type ChargingProfileStatus string

const (
	ChargingProfileStatusAccepted     ChargingProfileStatus = "Accepted"
	ChargingProfileStatusRejected     ChargingProfileStatus = "Rejected"
	ChargingProfileStatusNotSupported ChargingProfileStatus = "NotSupported"
)

type ChargingRateUnitType string

const (
	ChargingRateUnitWatts   ChargingRateUnitType = "W"
	ChargingRateUnitAmperes ChargingRateUnitType = "A"
)

type ClearCacheStatus string

const (
	ClearCacheStatusAccepted ClearCacheStatus = "Accepted"
	ClearCacheStatusRejected ClearCacheStatus = "Rejected"
)

type ClearChargingProfileStatus string

const (
	ClearChargingProfileStatusAccepted ClearChargingProfileStatus = "Accepted"
	ClearChargingProfileStatusUnknown  ClearChargingProfileStatus = "Unknown"
)

type ConfigurationStatus string

const (
	ConfigurationStatusAccepted       ConfigurationStatus = "Accepted"
	ConfigurationStatusRejected       ConfigurationStatus = "Rejected"
	ConfigurationStatusRebootRequired ConfigurationStatus = "RebootRequired"
	ConfigurationStatusNotSupported   ConfigurationStatus = "NotSupported"
)

type DataTransferStatus string

const (
	DataTransferStatusAccepted         DataTransferStatus = "Accepted"
	DataTransferStatusRejected         DataTransferStatus = "Rejected"
	DataTransferStatusUnknownMessageID DataTransferStatus = "UnknownMessageId"
	DataTransferStatusUnknownVendorID  DataTransferStatus = "UnknownVendorId"
)

type DeleteCertificateStatus string

const (
	DeleteCertificateStatusAccepted DeleteCertificateStatus = "Accepted"
	DeleteCertificateStatusFailed   DeleteCertificateStatus = "Failed"
	DeleteCertificateStatusNotFound DeleteCertificateStatus = "NotFound"
)

type DiagnosticsStatus string

const (
	DiagnosticsStatusIdle         DiagnosticsStatus = "Idle"
	DiagnosticsStatusUploaded     DiagnosticsStatus = "Uploaded"
	DiagnosticsStatusUploadFailed DiagnosticsStatus = "UploadFailed"
	DiagnosticsStatusUploading    DiagnosticsStatus = "Uploading"
)

type FirmwareStatus string

const (
	FirmwareStatusDownloaded         FirmwareStatus = "Downloaded"
	FirmwareStatusDownloadFailed     FirmwareStatus = "DownloadFailed"
	FirmwareStatusDownloading        FirmwareStatus = "Downloading"
	FirmwareStatusIdle               FirmwareStatus = "Idle"
	FirmwareStatusInstallationFailed FirmwareStatus = "InstallationFailed"
	FirmwareStatusInstalling         FirmwareStatus = "Installing"
	FirmwareStatusInstalled          FirmwareStatus = "Installed"
)

type GenericStatus string

const (
	GenericStatusAccepted GenericStatus = "Accepted"
	GenericStatusRejected GenericStatus = "Rejected"
)

type GetCompositeScheduleStatus string

const (
	GetCompositeScheduleStatusAccepted GetCompositeScheduleStatus = "Accepted"
	GetCompositeScheduleStatusRejected GetCompositeScheduleStatus = "Rejected"
)

type GetCertificateStatus string

const (
	GetCertificateStatusAccepted GetCertificateStatus = "Accepted"
	GetCertificateStatusFailed   GetCertificateStatus = "Failed"
)

type GetInstalledCertificateStatus string

const (
	GetInstalledCertificateStatusAccepted GetInstalledCertificateStatus = "Accepted"
	GetInstalledCertificateStatusFailed   GetInstalledCertificateStatus = "Failed"
)

type HashAlgorithmEnumType string

const (
	HashAlgorithmSHA256 HashAlgorithmEnumType = "SHA256"
	HashAlgorithmSHA384 HashAlgorithmEnumType = "SHA384"
	HashAlgorithmSHA512 HashAlgorithmEnumType = "SHA512"
)

type InstallCertificateStatus string

const (
	InstallCertificateStatusAccepted InstallCertificateStatus = "Accepted"
	InstallCertificateStatusFailed   InstallCertificateStatus = "Failed"
)

type Location string

const (
	LocationCable  Location = "Cable"
	LocationEV     Location = "EV"
	LocationInlet  Location = "Inlet"
	LocationOutlet Location = "Outlet"
	LocationBody   Location = "Body"
)

type LogStatus string

const (
	LogStatusAccepted         LogStatus = "Accepted"
	LogStatusRejected         LogStatus = "Rejected"
	LogStatusAcceptedCanceled LogStatus = "AcceptedCanceled"
)

type Measurand string

const (
	MeasurandCurrentExport                Measurand = "Current.Export"
	MeasurandCurrentImport                Measurand = "Current.Import"
	MeasurandCurrentOffered               Measurand = "Current.Offered"
	MeasurandEnergyActiveExportRegister   Measurand = "Energy.Active.Export.Register"
	MeasurandEnergyActiveImportRegister   Measurand = "Energy.Active.Import.Register"
	MeasurandEnergyReactiveExportRegister Measurand = "Energy.Reactive.Export.Register"
	MeasurandEnergyReactiveImportRegister Measurand = "Energy.Reactive.Import.Register"
	MeasurandEnergyActiveExportInterval   Measurand = "Energy.Active.Export.Interval"
	MeasurandEnergyActiveImportInterval   Measurand = "Energy.Active.Import.Interval"
	MeasurandEnergyReactiveExportInterval Measurand = "Energy.Reactive.Export.Interval"
	MeasurandEnergyReactiveImportInterval Measurand = "Energy.Reactive.Import.Interval"
	MeasurandFrequency                    Measurand = "Frequency"
	MeasurandPowerActiveExport            Measurand = "Power.Active.Export"
	MeasurandPowerActiveImport            Measurand = "Power.Active.Import"
	MeasurandPowerFactor                  Measurand = "Power.Factor"
	MeasurandPowerOffered                 Measurand = "Power.Offered"
	MeasurandPowerReactiveExport          Measurand = "Power.Reactive.Export"
	MeasurandPowerReactiveImport          Measurand = "Power.Reactive.Import"
	MeasurandSoC                          Measurand = "SoC"
	MeasurandTemperature                  Measurand = "Temperature"
	MeasurandVoltage                      Measurand = "Voltage"
)

type ReadingContext string

const (
	ReadingContextInterruptionBegin ReadingContext = "Interruption.Begin"
	ReadingContextInterruptionEnd   ReadingContext = "Interruption.End"
	ReadingContextOther             ReadingContext = "Other"
	ReadingContextSampleClock       ReadingContext = "Sample.Clock"
	ReadingContextSamplePeriodic    ReadingContext = "Sample.Periodic"
	ReadingContextTransactionBegin  ReadingContext = "Transaction.Begin"
	ReadingContextTransactionEnd    ReadingContext = "Transaction.End"
	ReadingContextTrigger           ReadingContext = "Trigger"
)

type RecurrencyKindType string

const (
	RecurrencyKindDaily  RecurrencyKindType = "Daily"
	RecurrencyKindWeekly RecurrencyKindType = "Weekly"
)

type RegistrationStatus string

const (
	RegistrationStatusAccepted RegistrationStatus = "Accepted"
	RegistrationStatusPending  RegistrationStatus = "Pending"
	RegistrationStatusRejected RegistrationStatus = "Rejected"
)

type RemoteStartStopStatus string

const (
	RemoteStartStopStatusAccepted RemoteStartStopStatus = "Accepted"
	RemoteStartStopStatusRejected RemoteStartStopStatus = "Rejected"
)

type ReservationStatus string

const (
	ReservationStatusAccepted    ReservationStatus = "Accepted"
	ReservationStatusFaulted     ReservationStatus = "Faulted"
	ReservationStatusOccupied    ReservationStatus = "Occupied"
	ReservationStatusRejected    ReservationStatus = "Rejected"
	ReservationStatusUnavailable ReservationStatus = "Unavailable"
)

type ResetStatus string

const (
	ResetStatusAccepted ResetStatus = "Accepted"
	ResetStatusRejected ResetStatus = "Rejected"
)

type ResetType string

const (
	ResetTypeHard ResetType = "Hard"
	ResetTypeSoft ResetType = "Soft"
)

type SendLocalListStatus string

const (
	SendLocalListStatusAccepted        SendLocalListStatus = "Accepted"
	SendLocalListStatusFailed          SendLocalListStatus = "Failed"
	SendLocalListStatusNotSupported    SendLocalListStatus = "NotSupported"
	SendLocalListStatusVersionMismatch SendLocalListStatus = "VersionMismatch"
)

type TriggerMessageStatus string

const (
	TriggerMessageStatusAccepted       TriggerMessageStatus = "Accepted"
	TriggerMessageStatusRejected       TriggerMessageStatus = "Rejected"
	TriggerMessageStatusNotImplemented TriggerMessageStatus = "NotImplemented"
)

type UnlockStatus string

const (
	UnlockStatusUnlocked     UnlockStatus = "Unlocked"
	UnlockStatusUnlockFailed UnlockStatus = "UnlockFailed"
	UnlockStatusNotSupported UnlockStatus = "NotSupported"
)

type UpdateStatus string

const (
	UpdateStatusAccepted        UpdateStatus = "Accepted"
	UpdateStatusFailed          UpdateStatus = "Failed"
	UpdateStatusNotSupported    UpdateStatus = "NotSupported"
	UpdateStatusVersionMismatch UpdateStatus = "VersionMismatch"
)

type UpdateType string

const (
	UpdateTypeDifferential UpdateType = "Differential"
	UpdateTypeFull         UpdateType = "Full"
)

type ValueFormat string

const (
	ValueFormatRaw        ValueFormat = "Raw"
	ValueFormatSignedData ValueFormat = "SignedData"
)

type ChargePointErrorCode string

const (
	ChargePointErrorCodeConnectorLockFailure ChargePointErrorCode = "ConnectorLockFailure"
	ChargePointErrorCodeEVCommunicationError ChargePointErrorCode = "EVCommunicationError"
	ChargePointErrorCodeGroundFailure        ChargePointErrorCode = "GroundFailure"
	ChargePointErrorCodeHighTemperature      ChargePointErrorCode = "HighTemperature"
	ChargePointErrorCodeInternalError        ChargePointErrorCode = "InternalError"
	ChargePointErrorCodeLocalListConflict    ChargePointErrorCode = "LocalListConflict"
	ChargePointErrorCodeNoError              ChargePointErrorCode = "NoError"
	ChargePointErrorCodeOtherError           ChargePointErrorCode = "OtherError"
	ChargePointErrorCodeOverCurrentFailure   ChargePointErrorCode = "OverCurrentFailure"
	ChargePointErrorCodeOverVoltage          ChargePointErrorCode = "OverVoltage"
	ChargePointErrorCodePowerMeterFailure    ChargePointErrorCode = "PowerMeterFailure"
	ChargePointErrorCodePowerSwitchFailure   ChargePointErrorCode = "PowerSwitchFailure"
	ChargePointErrorCodeReaderFailure        ChargePointErrorCode = "ReaderFailure"
	ChargePointErrorCodeResetFailure         ChargePointErrorCode = "ResetFailure"
	ChargePointErrorCodeUnderVoltage         ChargePointErrorCode = "UnderVoltage"
	ChargePointErrorCodeWeakSignal           ChargePointErrorCode = "WeakSignal"
)

type ChargePointStatus string

const (
	ChargePointStatusAvailable     ChargePointStatus = "Available"
	ChargePointStatusPreparing     ChargePointStatus = "Preparing"
	ChargePointStatusCharging      ChargePointStatus = "Charging"
	ChargePointStatusSuspendedEVSE ChargePointStatus = "SuspendedEVSE"
	ChargePointStatusSuspendedEV   ChargePointStatus = "SuspendedEV"
	ChargePointStatusFinishing     ChargePointStatus = "Finishing"
	ChargePointStatusReserved      ChargePointStatus = "Reserved"
	ChargePointStatusUnavailable   ChargePointStatus = "Unavailable"
	ChargePointStatusFaulted       ChargePointStatus = "Faulted"
)

type Reason string

const (
	ReasonDeAuthorized   Reason = "DeAuthorized"
	ReasonEmergencyStop  Reason = "EmergencyStop"
	ReasonEVDisconnected Reason = "EVDisconnected"
	ReasonHardReset      Reason = "HardReset"
	ReasonLocal          Reason = "Local"
	ReasonOther          Reason = "Other"
	ReasonPowerLoss      Reason = "PowerLoss"
	ReasonReboot         Reason = "Reboot"
	ReasonRemote         Reason = "Remote"
	ReasonSoftReset      Reason = "SoftReset"
	ReasonUnlockCommand  Reason = "UnlockCommand"
)

type MessageTrigger string

const (
	MessageTriggerBootNotification              = "BootNotification"
	MessageTriggerDiagnosticsStatusNotification = "DiagnosticsStatusNotification"
	MessageTriggerFirmwareStatusNotification    = "FirmwareStatusNotification"
	MessageTriggerHeartbeat                     = "Heartbeat"
	MessageTriggerMeterValues                   = "MeterValues"
	MessageTriggerStatusNotification            = "StatusNotification"
)
