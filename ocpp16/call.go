package ocpp16

import "time"

type AuthorizeRequest struct {
	IdTag string `json:"idTag" validate:"required,max=20"`
}

type BootNotificationRequest struct {
	ChargePointModel        string  `json:"chargePointModel" validate:"required,max=20"`
	ChargePointVendor       string  `json:"chargePointVendor" validate:"required,max=20"`
	ChargeBoxSerialNumber   *string `json:"chargeBoxSerialNumber,omitempty" validate:"max=25"`
	ChargePointSerialNumber *string `json:"chargePointSerialNumber,omitempty" validate:"max=25"`
	FirmwareVersion         *string `json:"firmwareVersion,omitempty" validate:"max=50"`
	Iccid                   *string `json:"iccid,omitempty" validate:"max=20"`
	Imsi                    *string `json:"imsi,omitempty" validate:"max=20"`
	MeterSerialNumber       *string `json:"meterSerialNumber,omitempty" validate:"max=25"`
	MeterType               *string `json:"meterType,omitempty" validate:"max=25"`
}

type ChangeAvailabilityRequest struct {
	ConnectorId int              `json:"connectorId" validate:"gte=0"`
	Type        AvailabilityType `json:"type" validate:"required"`
}

type ChangeConfigurationRequest struct {
	Key   string `json:"key" validate:"required,max=50"`
	Value string `json:"value" validate:"required,max=500"`
}

type ClearCacheRequest struct{}

type ClearChargingProfileRequest struct {
	ID                     *int                        `json:"id,omitempty"`
	ConnectorId            *int                        `json:"connectorId,omitempty" validate:"gte=0"`
	ChargingProfilePurpose *ChargingProfilePurposeType `json:"chargingProfilePurpose,omitempty"`
	StackLevel             *int                        `json:"stackLevel,omitempty" validate:"gte=0"`
}

type DataTransferRequest struct {
	VendorId  string      `json:"vendorId" validate:"required,max=255"`
	MessageId *string     `json:"messageId,omitempty" validate:"max=50"`
	Data      interface{} `json:"data,omitempty"`
}

type DiagnosticsStatusNotificationRequest struct {
	Status DiagnosticsStatus `json:"status" validate:"required"`
}

type FirmwareStatusNotificationRequest struct {
	Status FirmwareStatus `json:"status" validate:"required"`
}

type ExtendedTriggerMessageRequest struct {
	RequestedMessage MessageTrigger `json:"requestedMessage" validate:"required,oneof=BootNotification DiagnosticsStatusNotification FirmwareStatusNotification Heartbeat MeterValues StatusNotification"`
	ConnectorId      *int           `json:"connectorId,omitempty" validate:"gt=0"`
}

type GetConfigurationRequest struct {
	Key []string `json:"key,omitempty" validate:"max=50"`
}

type GetDiagnosticsRequest struct {
	Location      string     `json:"location" validate:"required,max=512"`
	Retries       *int       `json:"retries,omitempty" validate:"gte=0"`
	RetryInterval *int       `json:"retryInterval,omitempty" validate:"gte=0"`
	StartTime     *time.Time `json:"startTime,omitempty"`
	StopTime      *time.Time `json:"stopTime,omitempty"`
}

type GetLocalListVersionRequest struct{}

type HeartbeatRequest struct{}

type MeterValuesRequest struct {
	ConnectorId   int          `json:"connectorId" validate:"gte=0"`
	TransactionId *int         `json:"transactionId,omitempty"`
	MeterValue    []MeterValue `json:"meterValue" validate:"required,min=1"`
}

type RemoteStartTransactionRequest struct {
	ConnectorId     *int             `json:"connectorId,omitempty" validate:"gt=0"`
	IdTag           string           `json:"idTag" validate:"required,max=20"`
	ChargingProfile *ChargingProfile `json:"chargingProfile,omitempty"`
}

type RemoteStopTransactionRequest struct {
	TransactionId int `json:"transactionId"`
}

type ResetRequest struct {
	Type ResetType `json:"type" validate:"required"`
}

type SendLocalListRequest struct {
	ListVersion            int                 `json:"listVersion"`
	LocalAuthorizationList []AuthorizationData `json:"localAuthorizationList,omitempty"`
	UpdateType             UpdateType          `json:"updateType" validate:"required"`
}

type SetChargingProfileRequest struct {
	ConnectorId     int             `json:"connectorId" validate:"gte=0"`
	ChargingProfile ChargingProfile `json:"csChargingProfiles" validate:"required"`
}

type StartTransactionRequest struct {
	ConnectorId   int       `json:"connectorId" validate:"gt=0"`
	IdTag         string    `json:"idTag" validate:"required,max=20"`
	MeterStart    int       `json:"meterStart"`
	ReservationId *int      `json:"reservationId,omitempty"`
	Timestamp     time.Time `json:"timestamp"`
}

