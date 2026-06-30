// Package ocpp21 contains the data structures for OCPP 2.1 messages.
package ocpp21

// Defines all enumerated types used in the OCPP 2.1 specification.

type APNAuthenticationEnum string

const (
	APNAuthenticationEnumPAP  APNAuthenticationEnum = "PAP"
	APNAuthenticationEnumCHAP APNAuthenticationEnum = "CHAP"
	APNAuthenticationEnumNONE APNAuthenticationEnum = "NONE"
	APNAuthenticationEnumAUTO APNAuthenticationEnum = "AUTO"
)

type AttributeEnum string

const (
	AttributeEnumActual AttributeEnum = "Actual"
	AttributeEnumTarget AttributeEnum = "Target"
	AttributeEnumMinSet AttributeEnum = "MinSet"
	AttributeEnumMaxSet AttributeEnum = "MaxSet"
)

type AuthorizationStatusEnum string

const (
	AuthorizationStatusEnumAccepted           AuthorizationStatusEnum = "Accepted"
	AuthorizationStatusEnumBlocked            AuthorizationStatusEnum = "Blocked"
	AuthorizationStatusEnumConcurrentTx       AuthorizationStatusEnum = "ConcurrentTx"
	AuthorizationStatusEnumExpired            AuthorizationStatusEnum = "Expired"
	AuthorizationStatusEnumInvalid            AuthorizationStatusEnum = "Invalid"
	AuthorizationStatusEnumNoCredit           AuthorizationStatusEnum = "NoCredit"
	AuthorizationStatusEnumNotAllowedTypeEVSE AuthorizationStatusEnum = "NotAllowedTypeEVSE"
	AuthorizationStatusEnumNotAtThisLocation  AuthorizationStatusEnum = "NotAtThisLocation"
	AuthorizationStatusEnumNotAtThisTime      AuthorizationStatusEnum = "NotAtThisTime"
	AuthorizationStatusEnumUnknown            AuthorizationStatusEnum = "Unknown"
)

type AuthorizeCertificateStatusEnum string

const (
	AuthorizeCertificateStatusEnumAccepted               AuthorizeCertificateStatusEnum = "Accepted"
	AuthorizeCertificateStatusEnumSignatureError         AuthorizeCertificateStatusEnum = "SignatureError"
	AuthorizeCertificateStatusEnumCertificateExpired     AuthorizeCertificateStatusEnum = "CertificateExpired"
	AuthorizeCertificateStatusEnumCertificateRevoked     AuthorizeCertificateStatusEnum = "CertificateRevoked"
	AuthorizeCertificateStatusEnumNoCertificateAvailable AuthorizeCertificateStatusEnum = "NoCertificateAvailable"
	AuthorizeCertificateStatusEnumCertChainError         AuthorizeCertificateStatusEnum = "CertChainError"
	AuthorizeCertificateStatusEnumContractCancelled      AuthorizeCertificateStatusEnum = "ContractCancelled"
)

type BatterySwapEventEnum string

const (
	BatterySwapEventEnumBatteryIn         BatterySwapEventEnum = "BatteryIn"
	BatterySwapEventEnumBatteryOut        BatterySwapEventEnum = "BatteryOut"
	BatterySwapEventEnumBatteryOutTimeout BatterySwapEventEnum = "BatteryOutTimeout"
)

type BootReasonEnum string

const (
	BootReasonEnumApplicationReset BootReasonEnum = "ApplicationReset"
	BootReasonEnumFirmwareUpdate   BootReasonEnum = "FirmwareUpdate"
	BootReasonEnumLocalReset       BootReasonEnum = "LocalReset"
	BootReasonEnumPowerUp          BootReasonEnum = "PowerUp"
	BootReasonEnumRemoteReset      BootReasonEnum = "RemoteReset"
	BootReasonEnumScheduledReset   BootReasonEnum = "ScheduledReset"
	BootReasonEnumTriggered        BootReasonEnum = "Triggered"
	BootReasonEnumUnknown          BootReasonEnum = "Unknown"
	BootReasonEnumWatchdog         BootReasonEnum = "Watchdog"
)

type CancelReservationStatusEnum string

const (
	CancelReservationStatusEnumAccepted CancelReservationStatusEnum = "Accepted"
	CancelReservationStatusEnumRejected CancelReservationStatusEnum = "Rejected"
)

type CertificateActionEnum string

const (
	CertificateActionEnumInstall CertificateActionEnum = "Install"
	CertificateActionEnumUpdate  CertificateActionEnum = "Update"
)

type CertificateSignedStatusEnum string

const (
	CertificateSignedStatusEnumAccepted CertificateSignedStatusEnum = "Accepted"
	CertificateSignedStatusEnumRejected CertificateSignedStatusEnum = "Rejected"
)

type CertificateSigningUseEnum string

const (
	CertificateSigningUseEnumChargingStationCertificate CertificateSigningUseEnum = "ChargingStationCertificate"
	CertificateSigningUseEnumV2GCertificate             CertificateSigningUseEnum = "V2GCertificate"
	CertificateSigningUseEnumV2G20Certificate           CertificateSigningUseEnum = "V2G20Certificate"
)

type CertificateStatusEnum string

const (
	CertificateStatusEnumGood    CertificateStatusEnum = "Good"
	CertificateStatusEnumRevoked CertificateStatusEnum = "Revoked"
	CertificateStatusEnumUnknown CertificateStatusEnum = "Unknown"
	CertificateStatusEnumFailed  CertificateStatusEnum = "Failed"
)

type CertificateStatusSourceEnum string

const (
	CertificateStatusSourceEnumCRL  CertificateStatusSourceEnum = "CRL"
	CertificateStatusSourceEnumOCSP CertificateStatusSourceEnum = "OCSP"
)

type ChangeAvailabilityStatusEnum string

const (
	ChangeAvailabilityStatusEnumAccepted  ChangeAvailabilityStatusEnum = "Accepted"
	ChangeAvailabilityStatusEnumRejected  ChangeAvailabilityStatusEnum = "Rejected"
	ChangeAvailabilityStatusEnumScheduled ChangeAvailabilityStatusEnum = "Scheduled"
)

type ChargingProfileKindEnum string

const (
	ChargingProfileKindEnumAbsolute  ChargingProfileKindEnum = "Absolute"
	ChargingProfileKindEnumRecurring ChargingProfileKindEnum = "Recurring"
	ChargingProfileKindEnumRelative  ChargingProfileKindEnum = "Relative"
	ChargingProfileKindEnumDynamic   ChargingProfileKindEnum = "Dynamic"
)

type ChargingProfilePurposeEnum string

const (
	ChargingProfilePurposeEnumChargingStationExternalConstraints ChargingProfilePurposeEnum = "ChargingStationExternalConstraints"
	ChargingProfilePurposeEnumChargingStationMaxProfile          ChargingProfilePurposeEnum = "ChargingStationMaxProfile"
	ChargingProfilePurposeEnumTxDefaultProfile                   ChargingProfilePurposeEnum = "TxDefaultProfile"
	ChargingProfilePurposeEnumTxProfile                          ChargingProfilePurposeEnum = "TxProfile"
	ChargingProfilePurposeEnumPriorityCharging                   ChargingProfilePurposeEnum = "PriorityCharging"
	ChargingProfilePurposeEnumLocalGeneration                    ChargingProfilePurposeEnum = "LocalGeneration"
)

type ChargingProfileStatusEnum string

const (
	ChargingProfileStatusEnumAccepted ChargingProfileStatusEnum = "Accepted"
	ChargingProfileStatusEnumRejected ChargingProfileStatusEnum = "Rejected"
)

type ChargingRateUnitEnum string

const (
	ChargingRateUnitEnumW ChargingRateUnitEnum = "W"
	ChargingRateUnitEnumA ChargingRateUnitEnum = "A"
)

type ChargingStateEnum string

