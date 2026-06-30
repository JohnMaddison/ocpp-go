// Package ocpp201 contains the data structures for OCPP 2.0.1 messages.
package ocpp201

import "time"

type AuthorizeResponse struct {
	IDTokenInfo       IdTokenInfo                 `json:"idTokenInfo" validate:"required"`
	CertificateStatus *AuthorizeCertificateStatus `json:"certificateStatus,omitempty"`
	EvseID            []int                       `json:"evseId,omitempty"`
}

type BootNotificationResponse struct {
	CurrentTime time.Time          `json:"currentTime" validate:"required"`
	Interval    int                `json:"interval" validate:"required"`
	Status      RegistrationStatus `json:"status" validate:"required"`
	StatusInfo  *StatusInfo        `json:"statusInfo,omitempty"`
}

type CancelReservationResponse struct {
	Status     GenericStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo   `json:"statusInfo,omitempty"`
}

type CertificateSignedResponse struct {
	Status     CertificateSignedStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo             `json:"statusInfo,omitempty"`
}

type ChangeAvailabilityResponse struct {
	Status     ChangeAvailabilityStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type ClearCacheResponse struct {
	Status     ClearCacheStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo      `json:"statusInfo,omitempty"`
}

type ClearChargingProfileResponse struct {
	Status     ClearChargingProfileStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo                `json:"statusInfo,omitempty"`
}

type ClearDisplayMessageResponse struct {
	Status     ClearDisplayMessageStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo               `json:"statusInfo,omitempty"`
}

type ClearedChargingLimitResponse struct {
	// No fields
}

type ClearVariableMonitoringResponse struct {
	ClearMonitoringResult []ClearMonitoringResult `json:"clearMonitoringResult" validate:"required,min=1"`
}

type CostUpdatedResponse struct {
	// No fields
}

type CustomerInformationResponse struct {
	Status     CustomerInformationStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo               `json:"statusInfo,omitempty"`
}

type DataTransferResponse struct {
	Status     DataTransferStatus `json:"status" validate:"required"`
	Data       any                `json:"data,omitempty"`
	StatusInfo *StatusInfo        `json:"statusInfo,omitempty"`
}

type DeleteCertificateResponse struct {
	Status     DeleteCertificateStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo             `json:"statusInfo,omitempty"`
}

type FirmwareStatusNotificationResponse struct {
	// No fields
}

type Get15118EVCertificateResponse struct {
	Status      ISO15118EVCertificateStatus `json:"status" validate:"required"`
	ExiResponse string                      `json:"exiResponse" validate:"required,max=10000"`
	StatusInfo  *StatusInfo                 `json:"statusInfo,omitempty"`
}

type GetBaseReportResponse struct {
	Status     GenericDeviceModelStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type GetCertificateStatusResponse struct {
	Status     GetCertificateStatus `json:"status" validate:"required"`
	OCSPResult *string              `json:"ocspResult,omitempty" validate:"max=5500"`
	StatusInfo *StatusInfo          `json:"statusInfo,omitempty"`
}

type GetChargingProfilesResponse struct {
	Status     GetChargingProfileStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type GetCompositeScheduleResponse struct {
	Status            GenericStatus      `json:"status" validate:"required"`
	CompositeSchedule *CompositeSchedule `json:"compositeSchedule,omitempty"`
	StatusInfo        *StatusInfo        `json:"statusInfo,omitempty"`
}

type GetDisplayMessagesResponse struct {
	Status     GenericDeviceModelStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type GetInstalledCertificateIdsResponse struct {
	Status                   GenericDeviceModelStatus   `json:"status" validate:"required"`
	CertificateHashDataChain []CertificateHashDataChain `json:"certificateHashDataChain,omitempty"`
	StatusInfo               *StatusInfo                `json:"statusInfo,omitempty"`
}

type GetLocalListVersionResponse struct {
	VersionNumber int `json:"versionNumber" validate:"required"`
}

type GetLogResponse struct {
	Status     LogStatus   `json:"status" validate:"required"`
	Filename   *string     `json:"filename,omitempty" validate:"max=255"`
	StatusInfo *StatusInfo `json:"statusInfo,omitempty"`
}

type GetMonitoringReportResponse struct {
	Status     GenericDeviceModelStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type GetReportResponse struct {
	Status     GenericDeviceModelStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type GetTransactionStatusResponse struct {
	MessagesInQueue  bool  `json:"messagesInQueue" validate:"required"`
	OngoingIndicator *bool `json:"ongoingIndicator,omitempty"`
}

type GetVariablesResponse struct {
	GetVariableResult []GetVariableResult `json:"getVariableResult" validate:"required,min=1"`
}

type GetVariableResult struct {
	AttributeStatus *GetVariableStatus `json:"attributeStatus,omitempty"`
	Component       Component          `json:"component" validate:"required"`
	Variable        Variable           `json:"variable" validate:"required"`
	AttributeValue  *string            `json:"attributeValue,omitempty" validate:"max=2500"`
	StatusInfo      *StatusInfo        `json:"statusInfo,omitempty"`
}

type HeartbeatResponse struct {
	CurrentTime time.Time `json:"currentTime" validate:"required"`
}

type InstallCertificateResponse struct {
	Status     InstallCertificateStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type LogStatusNotificationResponse struct {
	// No fields
}

type MeterValuesResponse struct {
	// No fields
}

type NotifyChargingLimitResponse struct {
	// No fields
}

type NotifyCustomerInformationResponse struct {
	// No fields
}

type NotifyDisplayMessagesResponse struct {
	// No fields
}

type NotifyEVChargingNeedsResponse struct {
	Status     NotifyEVChargingNeedsStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo                 `json:"statusInfo,omitempty"`
}

type NotifyEVChargingScheduleResponse struct {
	Status     GenericStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo   `json:"statusInfo,omitempty"`
}

type NotifyEventResponse struct {
	// No fields
}

type NotifyMonitoringReportResponse struct {
	// No fields
}

type NotifyReportResponse struct {
	// No fields
}

type PublishFirmwareResponse struct {
	Status     GenericStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo   `json:"statusInfo,omitempty"`
}

type PublishFirmwareStatusNotificationResponse struct {
	// No fields
}

type ReportChargingProfilesResponse struct {
	// No fields
}

type RequestStartTransactionResponse struct {
	Status        RequestStartStopStatus `json:"status" validate:"required"`
	TransactionID *string                `json:"transactionId,omitempty" validate:"max=36"`
	StatusInfo    *StatusInfo            `json:"statusInfo,omitempty"`
}

type RequestStopTransactionResponse struct {
	Status     RequestStartStopStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo            `json:"statusInfo,omitempty"`
}

type ReservationStatusUpdateResponse struct {
	// No fields
}

type ReserveNowResponse struct {
	Status     ReserveNowStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo      `json:"statusInfo,omitempty"`
}

type ResetResponse struct {
	Status     GenericStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo   `json:"statusInfo,omitempty"`
}

type SecurityEventNotificationResponse struct {
	// No fields
}

type SendLocalListResponse struct {
	Status     SendLocalListStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo         `json:"statusInfo,omitempty"`
}

type SetChargingProfileResponse struct {
	Status     GenericStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo   `json:"statusInfo,omitempty"`
}

type SetClockResponse struct {
	// No fields
}

type SetDisplayMessageResponse struct {
	Status     DisplayMessageStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo          `json:"statusInfo,omitempty"`
}

type SetMonitoringBaseResponse struct {
	Status     GenericDeviceModelStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type SetMonitoringLevelResponse struct {
	Status     GenericStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo   `json:"statusInfo,omitempty"`
}

type SetNetworkProfileResponse struct {
	Status     SetNetworkProfileStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo             `json:"statusInfo,omitempty"`
}

type SetVariableMonitoringResponse struct {
	SetMonitoringResult []SetMonitoringResult `json:"setMonitoringResult" validate:"required,min=1"`
}

type SetMonitoringResult struct {
	ID         int                 `json:"id" validate:"required"`
	Status     SetMonitoringStatus `json:"status" validate:"required"`
	Type       Monitor             `json:"type" validate:"required"`
	Severity   int                 `json:"severity" validate:"required"`
	Component  Component           `json:"component" validate:"required"`
	Variable   Variable            `json:"variable" validate:"required"`
	StatusInfo *StatusInfo         `json:"statusInfo,omitempty"`
}

type SetVariablesResponse struct {
	SetVariableResult []SetVariableResult `json:"setVariableResult" validate:"required,min=1"`
}

type SetVariableResult struct {
	AttributeStatus *SetVariableStatus `json:"attributeStatus,omitempty"`
	Component       Component          `json:"component" validate:"required"`
	Variable        Variable           `json:"variable" validate:"required"`
	StatusInfo      *StatusInfo        `json:"statusInfo,omitempty"`
}

type SignCertificateResponse struct {
	Status     GenericStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo   `json:"statusInfo,omitempty"`
}

type StatusNotificationResponse struct {
	// No fields
}

type TransactionEventResponse struct {
	TotalCost              *float64        `json:"totalCost,omitempty"`
	ChargingPriority       *int            `json:"chargingPriority,omitempty"`
	IDTokenInfo            *IdTokenInfo    `json:"idTokenInfo,omitempty"`
	UpdatedPersonalMessage *MessageContent `json:"updatedPersonalMessage,omitempty"`
}

type TriggerMessageResponse struct {
	Status     TriggerMessageStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo          `json:"statusInfo,omitempty"`
}

type UnlockConnectorResponse struct {
	Status     UnlockStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo  `json:"statusInfo,omitempty"`
}

type UnpublishFirmwareResponse struct {
	Status UnpublishFirmwareStatus `json:"status" validate:"required"`
}

type UpdateFirmwareResponse struct {
	Status     UpdateFirmwareStatus `json:"status" validate:"required"`
	StatusInfo *StatusInfo          `json:"statusInfo,omitempty"`
}
