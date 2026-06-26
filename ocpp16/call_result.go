package ocpp16

import "time"

type AuthorizeResponse struct {
	IdTagInfo IdTagInfo `json:"idTagInfo" validate:"required"`
}

type BootNotificationResponse struct {
	CurrentTime time.Time          `json:"currentTime"`
	Interval    int                `json:"interval" validate:"gte=0"`
	Status      RegistrationStatus `json:"status" validate:"required"`
}

type ChangeAvailabilityResponse struct {
	Status AvailabilityStatus `json:"status" validate:"required"`
}

type ChangeConfigurationResponse struct {
	Status ConfigurationStatus `json:"status" validate:"required"`
}

type ClearCacheResponse struct {
	Status ClearCacheStatus `json:"status" validate:"required"`
}

type ClearChargingProfileResponse struct {
	Status ClearChargingProfileStatus `json:"status" validate:"required"`
}

type DataTransferResponse struct {
	Status DataTransferStatus `json:"status" validate:"required"`
	Data   interface{}        `json:"data,omitempty"`
}

type DiagnosticsStatusNotificationResponse struct{}

type ExtendedTriggerMessageResponse struct {
	Status TriggerMessageStatus `json:"status" validate:"required,oneof=Accepted Rejected NotImplemented"`
}

type FirmwareStatusNotificationResponse struct{}

type GetConfigurationResponse struct {
	ConfigurationKey []KeyValue `json:"configurationKey,omitempty"`
	UnknownKey       []string   `json:"unknownKey,omitempty" validate:"max=50"`
}

type GetDiagnosticsResponse struct {
	FileName *string `json:"fileName,omitempty" validate:"max=255"`
}

type GetLocalListVersionResponse struct {
	ListVersion int `json:"listVersion"`
}

type HeartbeatResponse struct {
	CurrentTime time.Time `json:"currentTime"`
}

type MeterValuesResponse struct{}

type RemoteStartTransactionResponse struct {
	Status RemoteStartStopStatus `json:"status" validate:"required"`
}

type RemoteStopTransactionResponse struct {
	Status RemoteStartStopStatus `json:"status" validate:"required"`
}

type ResetResponse struct {
	Status ResetStatus `json:"status" validate:"required"`
}

type SendLocalListResponse struct {
	Status UpdateStatus `json:"status" validate:"required"`
}

type SetChargingProfileResponse struct {
	Status ChargingProfileStatus `json:"status" validate:"required"`
}

type StartTransactionResponse struct {
	IdTagInfo     IdTagInfo `json:"idTagInfo" validate:"required"`
	TransactionId int       `json:"transactionId"`
}

type StatusNotificationResponse struct{}

type StopTransactionResponse struct {
	IdTagInfo *IdTagInfo `json:"idTagInfo,omitempty"`
}

type TriggerMessageResponse struct {
	Status TriggerMessageStatus `json:"status" validate:"required"`
}

type UnlockConnectorResponse struct {
	Status UnlockStatus `json:"status" validate:"required"`
}

type UpdateFirmwareResponse struct{}

type CancelReservationResponse struct {
	Status CancelReservationStatus `json:"status" validate:"required"`
}

type ReserveNowResponse struct {
	Status ReservationStatus `json:"status" validate:"required"`
}

type GetCompositeScheduleResponse struct {
	Status           GetCompositeScheduleStatus `json:"status" validate:"required"`
	ConnectorID      *int                       `json:"connectorId,omitempty"`
	ScheduleStart    *time.Time                 `json:"scheduleStart,omitempty"`
	ChargingSchedule *ChargingSchedule          `json:"chargingSchedule,omitempty"`
}

// Security Extension Responses
type CertificateSignedResponse struct {
	Status CertificateSignedStatus `json:"status" validate:"required"`
}

type DeleteCertificateResponse struct {
	Status DeleteCertificateStatus `json:"status" validate:"required"`
}

type GetInstalledCertificateIdsResponse struct {
	Status              GetCertificateStatus  `json:"status" validate:"required"`
	CertificateHashData []CertificateHashData `json:"certificateHashData,omitempty"`
}

type InstallCertificateResponse struct {
	Status InstallCertificateStatus `json:"status" validate:"required"`
}

type LogStatusNotificationResponse struct{}

type SecurityEventNotificationResponse struct{}

type SignCertificateResponse struct {
	Status GenericStatus `json:"status" validate:"required"`
}

type SignedUpdateFirmwareResponse struct {
	Status GenericStatus `json:"status" validate:"required"`
}

type GetLogResponse struct {
	Status   LogStatus `json:"status" validate:"required"`
	Filename *string   `json:"filename,omitempty" validate:"max=255"`
}

type SignedFirmwareStatusNotificationResponse struct{}