const (
	ChargingStateEnumEVConnected   ChargingStateEnum = "EVConnected"
	ChargingStateEnumCharging      ChargingStateEnum = "Charging"
	ChargingStateEnumSuspendedEV   ChargingStateEnum = "SuspendedEV"
	ChargingStateEnumSuspendedEVSE ChargingStateEnum = "SuspendedEVSE"
	ChargingStateEnumIdle          ChargingStateEnum = "Idle"
)

type ClearCacheStatusEnum string

const (
	ClearCacheStatusEnumAccepted ClearCacheStatusEnum = "Accepted"
	ClearCacheStatusEnumRejected ClearCacheStatusEnum = "Rejected"
)

type ClearChargingProfileStatusEnum string

const (
	ClearChargingProfileStatusEnumAccepted ClearChargingProfileStatusEnum = "Accepted"
	ClearChargingProfileStatusEnumUnknown  ClearChargingProfileStatusEnum = "Unknown"
)

type ClearMessageStatusEnum string

const (
	ClearMessageStatusEnumAccepted ClearMessageStatusEnum = "Accepted"
	ClearMessageStatusEnumUnknown  ClearMessageStatusEnum = "Unknown"
	ClearMessageStatusEnumRejected ClearMessageStatusEnum = "Rejected"
)

type ClearMonitoringStatusEnum string

const (
	ClearMonitoringStatusEnumAccepted ClearMonitoringStatusEnum = "Accepted"
	ClearMonitoringStatusEnumRejected ClearMonitoringStatusEnum = "Rejected"
	ClearMonitoringStatusEnumNotFound ClearMonitoringStatusEnum = "NotFound"
)

type ComponentCriterionEnum string

const (
	ComponentCriterionEnumActive    ComponentCriterionEnum = "Active"
	ComponentCriterionEnumAvailable ComponentCriterionEnum = "Available"
	ComponentCriterionEnumEnabled   ComponentCriterionEnum = "Enabled"
	ComponentCriterionEnumProblem   ComponentCriterionEnum = "Problem"
)

type ConnectorStatusEnum string

const (
	ConnectorStatusEnumAvailable   ConnectorStatusEnum = "Available"
	ConnectorStatusEnumOccupied    ConnectorStatusEnum = "Occupied"
	ConnectorStatusEnumReserved    ConnectorStatusEnum = "Reserved"
	ConnectorStatusEnumUnavailable ConnectorStatusEnum = "Unavailable"
	ConnectorStatusEnumFaulted     ConnectorStatusEnum = "Faulted"
)

type ControlModeEnum string

const (
	ControlModeEnumScheduledControl ControlModeEnum = "ScheduledControl"
	ControlModeEnumDynamicControl   ControlModeEnum = "DynamicControl"
)

type CostDimensionEnum string

const (
	CostDimensionEnumEnergy       CostDimensionEnum = "Energy"
	CostDimensionEnumMaxCurrent   CostDimensionEnum = "MaxCurrent"
	CostDimensionEnumMinCurrent   CostDimensionEnum = "MinCurrent"
	CostDimensionEnumMaxPower     CostDimensionEnum = "MaxPower"
	CostDimensionEnumMinPower     CostDimensionEnum = "MinPower"
	CostDimensionEnumIdleTIme     CostDimensionEnum = "IdleTIme"
	CostDimensionEnumChargingTime CostDimensionEnum = "ChargingTime"
)

type CostKindEnum string

const (
	CostKindEnumCarbonDioxideEmission         CostKindEnum = "CarbonDioxideEmission"
	CostKindEnumRelativePricePercentage       CostKindEnum = "RelativePricePercentage"
	CostKindEnumRenewableGenerationPercentage CostKindEnum = "RenewableGenerationPercentage"
)

type CustomerInformationStatusEnum string

const (
	CustomerInformationStatusEnumAccepted CustomerInformationStatusEnum = "Accepted"
	CustomerInformationStatusEnumRejected CustomerInformationStatusEnum = "Rejected"
	CustomerInformationStatusEnumInvalid  CustomerInformationStatusEnum = "Invalid"
)

type DERControlEnum string

const (
	DERControlEnumEnterService            DERControlEnum = "EnterService"
	DERControlEnumFreqDroop               DERControlEnum = "FreqDroop"
	DERControlEnumFreqWatt                DERControlEnum = "FreqWatt"
	DERControlEnumFixedPFAbsorb           DERControlEnum = "FixedPFAbsorb"
	DERControlEnumFixedPFInject           DERControlEnum = "FixedPFInject"
	DERControlEnumFixedVAR                DERControlEnum = "FixedVar"
	DERControlEnumGradients               DERControlEnum = "Gradients"
	DERControlEnumHFMustTrip              DERControlEnum = "HFMustTrip"
	DERControlEnumHFMayTrip               DERControlEnum = "HFMayTrip"
	DERControlEnumHVMustTrip              DERControlEnum = "HVMustTrip"
	DERControlEnumHVMomCess               DERControlEnum = "HVMomCess"
	DERControlEnumHVMayTrip               DERControlEnum = "HVMayTrip"
	DERControlEnumLimitMaxDischarge       DERControlEnum = "LimitMaxDischarge"
	DERControlEnumLFMustTrip              DERControlEnum = "LFMustTrip"
	DERControlEnumLVMustTrip              DERControlEnum = "LVMustTrip"
	DERControlEnumLVMomCess               DERControlEnum = "LVMomCess"
	DERControlEnumLVMayTrip               DERControlEnum = "LVMayTrip"
	DERControlEnumPowerMonitoringMustTrip DERControlEnum = "PowerMonitoringMustTrip"
	DERControlEnumVoltVAR                 DERControlEnum = "VoltVar"
	DERControlEnumVoltWatt                DERControlEnum = "VoltWatt"
	DERControlEnumWattPF                  DERControlEnum = "WattPF"
	DERControlEnumWattVAR                 DERControlEnum = "WattVar"
)

type DERControlStatusEnum string

const (
	DERControlStatusEnumAccepted     DERControlStatusEnum = "Accepted"
	DERControlStatusEnumRejected     DERControlStatusEnum = "Rejected"
	DERControlStatusEnumNotSupported DERControlStatusEnum = "NotSupported"
	DERControlStatusEnumNotFound     DERControlStatusEnum = "NotFound"
)

type DERUnitEnum string

const (
	DERUnitEnumNotApplicable DERUnitEnum = "Not_Applicable"
	DERUnitEnumPctMaxW       DERUnitEnum = "PctMaxW"
	DERUnitEnumPctMaxVAR     DERUnitEnum = "PctMaxVar"
	DERUnitEnumPctWAvail     DERUnitEnum = "PctWAvail"
	DERUnitEnumPctVARAvail   DERUnitEnum = "PctVarAvail"
	DERUnitEnumPctEffectiveV DERUnitEnum = "PctEffectiveV"
)

type DataEnum string

const (
	DataEnumString       DataEnum = "string"
	DataEnumDecimal      DataEnum = "decimal"
	DataEnumInteger      DataEnum = "integer"
	DataEnumDateTime     DataEnum = "dateTime"
	DataEnumBoolean      DataEnum = "boolean"
	DataEnumOptionList   DataEnum = "OptionList"
	DataEnumSequenceList DataEnum = "SequenceList"
	DataEnumMemberList   DataEnum = "MemberList"
)

type DataTransferStatusEnum string

const (
	DataTransferStatusEnumAccepted         DataTransferStatusEnum = "Accepted"
	DataTransferStatusEnumRejected         DataTransferStatusEnum = "Rejected"
	DataTransferStatusEnumUnknownMessageID DataTransferStatusEnum = "UnknownMessageId"
	DataTransferStatusEnumUnknownVendorID  DataTransferStatusEnum = "UnknownVendorId"
)

type DayOfWeekEnum string

