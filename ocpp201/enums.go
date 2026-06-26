// Package ocpp201 contains the data structures for OCPP 2.0.1 messages.
package ocpp201

// --- Enums ---

type APNAuthenticationType string

const (
	APNAuthenticationCHAP APNAuthenticationType = "CHAP"
	APNAuthenticationNONE APNAuthenticationType = "NONE"
	APNAuthenticationPAP  APNAuthenticationType = "PAP"
	APNAuthenticationAUTO APNAuthenticationType = "AUTO"
)

type AttributeType string

const (
	AttributeTypeActual AttributeType = "Actual"
	AttributeTypeTarget AttributeType = "Target"
	AttributeTypeMinSet AttributeType = "MinSet"
	AttributeTypeMaxSet AttributeType = "MaxSet"
)

type AuthorizeCertificateStatus string

const (
	AuthorizeCertificateStatusAccepted               AuthorizeCertificateStatus = "Accepted"
	AuthorizeCertificateStatusSignatureError         AuthorizeCertificateStatus = "SignatureError"
	AuthorizeCertificateStatusCertificateExpired     AuthorizeCertificateStatus = "CertificateExpired"
	AuthorizeCertificateStatusCertificateRevoked     AuthorizeCertificateStatus = "CertificateRevoked"
	AuthorizeCertificateStatusNoCertificateAvailable AuthorizeCertificateStatus = "NoCertificateAvailable"
	AuthorizeCertificateStatusCertChainError         AuthorizeCertificateStatus = "CertChainError"
	AuthorizeCertificateStatusContractCancelled      AuthorizeCertificateStatus = "ContractCancelled"
)

type AuthorizationStatus string

const (
	AuthorizationStatusAccepted           AuthorizationStatus = "Accepted"
	AuthorizationStatusBlocked            AuthorizationStatus = "Blocked"
	AuthorizationStatusConcurrentTx       AuthorizationStatus = "ConcurrentTx"
	AuthorizationStatusExpired            AuthorizationStatus = "Expired"
	AuthorizationStatusInvalid            AuthorizationStatus = "Invalid"
	AuthorizationStatusNoCredit           AuthorizationStatus = "NoCredit"
	AuthorizationStatusNotAllowedTypeEVSE AuthorizationStatus = "NotAllowedTypeEVSE"
	AuthorizationStatusNotAtThisLocation  AuthorizationStatus = "NotAtThisLocation"
	AuthorizationStatusNotAtThisTime      AuthorizationStatus = "NotAtThisTime"
	AuthorizationStatusUnknown            AuthorizationStatus = "Unknown"
)

type BootReason string

const (
	BootReasonApplicationReset BootReason = "ApplicationReset"
	BootReasonFirmwareUpdate   BootReason = "FirmwareUpdate"
	BootReasonLocalReset       BootReason = "LocalReset"
	BootReasonPowerUp          BootReason = "PowerUp"
	BootReasonRemoteReset      BootReason = "RemoteReset"
	BootReasonScheduledReset   BootReason = "ScheduledReset"
	BootReasonTriggered        BootReason = "Triggered"
	BootReasonUnknown          BootReason = "Unknown"
	BootReasonWatchdog         BootReason = "Watchdog"
)

type CertificateAction string

const (
	CertificateActionInstall CertificateAction = "Install"
	CertificateActionUpdate  CertificateAction = "Update"
)

type CertificateSignedStatus string

const (
	CertificateSignedStatusAccepted CertificateSignedStatus = "Accepted"
	CertificateSignedStatusRejected CertificateSignedStatus = "Rejected"
)

type CertificateSigningUse string

const (
	CertificateSigningUseChargingStationCertificate CertificateSigningUse = "ChargingStationCertificate"
	CertificateSigningUseV2GCertificate             CertificateSigningUse = "V2GCertificate"
)

type ChangeAvailabilityStatus string

