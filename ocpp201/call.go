// Package ocpp201 contains the data structures for OCPP 2.0.1 messages.
package ocpp201

import "time"

type AuthorizeRequest struct {
	IDToken                     IdToken               `json:"idToken" validate:"required"`
	Certificate                 *string               `json:"certificate,omitempty" validate:"max=800"`
	ISO15118CertificateHashData []CertificateHashData `json:"iso15118CertificateHashData,omitempty"`
}

type BootNotificationRequest struct {
	Reason          BootReason      `json:"reason" validate:"required"`
	ChargingStation ChargingStation `json:"chargingStation" validate:"required"`
}

type CancelReservationRequest struct {
	ReservationID int `json:"reservationId" validate:"required"`
}

type CertificateSignedRequest struct {
	CertificateChain string             `json:"certificateChain" validate:"required,max=10000"`
	CertificateType  *CertificateAction `json:"certificateType,omitempty"`
}

type ChangeAvailabilityRequest struct {
	OperationalStatus OperationalStatus `json:"operationalStatus" validate:"required"`
	Evse              *EVSE             `json:"evse,omitempty"`
}

type ClearCacheRequest struct {
	// No fields
}

type ClearChargingProfileRequest struct {
	ChargingProfileID    *int                  `json:"chargingProfileId,omitempty"`
	ClearChargingProfile *ClearChargingProfile `json:"clearChargingProfile,omitempty"`
}

type ClearChargingProfile struct {
	EvseID                 *int                    `json:"evseId,omitempty"`
	ChargingProfilePurpose *ChargingProfilePurpose `json:"chargingProfilePurpose,omitempty"`
	StackLevel             *int                    `json:"stackLevel,omitempty"`
}

type ClearDisplayMessageRequest struct {
	ID int `json:"id" validate:"required"`
}

type ClearedChargingLimitRequest struct {
	ChargingLimitSource string `json:"chargingLimitSource" validate:"required,max=20"`
	EvseID              *int   `json:"evseId,omitempty"`
}

type ClearVariableMonitoringRequest struct {
	ID []int `json:"id" validate:"required,min=1"`
}

type CostUpdatedRequest struct {
	TotalCost     float64 `json:"totalCost" validate:"required"`
	TransactionID string  `json:"transactionId" validate:"required,max=36"`
}

type CustomerInformationRequest struct {
	RequestID           int      `json:"requestId" validate:"required"`
	Report              bool     `json:"report"`
	Clear               bool     `json:"clear"`
	CustomerCertificate *string  `json:"customerCertificate,omitempty" validate:"max=5500"`
	IDToken             *IdToken `json:"idToken,omitempty"`
	CustomerIdentifier  *string  `json:"customerIdentifier,omitempty" validate:"max=64"`
}

type DataTransferRequest struct {
	VendorID  string      `json:"vendorId" validate:"required,max=255"`
	MessageID *string     `json:"messageId,omitempty" validate:"max=50"`
	Data      interface{} `json:"data,omitempty"`
}

type DeleteCertificateRequest struct {
	CertificateHashData CertificateHashData `json:"certificateHashData" validate:"required"`
}

type FirmwareStatusNotificationRequest struct {
	Status    FirmwareStatus `json:"status" validate:"required"`
	RequestID *int           `json:"requestId,omitempty"`
}

type Get15118EVCertificateRequest struct {
	ISO15118SchemaVersion string            `json:"iso15118SchemaVersion" validate:"required,max=50"`
	Action                CertificateAction `json:"action" validate:"required"`
	ExiRequest            string            `json:"exiRequest" validate:"required,max=10000"`
}

type GetBaseReportRequest struct {
	RequestID  int        `json:"requestId" validate:"required"`
	ReportBase ReportBase `json:"reportBase" validate:"required"`
}

type GetCertificateStatusRequest struct {
	OCSPRequestData OCSPRequestData `json:"ocspRequestData" validate:"required"`
}