const (
	DayOfWeekEnumMonday    DayOfWeekEnum = "Monday"
	DayOfWeekEnumTuesday   DayOfWeekEnum = "Tuesday"
	DayOfWeekEnumWednesday DayOfWeekEnum = "Wednesday"
	DayOfWeekEnumThursday  DayOfWeekEnum = "Thursday"
	DayOfWeekEnumFriday    DayOfWeekEnum = "Friday"
	DayOfWeekEnumSaturday  DayOfWeekEnum = "Saturday"
	DayOfWeekEnumSunday    DayOfWeekEnum = "Sunday"
)

type DeleteCertificateStatusEnum string

const (
	DeleteCertificateStatusEnumAccepted DeleteCertificateStatusEnum = "Accepted"
	DeleteCertificateStatusEnumFailed   DeleteCertificateStatusEnum = "Failed"
	DeleteCertificateStatusEnumNotFound DeleteCertificateStatusEnum = "NotFound"
)

type DisplayMessageStatusEnum string

const (
	DisplayMessageStatusEnumAccepted                  DisplayMessageStatusEnum = "Accepted"
	DisplayMessageStatusEnumNotSupportedMessageFormat DisplayMessageStatusEnum = "NotSupportedMessageFormat"
	DisplayMessageStatusEnumRejected                  DisplayMessageStatusEnum = "Rejected"
	DisplayMessageStatusEnumNotSupportedPriority      DisplayMessageStatusEnum = "NotSupportedPriority"
	DisplayMessageStatusEnumNotSupportedState         DisplayMessageStatusEnum = "NotSupportedState"
	DisplayMessageStatusEnumUnknownTransaction        DisplayMessageStatusEnum = "UnknownTransaction"
	DisplayMessageStatusEnumLanguageNotSupported      DisplayMessageStatusEnum = "LanguageNotSupported"
)

type EVSEKindEnum string

const (
	EVSEKindEnumAC EVSEKindEnum = "AC"
	EVSEKindEnumDC EVSEKindEnum = "DC"
)

type EnergyTransferModeEnum string

const (
	EnergyTransferModeEnumACSinglePhase EnergyTransferModeEnum = "AC_single_phase"
	EnergyTransferModeEnumACTwoPhase    EnergyTransferModeEnum = "AC_two_phase"
	EnergyTransferModeEnumACThreePhase  EnergyTransferModeEnum = "AC_three_phase"
	EnergyTransferModeEnumDC            EnergyTransferModeEnum = "DC"
	EnergyTransferModeEnumACBPT         EnergyTransferModeEnum = "AC_BPT"
	EnergyTransferModeEnumACBPTDER      EnergyTransferModeEnum = "AC_BPT_DER"
	EnergyTransferModeEnumACDER         EnergyTransferModeEnum = "AC_DER"
	EnergyTransferModeEnumDCBPT         EnergyTransferModeEnum = "DC_BPT"
	EnergyTransferModeEnumDCACDP        EnergyTransferModeEnum = "DC_ACDP"
	EnergyTransferModeEnumDCACDPBPT     EnergyTransferModeEnum = "DC_ACDP_BPT"
	EnergyTransferModeEnumWPT           EnergyTransferModeEnum = "WPT"
)

type EventNotificationEnum string

const (
	EventNotificationEnumHardWiredNotification EventNotificationEnum = "HardWiredNotification"
	EventNotificationEnumHardWiredMonitor      EventNotificationEnum = "HardWiredMonitor"
	EventNotificationEnumPreconfiguredMonitor  EventNotificationEnum = "PreconfiguredMonitor"
	EventNotificationEnumCustomMonitor         EventNotificationEnum = "CustomMonitor"
)

type EventTriggerEnum string

const (
	EventTriggerEnumAlerting EventTriggerEnum = "Alerting"
	EventTriggerEnumDelta    EventTriggerEnum = "Delta"
	EventTriggerEnumPeriodic EventTriggerEnum = "Periodic"
)

type FirmwareStatusEnum string

const (
	FirmwareStatusEnumDownloaded                FirmwareStatusEnum = "Downloaded"
	FirmwareStatusEnumDownloadFailed            FirmwareStatusEnum = "DownloadFailed"
	FirmwareStatusEnumDownloading               FirmwareStatusEnum = "Downloading"
	FirmwareStatusEnumDownloadScheduled         FirmwareStatusEnum = "DownloadScheduled"
	FirmwareStatusEnumDownloadPaused            FirmwareStatusEnum = "DownloadPaused"
	FirmwareStatusEnumIdle                      FirmwareStatusEnum = "Idle"
	FirmwareStatusEnumInstallationFailed        FirmwareStatusEnum = "InstallationFailed"
	FirmwareStatusEnumInstalling                FirmwareStatusEnum = "Installing"
	FirmwareStatusEnumInstalled                 FirmwareStatusEnum = "Installed"
	FirmwareStatusEnumInstallRebooting          FirmwareStatusEnum = "InstallRebooting"
	FirmwareStatusEnumInstallScheduled          FirmwareStatusEnum = "InstallScheduled"
	FirmwareStatusEnumInstallVerificationFailed FirmwareStatusEnum = "InstallVerificationFailed"
	FirmwareStatusEnumInvalidSignature          FirmwareStatusEnum = "InvalidSignature"
	FirmwareStatusEnumSignatureVerified         FirmwareStatusEnum = "SignatureVerified"
)

type GenericDeviceModelStatusEnum string

const (
	GenericDeviceModelStatusEnumAccepted       GenericDeviceModelStatusEnum = "Accepted"
	GenericDeviceModelStatusEnumRejected       GenericDeviceModelStatusEnum = "Rejected"
	GenericDeviceModelStatusEnumNotSupported   GenericDeviceModelStatusEnum = "NotSupported"
	GenericDeviceModelStatusEnumEmptyResultSet GenericDeviceModelStatusEnum = "EmptyResultSet"
)

type GenericStatusEnum string

const (
	GenericStatusEnumAccepted GenericStatusEnum = "Accepted"
	GenericStatusEnumRejected GenericStatusEnum = "Rejected"
)

type GetCertificateIDUseEnum string

const (
	GetCertificateIDUseEnumV2GRootCertificate          GetCertificateIDUseEnum = "V2GRootCertificate"
	GetCertificateIDUseEnumMORootCertificate           GetCertificateIDUseEnum = "MORootCertificate"
	GetCertificateIDUseEnumCSMSRootCertificate         GetCertificateIDUseEnum = "CSMSRootCertificate"
	GetCertificateIDUseEnumV2GCertificateChain         GetCertificateIDUseEnum = "V2GCertificateChain"
	GetCertificateIDUseEnumManufacturerRootCertificate GetCertificateIDUseEnum = "ManufacturerRootCertificate"
	GetCertificateIDUseEnumOEMRootCertificate          GetCertificateIDUseEnum = "OEMRootCertificate"
)

type GetCertificateStatusEnum string

const (
	GetCertificateStatusEnumAccepted GetCertificateStatusEnum = "Accepted"
	GetCertificateStatusEnumFailed   GetCertificateStatusEnum = "Failed"
)

type GetChargingProfileStatusEnum string

const (
	GetChargingProfileStatusEnumAccepted   GetChargingProfileStatusEnum = "Accepted"
	GetChargingProfileStatusEnumNoProfiles GetChargingProfileStatusEnum = "NoProfiles"
)

type GetDisplayMessagesStatusEnum string

const (
	GetDisplayMessagesStatusEnumAccepted GetDisplayMessagesStatusEnum = "Accepted"
	GetDisplayMessagesStatusEnumUnknown  GetDisplayMessagesStatusEnum = "Unknown"
)

type GetInstalledCertificateStatusEnum string

const (
	GetInstalledCertificateStatusEnumAccepted GetInstalledCertificateStatusEnum = "Accepted"
	GetInstalledCertificateStatusEnumNotFound GetInstalledCertificateStatusEnum = "NotFound"
)