type StatusNotificationRequest struct {
	ConnectorId     int                  `json:"connectorId" validate:"gte=0"`
	ErrorCode       ChargePointErrorCode `json:"errorCode" validate:"required"`
	Info            *string              `json:"info,omitempty" validate:"max=50"`
	Status          ChargePointStatus    `json:"status" validate:"required"`
	Timestamp       *time.Time           `json:"timestamp,omitempty"`
	VendorId        *string              `json:"vendorId,omitempty" validate:"max=255"`
	VendorErrorCode *string              `json:"vendorErrorCode,omitempty" validate:"max=50"`
}

type StopTransactionRequest struct {
	IdTag           *string      `json:"idTag,omitempty" validate:"max=20"`
	MeterStop       int          `json:"meterStop"`
	Timestamp       time.Time    `json:"timestamp"`
	TransactionId   int          `json:"transactionId"`
	Reason          *Reason      `json:"reason,omitempty"`
	TransactionData []MeterValue `json:"transactionData,omitempty"`
}

type TriggerMessageRequest struct {
	RequestedMessage MessageTrigger `json:"requestedMessage" validate:"required"`
	ConnectorId      *int           `json:"connectorId,omitempty" validate:"gt=0"`
}

type UnlockConnectorRequest struct {
	ConnectorId int `json:"connectorId" validate:"gt=0"`
}

type UpdateFirmwareRequest struct {
	Location      string    `json:"location" validate:"required,max=512"`
	Retries       *int      `json:"retries,omitempty" validate:"gte=0"`
	RetrieveDate  time.Time `json:"retrieveDate"`
	RetryInterval *int      `json:"retryInterval,omitempty" validate:"gte=0"`
}

type CancelReservationRequest struct {
	ReservationID int `json:"reservationId"`
}

type ReserveNowRequest struct {
	ConnectorID   int       `json:"connectorId"`
	ExpiryDate    time.Time `json:"expiryDate"`
	IDTag         string    `json:"idTag" validate:"required,max=20"`
	ParentIDTag   *string   `json:"parentIdTag,omitempty" validate:"max=20"`
	ReservationID int       `json:"reservationId"`
}

type GetCompositeScheduleRequest struct {
	ConnectorID      int                   `json:"connectorId"`
	Duration         int                   `json:"duration"`
	ChargingRateUnit *ChargingRateUnitType `json:"chargingRateUnit,omitempty"`
}

// Security Extension Requests
type CertificateSignedRequest struct {
	CertificateChain string `json:"certificateChain" validate:"required,max=10000"`
}

type DeleteCertificateRequest struct {
	CertificateHashData CertificateHashData `json:"certificateHashData" validate:"required"`
}

type GetInstalledCertificateIdsRequest struct {
	CertificateType string `json:"certificateType" validate:"required,oneof=V2GRootCertificate MORootCertificate CSOSubCA1 CSOSubCA2 ChargingStationCertificate"`
}

type InstallCertificateRequest struct {
	CertificateType string `json:"certificateType" validate:"required,oneof=V2GRootCertificate MORootCertificate CSOSubCA1 CSOSubCA2 ChargingStationCertificate"`
	Certificate     string `json:"certificate" validate:"required,max=800"`
}

type LogStatusNotificationRequest struct {
	Status    DiagnosticsStatus `json:"status" validate:"required"`
	RequestID *int              `json:"requestId,omitempty"`
}

type SecurityEventNotificationRequest struct {
	Type      string    `json:"type" validate:"required,max=50"`
	Timestamp time.Time `json:"timestamp"`
	TechInfo  *string   `json:"techInfo,omitempty" validate:"max=255"`
}

type SignCertificateRequest struct {
	Csr string `json:"csr" validate:"required,max=5500"`
}

type SignedUpdateFirmwareRequest struct {
	RequestID int      `json:"requestId"`
	Firmware  Firmware `json:"firmware" validate:"required"`
}

type Firmware struct {
	Location           string     `json:"location" validate:"required,max=512"`
	RetrieveTimestamp  time.Time  `json:"retrieveTimestamp"`
	InstallTimestamp   *time.Time `json:"installTimestamp,omitempty"`
	Signature          string     `json:"signature" validate:"required,max=800"`
	SigningCertificate string     `json:"signingCertificate" validate:"required,max=5500"`
}

type GetLogRequest struct {
	LogType       string        `json:"logType" validate:"required,oneof=DiagnosticsLog SecurityLog"`
	RequestID     int           `json:"requestId"`
	Retries       *int          `json:"retries,omitempty" validate:"gte=0"`
	RetryInterval *int          `json:"retryInterval,omitempty" validate:"gte=0"`
	Log           LogParameters `json:"log" validate:"required"`
}

type SignedFirmwareStatusNotificationRequest struct {
	Status    FirmwareStatus `json:"status" validate:"required,oneof=Downloaded DownloadFailed Downloading Idle InstallationFailed Installing Installed SignatureVerified SignVerified VerificationFailed"`
	RequestId *int           `json:"requestId,omitempty"`
}