type GetChargingProfilesRequest struct {
	RequestID       int             `json:"requestId" validate:"required"`
	EvseID          *int            `json:"evseId,omitempty"`
	ChargingProfile ChargingProfile `json:"chargingProfile" validate:"required"`
}

type GetCompositeScheduleRequest struct {
	Duration         int               `json:"duration" validate:"required"`
	EvseID           int               `json:"evseId" validate:"required"`
	ChargingRateUnit *ChargingRateUnit `json:"chargingRateUnit,omitempty"`
}

type GetDisplayMessagesRequest struct {
	ID        []int            `json:"id,omitempty"`
	RequestID int              `json:"requestId" validate:"required"`
	Priority  *MessagePriority `json:"priority,omitempty"`
	State     *MessageState    `json:"state,omitempty"`
}

type GetInstalledCertificateIdsRequest struct {
	CertificateType []GetCertificateIdUse `json:"certificateType,omitempty"`
}

type GetLocalListVersionRequest struct {
	// No fields
}

type GetLogRequest struct {
	LogType       Log           `json:"logType" validate:"required"`
	RequestID     int           `json:"requestId" validate:"required"`
	Retries       *int          `json:"retries,omitempty"`
	RetryInterval *int          `json:"retryInterval,omitempty"`
	Log           LogParameters `json:"log" validate:"required"`
}

type GetMonitoringReportRequest struct {
	RequestID          int                 `json:"requestId" validate:"required"`
	MonitoringCriteria []string            `json:"monitoringCriteria,omitempty"`
	ComponentVariable  []ComponentVariable `json:"componentVariable,omitempty"`
}

type GetReportRequest struct {
	RequestID         int                 `json:"requestId" validate:"required"`
	ComponentCriteria []string            `json:"componentCriteria,omitempty"`
	ComponentVariable []ComponentVariable `json:"componentVariable,omitempty"`
}

type GetTransactionStatusRequest struct {
	TransactionID *string `json:"transactionId,omitempty" validate:"max=36"`
}

type GetVariablesRequest struct {
	GetVariableData []GetVariableData `json:"getVariableData" validate:"required,min=1"`
}

type HeartbeatRequest struct {
	// No fields
}

type InstallCertificateRequest struct {
	CertificateType InstallCertificateUse `json:"certificateType" validate:"required"`
	Certificate     string                `json:"certificate" validate:"required,max=10000"`
}

type LogStatusNotificationRequest struct {
	Status    UploadLogStatus `json:"status" validate:"required"`
	RequestID *int            `json:"requestId,omitempty"`
}

type MeterValuesRequest struct {
	EvseID     int          `json:"evseId" validate:"required"`
	MeterValue []MeterValue `json:"meterValue" validate:"required,min=1"`
}

type NotifyChargingLimitRequest struct {
	ChargingLimit    ChargingLimit      `json:"chargingLimit" validate:"required"`
	ChargingSchedule []ChargingSchedule `json:"chargingSchedule,omitempty"`
	EvseID           *int               `json:"evseId,omitempty"`
}

type NotifyCustomerInformationRequest struct {
	Data        string    `json:"data" validate:"required,max=512"`
	TBC         bool      `json:"tbc,omitempty"`
	SeqNo       int       `json:"seqNo" validate:"required"`
	GeneratedAt time.Time `json:"generatedAt" validate:"required"`
	RequestID   int       `json:"requestId" validate:"required"`
}

type NotifyDisplayMessagesRequest struct {
	RequestID   int           `json:"requestId" validate:"required"`
	TBC         bool          `json:"tbc,omitempty"`
	MessageInfo []MessageInfo `json:"messageInfo,omitempty"`
}

type NotifyEVChargingNeedsRequest struct {
	MaxScheduleTuples *int          `json:"maxScheduleTuples,omitempty"`
	ChargingNeeds     ChargingNeeds `json:"chargingNeeds" validate:"required"`
	EvseID            int           `json:"evseId" validate:"required"`
}