type GetVariableStatusEnum string

const (
	GetVariableStatusEnumAccepted                  GetVariableStatusEnum = "Accepted"
	GetVariableStatusEnumRejected                  GetVariableStatusEnum = "Rejected"
	GetVariableStatusEnumUnknownComponent          GetVariableStatusEnum = "UnknownComponent"
	GetVariableStatusEnumUnknownVariable           GetVariableStatusEnum = "UnknownVariable"
	GetVariableStatusEnumNotSupportedAttributeType GetVariableStatusEnum = "NotSupportedAttributeType"
)

type GridEventFaultEnum string

const (
	GridEventFaultEnumCurrentImbalance GridEventFaultEnum = "CurrentImbalance"
	GridEventFaultEnumLocalEmergency   GridEventFaultEnum = "LocalEmergency"
	GridEventFaultEnumLowInputPower    GridEventFaultEnum = "LowInputPower"
	GridEventFaultEnumOverCurrent      GridEventFaultEnum = "OverCurrent"
	GridEventFaultEnumOverFrequency    GridEventFaultEnum = "OverFrequency"
	GridEventFaultEnumOverVoltage      GridEventFaultEnum = "OverVoltage"
	GridEventFaultEnumPhaseRotation    GridEventFaultEnum = "PhaseRotation"
	GridEventFaultEnumRemoteEmergency  GridEventFaultEnum = "RemoteEmergency"
	GridEventFaultEnumUnderFrequency   GridEventFaultEnum = "UnderFrequency"
	GridEventFaultEnumUnderVoltage     GridEventFaultEnum = "UnderVoltage"
	GridEventFaultEnumVoltageImbalance GridEventFaultEnum = "VoltageImbalance"
)

type HashAlgorithmEnum string

const (
	HashAlgorithmEnumSHA256 HashAlgorithmEnum = "SHA256"
	HashAlgorithmEnumSHA384 HashAlgorithmEnum = "SHA384"
	HashAlgorithmEnumSHA512 HashAlgorithmEnum = "SHA512"
)

type InstallCertificateStatusEnum string

const (
	InstallCertificateStatusEnumAccepted InstallCertificateStatusEnum = "Accepted"
	InstallCertificateStatusEnumRejected InstallCertificateStatusEnum = "Rejected"
	InstallCertificateStatusEnumFailed   InstallCertificateStatusEnum = "Failed"
)

type InstallCertificateUseEnum string

const (
	InstallCertificateUseEnumV2GRootCertificate          InstallCertificateUseEnum = "V2GRootCertificate"
	InstallCertificateUseEnumMORootCertificate           InstallCertificateUseEnum = "MORootCertificate"
	InstallCertificateUseEnumManufacturerRootCertificate InstallCertificateUseEnum = "ManufacturerRootCertificate"
	InstallCertificateUseEnumCSMSRootCertificate         InstallCertificateUseEnum = "CSMSRootCertificate"
	InstallCertificateUseEnumOEMRootCertificate          InstallCertificateUseEnum = "OEMRootCertificate"
)

type IslandingDetectionEnum string

const (
	IslandingDetectionEnumNoAntiIslandingSupport IslandingDetectionEnum = "NoAntiIslandingSupport"
	IslandingDetectionEnumRoCoF                  IslandingDetectionEnum = "RoCoF"
	IslandingDetectionEnumUVPOVP                 IslandingDetectionEnum = "UVP_OVP"
	IslandingDetectionEnumUFPOFP                 IslandingDetectionEnum = "UFP_OFP"
	IslandingDetectionEnumVoltageVectorShift     IslandingDetectionEnum = "VoltageVectorShift"
	IslandingDetectionEnumZeroCrossingDetection  IslandingDetectionEnum = "ZeroCrossingDetection"
	IslandingDetectionEnumOtherPassive           IslandingDetectionEnum = "OtherPassive"
	IslandingDetectionEnumImpedanceMeasurement   IslandingDetectionEnum = "ImpedanceMeasurement"
	IslandingDetectionEnumImpedanceAtFrequency   IslandingDetectionEnum = "ImpedanceAtFrequency"
	IslandingDetectionEnumSlipModeFrequencyShift IslandingDetectionEnum = "SlipModeFrequencyShift"
	IslandingDetectionEnumSandiaFrequencyShift   IslandingDetectionEnum = "SandiaFrequencyShift"
	IslandingDetectionEnumSandiaVoltageShift     IslandingDetectionEnum = "SandiaVoltageShift"
	IslandingDetectionEnumFrequencyJump          IslandingDetectionEnum = "FrequencyJump"
	IslandingDetectionEnumRCLQFactor             IslandingDetectionEnum = "RCLQFactor"
	IslandingDetectionEnumOtherActive            IslandingDetectionEnum = "OtherActive"
)

type Iso15118EVCertificateStatusEnum string

const (
	Iso15118EVCertificateStatusEnumAccepted Iso15118EVCertificateStatusEnum = "Accepted"
	Iso15118EVCertificateStatusEnumFailed   Iso15118EVCertificateStatusEnum = "Failed"
)

type LocationEnum string

const (
	LocationEnumBody     LocationEnum = "Body"
	LocationEnumCable    LocationEnum = "Cable"
	LocationEnumEV       LocationEnum = "EV"
	LocationEnumInlet    LocationEnum = "Inlet"
	LocationEnumOutlet   LocationEnum = "Outlet"
	LocationEnumUpstream LocationEnum = "Upstream"
)

type LogEnum string

const (
	LogEnumDiagnosticsLog   LogEnum = "DiagnosticsLog"
	LogEnumSecurityLog      LogEnum = "SecurityLog"
	LogEnumDataCollectorLog LogEnum = "DataCollectorLog"
)

type LogStatusEnum string

const (
	LogStatusEnumAccepted         LogStatusEnum = "Accepted"
	LogStatusEnumRejected         LogStatusEnum = "Rejected"
	LogStatusEnumAcceptedCanceled LogStatusEnum = "AcceptedCanceled"
)

type MeasurandEnum string

