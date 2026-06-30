// Package ocpp21 contains the data structures for OCPP 2.1 messages.
package ocpp21

import "time"

type AFRRSignalRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	Signal     int        `json:"signal" validate:"required"`
	Timestamp  time.Time  `json:"timestamp" validate:"required"`
}

type AdjustPeriodicEventStreamRequest struct {
	CustomData CustomData                `json:"customData,omitempty"`
	ID         int                       `json:"id" validate:"required,gte=0.0"`
	Params     PeriodicEventStreamParams `json:"params" validate:"required"`
}

type AuthorizeRequest struct {
	Certificate                 *string           `json:"certificate,omitempty" validate:"max=10000"`
	CustomData                  CustomData        `json:"customData,omitempty"`
	IDToken                     IDToken           `json:"idToken" validate:"required"`
	Iso15118CertificateHashData []OCSPRequestData `json:"iso15118CertificateHashData,omitempty" validate:"min=1,max=4"`
}

type BatterySwapRequest struct {
	BatteryData []BatteryData        `json:"batteryData" validate:"required,min=1"`
	CustomData  CustomData           `json:"customData,omitempty"`
	EventType   BatterySwapEventEnum `json:"eventType" validate:"required"`
	IDToken     IDToken              `json:"idToken" validate:"required"`
	RequestID   int                  `json:"requestId" validate:"required"`
}

type BootNotificationRequest struct {
	ChargingStation ChargingStation `json:"chargingStation" validate:"required"`
	CustomData      CustomData      `json:"customData,omitempty"`
	Reason          BootReasonEnum  `json:"reason" validate:"required"`
}

type CancelReservationRequest struct {
	CustomData    CustomData `json:"customData,omitempty"`
	ReservationID int        `json:"reservationId" validate:"required,gte=0.0"`
}

type CertificateSignedRequest struct {
	CertificateChain string                     `json:"certificateChain" validate:"required,max=10000"`
	CertificateType  *CertificateSigningUseEnum `json:"certificateType,omitempty"`
	CustomData       CustomData                 `json:"customData,omitempty"`
	RequestID        *int                       `json:"requestId,omitempty"`
}

type ChangeAvailabilityRequest struct {
	CustomData        CustomData            `json:"customData,omitempty"`
	EVSE              *EVSE                 `json:"evse,omitempty"`
	OperationalStatus OperationalStatusEnum `json:"operationalStatus" validate:"required"`
}

type ChangeTransactionTariffRequest struct {
	CustomData    CustomData `json:"customData,omitempty"`
	Tariff        Tariff     `json:"tariff" validate:"required"`
	TransactionID string     `json:"transactionId" validate:"required,max=36"`
}

type ClearCacheRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type ClearChargingProfileRequest struct {
	ChargingProfileCriteria *ClearChargingProfile `json:"chargingProfileCriteria,omitempty"`
	ChargingProfileID       *int                  `json:"chargingProfileId,omitempty"`
	CustomData              CustomData            `json:"customData,omitempty"`
}

type ClearDERControlRequest struct {
	ControlID   *string         `json:"controlId,omitempty" validate:"max=36"`
	ControlType *DERControlEnum `json:"controlType,omitempty"`
	CustomData  CustomData      `json:"customData,omitempty"`
	IsDefault   bool            `json:"isDefault" validate:"required"`
}

type ClearDisplayMessageRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	ID         int        `json:"id" validate:"required,gte=0.0"`
}

type ClearTariffsRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	EVSEID     *int       `json:"evseId,omitempty" validate:"gte=0.0"`
	TariffIds  []string   `json:"tariffIds,omitempty" validate:"min=1"`
}

type ClearVariableMonitoringRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	ID         []int      `json:"id" validate:"required,min=1"`
}

type ClearedChargingLimitRequest struct {
	ChargingLimitSource string     `json:"chargingLimitSource" validate:"required,max=20"`
	CustomData          CustomData `json:"customData,omitempty"`
	EVSEID              *int       `json:"evseId,omitempty" validate:"gte=0.0"`
}

type ClosePeriodicEventStreamRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	ID         int        `json:"id" validate:"required,gte=0.0"`
}

type CostUpdatedRequest struct {
	CustomData    CustomData `json:"customData,omitempty"`
	TotalCost     float64    `json:"totalCost" validate:"required"`
	TransactionID string     `json:"transactionId" validate:"required,max=36"`
}