type NotifyEVChargingScheduleRequest struct {
	TimeBase         time.Time        `json:"timeBase" validate:"required"`
	ChargingSchedule ChargingSchedule `json:"chargingSchedule" validate:"required"`
	EvseID           int              `json:"evseId" validate:"required"`
}

type NotifyEventRequest struct {
	GeneratedAt time.Time   `json:"generatedAt" validate:"required"`
	TBC         bool        `json:"tbc,omitempty"`
	SeqNo       int         `json:"seqNo" validate:"required"`
	EventData   []EventData `json:"eventData" validate:"required,min=1"`
}

type NotifyMonitoringReportRequest struct {
	RequestID   int              `json:"requestId" validate:"required"`
	TBC         bool             `json:"tbc,omitempty"`
	SeqNo       int              `json:"seqNo" validate:"required"`
	GeneratedAt time.Time        `json:"generatedAt" validate:"required"`
	Monitor     []MonitoringData `json:"monitor,omitempty"`
}

type NotifyReportRequest struct {
	RequestID   int       `json:"requestId" validate:"required"`
	GeneratedAt time.Time `json:"generatedAt" validate:"required"`
	TBC         bool      `json:"tbc,omitempty"`
	SeqNo       int       `json:"seqNo" validate:"required"`
}

type PublishFirmwareRequest struct {
	Location      string `json:"location" validate:"required,max=512"`
	Retries       *int   `json:"retries,omitempty"`
	Checksum      string `json:"checksum" validate:"required,max=32"`
	RequestID     int    `json:"requestId" validate:"required"`
	RetryInterval *int   `json:"retryInterval,omitempty"`
}

type PublishFirmwareStatusNotificationRequest struct {
	Status    PublishFirmwareStatus `json:"status" validate:"required"`
	Location  []string              `json:"location,omitempty"`
	RequestID *int                  `json:"requestId,omitempty"`
}

type ReportChargingProfilesRequest struct {
	RequestID           int               `json:"requestId" validate:"required"`
	ChargingLimitSource string            `json:"chargingLimitSource" validate:"required,max=20"`
	ChargingProfile     []ChargingProfile `json:"chargingProfile" validate:"required,min=1"`
	EvseID              int               `json:"evseId" validate:"required"`
	TBC                 bool              `json:"tbc,omitempty"`
}

type RequestStartTransactionRequest struct {
	RemoteStartID   int              `json:"remoteStartId" validate:"required"`
	IDToken         IdToken          `json:"idToken" validate:"required"`
	EvseID          *int             `json:"evseId,omitempty"`
	ChargingProfile *ChargingProfile `json:"chargingProfile,omitempty"`
	GroupIDToken    *IdToken         `json:"groupIdToken,omitempty"`
}

type RequestStopTransactionRequest struct {
	TransactionID string `json:"transactionId" validate:"required,max=36"`
}

type ReservationStatusUpdateRequest struct {
	ReservationID           int                     `json:"reservationId" validate:"required"`
	ReservationUpdateStatus ReservationUpdateStatus `json:"reservationUpdateStatus" validate:"required"`
}

type ReserveNowRequest struct {
	ID             int       `json:"id" validate:"required"`
	ExpiryDateTime time.Time `json:"expiryDateTime" validate:"required"`
	ConnectorType  *string   `json:"connectorType,omitempty" validate:"max=16"`
	EvseID         *int      `json:"evseId,omitempty"`
	IDToken        IdToken   `json:"idToken" validate:"required"`
	GroupIDToken   *IdToken  `json:"groupIdToken,omitempty"`
}

type ResetRequest struct {
	Type   Reset `json:"type" validate:"required"`
	EvseID *int  `json:"evseId,omitempty"`
}

type SecurityEventNotificationRequest struct {
	Type      string    `json:"type" validate:"required,max=50"`
	Timestamp time.Time `json:"timestamp" validate:"required"`
	TechInfo  *string   `json:"techInfo,omitempty" validate:"max=255"`
}