const (
	ChangeAvailabilityStatusAccepted  ChangeAvailabilityStatus = "Accepted"
	ChangeAvailabilityStatusRejected  ChangeAvailabilityStatus = "Rejected"
	ChangeAvailabilityStatusScheduled ChangeAvailabilityStatus = "Scheduled"
)

type ChargingProfileKind string

const (
	ChargingProfileKindAbsolute  ChargingProfileKind = "Absolute"
	ChargingProfileKindRecurring ChargingProfileKind = "Recurring"
	ChargingProfileKindRelative  ChargingProfileKind = "Relative"
)

type ChargingProfilePurpose string

const (
	ChargingProfilePurposeChargingStationExternalConstraints ChargingProfilePurpose = "ChargingStationExternalConstraints"
	ChargingProfilePurposeChargingStationMaxProfile          ChargingProfilePurpose = "ChargingStationMaxProfile"
	ChargingProfilePurposeTxDefaultProfile                   ChargingProfilePurpose = "TxDefaultProfile"
	ChargingProfilePurposeTxProfile                          ChargingProfilePurpose = "TxProfile"
)

type ChargingRateUnit string

const (
	ChargingRateUnitWatts   ChargingRateUnit = "W"
	ChargingRateUnitAmperes ChargingRateUnit = "A"
)

type ChargingState string