type CustomerInformationRequest struct {
	Clear               bool                 `json:"clear" validate:"required"`
	CustomData          CustomData           `json:"customData,omitempty"`
	CustomerCertificate *CertificateHashData `json:"customerCertificate,omitempty"`
	CustomerIdentifier  *string              `json:"customerIdentifier,omitempty" validate:"max=64"`
	IDToken             *IDToken             `json:"idToken,omitempty"`
	Report              bool                 `json:"report" validate:"required"`
	RequestID           int                  `json:"requestId" validate:"required,gte=0.0"`
}

type DataTransferRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	Data       any        `json:"data,omitempty"`
	MessageID  *string    `json:"messageId,omitempty" validate:"max=50"`
	VendorID   string     `json:"vendorId" validate:"required,max=255"`
}

type DeleteCertificateRequest struct {
	CertificateHashData CertificateHashData `json:"certificateHashData" validate:"required"`
	CustomData          CustomData          `json:"customData,omitempty"`
}

type FirmwareStatusNotificationRequest struct {
	CustomData CustomData         `json:"customData,omitempty"`
	RequestID  *int               `json:"requestId,omitempty"`
	Status     FirmwareStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo        `json:"statusInfo,omitempty"`
}

type Get15118EVCertificateRequest struct {
	Action                           CertificateActionEnum `json:"action" validate:"required"`
	CustomData                       CustomData            `json:"customData,omitempty"`
	ExiRequest                       string                `json:"exiRequest" validate:"required,max=11000"`
	Iso15118SchemaVersion            string                `json:"iso15118SchemaVersion" validate:"required,max=50"`
	MaximumContractCertificateChains *int                  `json:"maximumContractCertificateChains,omitempty" validate:"gte=0.0"`
	PrioritizedEMAIDs                []string              `json:"prioritizedEMAIDs,omitempty" validate:"min=1,max=8"`
}

type GetBaseReportRequest struct {
	CustomData CustomData     `json:"customData,omitempty"`
	ReportBase ReportBaseEnum `json:"reportBase" validate:"required"`
	RequestID  int            `json:"requestId" validate:"required"`
}

type GetCertificateChainStatusRequest struct {
	CertificateStatusRequests []CertificateStatusRequestInfo `json:"certificateStatusRequests" validate:"required,min=1,max=4"`
	CustomData                CustomData                     `json:"customData,omitempty"`
}

type GetCertificateStatusRequest struct {
	CustomData      CustomData      `json:"customData,omitempty"`
	OCSPRequestData OCSPRequestData `json:"ocspRequestData" validate:"required"`
}

type GetChargingProfilesRequest struct {
	ChargingProfile ChargingProfileCriterion `json:"chargingProfile" validate:"required"`
	CustomData      CustomData               `json:"customData,omitempty"`
	EVSEID          *int                     `json:"evseId,omitempty" validate:"gte=0.0"`
	RequestID       int                      `json:"requestId" validate:"required"`
}

type GetCompositeScheduleRequest struct {
	ChargingRateUnit *ChargingRateUnitEnum `json:"chargingRateUnit,omitempty"`
	CustomData       CustomData            `json:"customData,omitempty"`
	Duration         int                   `json:"duration" validate:"required"`
	EVSEID           int                   `json:"evseId" validate:"required,gte=0.0"`
}

type GetDERControlRequest struct {
	ControlID   *string         `json:"controlId,omitempty" validate:"max=36"`
	ControlType *DERControlEnum `json:"controlType,omitempty"`
	CustomData  CustomData      `json:"customData,omitempty"`
	IsDefault   *bool           `json:"isDefault,omitempty"`
	RequestID   int             `json:"requestId" validate:"required"`
}

type GetDisplayMessagesRequest struct {
	CustomData CustomData           `json:"customData,omitempty"`
	ID         []int                `json:"id,omitempty" validate:"min=1"`
	Priority   *MessagePriorityEnum `json:"priority,omitempty"`
	RequestID  int                  `json:"requestId" validate:"required"`
	State      *MessageStateEnum    `json:"state,omitempty"`
}

type GetInstalledCertificateIdsRequest struct {
	CertificateType []GetCertificateIDUseEnum `json:"certificateType,omitempty" validate:"min=1"`
	CustomData      CustomData                `json:"customData,omitempty"`
}

type GetLocalListVersionRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type GetLogRequest struct {
	CustomData    CustomData    `json:"customData,omitempty"`
	Log           LogParameters `json:"log" validate:"required"`
	LogType       LogEnum       `json:"logType" validate:"required"`
	RequestID     int           `json:"requestId" validate:"required"`
	Retries       *int          `json:"retries,omitempty" validate:"gte=0.0"`
	RetryInterval *int          `json:"retryInterval,omitempty"`
}

type GetMonitoringReportRequest struct {
	ComponentVariable  []ComponentVariable       `json:"componentVariable,omitempty" validate:"min=1"`
	CustomData         CustomData                `json:"customData,omitempty"`
	MonitoringCriteria []MonitoringCriterionEnum `json:"monitoringCriteria,omitempty" validate:"min=1,max=3"`
	RequestID          int                       `json:"requestId" validate:"required"`
}

type GetPeriodicEventStreamRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type GetReportRequest struct {
	ComponentCriteria []ComponentCriterionEnum `json:"componentCriteria,omitempty" validate:"min=1,max=4"`
	ComponentVariable []ComponentVariable      `json:"componentVariable,omitempty" validate:"min=1"`
	CustomData        CustomData               `json:"customData,omitempty"`
	RequestID         int                      `json:"requestId" validate:"required"`
}

type GetTariffsRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	EVSEID     int        `json:"evseId" validate:"required,gte=0.0"`
}

type GetTransactionStatusRequest struct {
	CustomData    CustomData `json:"customData,omitempty"`
	TransactionID *string    `json:"transactionId,omitempty" validate:"max=36"`
}

type GetVariablesRequest struct {
	CustomData      CustomData        `json:"customData,omitempty"`
	GetVariableData []GetVariableData `json:"getVariableData" validate:"required,min=1"`
}

type HeartbeatRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type InstallCertificateRequest struct {
	Certificate     string                    `json:"certificate" validate:"required,max=10000"`
	CertificateType InstallCertificateUseEnum `json:"certificateType" validate:"required"`
	CustomData      CustomData                `json:"customData,omitempty"`
}

type LogStatusNotificationRequest struct {
	CustomData CustomData          `json:"customData,omitempty"`
	RequestID  *int                `json:"requestId,omitempty"`
	Status     UploadLogStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo         `json:"statusInfo,omitempty"`
}

type MeterValuesRequest struct {
	CustomData CustomData   `json:"customData,omitempty"`
	EVSEID     int          `json:"evseId" validate:"required,gte=0.0"`
	MeterValue []MeterValue `json:"meterValue" validate:"required,min=1"`
}

type NotifyAllowedEnergyTransferRequest struct {
	AllowedEnergyTransfer []EnergyTransferModeEnum `json:"allowedEnergyTransfer" validate:"required,min=1"`
	CustomData            CustomData               `json:"customData,omitempty"`
	TransactionID         string                   `json:"transactionId" validate:"required,max=36"`
}

type NotifyChargingLimitRequest struct {
	ChargingLimit    ChargingLimit      `json:"chargingLimit" validate:"required"`
	ChargingSchedule []ChargingSchedule `json:"chargingSchedule,omitempty" validate:"min=1"`
	CustomData       CustomData         `json:"customData,omitempty"`
	EVSEID           *int               `json:"evseId,omitempty" validate:"gte=0.0"`
}

type NotifyCustomerInformationRequest struct {
	CustomData  CustomData `json:"customData,omitempty"`
	Data        string     `json:"data" validate:"required,max=512"`
	GeneratedAt time.Time  `json:"generatedAt" validate:"required"`
	RequestID   int        `json:"requestId" validate:"required,gte=0.0"`
	SeqNo       int        `json:"seqNo" validate:"required,gte=0.0"`
	Tbc         *bool      `json:"tbc,omitempty"`
}

type NotifyDERAlarmRequest struct {
	AlarmEnded     *bool               `json:"alarmEnded,omitempty"`
	ControlType    DERControlEnum      `json:"controlType" validate:"required"`
	CustomData     CustomData          `json:"customData,omitempty"`
	ExtraInfo      *string             `json:"extraInfo,omitempty" validate:"max=200"`
	GridEventFault *GridEventFaultEnum `json:"gridEventFault,omitempty"`
	Timestamp      time.Time           `json:"timestamp" validate:"required"`
}