const (
	MeasurandEnumCurrentExport                             MeasurandEnum = "Current.Export"
	MeasurandEnumCurrentExportOffered                      MeasurandEnum = "Current.Export.Offered"
	MeasurandEnumCurrentExportMinimum                      MeasurandEnum = "Current.Export.Minimum"
	MeasurandEnumCurrentImport                             MeasurandEnum = "Current.Import"
	MeasurandEnumCurrentImportOffered                      MeasurandEnum = "Current.Import.Offered"
	MeasurandEnumCurrentImportMinimum                      MeasurandEnum = "Current.Import.Minimum"
	MeasurandEnumCurrentOffered                            MeasurandEnum = "Current.Offered"
	MeasurandEnumDisplayPresentSoC                         MeasurandEnum = "Display.PresentSOC"
	MeasurandEnumDisplayMinimumSoC                         MeasurandEnum = "Display.MinimumSOC"
	MeasurandEnumDisplayTargetSoC                          MeasurandEnum = "Display.TargetSOC"
	MeasurandEnumDisplayMaximumSoC                         MeasurandEnum = "Display.MaximumSOC"
	MeasurandEnumDisplayRemainingTimeToMinimumSoC          MeasurandEnum = "Display.RemainingTimeToMinimumSOC"
	MeasurandEnumDisplayRemainingTimeToTargetSoC           MeasurandEnum = "Display.RemainingTimeToTargetSOC"
	MeasurandEnumDisplayRemainingTimeToMaximumSoC          MeasurandEnum = "Display.RemainingTimeToMaximumSOC"
	MeasurandEnumDisplayChargingComplete                   MeasurandEnum = "Display.ChargingComplete"
	MeasurandEnumDisplayBatteryEnergyCapacity              MeasurandEnum = "Display.BatteryEnergyCapacity"
	MeasurandEnumDisplayInletHot                           MeasurandEnum = "Display.InletHot"
	MeasurandEnumEnergyActiveExportInterval                MeasurandEnum = "Energy.Active.Export.Interval"
	MeasurandEnumEnergyActiveExportRegister                MeasurandEnum = "Energy.Active.Export.Register"
	MeasurandEnumEnergyActiveImportInterval                MeasurandEnum = "Energy.Active.Import.Interval"
	MeasurandEnumEnergyActiveImportRegister                MeasurandEnum = "Energy.Active.Import.Register"
	MeasurandEnumEnergyActiveImportCableLoss               MeasurandEnum = "Energy.Active.Import.CableLoss"
	MeasurandEnumEnergyActiveImportLocalGenerationRegister MeasurandEnum = "Energy.Active.Import.LocalGeneration.Register"
	MeasurandEnumEnergyActiveNet                           MeasurandEnum = "Energy.Active.Net"
	MeasurandEnumEnergyActiveSetpointInterval              MeasurandEnum = "Energy.Active.Setpoint.Interval"
	MeasurandEnumEnergyApparentExport                      MeasurandEnum = "Energy.Apparent.Export"
	MeasurandEnumEnergyApparentImport                      MeasurandEnum = "Energy.Apparent.Import"
	MeasurandEnumEnergyApparentNet                         MeasurandEnum = "Energy.Apparent.Net"
	MeasurandEnumEnergyReactiveExportInterval              MeasurandEnum = "Energy.Reactive.Export.Interval"
	MeasurandEnumEnergyReactiveExportRegister              MeasurandEnum = "Energy.Reactive.Export.Register"
	MeasurandEnumEnergyReactiveImportInterval              MeasurandEnum = "Energy.Reactive.Import.Interval"
	MeasurandEnumEnergyReactiveImportRegister              MeasurandEnum = "Energy.Reactive.Import.Register"
	MeasurandEnumEnergyReactiveNet                         MeasurandEnum = "Energy.Reactive.Net"
	MeasurandEnumEnergyRequestTarget                       MeasurandEnum = "EnergyRequest.Target"
	MeasurandEnumEnergyRequestMinimum                      MeasurandEnum = "EnergyRequest.Minimum"
	MeasurandEnumEnergyRequestMaximum                      MeasurandEnum = "EnergyRequest.Maximum"
	MeasurandEnumEnergyRequestMinimumV2X                   MeasurandEnum = "EnergyRequest.Minimum.V2X"
	MeasurandEnumEnergyRequestMaximumV2X                   MeasurandEnum = "EnergyRequest.Maximum.V2X"
	MeasurandEnumEnergyRequestBulk                         MeasurandEnum = "EnergyRequest.Bulk"
	MeasurandEnumFrequency                                 MeasurandEnum = "Frequency"
	MeasurandEnumPowerActiveExport                         MeasurandEnum = "Power.Active.Export"
	MeasurandEnumPowerActiveImport                         MeasurandEnum = "Power.Active.Import"
	MeasurandEnumPowerActiveSetpoint                       MeasurandEnum = "Power.Active.Setpoint"
	MeasurandEnumPowerActiveResidual                       MeasurandEnum = "Power.Active.Residual"
	MeasurandEnumPowerExportMinimum                        MeasurandEnum = "Power.Export.Minimum"
	MeasurandEnumPowerExportOffered                        MeasurandEnum = "Power.Export.Offered"
	MeasurandEnumPowerFactor                               MeasurandEnum = "Power.Factor"
	MeasurandEnumPowerImportOffered                        MeasurandEnum = "Power.Import.Offered"
	MeasurandEnumPowerImportMinimum                        MeasurandEnum = "Power.Import.Minimum"
	MeasurandEnumPowerOffered                              MeasurandEnum = "Power.Offered"
	MeasurandEnumPowerReactiveExport                       MeasurandEnum = "Power.Reactive.Export"
	MeasurandEnumPowerReactiveImport                       MeasurandEnum = "Power.Reactive.Import"
	MeasurandEnumSoC                                       MeasurandEnum = "SoC"
	MeasurandEnumVoltage                                   MeasurandEnum = "Voltage"
	MeasurandEnumVoltageMinimum                            MeasurandEnum = "Voltage.Minimum"
	MeasurandEnumVoltageMaximum                            MeasurandEnum = "Voltage.Maximum"
)

type MessageFormatEnum string

const (
	MessageFormatEnumASCII  MessageFormatEnum = "ASCII"
	MessageFormatEnumHTML   MessageFormatEnum = "HTML"
	MessageFormatEnumURI    MessageFormatEnum = "URI"
	MessageFormatEnumUTF8   MessageFormatEnum = "UTF8"
	MessageFormatEnumQRCODE MessageFormatEnum = "QRCODE"
)

type MessagePriorityEnum string

const (
	MessagePriorityEnumAlwaysFront MessagePriorityEnum = "AlwaysFront"
	MessagePriorityEnumInFront     MessagePriorityEnum = "InFront"
	MessagePriorityEnumNormalCycle MessagePriorityEnum = "NormalCycle"
)

type MessageStateEnum string

const (
	MessageStateEnumCharging    MessageStateEnum = "Charging"
	MessageStateEnumFaulted     MessageStateEnum = "Faulted"
	MessageStateEnumIdle        MessageStateEnum = "Idle"
	MessageStateEnumUnavailable MessageStateEnum = "Unavailable"
	MessageStateEnumSuspended   MessageStateEnum = "Suspended"
	MessageStateEnumDischarging MessageStateEnum = "Discharging"
)

type MessageTriggerEnum string

const (
	MessageTriggerEnumBootNotification                  MessageTriggerEnum = "BootNotification"
	MessageTriggerEnumLogStatusNotification             MessageTriggerEnum = "LogStatusNotification"
	MessageTriggerEnumFirmwareStatusNotification        MessageTriggerEnum = "FirmwareStatusNotification"
	MessageTriggerEnumHeartbeat                         MessageTriggerEnum = "Heartbeat"
	MessageTriggerEnumMeterValues                       MessageTriggerEnum = "MeterValues"
	MessageTriggerEnumSignChargingStationCertificate    MessageTriggerEnum = "SignChargingStationCertificate"
	MessageTriggerEnumSignV2GCertificate                MessageTriggerEnum = "SignV2GCertificate"
	MessageTriggerEnumSignV2G20Certificate              MessageTriggerEnum = "SignV2G20Certificate"
	MessageTriggerEnumStatusNotification                MessageTriggerEnum = "StatusNotification"
	MessageTriggerEnumTransactionEvent                  MessageTriggerEnum = "TransactionEvent"
	MessageTriggerEnumSignCombinedCertificate           MessageTriggerEnum = "SignCombinedCertificate"
	MessageTriggerEnumPublishFirmwareStatusNotification MessageTriggerEnum = "PublishFirmwareStatusNotification"
	MessageTriggerEnumCustomTrigger                     MessageTriggerEnum = "CustomTrigger"
)

type MobilityNeedsModeEnum string

const (
	MobilityNeedsModeEnumEVCC     MobilityNeedsModeEnum = "EVCC"
	MobilityNeedsModeEnumEVCCSECC MobilityNeedsModeEnum = "EVCC_SECC"
)

type MonitorEnum string

const (
	MonitorEnumUpperThreshold       MonitorEnum = "UpperThreshold"
	MonitorEnumLowerThreshold       MonitorEnum = "LowerThreshold"
	MonitorEnumDelta                MonitorEnum = "Delta"
	MonitorEnumPeriodic             MonitorEnum = "Periodic"
	MonitorEnumPeriodicClockAligned MonitorEnum = "PeriodicClockAligned"
	MonitorEnumTargetDelta          MonitorEnum = "TargetDelta"
	MonitorEnumTargetDeltaRelative  MonitorEnum = "TargetDeltaRelative"
)

type MonitoringBaseEnum string