const (
	ChargingStateCharging      ChargingState = "Charging"
	ChargingStateEVConnected   ChargingState = "EVConnected"
	ChargingStateSuspendedEV   ChargingState = "SuspendedEV"
	ChargingStateSuspendedEVSE ChargingState = "SuspendedEVSE"
	ChargingStateIdle          ChargingState = "Idle"
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

type ClearDisplayMessageStatus string

const (
	ClearDisplayMessageStatusAccepted ClearDisplayMessageStatus = "Accepted"
	ClearDisplayMessageStatusUnknown  ClearDisplayMessageStatus = "Unknown"
)

type ClearMonitoringStatus string

const (
	ClearMonitoringStatusAccepted ClearMonitoringStatus = "Accepted"
	ClearMonitoringStatusRejected ClearMonitoringStatus = "Rejected"
	ClearMonitoringStatusNotFound ClearMonitoringStatus = "NotFound"
)

type ConnectorStatus string

const (
	ConnectorStatusAvailable   ConnectorStatus = "Available"
	ConnectorStatusOccupied    ConnectorStatus = "Occupied"
	ConnectorStatusReserved    ConnectorStatus = "Reserved"
	ConnectorStatusUnavailable ConnectorStatus = "Unavailable"
	ConnectorStatusFaulted     ConnectorStatus = "Faulted"
)

type CostKind string

const (
	CostKindCarbonDioxideEmissions        CostKind = "CarbonDioxideEmissions"
	CostKindRelativePricePercentage       CostKind = "RelativePricePercentage"
	CostKindRenewableGenerationPercentage CostKind = "RenewableGenerationPercentage"
)

type CustomerInformationStatus string

const (
	CustomerInformationStatusAccepted CustomerInformationStatus = "Accepted"
	CustomerInformationStatusRejected CustomerInformationStatus = "Rejected"
	CustomerInformationStatusInvalid  CustomerInformationStatus = "Invalid"
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

type DisplayMessageStatus string

const (
	DisplayMessageStatusAccepted                  DisplayMessageStatus = "Accepted"
	DisplayMessageStatusNotSupportedMessageFormat DisplayMessageStatus = "NotSupportedMessageFormat"
	DisplayMessageStatusRejected                  DisplayMessageStatus = "Rejected"
	DisplayMessageStatusNotSupportedPriority      DisplayMessageStatus = "NotSupportedPriority"
	DisplayMessageStatusNotSupportedState         DisplayMessageStatus = "NotSupportedState"
	DisplayMessageStatusUnknownTransaction        DisplayMessageStatus = "UnknownTransaction"
)

type EnergyTransferMode string

const (
	EnergyTransferModeDC            EnergyTransferMode = "DC"
	EnergyTransferModeACSinglePhase EnergyTransferMode = "AC_single_phase"
	EnergyTransferModeACThreePhase  EnergyTransferMode = "AC_three_phase"
)

type EventNotification string

const (
	EventNotificationHardWiredNotification EventNotification = "HardWiredNotification"
	EventNotificationHardWiredMonitor      EventNotification = "HardWiredMonitor"
	EventNotificationPull                  EventNotification = "Pull"
	EventNotificationPeriodic              EventNotification = "Periodic"
)

type EventTrigger string

const (
	EventTriggerAlerting EventTrigger = "Alerting"
	EventTriggerDelta    EventTrigger = "Delta"
	EventTriggerPeriodic EventTrigger = "Periodic"
)

type FirmwareStatus string

const (
	FirmwareStatusDownloaded                FirmwareStatus = "Downloaded"
	FirmwareStatusDownloadFailed            FirmwareStatus = "DownloadFailed"
	FirmwareStatusDownloading               FirmwareStatus = "Downloading"
	FirmwareStatusDownloadScheduled         FirmwareStatus = "DownloadScheduled"
	FirmwareStatusDownloadPaused            FirmwareStatus = "DownloadPaused"
	FirmwareStatusIdle                      FirmwareStatus = "Idle"
	FirmwareStatusInstallationFailed        FirmwareStatus = "InstallationFailed"
	FirmwareStatusInstalling                FirmwareStatus = "Installing"
	FirmwareStatusInstalled                 FirmwareStatus = "Installed"
	FirmwareStatusInstallRebooting          FirmwareStatus = "InstallRebooting"
	FirmwareStatusInstallScheduled          FirmwareStatus = "InstallScheduled"
	FirmwareStatusInstallVerificationFailed FirmwareStatus = "InstallVerificationFailed"
	FirmwareStatusInvalidSignature          FirmwareStatus = "InvalidSignature"
	FirmwareStatusSignatureVerified         FirmwareStatus = "SignatureVerified"
)

type GenericDeviceModelStatus string

const (
	GenericDeviceModelStatusAccepted     GenericDeviceModelStatus = "Accepted"
	GenericDeviceModelStatusRejected     GenericDeviceModelStatus = "Rejected"
	GenericDeviceModelStatusNotSupported GenericDeviceModelStatus = "NotSupported"
	GenericDeviceModelStatusEmptyResult  GenericDeviceModelStatus = "EmptyResult"
)

type GenericStatus string

const (
	GenericStatusAccepted GenericStatus = "Accepted"
	GenericStatusRejected GenericStatus = "Rejected"
)

type GetCertificateIdUse string

const (
	GetCertificateIdUseV2GRootCertificate          GetCertificateIdUse = "V2GRootCertificate"
	GetCertificateIdUseMORootCertificate           GetCertificateIdUse = "MORootCertificate"
	GetCertificateIdUseCSMSRootCertificate         GetCertificateIdUse = "CSMSRootCertificate"
	GetCertificateIdUseChargingStationCertificate  GetCertificateIdUse = "ChargingStationCertificate"
	GetCertificateIdUseManufacturerRootCertificate GetCertificateIdUse = "ManufacturerRootCertificate"
)

type GetCertificateStatus string

const (
	GetCertificateStatusAccepted GetCertificateStatus = "Accepted"
	GetCertificateStatusFailed   GetCertificateStatus = "Failed"
)

type GetChargingProfileStatus string

const (
	GetChargingProfileStatusAccepted   GetChargingProfileStatus = "Accepted"
	GetChargingProfileStatusNoProfiles GetChargingProfileStatus = "NoProfiles"
)

type GetVariableStatus string

const (
	GetVariableStatusAccepted                  GetVariableStatus = "Accepted"
	GetVariableStatusRejected                  GetVariableStatus = "Rejected"
	GetVariableStatusUnknownComponent          GetVariableStatus = "UnknownComponent"
	GetVariableStatusUnknownVariable           GetVariableStatus = "UnknownVariable"
	GetVariableStatusNotSupportedAttributeType GetVariableStatus = "NotSupportedAttributeType"
)

type HashAlgorithm string

const (
	HashAlgorithmSHA256 HashAlgorithm = "SHA256"
	HashAlgorithmSHA384 HashAlgorithm = "SHA384"
	HashAlgorithmSHA512 HashAlgorithm = "SHA512"
)

type IdTokenType string

const (
	IdTokenTypeCentral         IdTokenType = "Central"
	IdTokenTypeEmaid           IdTokenType = "eMAID"
	IdTokenTypeISO14443        IdTokenType = "ISO14443"
	IdTokenTypeISO15693        IdTokenType = "ISO15693"
	IdTokenTypeKeyCode         IdTokenType = "KeyCode"
	IdTokenTypeLocal           IdTokenType = "Local"
	IdTokenTypeMacAddress      IdTokenType = "MacAddress"
	IdTokenTypeNoAuthorization IdTokenType = "NoAuthorization"
)

type InstallCertificateStatus string

const (
	InstallCertificateStatusAccepted InstallCertificateStatus = "Accepted"
	InstallCertificateStatusRejected InstallCertificateStatus = "Rejected"
	InstallCertificateStatusFailed   InstallCertificateStatus = "Failed"
	InstallCertificateStatusPending  InstallCertificateStatus = "Pending"
)

type InstallCertificateUse string

const (
	InstallCertificateUseV2GRootCertificate          InstallCertificateUse = "V2GRootCertificate"
	InstallCertificateUseMORootCertificate           InstallCertificateUse = "MORootCertificate"
	InstallCertificateUseCSMSRootCertificate         InstallCertificateUse = "CSMSRootCertificate"
	InstallCertificateUseChargingStationCertificate  InstallCertificateUse = "ChargingStationCertificate"
	InstallCertificateUseManufacturerRootCertificate InstallCertificateUse = "ManufacturerRootCertificate"
)

type ISO15118EVCertificateStatus string

const (
	ISO15118EVCertificateStatusAccepted ISO15118EVCertificateStatus = "Accepted"
	ISO15118EVCertificateStatusFailed   ISO15118EVCertificateStatus = "Failed"
)

type Location string

const (
	LocationBody   Location = "Body"
	LocationCable  Location = "Cable"
	LocationEV     Location = "EV"
	LocationInlet  Location = "Inlet"
	LocationOutlet Location = "Outlet"
)

type Log string

const (
	LogDiagnosticsLog Log = "DiagnosticsLog"
	LogSecurityLog    Log = "SecurityLog"
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

type MessageFormat string

const (
	MessageFormatASCII MessageFormat = "ASCII"
	MessageFormatHTML  MessageFormat = "HTML"
	MessageFormatURI   MessageFormat = "URI"
	MessageFormatUTF8  MessageFormat = "UTF8"
)

type MessagePriority string

const (
	MessagePriorityAlwaysFront MessagePriority = "AlwaysFront"
	MessagePriorityInFront     MessagePriority = "InFront"
	MessagePriorityNormalCycle MessagePriority = "NormalCycle"
)

type MessageState string

const (
	MessageStateCharging MessageState = "Charging"
	MessageStateFaulted  MessageState = "Faulted"
	MessageStateIdle     MessageState = "Idle"
)

type Monitor string

const (
	MonitorUpperThreshold       Monitor = "UpperThreshold"
	MonitorLowerThreshold       Monitor = "LowerThreshold"
	MonitorDelta                Monitor = "Delta"
	MonitorPeriodic             Monitor = "Periodic"
	MonitorPeriodicClockAligned Monitor = "PeriodicClockAligned"
)

type NotifyEVChargingNeedsStatus string

const (
	NotifyEVChargingNeedsStatusAccepted   NotifyEVChargingNeedsStatus = "Accepted"
	NotifyEVChargingNeedsStatusRejected   NotifyEVChargingNeedsStatus = "Rejected"
	NotifyEVChargingNeedsStatusProcessing NotifyEVChargingNeedsStatus = "Processing"
)

type OperationalStatus string

const (
	OperationalStatusInoperative OperationalStatus = "Inoperative"
	OperationalStatusOperative   OperationalStatus = "Operative"
)

type Phase string

const (
	PhaseL1   Phase = "L1"
	PhaseL2   Phase = "L2"
	PhaseL3   Phase = "L3"
	PhaseN    Phase = "N"
	PhaseL1N  Phase = "L1-N"
	PhaseL2N  Phase = "L2-N"
	PhaseL3N  Phase = "L3-N"
	PhaseL1L2 Phase = "L1-L2"
	PhaseL2L3 Phase = "L2-L3"
	PhaseL3L1 Phase = "L3-L1"
)

type PublishFirmwareStatus string

const (
	PublishFirmwareStatusIdle              PublishFirmwareStatus = "Idle"
	PublishFirmwareStatusDownloadScheduled PublishFirmwareStatus = "DownloadScheduled"
	PublishFirmwareStatusDownloading       PublishFirmwareStatus = "Downloading"
	PublishFirmwareStatusDownloaded        PublishFirmwareStatus = "Downloaded"
	PublishFirmwareStatusPublished         PublishFirmwareStatus = "Published"
	PublishFirmwareStatusDownloadFailed    PublishFirmwareStatus = "DownloadFailed"
	PublishFirmwareStatusDownloadPaused    PublishFirmwareStatus = "DownloadPaused"
	PublishFirmwareStatusInvalidChecksum   PublishFirmwareStatus = "InvalidChecksum"
	PublishFirmwareStatusChecksumVerified  PublishFirmwareStatus = "ChecksumVerified"
	PublishFirmwareStatusPublishFailed     PublishFirmwareStatus = "PublishFailed"
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

type Reason string

const (
	ReasonDeAuthorized       Reason = "DeAuthorized"
	ReasonEmergencyStop      Reason = "EmergencyStop"
	ReasonEnergyLimitReached Reason = "EnergyLimitReached"
	ReasonEVDisconnected     Reason = "EVDisconnected"
	ReasonGroundFault        Reason = "GroundFault"
	ReasonImmediateReset     Reason = "ImmediateReset"
	ReasonLocal              Reason = "Local"
	ReasonLocalOutOfCredit   Reason = "LocalOutOfCredit"
	ReasonMasterPass         Reason = "MasterPass"
	ReasonOther              Reason = "Other"
	ReasonOvercurrentFault   Reason = "OvercurrentFault"
	ReasonPowerLoss          Reason = "PowerLoss"
	ReasonPowerQuality       Reason = "PowerQuality"
	ReasonReboot             Reason = "Reboot"
	ReasonRemote             Reason = "Remote"
	ReasonSOCLimitReached    Reason = "SOCLimitReached"
	ReasonStoppedByEV        Reason = "StoppedByEV"
	ReasonTimeLimitReached   Reason = "TimeLimitReached"
	ReasonTimeout            Reason = "Timeout"
)

type RecurrencyKind string

const (
	RecurrencyKindDaily  RecurrencyKind = "Daily"
	RecurrencyKindWeekly RecurrencyKind = "Weekly"
)

type RegistrationStatus string

const (
	RegistrationStatusAccepted RegistrationStatus = "Accepted"
	RegistrationStatusPending  RegistrationStatus = "Pending"
	RegistrationStatusRejected RegistrationStatus = "Rejected"
)

type ReportBase string

const (
	ReportBaseConfigurationInventory ReportBase = "ConfigurationInventory"
	ReportBaseFullInventory          ReportBase = "FullInventory"
	ReportBaseSummaryInventory       ReportBase = "SummaryInventory"
)

type RequestStartStopStatus string

const (
	RequestStartStopStatusAccepted RequestStartStopStatus = "Accepted"
	RequestStartStopStatusRejected RequestStartStopStatus = "Rejected"
)

type ReservationUpdateStatus string

const (
	ReservationUpdateStatusExpired ReservationUpdateStatus = "Expired"
	ReservationUpdateStatusRemoved ReservationUpdateStatus = "Removed"
)

type ReserveNowStatus string

const (
	ReserveNowStatusAccepted    ReserveNowStatus = "Accepted"
	ReserveNowStatusFaulted     ReserveNowStatus = "Faulted"
	ReserveNowStatusOccupied    ReserveNowStatus = "Occupied"
	ReserveNowStatusRejected    ReserveNowStatus = "Rejected"
	ReserveNowStatusUnavailable ReserveNowStatus = "Unavailable"
)

type Reset string

const (
	ResetImmediate Reset = "Immediate"
	ResetOnIdle    Reset = "OnIdle"
)

type SendLocalListStatus string

const (
	SendLocalListStatusAccepted        SendLocalListStatus = "Accepted"
	SendLocalListStatusFailed          SendLocalListStatus = "Failed"
	SendLocalListStatusVersionMismatch SendLocalListStatus = "VersionMismatch"
)

type SetMonitoringStatus string

const (
	SetMonitoringStatusAccepted               SetMonitoringStatus = "Accepted"
	SetMonitoringStatusUnknownComponent       SetMonitoringStatus = "UnknownComponent"
	SetMonitoringStatusUnknownVariable        SetMonitoringStatus = "UnknownVariable"
	SetMonitoringStatusUnsupportedMonitorType SetMonitoringStatus = "UnsupportedMonitorType"
	SetMonitoringStatusRejected               SetMonitoringStatus = "Rejected"
	SetMonitoringStatusDuplicate              SetMonitoringStatus = "Duplicate"
)

type SetNetworkProfileStatus string

const (
	SetNetworkProfileStatusAccepted SetNetworkProfileStatus = "Accepted"
	SetNetworkProfileStatusRejected SetNetworkProfileStatus = "Rejected"
	SetNetworkProfileStatusFailed   SetNetworkProfileStatus = "Failed"
)

type SetVariableStatus string

const (
	SetVariableStatusAccepted                  SetVariableStatus = "Accepted"
	SetVariableStatusRejected                  SetVariableStatus = "Rejected"
	SetVariableStatusUnknownComponent          SetVariableStatus = "UnknownComponent"
	SetVariableStatusUnknownVariable           SetVariableStatus = "UnknownVariable"
	SetVariableStatusNotSupportedAttributeType SetVariableStatus = "NotSupportedAttributeType"
	SetVariableStatusRebootRequired            SetVariableStatus = "RebootRequired"
)

type TransactionEvent string

const (
	TransactionEventEnded   TransactionEvent = "Ended"
	TransactionEventStarted TransactionEvent = "Started"
	TransactionEventUpdated TransactionEvent = "Updated"
)

type TriggerMessageStatus string

const (
	TriggerMessageStatusAccepted       TriggerMessageStatus = "Accepted"
	TriggerMessageStatusRejected       TriggerMessageStatus = "Rejected"
	TriggerMessageStatusNotImplemented TriggerMessageStatus = "NotImplemented"
)

type TriggerReason string

const (
	TriggerReasonAuthorized           TriggerReason = "Authorized"
	TriggerReasonCablePluggedIn       TriggerReason = "CablePluggedIn"
	TriggerReasonChargingRateChanged  TriggerReason = "ChargingRateChanged"
	TriggerReasonChargingStateChanged TriggerReason = "ChargingStateChanged"
	TriggerReasonDeauthorized         TriggerReason = "Deauthorized"
	TriggerReasonEnergyLimitReached   TriggerReason = "EnergyLimitReached"
	TriggerReasonEVCommunicationLost  TriggerReason = "EVCommunicationLost"
	TriggerReasonEVConnectTimeout     TriggerReason = "EVConnectTimeout"
	TriggerReasonMeterValueClock      TriggerReason = "MeterValueClock"
	TriggerReasonMeterValuePeriodic   TriggerReason = "MeterValuePeriodic"
	TriggerReasonTimeLimitReached     TriggerReason = "TimeLimitReached"
	TriggerReasonTrigger              TriggerReason = "Trigger"
	TriggerReasonUnlockCommand        TriggerReason = "UnlockCommand"
	TriggerReasonStopAuthorized       TriggerReason = "StopAuthorized"
	TriggerReasonEVDeparted           TriggerReason = "EVDeparted"
	TriggerReasonEVDetected           TriggerReason = "EVDetected"
	TriggerReasonRemoteStop           TriggerReason = "RemoteStop"
	TriggerReasonRemoteStart          TriggerReason = "RemoteStart"
	TriggerReasonAbnormalCondition    TriggerReason = "AbnormalCondition"
	TriggerReasonSignedDataReceived   TriggerReason = "SignedDataReceived"
	TriggerReasonResetCommand         TriggerReason = "ResetCommand"
)

type UnlockStatus string

const (
	UnlockStatusUnlocked                     UnlockStatus = "Unlocked"
	UnlockStatusUnlockFailed                 UnlockStatus = "UnlockFailed"
	UnlockStatusOngoingAuthorizedTransaction UnlockStatus = "OngoingAuthorizedTransaction"
	UnlockStatusUnknownConnector             UnlockStatus = "UnknownConnector"
)

type UnpublishFirmwareStatus string

const (
	UnpublishFirmwareStatusDownloadOngoing UnpublishFirmwareStatus = "DownloadOngoing"
	UnpublishFirmwareStatusNoFirmware      UnpublishFirmwareStatus = "NoFirmware"
	UnpublishFirmwareStatusUnpublished     UnpublishFirmwareStatus = "Unpublished"
)

type UpdateFirmwareStatus string

const (
	UpdateFirmwareStatusAccepted           UpdateFirmwareStatus = "Accepted"
	UpdateFirmwareStatusRejected           UpdateFirmwareStatus = "Rejected"
	UpdateFirmwareStatusAcceptedCanceled   UpdateFirmwareStatus = "AcceptedCanceled"
	UpdateFirmwareStatusInvalidCertificate UpdateFirmwareStatus = "InvalidCertificate"
	UpdateFirmwareStatusRevokedCertificate UpdateFirmwareStatus = "RevokedCertificate"
)

type Update string

const (
	UpdateDifferential Update = "Differential"
	UpdateFull         Update = "Full"
)

type UploadLogStatus string

const (
	UploadLogStatusBadMessage            UploadLogStatus = "BadMessage"
	UploadLogStatusIdle                  UploadLogStatus = "Idle"
	UploadLogStatusNotSupportedOperation UploadLogStatus = "NotSupportedOperation"
	UploadLogStatusPermissionDenied      UploadLogStatus = "PermissionDenied"
	UploadLogStatusUploaded              UploadLogStatus = "Uploaded"
	UploadLogStatusUploadFailure         UploadLogStatus = "UploadFailure"
	UploadLogStatusUploading             UploadLogStatus = "Uploading"
	UploadLogStatusAcceptedCanceled      UploadLogStatus = "AcceptedCanceled"
)

type VPNType string

const (
	VPNIKEv2 VPNType = "IKEv2"
	VPNIPSec VPNType = "IPSec"
	VPNL2TP  VPNType = "L2TP"
	VPNPPTP  VPNType = "PPTP"
)