type NotifyDERStartStopRequest struct {
	ControlID     string     `json:"controlId" validate:"required,max=36"`
	CustomData    CustomData `json:"customData,omitempty"`
	Started       bool       `json:"started" validate:"required"`
	SupersededIds []string   `json:"supersededIds,omitempty" validate:"min=1,max=24"`
	Timestamp     time.Time  `json:"timestamp" validate:"required"`
}

type NotifyDisplayMessagesRequest struct {
	CustomData  CustomData    `json:"customData,omitempty"`
	MessageInfo []MessageInfo `json:"messageInfo,omitempty" validate:"min=1"`
	RequestID   int           `json:"requestId" validate:"required"`
	Tbc         *bool         `json:"tbc,omitempty"`
}

type NotifyEVChargingNeedsRequest struct {
	ChargingNeeds     ChargingNeeds `json:"chargingNeeds" validate:"required"`
	CustomData        CustomData    `json:"customData,omitempty"`
	EVSEID            int           `json:"evseId" validate:"required,gte=1.0"`
	MaxScheduleTuples *int          `json:"maxScheduleTuples,omitempty" validate:"gte=0.0"`
	Timestamp         *time.Time    `json:"timestamp,omitempty"`
}

type NotifyEVChargingScheduleRequest struct {
	ChargingSchedule           ChargingSchedule `json:"chargingSchedule" validate:"required"`
	CustomData                 CustomData       `json:"customData,omitempty"`
	EVSEID                     int              `json:"evseId" validate:"required,gte=1.0"`
	PowerToleranceAcceptance   *bool            `json:"powerToleranceAcceptance,omitempty"`
	SelectedChargingScheduleID *int             `json:"selectedChargingScheduleId,omitempty" validate:"gte=0.0"`
	TimeBase                   time.Time        `json:"timeBase" validate:"required"`
}

type NotifyEventRequest struct {
	CustomData  CustomData  `json:"customData,omitempty"`
	EventData   []EventData `json:"eventData" validate:"required,min=1"`
	GeneratedAt time.Time   `json:"generatedAt" validate:"required"`
	SeqNo       int         `json:"seqNo" validate:"required,gte=0.0"`
	Tbc         *bool       `json:"tbc,omitempty"`
}

type NotifyMonitoringReportRequest struct {
	CustomData  CustomData       `json:"customData,omitempty"`
	GeneratedAt time.Time        `json:"generatedAt" validate:"required"`
	Monitor     []MonitoringData `json:"monitor,omitempty" validate:"min=1"`
	RequestID   int              `json:"requestId" validate:"required"`
	SeqNo       int              `json:"seqNo" validate:"required,gte=0.0"`
	Tbc         *bool            `json:"tbc,omitempty"`
}

type NotifyPeriodicEventStream struct {
	Basetime   time.Time           `json:"basetime" validate:"required"`
	CustomData CustomData          `json:"customData,omitempty"`
	Data       []StreamDataElement `json:"data" validate:"required,min=1"`
	ID         int                 `json:"id" validate:"required,gte=0.0"`
	Pending    int                 `json:"pending" validate:"required,gte=0.0"`
}

type NotifyPeriodicEventStreamRequest = NotifyPeriodicEventStream

type NotifyPriorityChargingRequest struct {
	Activated     bool       `json:"activated" validate:"required"`
	CustomData    CustomData `json:"customData,omitempty"`
	TransactionID string     `json:"transactionId" validate:"required,max=36"`
}

type NotifyReportRequest struct {
	CustomData  CustomData   `json:"customData,omitempty"`
	GeneratedAt time.Time    `json:"generatedAt" validate:"required"`
	ReportData  []ReportData `json:"reportData,omitempty" validate:"min=1"`
	RequestID   int          `json:"requestId" validate:"required"`
	SeqNo       int          `json:"seqNo" validate:"required,gte=0.0"`
	Tbc         *bool        `json:"tbc,omitempty"`
}

type NotifySettlementRequest struct {
	CustomData       CustomData        `json:"customData,omitempty"`
	PspRef           string            `json:"pspRef" validate:"required,max=255"`
	ReceiptID        *string           `json:"receiptId,omitempty" validate:"max=50"`
	ReceiptURL       *string           `json:"receiptUrl,omitempty" validate:"max=2000"`
	SettlementAmount float64           `json:"settlementAmount" validate:"required"`
	SettlementTime   time.Time         `json:"settlementTime" validate:"required"`
	Status           PaymentStatusEnum `json:"status" validate:"required"`
	StatusInfo       *string           `json:"statusInfo,omitempty" validate:"max=500"`
	TransactionID    *string           `json:"transactionId,omitempty" validate:"max=36"`
	VatCompany       *Address          `json:"vatCompany,omitempty"`
	VatNumber        *string           `json:"vatNumber,omitempty" validate:"max=20"`
}

type NotifyWebPaymentStartedRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	EVSEID     int        `json:"evseId" validate:"required,gte=0.0"`
	Timeout    int        `json:"timeout" validate:"required"`
}

type OpenPeriodicEventStreamRequest struct {
	ConstantStreamData ConstantStreamData `json:"constantStreamData" validate:"required"`
	CustomData         CustomData         `json:"customData,omitempty"`
}

type PublishFirmwareRequest struct {
	Checksum      string     `json:"checksum" validate:"required,max=32"`
	CustomData    CustomData `json:"customData,omitempty"`
	Location      string     `json:"location" validate:"required,max=2000"`
	RequestID     int        `json:"requestId" validate:"required,gte=0.0"`
	Retries       *int       `json:"retries,omitempty" validate:"gte=0.0"`
	RetryInterval *int       `json:"retryInterval,omitempty" validate:"gte=0.0"`
}

type PublishFirmwareStatusNotificationRequest struct {
	CustomData CustomData                `json:"customData,omitempty"`
	Location   []string                  `json:"location,omitempty" validate:"min=1"`
	RequestID  *int                      `json:"requestId,omitempty" validate:"gte=0.0"`
	Status     PublishFirmwareStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo               `json:"statusInfo,omitempty"`
}

type PullDynamicScheduleUpdateRequest struct {
	ChargingProfileID int        `json:"chargingProfileId" validate:"required"`
	CustomData        CustomData `json:"customData,omitempty"`
}

type ReportChargingProfilesRequest struct {
	ChargingLimitSource string            `json:"chargingLimitSource" validate:"required,max=20"`
	ChargingProfile     []ChargingProfile `json:"chargingProfile" validate:"required,min=1"`
	CustomData          CustomData        `json:"customData,omitempty"`
	EVSEID              int               `json:"evseId" validate:"required,gte=0.0"`
	RequestID           int               `json:"requestId" validate:"required"`
	Tbc                 *bool             `json:"tbc,omitempty"`
}

type ReportDERControlRequest struct {
	Curve             []DERCurveGet          `json:"curve,omitempty" validate:"min=1,max=24"`
	CustomData        CustomData             `json:"customData,omitempty"`
	EnterService      []EnterServiceGet      `json:"enterService,omitempty" validate:"min=1,max=24"`
	FixedPFAbsorb     []FixedPFGet           `json:"fixedPFAbsorb,omitempty" validate:"min=1,max=24"`
	FixedPFInject     []FixedPFGet           `json:"fixedPFInject,omitempty" validate:"min=1,max=24"`
	FixedVAR          []FixedVARGet          `json:"fixedVar,omitempty" validate:"min=1,max=24"`
	FreqDroop         []FreqDroopGet         `json:"freqDroop,omitempty" validate:"min=1,max=24"`
	Gradient          []GradientGet          `json:"gradient,omitempty" validate:"min=1,max=24"`
	LimitMaxDischarge []LimitMaxDischargeGet `json:"limitMaxDischarge,omitempty" validate:"min=1,max=24"`
	RequestID         int                    `json:"requestId" validate:"required"`
	Tbc               *bool                  `json:"tbc,omitempty"`
}

type RequestBatterySwapRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	IDToken    IDToken    `json:"idToken" validate:"required"`
	RequestID  int        `json:"requestId" validate:"required"`
}

type RequestStartTransactionRequest struct {
	ChargingProfile *ChargingProfile `json:"chargingProfile,omitempty"`
	CustomData      CustomData       `json:"customData,omitempty"`
	EVSEID          *int             `json:"evseId,omitempty" validate:"gte=1.0"`
	GroupIDToken    *IDToken         `json:"groupIdToken,omitempty"`
	IDToken         IDToken          `json:"idToken" validate:"required"`
	RemoteStartID   int              `json:"remoteStartId" validate:"required"`
}

type RequestStopTransactionRequest struct {
	CustomData    CustomData `json:"customData,omitempty"`
	TransactionID string     `json:"transactionId" validate:"required,max=36"`
}