const (
	MonitoringBaseEnumAll            MonitoringBaseEnum = "All"
	MonitoringBaseEnumFactoryDefault MonitoringBaseEnum = "FactoryDefault"
	MonitoringBaseEnumHardWiredOnly  MonitoringBaseEnum = "HardWiredOnly"
)

type MonitoringCriterionEnum string

const (
	MonitoringCriterionEnumThresholdMonitoring MonitoringCriterionEnum = "ThresholdMonitoring"
	MonitoringCriterionEnumDeltaMonitoring     MonitoringCriterionEnum = "DeltaMonitoring"
	MonitoringCriterionEnumPeriodicMonitoring  MonitoringCriterionEnum = "PeriodicMonitoring"
)

type MutabilityEnum string

const (
	MutabilityEnumReadOnly  MutabilityEnum = "ReadOnly"
	MutabilityEnumWriteOnly MutabilityEnum = "WriteOnly"
	MutabilityEnumReadWrite MutabilityEnum = "ReadWrite"
)

type NotifyAllowedEnergyTransferStatusEnum string

const (
	NotifyAllowedEnergyTransferStatusEnumAccepted NotifyAllowedEnergyTransferStatusEnum = "Accepted"
	NotifyAllowedEnergyTransferStatusEnumRejected NotifyAllowedEnergyTransferStatusEnum = "Rejected"
)

type NotifyEVChargingNeedsStatusEnum string

const (
	NotifyEVChargingNeedsStatusEnumAccepted          NotifyEVChargingNeedsStatusEnum = "Accepted"
	NotifyEVChargingNeedsStatusEnumRejected          NotifyEVChargingNeedsStatusEnum = "Rejected"
	NotifyEVChargingNeedsStatusEnumProcessing        NotifyEVChargingNeedsStatusEnum = "Processing"
	NotifyEVChargingNeedsStatusEnumNoChargingProfile NotifyEVChargingNeedsStatusEnum = "NoChargingProfile"
)

type OCPPInterfaceEnum string

const (
	OCPPInterfaceEnumWired0    OCPPInterfaceEnum = "Wired0"
	OCPPInterfaceEnumWired1    OCPPInterfaceEnum = "Wired1"
	OCPPInterfaceEnumWired2    OCPPInterfaceEnum = "Wired2"
	OCPPInterfaceEnumWired3    OCPPInterfaceEnum = "Wired3"
	OCPPInterfaceEnumWireless0 OCPPInterfaceEnum = "Wireless0"
	OCPPInterfaceEnumWireless1 OCPPInterfaceEnum = "Wireless1"
	OCPPInterfaceEnumWireless2 OCPPInterfaceEnum = "Wireless2"
	OCPPInterfaceEnumWireless3 OCPPInterfaceEnum = "Wireless3"
	OCPPInterfaceEnumAny       OCPPInterfaceEnum = "Any"
)

type OCPPTransportEnum string

const (
	OCPPTransportEnumSOAP OCPPTransportEnum = "SOAP"
	OCPPTransportEnumJSON OCPPTransportEnum = "JSON"
)

type OCPPVersionEnum string

const (
	OCPPVersionEnumOCPP12  OCPPVersionEnum = "OCPP12"
	OCPPVersionEnumOCPP15  OCPPVersionEnum = "OCPP15"
	OCPPVersionEnumOCPP16  OCPPVersionEnum = "OCPP16"
	OCPPVersionEnumOCPP20  OCPPVersionEnum = "OCPP20"
	OCPPVersionEnumOCPP201 OCPPVersionEnum = "OCPP201"
	OCPPVersionEnumOCPP21  OCPPVersionEnum = "OCPP21"
)

type OperationModeEnum string

const (
	OperationModeEnumIdle               OperationModeEnum = "Idle"
	OperationModeEnumChargingOnly       OperationModeEnum = "ChargingOnly"
	OperationModeEnumCentralSetpoint    OperationModeEnum = "CentralSetpoint"
	OperationModeEnumExternalSetpoint   OperationModeEnum = "ExternalSetpoint"
	OperationModeEnumExternalLimits     OperationModeEnum = "ExternalLimits"
	OperationModeEnumCentralFrequency   OperationModeEnum = "CentralFrequency"
	OperationModeEnumLocalFrequency     OperationModeEnum = "LocalFrequency"
	OperationModeEnumLocalLoadBalancing OperationModeEnum = "LocalLoadBalancing"
)

type OperationalStatusEnum string

const (
	OperationalStatusEnumInoperative OperationalStatusEnum = "Inoperative"
	OperationalStatusEnumOperative   OperationalStatusEnum = "Operative"
)

type PaymentStatusEnum string

const (
	PaymentStatusEnumSettled  PaymentStatusEnum = "Settled"
	PaymentStatusEnumCanceled PaymentStatusEnum = "Canceled"
	PaymentStatusEnumRejected PaymentStatusEnum = "Rejected"
	PaymentStatusEnumFailed   PaymentStatusEnum = "Failed"
)

type PhaseEnum string

const (
	PhaseEnumL1   PhaseEnum = "L1"
	PhaseEnumL2   PhaseEnum = "L2"
	PhaseEnumL3   PhaseEnum = "L3"
	PhaseEnumN    PhaseEnum = "N"
	PhaseEnumL1N  PhaseEnum = "L1-N"
	PhaseEnumL2N  PhaseEnum = "L2-N"
	PhaseEnumL3N  PhaseEnum = "L3-N"
	PhaseEnumL1L2 PhaseEnum = "L1-L2"
	PhaseEnumL2L3 PhaseEnum = "L2-L3"
	PhaseEnumL3L1 PhaseEnum = "L3-L1"
)

type PowerDuringCessationEnum string

const (
	PowerDuringCessationEnumActive   PowerDuringCessationEnum = "Active"
	PowerDuringCessationEnumReactive PowerDuringCessationEnum = "Reactive"
)

type PreconditioningStatusEnum string

const (
	PreconditioningStatusEnumUnknown         PreconditioningStatusEnum = "Unknown"
	PreconditioningStatusEnumReady           PreconditioningStatusEnum = "Ready"
	PreconditioningStatusEnumNotReady        PreconditioningStatusEnum = "NotReady"
	PreconditioningStatusEnumPreconditioning PreconditioningStatusEnum = "Preconditioning"
)

type PriorityChargingStatusEnum string

const (
	PriorityChargingStatusEnumAccepted  PriorityChargingStatusEnum = "Accepted"
	PriorityChargingStatusEnumRejected  PriorityChargingStatusEnum = "Rejected"
	PriorityChargingStatusEnumNoProfile PriorityChargingStatusEnum = "NoProfile"
)

type PublishFirmwareStatusEnum string

const (
	PublishFirmwareStatusEnumIdle              PublishFirmwareStatusEnum = "Idle"
	PublishFirmwareStatusEnumDownloadScheduled PublishFirmwareStatusEnum = "DownloadScheduled"
	PublishFirmwareStatusEnumDownloading       PublishFirmwareStatusEnum = "Downloading"
	PublishFirmwareStatusEnumDownloaded        PublishFirmwareStatusEnum = "Downloaded"
	PublishFirmwareStatusEnumPublished         PublishFirmwareStatusEnum = "Published"
	PublishFirmwareStatusEnumDownloadFailed    PublishFirmwareStatusEnum = "DownloadFailed"
	PublishFirmwareStatusEnumDownloadPaused    PublishFirmwareStatusEnum = "DownloadPaused"
	PublishFirmwareStatusEnumInvalidChecksum   PublishFirmwareStatusEnum = "InvalidChecksum"
	PublishFirmwareStatusEnumChecksumVerified  PublishFirmwareStatusEnum = "ChecksumVerified"
	PublishFirmwareStatusEnumPublishFailed     PublishFirmwareStatusEnum = "PublishFailed"
)

type ReadingContextEnum string

