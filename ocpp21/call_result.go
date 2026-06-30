// Package ocpp21 contains the data structures for OCPP 2.1 messages.
package ocpp21

import "time"

type AFRRSignalResponse struct {
	CustomData CustomData        `json:"customData,omitempty"`
	Status     GenericStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo       `json:"statusInfo,omitempty"`
}

type AdjustPeriodicEventStreamResponse struct {
	CustomData CustomData        `json:"customData,omitempty"`
	Status     GenericStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo       `json:"statusInfo,omitempty"`
}

type AuthorizeResponse struct {
	AllowedEnergyTransfer []EnergyTransferModeEnum        `json:"allowedEnergyTransfer,omitempty" validate:"min=1"`
	CertificateStatus     *AuthorizeCertificateStatusEnum `json:"certificateStatus,omitempty"`
	CustomData            CustomData                      `json:"customData,omitempty"`
	IDTokenInfo           IDTokenInfo                     `json:"idTokenInfo" validate:"required"`
	Tariff                *Tariff                         `json:"tariff,omitempty"`
}

type BatterySwapResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type BootNotificationResponse struct {
	CurrentTime time.Time              `json:"currentTime" validate:"required"`
	CustomData  CustomData             `json:"customData,omitempty"`
	Interval    int                    `json:"interval" validate:"required"`
	Status      RegistrationStatusEnum `json:"status" validate:"required"`
	StatusInfo  *StatusInfo            `json:"statusInfo,omitempty"`
}

type CancelReservationResponse struct {
	CustomData CustomData                  `json:"customData,omitempty"`
	Status     CancelReservationStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                 `json:"statusInfo,omitempty"`
}

type CertificateSignedResponse struct {
	CustomData CustomData                  `json:"customData,omitempty"`
	Status     CertificateSignedStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                 `json:"statusInfo,omitempty"`
}

type ChangeAvailabilityResponse struct {
	CustomData CustomData                   `json:"customData,omitempty"`
	Status     ChangeAvailabilityStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                  `json:"statusInfo,omitempty"`
}

type ChangeTransactionTariffResponse struct {
	CustomData CustomData             `json:"customData,omitempty"`
	Status     TariffChangeStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo            `json:"statusInfo,omitempty"`
}

type ClearCacheResponse struct {
	CustomData CustomData           `json:"customData,omitempty"`
	Status     ClearCacheStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo          `json:"statusInfo,omitempty"`
}

type ClearChargingProfileResponse struct {
	CustomData CustomData                     `json:"customData,omitempty"`
	Status     ClearChargingProfileStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                    `json:"statusInfo,omitempty"`
}

type ClearDERControlResponse struct {
	CustomData CustomData           `json:"customData,omitempty"`
	Status     DERControlStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo          `json:"statusInfo,omitempty"`
}

type ClearDisplayMessageResponse struct {
	CustomData CustomData             `json:"customData,omitempty"`
	Status     ClearMessageStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo            `json:"statusInfo,omitempty"`
}

type ClearTariffsResponse struct {
	ClearTariffsResult []ClearTariffsResult `json:"clearTariffsResult" validate:"required,min=1"`
	CustomData         CustomData           `json:"customData,omitempty"`
}

type ClearVariableMonitoringResponse struct {
	ClearMonitoringResult []ClearMonitoringResult `json:"clearMonitoringResult" validate:"required,min=1"`
	CustomData            CustomData              `json:"customData,omitempty"`
}

type ClearedChargingLimitResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type ClosePeriodicEventStreamResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type CostUpdatedResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type CustomerInformationResponse struct {
	CustomData CustomData                    `json:"customData,omitempty"`
	Status     CustomerInformationStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                   `json:"statusInfo,omitempty"`
}

type DataTransferResponse struct {
	CustomData CustomData             `json:"customData,omitempty"`
	Data       any                    `json:"data,omitempty"`
	Status     DataTransferStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo            `json:"statusInfo,omitempty"`
}

type DeleteCertificateResponse struct {
	CustomData CustomData                  `json:"customData,omitempty"`
	Status     DeleteCertificateStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                 `json:"statusInfo,omitempty"`
}

type FirmwareStatusNotificationResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type Get15118EVCertificateResponse struct {
	CustomData         CustomData                      `json:"customData,omitempty"`
	ExiResponse        string                          `json:"exiResponse" validate:"required,max=17000"`
	RemainingContracts *int                            `json:"remainingContracts,omitempty" validate:"gte=0.0"`
	Status             Iso15118EVCertificateStatusEnum `json:"status" validate:"required"`
	StatusInfo         *StatusInfo                     `json:"statusInfo,omitempty"`
}

type GetBaseReportResponse struct {
	CustomData CustomData                   `json:"customData,omitempty"`
	Status     GenericDeviceModelStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                  `json:"statusInfo,omitempty"`
}

type GetCertificateChainStatusResponse struct {
	CertificateStatus []CertificateStatus `json:"certificateStatus" validate:"required,min=1,max=4"`
	CustomData        CustomData          `json:"customData,omitempty"`
}

type GetCertificateStatusResponse struct {
	CustomData CustomData               `json:"customData,omitempty"`
	OCSPResult *string                  `json:"ocspResult,omitempty" validate:"max=18000"`
	Status     GetCertificateStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type GetChargingProfilesResponse struct {
	CustomData CustomData                   `json:"customData,omitempty"`
	Status     GetChargingProfileStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                  `json:"statusInfo,omitempty"`
}

type GetCompositeScheduleResponse struct {
	CustomData CustomData         `json:"customData,omitempty"`
	Schedule   *CompositeSchedule `json:"schedule,omitempty"`
	Status     GenericStatusEnum  `json:"status" validate:"required"`
	StatusInfo *StatusInfo        `json:"statusInfo,omitempty"`
}

type GetDERControlResponse struct {
	CustomData CustomData           `json:"customData,omitempty"`
	Status     DERControlStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo          `json:"statusInfo,omitempty"`
}

type GetDisplayMessagesResponse struct {
	CustomData CustomData                   `json:"customData,omitempty"`
	Status     GetDisplayMessagesStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                  `json:"statusInfo,omitempty"`
}

type GetInstalledCertificateIdsResponse struct {
	CertificateHashDataChain []CertificateHashDataChain        `json:"certificateHashDataChain,omitempty" validate:"min=1"`
	CustomData               CustomData                        `json:"customData,omitempty"`
	Status                   GetInstalledCertificateStatusEnum `json:"status" validate:"required"`
	StatusInfo               *StatusInfo                       `json:"statusInfo,omitempty"`
}

type GetLocalListVersionResponse struct {
	CustomData    CustomData `json:"customData,omitempty"`
	VersionNumber int        `json:"versionNumber" validate:"required"`
}

type GetLogResponse struct {
	CustomData CustomData    `json:"customData,omitempty"`
	Filename   *string       `json:"filename,omitempty" validate:"max=255"`
	Status     LogStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo   `json:"statusInfo,omitempty"`
}

type GetMonitoringReportResponse struct {
	CustomData CustomData                   `json:"customData,omitempty"`
	Status     GenericDeviceModelStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                  `json:"statusInfo,omitempty"`
}

type GetPeriodicEventStreamResponse struct {
	ConstantStreamData []ConstantStreamData `json:"constantStreamData,omitempty" validate:"min=1"`
	CustomData         CustomData           `json:"customData,omitempty"`
}

type GetReportResponse struct {
	CustomData CustomData                   `json:"customData,omitempty"`
	Status     GenericDeviceModelStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                  `json:"statusInfo,omitempty"`
}

type GetTariffsResponse struct {
	CustomData        CustomData          `json:"customData,omitempty"`
	Status            TariffGetStatusEnum `json:"status" validate:"required"`
	StatusInfo        *StatusInfo         `json:"statusInfo,omitempty"`
	TariffAssignments []TariffAssignment  `json:"tariffAssignments,omitempty" validate:"min=1"`
}

type GetTransactionStatusResponse struct {
	CustomData       CustomData `json:"customData,omitempty"`
	MessagesInQueue  bool       `json:"messagesInQueue" validate:"required"`
	OngoingIndicator *bool      `json:"ongoingIndicator,omitempty"`
}

type GetVariablesResponse struct {
	CustomData        CustomData          `json:"customData,omitempty"`
	GetVariableResult []GetVariableResult `json:"getVariableResult" validate:"required,min=1"`
}

type HeartbeatResponse struct {
	CurrentTime time.Time  `json:"currentTime" validate:"required"`
	CustomData  CustomData `json:"customData,omitempty"`
}

type InstallCertificateResponse struct {
	CustomData CustomData                   `json:"customData,omitempty"`
	Status     InstallCertificateStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                  `json:"statusInfo,omitempty"`
}

type LogStatusNotificationResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type MeterValuesResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type NotifyAllowedEnergyTransferResponse struct {
	CustomData CustomData                            `json:"customData,omitempty"`
	Status     NotifyAllowedEnergyTransferStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                           `json:"statusInfo,omitempty"`
}

type NotifyChargingLimitResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type NotifyCustomerInformationResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type NotifyDERAlarmResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type NotifyDERStartStopResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type NotifyDisplayMessagesResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsResponse struct {
	CustomData CustomData                      `json:"customData,omitempty"`
	Status     NotifyEVChargingNeedsStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                     `json:"statusInfo,omitempty"`
}

type NotifyEVChargingScheduleResponse struct {
	CustomData CustomData        `json:"customData,omitempty"`
	Status     GenericStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo       `json:"statusInfo,omitempty"`
}

type NotifyEventResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type NotifyMonitoringReportResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type NotifyPeriodicEventStreamResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type NotifyPriorityChargingResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type NotifyReportResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type NotifySettlementResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
	ReceiptID  *string    `json:"receiptId,omitempty" validate:"max=50"`
	ReceiptURL *string    `json:"receiptUrl,omitempty" validate:"max=2000"`
}

type NotifyWebPaymentStartedResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type OpenPeriodicEventStreamResponse struct {
	CustomData CustomData        `json:"customData,omitempty"`
	Status     GenericStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo       `json:"statusInfo,omitempty"`
}

type PublishFirmwareResponse struct {
	CustomData CustomData        `json:"customData,omitempty"`
	Status     GenericStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo       `json:"statusInfo,omitempty"`
}

type PublishFirmwareStatusNotificationResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type PullDynamicScheduleUpdateResponse struct {
	CustomData     CustomData                `json:"customData,omitempty"`
	ScheduleUpdate *ChargingScheduleUpdate   `json:"scheduleUpdate,omitempty"`
	Status         ChargingProfileStatusEnum `json:"status" validate:"required"`
	StatusInfo     *StatusInfo               `json:"statusInfo,omitempty"`
}

type ReportChargingProfilesResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type ReportDERControlResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type RequestBatterySwapResponse struct {
	CustomData CustomData        `json:"customData,omitempty"`
	Status     GenericStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo       `json:"statusInfo,omitempty"`
}

type RequestStartTransactionResponse struct {
	CustomData    CustomData                 `json:"customData,omitempty"`
	Status        RequestStartStopStatusEnum `json:"status" validate:"required"`
	StatusInfo    *StatusInfo                `json:"statusInfo,omitempty"`
	TransactionID *string                    `json:"transactionId,omitempty" validate:"max=36"`
}

type RequestStopTransactionResponse struct {
	CustomData CustomData                 `json:"customData,omitempty"`
	Status     RequestStartStopStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                `json:"statusInfo,omitempty"`
}

type ReservationStatusUpdateResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type ReserveNowResponse struct {
	CustomData CustomData           `json:"customData,omitempty"`
	Status     ReserveNowStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo          `json:"statusInfo,omitempty"`
}

type ResetResponse struct {
	CustomData CustomData      `json:"customData,omitempty"`
	Status     ResetStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo     `json:"statusInfo,omitempty"`
}

type SecurityEventNotificationResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type SendLocalListResponse struct {
	CustomData CustomData              `json:"customData,omitempty"`
	Status     SendLocalListStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo             `json:"statusInfo,omitempty"`
}

type SetChargingProfileResponse struct {
	CustomData CustomData                `json:"customData,omitempty"`
	Status     ChargingProfileStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo               `json:"statusInfo,omitempty"`
}