type SendLocalListRequest struct {
	VersionNumber          int                 `json:"versionNumber" validate:"required"`
	UpdateType             Update              `json:"updateType" validate:"required"`
	LocalAuthorizationList []AuthorizationData `json:"localAuthorizationList,omitempty"`
}

type AuthorizationData struct {
	IDToken     IdToken      `json:"idToken" validate:"required"`
	IDTokenInfo *IdTokenInfo `json:"idTokenInfo,omitempty"`
}

type SetChargingProfileRequest struct {
	EvseID          int             `json:"evseId" validate:"required"`
	ChargingProfile ChargingProfile `json:"chargingProfile" validate:"required"`
}

type SetClockRequest struct {
	CurrentTime time.Time `json:"currentTime" validate:"required"`
}

type SetDisplayMessageRequest struct {
	Message MessageInfo `json:"message" validate:"required"`
}

type SetMonitoringBaseRequest struct {
	MonitoringBase ReportBase `json:"monitoringBase" validate:"required"`
}

type SetMonitoringLevelRequest struct {
	Severity int `json:"severity" validate:"required"`
}

type SetNetworkProfileRequest struct {
	ConfigurationSlot int                      `json:"configurationSlot" validate:"required"`
	ConnectionData    NetworkConnectionProfile `json:"connectionData" validate:"required"`
}

type SetVariableMonitoringRequest struct {
	SetMonitoringData []MonitoringData `json:"setMonitoringData" validate:"required,min=1"`
}

type SetVariablesRequest struct {
	SetVariableData []SetVariableData `json:"setVariableData" validate:"required,min=1"`
}

type SignCertificateRequest struct {
	Csr             string                 `json:"csr" validate:"required,max=5500"`
	CertificateType *CertificateSigningUse `json:"certificateType,omitempty"`
}

type StatusNotificationRequest struct {
	Timestamp       time.Time       `json:"timestamp" validate:"required"`
	ConnectorStatus ConnectorStatus `json:"connectorStatus" validate:"required"`
	EvseID          int             `json:"evseId" validate:"required,gte=0"`
	ConnectorID     int             `json:"connectorId" validate:"required,gte=0"`
}

type TransactionEventRequest struct {
	EventType          TransactionEvent `json:"eventType" validate:"required"`
	Timestamp          time.Time        `json:"timestamp" validate:"required"`
	TriggerReason      TriggerReason    `json:"triggerReason" validate:"required"`
	SeqNo              int              `json:"seqNo" validate:"required"`
	TransactionInfo    Transaction      `json:"transactionInfo" validate:"required"`
	IdToken            *IdToken         `json:"idToken,omitempty"`
	Evse               *EVSE            `json:"evse,omitempty"`
	MeterValue         []MeterValue     `json:"meterValue,omitempty"`
	NumberOfPhasesUsed *int             `json:"numberOfPhasesUsed,omitempty"`
	CableMaxCurrent    *int             `json:"cableMaxCurrent,omitempty"`
	ReservationID      *int             `json:"reservationId,omitempty"`
	Offline            bool             `json:"offline"`
}

type TriggerMessageRequest struct {
	RequestedMessage TriggerMessage `json:"requestedMessage" validate:"required"`
	Evse             *EVSE          `json:"evse,omitempty"`
}

type TriggerMessage struct {
	// Enum type for TriggerMessageRequest.RequestedMessage
}

type UnlockConnectorRequest struct {
	EvseID      int `json:"evseId" validate:"required"`
	ConnectorID int `json:"connectorId" validate:"required"`
}

type UnpublishFirmwareRequest struct {
	Checksum string `json:"checksum" validate:"required,max=32"`
}

type UpdateFirmwareRequest struct {
	RequestID     int      `json:"requestId" validate:"required"`
	Firmware      Firmware `json:"firmware" validate:"required"`
	Retries       *int     `json:"retries,omitempty"`
	RetryInterval *int     `json:"retryInterval,omitempty"`
}