const (
	ReadingContextEnumInterruptionBegin ReadingContextEnum = "Interruption.Begin"
	ReadingContextEnumInterruptionEnd   ReadingContextEnum = "Interruption.End"
	ReadingContextEnumOther             ReadingContextEnum = "Other"
	ReadingContextEnumSampleClock       ReadingContextEnum = "Sample.Clock"
	ReadingContextEnumSamplePeriodic    ReadingContextEnum = "Sample.Periodic"
	ReadingContextEnumTransactionBegin  ReadingContextEnum = "Transaction.Begin"
	ReadingContextEnumTransactionEnd    ReadingContextEnum = "Transaction.End"
	ReadingContextEnumTrigger           ReadingContextEnum = "Trigger"
)

type ReasonEnum string

const (
	ReasonEnumDeAuthorized              ReasonEnum = "DeAuthorized"
	ReasonEnumEmergencyStop             ReasonEnum = "EmergencyStop"
	ReasonEnumEnergyLimitReached        ReasonEnum = "EnergyLimitReached"
	ReasonEnumEVDisconnected            ReasonEnum = "EVDisconnected"
	ReasonEnumGroundFault               ReasonEnum = "GroundFault"
	ReasonEnumImmediateReset            ReasonEnum = "ImmediateReset"
	ReasonEnumMasterPass                ReasonEnum = "MasterPass"
	ReasonEnumLocal                     ReasonEnum = "Local"
	ReasonEnumLocalOutOfCredit          ReasonEnum = "LocalOutOfCredit"
	ReasonEnumOther                     ReasonEnum = "Other"
	ReasonEnumOvercurrentFault          ReasonEnum = "OvercurrentFault"
	ReasonEnumPowerLoss                 ReasonEnum = "PowerLoss"
	ReasonEnumPowerQuality              ReasonEnum = "PowerQuality"
	ReasonEnumReboot                    ReasonEnum = "Reboot"
	ReasonEnumRemote                    ReasonEnum = "Remote"
	ReasonEnumSoCLimitReached           ReasonEnum = "SOCLimitReached"
	ReasonEnumStoppedByEV               ReasonEnum = "StoppedByEV"
	ReasonEnumTimeLimitReached          ReasonEnum = "TimeLimitReached"
	ReasonEnumTimeout                   ReasonEnum = "Timeout"
	ReasonEnumReqEnergyTransferRejected ReasonEnum = "ReqEnergyTransferRejected"
)

type RecurrencyKindEnum string

const (
	RecurrencyKindEnumDaily  RecurrencyKindEnum = "Daily"
	RecurrencyKindEnumWeekly RecurrencyKindEnum = "Weekly"
)

type RegistrationStatusEnum string

const (
	RegistrationStatusEnumAccepted RegistrationStatusEnum = "Accepted"
	RegistrationStatusEnumPending  RegistrationStatusEnum = "Pending"
	RegistrationStatusEnumRejected RegistrationStatusEnum = "Rejected"
)

type ReportBaseEnum string

const (
	ReportBaseEnumConfigurationInventory ReportBaseEnum = "ConfigurationInventory"
	ReportBaseEnumFullInventory          ReportBaseEnum = "FullInventory"
	ReportBaseEnumSummaryInventory       ReportBaseEnum = "SummaryInventory"
)

type RequestStartStopStatusEnum string

const (
	RequestStartStopStatusEnumAccepted RequestStartStopStatusEnum = "Accepted"
	RequestStartStopStatusEnumRejected RequestStartStopStatusEnum = "Rejected"
)

type ReservationUpdateStatusEnum string

const (
	ReservationUpdateStatusEnumExpired       ReservationUpdateStatusEnum = "Expired"
	ReservationUpdateStatusEnumRemoved       ReservationUpdateStatusEnum = "Removed"
	ReservationUpdateStatusEnumNoTransaction ReservationUpdateStatusEnum = "NoTransaction"
)

type ReserveNowStatusEnum string

const (
	ReserveNowStatusEnumAccepted    ReserveNowStatusEnum = "Accepted"
	ReserveNowStatusEnumFaulted     ReserveNowStatusEnum = "Faulted"
	ReserveNowStatusEnumOccupied    ReserveNowStatusEnum = "Occupied"
	ReserveNowStatusEnumRejected    ReserveNowStatusEnum = "Rejected"
	ReserveNowStatusEnumUnavailable ReserveNowStatusEnum = "Unavailable"
)

type ResetEnum string

const (
	ResetEnumImmediate          ResetEnum = "Immediate"
	ResetEnumOnIdle             ResetEnum = "OnIdle"
	ResetEnumImmediateAndResume ResetEnum = "ImmediateAndResume"
)

type ResetStatusEnum string

const (
	ResetStatusEnumAccepted  ResetStatusEnum = "Accepted"
	ResetStatusEnumRejected  ResetStatusEnum = "Rejected"
	ResetStatusEnumScheduled ResetStatusEnum = "Scheduled"
)

type SendLocalListStatusEnum string

const (
	SendLocalListStatusEnumAccepted        SendLocalListStatusEnum = "Accepted"
	SendLocalListStatusEnumFailed          SendLocalListStatusEnum = "Failed"
	SendLocalListStatusEnumVersionMismatch SendLocalListStatusEnum = "VersionMismatch"
)

type SetMonitoringStatusEnum string

const (
	SetMonitoringStatusEnumAccepted               SetMonitoringStatusEnum = "Accepted"
	SetMonitoringStatusEnumUnknownComponent       SetMonitoringStatusEnum = "UnknownComponent"
	SetMonitoringStatusEnumUnknownVariable        SetMonitoringStatusEnum = "UnknownVariable"
	SetMonitoringStatusEnumUnsupportedMonitorType SetMonitoringStatusEnum = "UnsupportedMonitorType"
	SetMonitoringStatusEnumRejected               SetMonitoringStatusEnum = "Rejected"
	SetMonitoringStatusEnumDuplicate              SetMonitoringStatusEnum = "Duplicate"
)

type SetNetworkProfileStatusEnum string

const (
	SetNetworkProfileStatusEnumAccepted SetNetworkProfileStatusEnum = "Accepted"
	SetNetworkProfileStatusEnumRejected SetNetworkProfileStatusEnum = "Rejected"
	SetNetworkProfileStatusEnumFailed   SetNetworkProfileStatusEnum = "Failed"
)

type SetVariableStatusEnum string

const (
	SetVariableStatusEnumAccepted                  SetVariableStatusEnum = "Accepted"
	SetVariableStatusEnumRejected                  SetVariableStatusEnum = "Rejected"
	SetVariableStatusEnumUnknownComponent          SetVariableStatusEnum = "UnknownComponent"
	SetVariableStatusEnumUnknownVariable           SetVariableStatusEnum = "UnknownVariable"
	SetVariableStatusEnumNotSupportedAttributeType SetVariableStatusEnum = "NotSupportedAttributeType"
	SetVariableStatusEnumRebootRequired            SetVariableStatusEnum = "RebootRequired"
)

type TariffChangeStatusEnum string

const (
	TariffChangeStatusEnumAccepted              TariffChangeStatusEnum = "Accepted"
	TariffChangeStatusEnumRejected              TariffChangeStatusEnum = "Rejected"
	TariffChangeStatusEnumTooManyElements       TariffChangeStatusEnum = "TooManyElements"
	TariffChangeStatusEnumConditionNotSupported TariffChangeStatusEnum = "ConditionNotSupported"
	TariffChangeStatusEnumTxNotFound            TariffChangeStatusEnum = "TxNotFound"
	TariffChangeStatusEnumNoCurrencyChange      TariffChangeStatusEnum = "NoCurrencyChange"
)

type TariffClearStatusEnum string

const (
	TariffClearStatusEnumAccepted TariffClearStatusEnum = "Accepted"
	TariffClearStatusEnumRejected TariffClearStatusEnum = "Rejected"
	TariffClearStatusEnumNoTariff TariffClearStatusEnum = "NoTariff"
)