type ReservationStatusUpdateRequest struct {
	CustomData              CustomData                  `json:"customData,omitempty"`
	ReservationID           int                         `json:"reservationId" validate:"required,gte=0.0"`
	ReservationUpdateStatus ReservationUpdateStatusEnum `json:"reservationUpdateStatus" validate:"required"`
}

type ReserveNowRequest struct {
	ConnectorType  *string    `json:"connectorType,omitempty" validate:"max=20"`
	CustomData     CustomData `json:"customData,omitempty"`
	EVSEID         *int       `json:"evseId,omitempty" validate:"gte=0.0"`
	ExpiryDateTime time.Time  `json:"expiryDateTime" validate:"required"`
	GroupIDToken   *IDToken   `json:"groupIdToken,omitempty"`
	ID             int        `json:"id" validate:"required,gte=0.0"`
	IDToken        IDToken    `json:"idToken" validate:"required"`
}

type ResetRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	EVSEID     *int       `json:"evseId,omitempty" validate:"gte=0.0"`
	TypeValue  ResetEnum  `json:"type" validate:"required"`
}

type SecurityEventNotificationRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	TechInfo   *string    `json:"techInfo,omitempty" validate:"max=255"`
	Timestamp  time.Time  `json:"timestamp" validate:"required"`
	TypeValue  string     `json:"type" validate:"required,max=50"`
}

type SendLocalListRequest struct {
	CustomData             CustomData          `json:"customData,omitempty"`
	LocalAuthorizationList []AuthorizationData `json:"localAuthorizationList,omitempty" validate:"min=1"`
	UpdateType             UpdateEnum          `json:"updateType" validate:"required"`
	VersionNumber          int                 `json:"versionNumber" validate:"required"`
}

type SetChargingProfileRequest struct {
	ChargingProfile ChargingProfile `json:"chargingProfile" validate:"required"`
	CustomData      CustomData      `json:"customData,omitempty"`
	EVSEID          int             `json:"evseId" validate:"required,gte=0.0"`
}

type SetDERControlRequest struct {
	ControlID         string             `json:"controlId" validate:"required,max=36"`
	ControlType       DERControlEnum     `json:"controlType" validate:"required"`
	Curve             *DERCurve          `json:"curve,omitempty"`
	CustomData        CustomData         `json:"customData,omitempty"`
	EnterService      *EnterService      `json:"enterService,omitempty"`
	FixedPFAbsorb     *FixedPF           `json:"fixedPFAbsorb,omitempty"`
	FixedPFInject     *FixedPF           `json:"fixedPFInject,omitempty"`
	FixedVAR          *FixedVAR          `json:"fixedVar,omitempty"`
	FreqDroop         *FreqDroop         `json:"freqDroop,omitempty"`
	Gradient          *Gradient          `json:"gradient,omitempty"`
	IsDefault         bool               `json:"isDefault" validate:"required"`
	LimitMaxDischarge *LimitMaxDischarge `json:"limitMaxDischarge,omitempty"`
}

type SetDefaultTariffRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	EVSEID     int        `json:"evseId" validate:"required,gte=0.0"`
	Tariff     Tariff     `json:"tariff" validate:"required"`
}

type SetDisplayMessageRequest struct {
	CustomData CustomData  `json:"customData,omitempty"`
	Message    MessageInfo `json:"message" validate:"required"`
}

type SetMonitoringBaseRequest struct {
	CustomData     CustomData         `json:"customData,omitempty"`
	MonitoringBase MonitoringBaseEnum `json:"monitoringBase" validate:"required"`
}

type SetMonitoringLevelRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	Severity   int        `json:"severity" validate:"required,gte=0.0"`
}

type SetNetworkProfileRequest struct {
	ConfigurationSlot int                      `json:"configurationSlot" validate:"required"`
	ConnectionData    NetworkConnectionProfile `json:"connectionData" validate:"required"`
	CustomData        CustomData               `json:"customData,omitempty"`
}

type SetVariableMonitoringRequest struct {
	CustomData        CustomData          `json:"customData,omitempty"`
	SetMonitoringData []SetMonitoringData `json:"setMonitoringData" validate:"required,min=1"`
}

type SetVariablesRequest struct {
	CustomData      CustomData        `json:"customData,omitempty"`
	SetVariableData []SetVariableData `json:"setVariableData" validate:"required,min=1"`
}