type SetDERControlResponse struct {
	CustomData    CustomData           `json:"customData,omitempty"`
	Status        DERControlStatusEnum `json:"status" validate:"required"`
	StatusInfo    *StatusInfo          `json:"statusInfo,omitempty"`
	SupersededIds []string             `json:"supersededIds,omitempty" validate:"min=1,max=24"`
}

type SetDefaultTariffResponse struct {
	CustomData CustomData          `json:"customData,omitempty"`
	Status     TariffSetStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo         `json:"statusInfo,omitempty"`
}

type SetDisplayMessageResponse struct {
	CustomData CustomData               `json:"customData,omitempty"`
	Status     DisplayMessageStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type SetMonitoringBaseResponse struct {
	CustomData CustomData                   `json:"customData,omitempty"`
	Status     GenericDeviceModelStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                  `json:"statusInfo,omitempty"`
}

type SetMonitoringLevelResponse struct {
	CustomData CustomData        `json:"customData,omitempty"`
	Status     GenericStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo       `json:"statusInfo,omitempty"`
}

type SetNetworkProfileResponse struct {
	CustomData CustomData                  `json:"customData,omitempty"`
	Status     SetNetworkProfileStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                 `json:"statusInfo,omitempty"`
}

type SetVariableMonitoringResponse struct {
	CustomData          CustomData            `json:"customData,omitempty"`
	SetMonitoringResult []SetMonitoringResult `json:"setMonitoringResult" validate:"required,min=1"`
}

type SetVariablesResponse struct {
	CustomData        CustomData          `json:"customData,omitempty"`
	SetVariableResult []SetVariableResult `json:"setVariableResult" validate:"required,min=1"`
}

type SignCertificateResponse struct {
	CustomData CustomData        `json:"customData,omitempty"`
	Status     GenericStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo       `json:"statusInfo,omitempty"`
}

type StatusNotificationResponse struct {
	CustomData CustomData `json:"customData,omitempty"`
}

type TransactionEventResponse struct {
	ChargingPriority            *int              `json:"chargingPriority,omitempty"`
	CustomData                  CustomData        `json:"customData,omitempty"`
	IDTokenInfo                 *IDTokenInfo      `json:"idTokenInfo,omitempty"`
	TotalCost                   *float64          `json:"totalCost,omitempty"`
	TransactionLimit            *TransactionLimit `json:"transactionLimit,omitempty"`
	UpdatedPersonalMessage      *MessageContent   `json:"updatedPersonalMessage,omitempty"`
	UpdatedPersonalMessageExtra []MessageContent  `json:"updatedPersonalMessageExtra,omitempty" validate:"min=1,max=4"`
}

type TriggerMessageResponse struct {
	CustomData CustomData               `json:"customData,omitempty"`
	Status     TriggerMessageStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type UnlockConnectorResponse struct {
	CustomData CustomData       `json:"customData,omitempty"`
	Status     UnlockStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo      `json:"statusInfo,omitempty"`
}

type UnpublishFirmwareResponse struct {
	CustomData CustomData                  `json:"customData,omitempty"`
	Status     UnpublishFirmwareStatusEnum `json:"status" validate:"required"`
}

type UpdateDynamicScheduleResponse struct {
	CustomData CustomData                `json:"customData,omitempty"`
	Status     ChargingProfileStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo               `json:"statusInfo,omitempty"`
}

type UpdateFirmwareResponse struct {
	CustomData CustomData               `json:"customData,omitempty"`
	Status     UpdateFirmwareStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo              `json:"statusInfo,omitempty"`
}

type UsePriorityChargingResponse struct {
	CustomData CustomData                 `json:"customData,omitempty"`
	Status     PriorityChargingStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo                `json:"statusInfo,omitempty"`
}

type VatNumberValidationResponse struct {
	Company    *Address          `json:"company,omitempty"`
	CustomData CustomData        `json:"customData,omitempty"`
	EVSEID     *int              `json:"evseId,omitempty" validate:"gte=0.0"`
	Status     GenericStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo       `json:"statusInfo,omitempty"`
	VatNumber  string            `json:"vatNumber" validate:"required,max=20"`
}