type TariffCostEnum string

const (
	TariffCostEnumNormalCost TariffCostEnum = "NormalCost"
	TariffCostEnumMinCost    TariffCostEnum = "MinCost"
	TariffCostEnumMaxCost    TariffCostEnum = "MaxCost"
)

type TariffGetStatusEnum string

const (
	TariffGetStatusEnumAccepted TariffGetStatusEnum = "Accepted"
	TariffGetStatusEnumRejected TariffGetStatusEnum = "Rejected"
	TariffGetStatusEnumNoTariff TariffGetStatusEnum = "NoTariff"
)

type TariffKindEnum string

const (
	TariffKindEnumDefaultTariff TariffKindEnum = "DefaultTariff"
	TariffKindEnumDriverTariff  TariffKindEnum = "DriverTariff"
)

type TariffSetStatusEnum string

const (
	TariffSetStatusEnumAccepted              TariffSetStatusEnum = "Accepted"
	TariffSetStatusEnumRejected              TariffSetStatusEnum = "Rejected"
	TariffSetStatusEnumTooManyElements       TariffSetStatusEnum = "TooManyElements"
	TariffSetStatusEnumConditionNotSupported TariffSetStatusEnum = "ConditionNotSupported"
	TariffSetStatusEnumDuplicateTariffID     TariffSetStatusEnum = "DuplicateTariffId"
)

type TransactionEventEnum string

const (
	TransactionEventEnumEnded   TransactionEventEnum = "Ended"
	TransactionEventEnumStarted TransactionEventEnum = "Started"
	TransactionEventEnumUpdated TransactionEventEnum = "Updated"
)

type TriggerMessageStatusEnum string

const (
	TriggerMessageStatusEnumAccepted       TriggerMessageStatusEnum = "Accepted"
	TriggerMessageStatusEnumRejected       TriggerMessageStatusEnum = "Rejected"
	TriggerMessageStatusEnumNotImplemented TriggerMessageStatusEnum = "NotImplemented"
)

type TriggerReasonEnum string

const (
	TriggerReasonEnumAbnormalCondition    TriggerReasonEnum = "AbnormalCondition"
	TriggerReasonEnumAuthorized           TriggerReasonEnum = "Authorized"
	TriggerReasonEnumCablePluggedIn       TriggerReasonEnum = "CablePluggedIn"
	TriggerReasonEnumChargingRateChanged  TriggerReasonEnum = "ChargingRateChanged"
	TriggerReasonEnumChargingStateChanged TriggerReasonEnum = "ChargingStateChanged"
	TriggerReasonEnumCostLimitReached     TriggerReasonEnum = "CostLimitReached"
	TriggerReasonEnumDeauthorized         TriggerReasonEnum = "Deauthorized"
	TriggerReasonEnumEnergyLimitReached   TriggerReasonEnum = "EnergyLimitReached"
	TriggerReasonEnumEVCommunicationLost  TriggerReasonEnum = "EVCommunicationLost"
	TriggerReasonEnumEVConnectTimeout     TriggerReasonEnum = "EVConnectTimeout"
	TriggerReasonEnumEVDeparted           TriggerReasonEnum = "EVDeparted"
	TriggerReasonEnumEVDetected           TriggerReasonEnum = "EVDetected"
	TriggerReasonEnumLimitSet             TriggerReasonEnum = "LimitSet"
	TriggerReasonEnumMeterValueClock      TriggerReasonEnum = "MeterValueClock"
	TriggerReasonEnumMeterValuePeriodic   TriggerReasonEnum = "MeterValuePeriodic"
	TriggerReasonEnumOperationModeChanged TriggerReasonEnum = "OperationModeChanged"
	TriggerReasonEnumRemoteStart          TriggerReasonEnum = "RemoteStart"
	TriggerReasonEnumRemoteStop           TriggerReasonEnum = "RemoteStop"
	TriggerReasonEnumResetCommand         TriggerReasonEnum = "ResetCommand"
	TriggerReasonEnumRunningCost          TriggerReasonEnum = "RunningCost"
	TriggerReasonEnumSignedDataReceived   TriggerReasonEnum = "SignedDataReceived"
	TriggerReasonEnumSoCLimitReached      TriggerReasonEnum = "SoCLimitReached"
	TriggerReasonEnumStopAuthorized       TriggerReasonEnum = "StopAuthorized"
	TriggerReasonEnumTariffChanged        TriggerReasonEnum = "TariffChanged"
	TriggerReasonEnumTariffNotAccepted    TriggerReasonEnum = "TariffNotAccepted"
	TriggerReasonEnumTimeLimitReached     TriggerReasonEnum = "TimeLimitReached"
	TriggerReasonEnumTrigger              TriggerReasonEnum = "Trigger"
	TriggerReasonEnumTxResumed            TriggerReasonEnum = "TxResumed"
	TriggerReasonEnumUnlockCommand        TriggerReasonEnum = "UnlockCommand"
)

type UnlockStatusEnum string

const (
	UnlockStatusEnumUnlocked                     UnlockStatusEnum = "Unlocked"
	UnlockStatusEnumUnlockFailed                 UnlockStatusEnum = "UnlockFailed"
	UnlockStatusEnumOngoingAuthorizedTransaction UnlockStatusEnum = "OngoingAuthorizedTransaction"
	UnlockStatusEnumUnknownConnector             UnlockStatusEnum = "UnknownConnector"
)

type UnpublishFirmwareStatusEnum string

const (
	UnpublishFirmwareStatusEnumDownloadOngoing UnpublishFirmwareStatusEnum = "DownloadOngoing"
	UnpublishFirmwareStatusEnumNoFirmware      UnpublishFirmwareStatusEnum = "NoFirmware"
	UnpublishFirmwareStatusEnumUnpublished     UnpublishFirmwareStatusEnum = "Unpublished"
)

type UpdateEnum string

const (
	UpdateEnumDifferential UpdateEnum = "Differential"
	UpdateEnumFull         UpdateEnum = "Full"
)

type UpdateFirmwareStatusEnum string

const (
	UpdateFirmwareStatusEnumAccepted           UpdateFirmwareStatusEnum = "Accepted"
	UpdateFirmwareStatusEnumRejected           UpdateFirmwareStatusEnum = "Rejected"
	UpdateFirmwareStatusEnumAcceptedCanceled   UpdateFirmwareStatusEnum = "AcceptedCanceled"
	UpdateFirmwareStatusEnumInvalidCertificate UpdateFirmwareStatusEnum = "InvalidCertificate"
	UpdateFirmwareStatusEnumRevokedCertificate UpdateFirmwareStatusEnum = "RevokedCertificate"
)

type UploadLogStatusEnum string

const (
	UploadLogStatusEnumBadMessage            UploadLogStatusEnum = "BadMessage"
	UploadLogStatusEnumIdle                  UploadLogStatusEnum = "Idle"
	UploadLogStatusEnumNotSupportedOperation UploadLogStatusEnum = "NotSupportedOperation"
	UploadLogStatusEnumPermissionDenied      UploadLogStatusEnum = "PermissionDenied"
	UploadLogStatusEnumUploaded              UploadLogStatusEnum = "Uploaded"
	UploadLogStatusEnumUploadFailure         UploadLogStatusEnum = "UploadFailure"
	UploadLogStatusEnumUploading             UploadLogStatusEnum = "Uploading"
	UploadLogStatusEnumAcceptedCanceled      UploadLogStatusEnum = "AcceptedCanceled"
)

type VPNEnum string

const (
	VPNEnumIKEv2 VPNEnum = "IKEv2"
	VPNEnumIPSec VPNEnum = "IPSec"
	VPNEnumL2TP  VPNEnum = "L2TP"
	VPNEnumPPTP  VPNEnum = "PPTP"
)