type SignCertificateRequest struct {
	CertificateType     *CertificateSigningUseEnum `json:"certificateType,omitempty"`
	Csr                 string                     `json:"csr" validate:"required,max=5500"`
	CustomData          CustomData                 `json:"customData,omitempty"`
	HashRootCertificate *CertificateHashData       `json:"hashRootCertificate,omitempty"`
	RequestID           *int                       `json:"requestId,omitempty"`
}

type StatusNotificationRequest struct {
	ConnectorID     int                 `json:"connectorId" validate:"required,gte=0.0"`
	ConnectorStatus ConnectorStatusEnum `json:"connectorStatus" validate:"required"`
	CustomData      CustomData          `json:"customData,omitempty"`
	EVSEID          int                 `json:"evseId" validate:"required,gte=0.0"`
	Timestamp       time.Time           `json:"timestamp" validate:"required"`
}

type TransactionEventRequest struct {
	CableMaxCurrent       *int                       `json:"cableMaxCurrent,omitempty"`
	CostDetails           *CostDetails               `json:"costDetails,omitempty"`
	CustomData            CustomData                 `json:"customData,omitempty"`
	EventType             TransactionEventEnum       `json:"eventType" validate:"required"`
	EVSE                  *EVSE                      `json:"evse,omitempty"`
	EVSESleep             *bool                      `json:"evseSleep,omitempty"`
	IDToken               *IDToken                   `json:"idToken,omitempty"`
	MeterValue            []MeterValue               `json:"meterValue,omitempty" validate:"min=1"`
	NumberOfPhasesUsed    *int                       `json:"numberOfPhasesUsed,omitempty" validate:"gte=0.0,lte=3.0"`
	Offline               *bool                      `json:"offline,omitempty"`
	PreconditioningStatus *PreconditioningStatusEnum `json:"preconditioningStatus,omitempty"`
	ReservationID         *int                       `json:"reservationId,omitempty" validate:"gte=0.0"`
	SeqNo                 int                        `json:"seqNo" validate:"required,gte=0.0"`
	Timestamp             time.Time                  `json:"timestamp" validate:"required"`
	TransactionInfo       Transaction                `json:"transactionInfo" validate:"required"`
	TriggerReason         TriggerReasonEnum          `json:"triggerReason" validate:"required"`
}

type TriggerMessageRequest struct {
	CustomData       CustomData         `json:"customData,omitempty"`
	CustomTrigger    *string            `json:"customTrigger,omitempty" validate:"max=50"`
	EVSE             *EVSE              `json:"evse,omitempty"`
	RequestedMessage MessageTriggerEnum `json:"requestedMessage" validate:"required"`
}

type UnlockConnectorRequest struct {
	ConnectorID int        `json:"connectorId" validate:"required,gte=0.0"`
	CustomData  CustomData `json:"customData,omitempty"`
	EVSEID      int        `json:"evseId" validate:"required,gte=0.0"`
}

type UnpublishFirmwareRequest struct {
	Checksum   string     `json:"checksum" validate:"required,max=32"`
	CustomData CustomData `json:"customData,omitempty"`
}

type UpdateDynamicScheduleRequest struct {
	ChargingProfileID int                    `json:"chargingProfileId" validate:"required"`
	CustomData        CustomData             `json:"customData,omitempty"`
	ScheduleUpdate    ChargingScheduleUpdate `json:"scheduleUpdate" validate:"required"`
}

type UpdateFirmwareRequest struct {
	CustomData    CustomData `json:"customData,omitempty"`
	Firmware      Firmware   `json:"firmware" validate:"required"`
	RequestID     int        `json:"requestId" validate:"required"`
	Retries       *int       `json:"retries,omitempty" validate:"gte=0.0"`
	RetryInterval *int       `json:"retryInterval,omitempty"`
}

type UsePriorityChargingRequest struct {
	Activate      bool       `json:"activate" validate:"required"`
	CustomData    CustomData `json:"customData,omitempty"`
	TransactionID string     `json:"transactionId" validate:"required,max=36"`
}

type VatNumberValidationRequest struct {
	CustomData CustomData `json:"customData,omitempty"`
	EVSEID     *int       `json:"evseId,omitempty" validate:"gte=0.0"`
	VatNumber  string     `json:"vatNumber" validate:"required,max=20"`
}
