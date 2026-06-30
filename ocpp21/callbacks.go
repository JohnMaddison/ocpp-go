package ocpp21

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	NotSupported string = "NotSupported"
)

const (
	ActionNotRecognized string = "Requested Action is recognized but not supported by the receiver"
)

type callHandler struct {
	requestType  reflect.Type
	responseType reflect.Type
	callback     any
}

type OCPPError struct {
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	ErrorDetails     any    `json:"errorDetails,omitempty"`
}

type AFRRSignalCallback func(ctx *OCPPContext, request AFRRSignalRequest) (*AFRRSignalResponse, *OCPPError)
type AdjustPeriodicEventStreamCallback func(ctx *OCPPContext, request AdjustPeriodicEventStreamRequest) (*AdjustPeriodicEventStreamResponse, *OCPPError)
type AuthorizeCallback func(ctx *OCPPContext, request AuthorizeRequest) (*AuthorizeResponse, *OCPPError)
type BatterySwapCallback func(ctx *OCPPContext, request BatterySwapRequest) (*BatterySwapResponse, *OCPPError)
type BootNotificationCallback func(ctx *OCPPContext, request BootNotificationRequest) (*BootNotificationResponse, *OCPPError)
type CancelReservationCallback func(ctx *OCPPContext, request CancelReservationRequest) (*CancelReservationResponse, *OCPPError)
type CertificateSignedCallback func(ctx *OCPPContext, request CertificateSignedRequest) (*CertificateSignedResponse, *OCPPError)
type ChangeAvailabilityCallback func(ctx *OCPPContext, request ChangeAvailabilityRequest) (*ChangeAvailabilityResponse, *OCPPError)
type ChangeTransactionTariffCallback func(ctx *OCPPContext, request ChangeTransactionTariffRequest) (*ChangeTransactionTariffResponse, *OCPPError)
type ClearCacheCallback func(ctx *OCPPContext, request ClearCacheRequest) (*ClearCacheResponse, *OCPPError)
type ClearChargingProfileCallback func(ctx *OCPPContext, request ClearChargingProfileRequest) (*ClearChargingProfileResponse, *OCPPError)
type ClearDERControlCallback func(ctx *OCPPContext, request ClearDERControlRequest) (*ClearDERControlResponse, *OCPPError)
type ClearDisplayMessageCallback func(ctx *OCPPContext, request ClearDisplayMessageRequest) (*ClearDisplayMessageResponse, *OCPPError)
type ClearTariffsCallback func(ctx *OCPPContext, request ClearTariffsRequest) (*ClearTariffsResponse, *OCPPError)
type ClearVariableMonitoringCallback func(ctx *OCPPContext, request ClearVariableMonitoringRequest) (*ClearVariableMonitoringResponse, *OCPPError)
type ClearedChargingLimitCallback func(ctx *OCPPContext, request ClearedChargingLimitRequest) (*ClearedChargingLimitResponse, *OCPPError)
type ClosePeriodicEventStreamCallback func(ctx *OCPPContext, request ClosePeriodicEventStreamRequest) (*ClosePeriodicEventStreamResponse, *OCPPError)
type CostUpdatedCallback func(ctx *OCPPContext, request CostUpdatedRequest) (*CostUpdatedResponse, *OCPPError)
type CustomerInformationCallback func(ctx *OCPPContext, request CustomerInformationRequest) (*CustomerInformationResponse, *OCPPError)
type DataTransferCallback func(ctx *OCPPContext, request DataTransferRequest) (*DataTransferResponse, *OCPPError)
type DeleteCertificateCallback func(ctx *OCPPContext, request DeleteCertificateRequest) (*DeleteCertificateResponse, *OCPPError)
type FirmwareStatusNotificationCallback func(ctx *OCPPContext, request FirmwareStatusNotificationRequest) (*FirmwareStatusNotificationResponse, *OCPPError)
type Get15118EVCertificateCallback func(ctx *OCPPContext, request Get15118EVCertificateRequest) (*Get15118EVCertificateResponse, *OCPPError)
type GetBaseReportCallback func(ctx *OCPPContext, request GetBaseReportRequest) (*GetBaseReportResponse, *OCPPError)
type GetCertificateChainStatusCallback func(ctx *OCPPContext, request GetCertificateChainStatusRequest) (*GetCertificateChainStatusResponse, *OCPPError)
type GetCertificateStatusCallback func(ctx *OCPPContext, request GetCertificateStatusRequest) (*GetCertificateStatusResponse, *OCPPError)
type GetChargingProfilesCallback func(ctx *OCPPContext, request GetChargingProfilesRequest) (*GetChargingProfilesResponse, *OCPPError)
type GetCompositeScheduleCallback func(ctx *OCPPContext, request GetCompositeScheduleRequest) (*GetCompositeScheduleResponse, *OCPPError)
type GetDERControlCallback func(ctx *OCPPContext, request GetDERControlRequest) (*GetDERControlResponse, *OCPPError)
type GetDisplayMessagesCallback func(ctx *OCPPContext, request GetDisplayMessagesRequest) (*GetDisplayMessagesResponse, *OCPPError)
type GetInstalledCertificateIdsCallback func(ctx *OCPPContext, request GetInstalledCertificateIdsRequest) (*GetInstalledCertificateIdsResponse, *OCPPError)
type GetLocalListVersionCallback func(ctx *OCPPContext, request GetLocalListVersionRequest) (*GetLocalListVersionResponse, *OCPPError)
type GetLogCallback func(ctx *OCPPContext, request GetLogRequest) (*GetLogResponse, *OCPPError)
type GetMonitoringReportCallback func(ctx *OCPPContext, request GetMonitoringReportRequest) (*GetMonitoringReportResponse, *OCPPError)
type GetPeriodicEventStreamCallback func(ctx *OCPPContext, request GetPeriodicEventStreamRequest) (*GetPeriodicEventStreamResponse, *OCPPError)
type GetReportCallback func(ctx *OCPPContext, request GetReportRequest) (*GetReportResponse, *OCPPError)
type GetTariffsCallback func(ctx *OCPPContext, request GetTariffsRequest) (*GetTariffsResponse, *OCPPError)
type GetTransactionStatusCallback func(ctx *OCPPContext, request GetTransactionStatusRequest) (*GetTransactionStatusResponse, *OCPPError)
type GetVariablesCallback func(ctx *OCPPContext, request GetVariablesRequest) (*GetVariablesResponse, *OCPPError)
type HeartbeatCallback func(ctx *OCPPContext, request HeartbeatRequest) (*HeartbeatResponse, *OCPPError)
type InstallCertificateCallback func(ctx *OCPPContext, request InstallCertificateRequest) (*InstallCertificateResponse, *OCPPError)
type LogStatusNotificationCallback func(ctx *OCPPContext, request LogStatusNotificationRequest) (*LogStatusNotificationResponse, *OCPPError)
type MeterValuesCallback func(ctx *OCPPContext, request MeterValuesRequest) (*MeterValuesResponse, *OCPPError)
type NotifyAllowedEnergyTransferCallback func(ctx *OCPPContext, request NotifyAllowedEnergyTransferRequest) (*NotifyAllowedEnergyTransferResponse, *OCPPError)
type NotifyChargingLimitCallback func(ctx *OCPPContext, request NotifyChargingLimitRequest) (*NotifyChargingLimitResponse, *OCPPError)
type NotifyCustomerInformationCallback func(ctx *OCPPContext, request NotifyCustomerInformationRequest) (*NotifyCustomerInformationResponse, *OCPPError)
type NotifyDERAlarmCallback func(ctx *OCPPContext, request NotifyDERAlarmRequest) (*NotifyDERAlarmResponse, *OCPPError)
type NotifyDERStartStopCallback func(ctx *OCPPContext, request NotifyDERStartStopRequest) (*NotifyDERStartStopResponse, *OCPPError)
type NotifyDisplayMessagesCallback func(ctx *OCPPContext, request NotifyDisplayMessagesRequest) (*NotifyDisplayMessagesResponse, *OCPPError)
type NotifyEVChargingNeedsCallback func(ctx *OCPPContext, request NotifyEVChargingNeedsRequest) (*NotifyEVChargingNeedsResponse, *OCPPError)
type NotifyEVChargingScheduleCallback func(ctx *OCPPContext, request NotifyEVChargingScheduleRequest) (*NotifyEVChargingScheduleResponse, *OCPPError)
type NotifyEventCallback func(ctx *OCPPContext, request NotifyEventRequest) (*NotifyEventResponse, *OCPPError)
type NotifyMonitoringReportCallback func(ctx *OCPPContext, request NotifyMonitoringReportRequest) (*NotifyMonitoringReportResponse, *OCPPError)
type NotifyPeriodicEventStreamCallback func(ctx *OCPPContext, request NotifyPeriodicEventStreamRequest) (*NotifyPeriodicEventStreamResponse, *OCPPError)
type NotifyPriorityChargingCallback func(ctx *OCPPContext, request NotifyPriorityChargingRequest) (*NotifyPriorityChargingResponse, *OCPPError)
type NotifyReportCallback func(ctx *OCPPContext, request NotifyReportRequest) (*NotifyReportResponse, *OCPPError)
type NotifySettlementCallback func(ctx *OCPPContext, request NotifySettlementRequest) (*NotifySettlementResponse, *OCPPError)
type NotifyWebPaymentStartedCallback func(ctx *OCPPContext, request NotifyWebPaymentStartedRequest) (*NotifyWebPaymentStartedResponse, *OCPPError)
type OpenPeriodicEventStreamCallback func(ctx *OCPPContext, request OpenPeriodicEventStreamRequest) (*OpenPeriodicEventStreamResponse, *OCPPError)
type PublishFirmwareCallback func(ctx *OCPPContext, request PublishFirmwareRequest) (*PublishFirmwareResponse, *OCPPError)
type PublishFirmwareStatusNotificationCallback func(ctx *OCPPContext, request PublishFirmwareStatusNotificationRequest) (*PublishFirmwareStatusNotificationResponse, *OCPPError)
type PullDynamicScheduleUpdateCallback func(ctx *OCPPContext, request PullDynamicScheduleUpdateRequest) (*PullDynamicScheduleUpdateResponse, *OCPPError)
type ReportChargingProfilesCallback func(ctx *OCPPContext, request ReportChargingProfilesRequest) (*ReportChargingProfilesResponse, *OCPPError)
type ReportDERControlCallback func(ctx *OCPPContext, request ReportDERControlRequest) (*ReportDERControlResponse, *OCPPError)
type RequestBatterySwapCallback func(ctx *OCPPContext, request RequestBatterySwapRequest) (*RequestBatterySwapResponse, *OCPPError)
type RequestStartTransactionCallback func(ctx *OCPPContext, request RequestStartTransactionRequest) (*RequestStartTransactionResponse, *OCPPError)
type RequestStopTransactionCallback func(ctx *OCPPContext, request RequestStopTransactionRequest) (*RequestStopTransactionResponse, *OCPPError)
type ReservationStatusUpdateCallback func(ctx *OCPPContext, request ReservationStatusUpdateRequest) (*ReservationStatusUpdateResponse, *OCPPError)
type ReserveNowCallback func(ctx *OCPPContext, request ReserveNowRequest) (*ReserveNowResponse, *OCPPError)
type ResetCallback func(ctx *OCPPContext, request ResetRequest) (*ResetResponse, *OCPPError)
type SecurityEventNotificationCallback func(ctx *OCPPContext, request SecurityEventNotificationRequest) (*SecurityEventNotificationResponse, *OCPPError)
type SendLocalListCallback func(ctx *OCPPContext, request SendLocalListRequest) (*SendLocalListResponse, *OCPPError)
type SetChargingProfileCallback func(ctx *OCPPContext, request SetChargingProfileRequest) (*SetChargingProfileResponse, *OCPPError)
type SetDERControlCallback func(ctx *OCPPContext, request SetDERControlRequest) (*SetDERControlResponse, *OCPPError)
type SetDefaultTariffCallback func(ctx *OCPPContext, request SetDefaultTariffRequest) (*SetDefaultTariffResponse, *OCPPError)
type SetDisplayMessageCallback func(ctx *OCPPContext, request SetDisplayMessageRequest) (*SetDisplayMessageResponse, *OCPPError)
type SetMonitoringBaseCallback func(ctx *OCPPContext, request SetMonitoringBaseRequest) (*SetMonitoringBaseResponse, *OCPPError)
type SetMonitoringLevelCallback func(ctx *OCPPContext, request SetMonitoringLevelRequest) (*SetMonitoringLevelResponse, *OCPPError)
type SetNetworkProfileCallback func(ctx *OCPPContext, request SetNetworkProfileRequest) (*SetNetworkProfileResponse, *OCPPError)
type SetVariableMonitoringCallback func(ctx *OCPPContext, request SetVariableMonitoringRequest) (*SetVariableMonitoringResponse, *OCPPError)
type SetVariablesCallback func(ctx *OCPPContext, request SetVariablesRequest) (*SetVariablesResponse, *OCPPError)
type SignCertificateCallback func(ctx *OCPPContext, request SignCertificateRequest) (*SignCertificateResponse, *OCPPError)
type StatusNotificationCallback func(ctx *OCPPContext, request StatusNotificationRequest) (*StatusNotificationResponse, *OCPPError)
type TransactionEventCallback func(ctx *OCPPContext, request TransactionEventRequest) (*TransactionEventResponse, *OCPPError)
type TriggerMessageCallback func(ctx *OCPPContext, request TriggerMessageRequest) (*TriggerMessageResponse, *OCPPError)
type UnlockConnectorCallback func(ctx *OCPPContext, request UnlockConnectorRequest) (*UnlockConnectorResponse, *OCPPError)
type UnpublishFirmwareCallback func(ctx *OCPPContext, request UnpublishFirmwareRequest) (*UnpublishFirmwareResponse, *OCPPError)
type UpdateDynamicScheduleCallback func(ctx *OCPPContext, request UpdateDynamicScheduleRequest) (*UpdateDynamicScheduleResponse, *OCPPError)
type UpdateFirmwareCallback func(ctx *OCPPContext, request UpdateFirmwareRequest) (*UpdateFirmwareResponse, *OCPPError)
type UsePriorityChargingCallback func(ctx *OCPPContext, request UsePriorityChargingRequest) (*UsePriorityChargingResponse, *OCPPError)
type VatNumberValidationCallback func(ctx *OCPPContext, request VatNumberValidationRequest) (*VatNumberValidationResponse, *OCPPError)

type OCPPCallbacks struct {
	handlers map[Action]callHandler

	AFRRSignal                        AFRRSignalCallback
	AdjustPeriodicEventStream         AdjustPeriodicEventStreamCallback
	Authorize                         AuthorizeCallback
	BatterySwap                       BatterySwapCallback
	BootNotification                  BootNotificationCallback
	CancelReservation                 CancelReservationCallback
	CertificateSigned                 CertificateSignedCallback
	ChangeAvailability                ChangeAvailabilityCallback
	ChangeTransactionTariff           ChangeTransactionTariffCallback
	ClearCache                        ClearCacheCallback
	ClearChargingProfile              ClearChargingProfileCallback
	ClearDERControl                   ClearDERControlCallback
	ClearDisplayMessage               ClearDisplayMessageCallback
	ClearTariffs                      ClearTariffsCallback
	ClearVariableMonitoring           ClearVariableMonitoringCallback
	ClearedChargingLimit              ClearedChargingLimitCallback
	ClosePeriodicEventStream          ClosePeriodicEventStreamCallback
	CostUpdated                       CostUpdatedCallback
	CustomerInformation               CustomerInformationCallback
	DataTransfer                      DataTransferCallback
	DeleteCertificate                 DeleteCertificateCallback
	FirmwareStatusNotification        FirmwareStatusNotificationCallback
	Get15118EVCertificate             Get15118EVCertificateCallback
	GetBaseReport                     GetBaseReportCallback
	GetCertificateChainStatus         GetCertificateChainStatusCallback
	GetCertificateStatus              GetCertificateStatusCallback
	GetChargingProfiles               GetChargingProfilesCallback
	GetCompositeSchedule              GetCompositeScheduleCallback
	GetDERControl                     GetDERControlCallback
	GetDisplayMessages                GetDisplayMessagesCallback
	GetInstalledCertificateIds        GetInstalledCertificateIdsCallback
	GetLocalListVersion               GetLocalListVersionCallback
	GetLog                            GetLogCallback
	GetMonitoringReport               GetMonitoringReportCallback
	GetPeriodicEventStream            GetPeriodicEventStreamCallback
	GetReport                         GetReportCallback
	GetTariffs                        GetTariffsCallback
	GetTransactionStatus              GetTransactionStatusCallback
	GetVariables                      GetVariablesCallback
	Heartbeat                         HeartbeatCallback
	InstallCertificate                InstallCertificateCallback
	LogStatusNotification             LogStatusNotificationCallback
	MeterValues                       MeterValuesCallback
	NotifyAllowedEnergyTransfer       NotifyAllowedEnergyTransferCallback
	NotifyChargingLimit               NotifyChargingLimitCallback
	NotifyCustomerInformation         NotifyCustomerInformationCallback
	NotifyDERAlarm                    NotifyDERAlarmCallback
	NotifyDERStartStop                NotifyDERStartStopCallback
	NotifyDisplayMessages             NotifyDisplayMessagesCallback
	NotifyEVChargingNeeds             NotifyEVChargingNeedsCallback
	NotifyEVChargingSchedule          NotifyEVChargingScheduleCallback
	NotifyEvent                       NotifyEventCallback
	NotifyMonitoringReport            NotifyMonitoringReportCallback
	NotifyPeriodicEventStream         NotifyPeriodicEventStreamCallback
	NotifyPriorityCharging            NotifyPriorityChargingCallback
	NotifyReport                      NotifyReportCallback
	NotifySettlement                  NotifySettlementCallback
	NotifyWebPaymentStarted           NotifyWebPaymentStartedCallback
	OpenPeriodicEventStream           OpenPeriodicEventStreamCallback
	PublishFirmware                   PublishFirmwareCallback
	PublishFirmwareStatusNotification PublishFirmwareStatusNotificationCallback
	PullDynamicScheduleUpdate         PullDynamicScheduleUpdateCallback
	ReportChargingProfiles            ReportChargingProfilesCallback
	ReportDERControl                  ReportDERControlCallback
	RequestBatterySwap                RequestBatterySwapCallback
	RequestStartTransaction           RequestStartTransactionCallback
	RequestStopTransaction            RequestStopTransactionCallback
	ReservationStatusUpdate           ReservationStatusUpdateCallback
	ReserveNow                        ReserveNowCallback
	Reset                             ResetCallback
	SecurityEventNotification         SecurityEventNotificationCallback
	SendLocalList                     SendLocalListCallback
	SetChargingProfile                SetChargingProfileCallback
	SetDERControl                     SetDERControlCallback
	SetDefaultTariff                  SetDefaultTariffCallback
	SetDisplayMessage                 SetDisplayMessageCallback
	SetMonitoringBase                 SetMonitoringBaseCallback
	SetMonitoringLevel                SetMonitoringLevelCallback
	SetNetworkProfile                 SetNetworkProfileCallback
	SetVariableMonitoring             SetVariableMonitoringCallback
	SetVariables                      SetVariablesCallback
	SignCertificate                   SignCertificateCallback
	StatusNotification                StatusNotificationCallback
	TransactionEvent                  TransactionEventCallback
	TriggerMessage                    TriggerMessageCallback
	UnlockConnector                   UnlockConnectorCallback
	UnpublishFirmware                 UnpublishFirmwareCallback
	UpdateDynamicSchedule             UpdateDynamicScheduleCallback
	UpdateFirmware                    UpdateFirmwareCallback
	UsePriorityCharging               UsePriorityChargingCallback
	VatNumberValidation               VatNumberValidationCallback
}

func (o *OCPPCallbacks) InitHandlers() {
	o.handlers = map[Action]callHandler{
		ActionAFRRSignal:                        newCallHandler[AFRRSignalRequest, AFRRSignalResponse](o.AFRRSignal),
		ActionAdjustPeriodicEventStream:         newCallHandler[AdjustPeriodicEventStreamRequest, AdjustPeriodicEventStreamResponse](o.AdjustPeriodicEventStream),
		ActionAuthorize:                         newCallHandler[AuthorizeRequest, AuthorizeResponse](o.Authorize),
		ActionBatterySwap:                       newCallHandler[BatterySwapRequest, BatterySwapResponse](o.BatterySwap),
		ActionBootNotification:                  newCallHandler[BootNotificationRequest, BootNotificationResponse](o.BootNotification),
		ActionCancelReservation:                 newCallHandler[CancelReservationRequest, CancelReservationResponse](o.CancelReservation),
		ActionCertificateSigned:                 newCallHandler[CertificateSignedRequest, CertificateSignedResponse](o.CertificateSigned),
		ActionChangeAvailability:                newCallHandler[ChangeAvailabilityRequest, ChangeAvailabilityResponse](o.ChangeAvailability),
		ActionChangeTransactionTariff:           newCallHandler[ChangeTransactionTariffRequest, ChangeTransactionTariffResponse](o.ChangeTransactionTariff),
		ActionClearCache:                        newCallHandler[ClearCacheRequest, ClearCacheResponse](o.ClearCache),
		ActionClearChargingProfile:              newCallHandler[ClearChargingProfileRequest, ClearChargingProfileResponse](o.ClearChargingProfile),
		ActionClearDERControl:                   newCallHandler[ClearDERControlRequest, ClearDERControlResponse](o.ClearDERControl),
		ActionClearDisplayMessage:               newCallHandler[ClearDisplayMessageRequest, ClearDisplayMessageResponse](o.ClearDisplayMessage),
		ActionClearTariffs:                      newCallHandler[ClearTariffsRequest, ClearTariffsResponse](o.ClearTariffs),
		ActionClearVariableMonitoring:           newCallHandler[ClearVariableMonitoringRequest, ClearVariableMonitoringResponse](o.ClearVariableMonitoring),
		ActionClearedChargingLimit:              newCallHandler[ClearedChargingLimitRequest, ClearedChargingLimitResponse](o.ClearedChargingLimit),
		ActionClosePeriodicEventStream:          newCallHandler[ClosePeriodicEventStreamRequest, ClosePeriodicEventStreamResponse](o.ClosePeriodicEventStream),
		ActionCostUpdated:                       newCallHandler[CostUpdatedRequest, CostUpdatedResponse](o.CostUpdated),
		ActionCustomerInformation:               newCallHandler[CustomerInformationRequest, CustomerInformationResponse](o.CustomerInformation),
		ActionDataTransfer:                      newCallHandler[DataTransferRequest, DataTransferResponse](o.DataTransfer),
		ActionDeleteCertificate:                 newCallHandler[DeleteCertificateRequest, DeleteCertificateResponse](o.DeleteCertificate),
		ActionFirmwareStatusNotification:        newCallHandler[FirmwareStatusNotificationRequest, FirmwareStatusNotificationResponse](o.FirmwareStatusNotification),
		ActionGet15118EVCertificate:             newCallHandler[Get15118EVCertificateRequest, Get15118EVCertificateResponse](o.Get15118EVCertificate),
		ActionGetBaseReport:                     newCallHandler[GetBaseReportRequest, GetBaseReportResponse](o.GetBaseReport),
		ActionGetCertificateChainStatus:         newCallHandler[GetCertificateChainStatusRequest, GetCertificateChainStatusResponse](o.GetCertificateChainStatus),
		ActionGetCertificateStatus:              newCallHandler[GetCertificateStatusRequest, GetCertificateStatusResponse](o.GetCertificateStatus),
		ActionGetChargingProfiles:               newCallHandler[GetChargingProfilesRequest, GetChargingProfilesResponse](o.GetChargingProfiles),
		ActionGetCompositeSchedule:              newCallHandler[GetCompositeScheduleRequest, GetCompositeScheduleResponse](o.GetCompositeSchedule),
		ActionGetDERControl:                     newCallHandler[GetDERControlRequest, GetDERControlResponse](o.GetDERControl),
		ActionGetDisplayMessages:                newCallHandler[GetDisplayMessagesRequest, GetDisplayMessagesResponse](o.GetDisplayMessages),
		ActionGetInstalledCertificateIds:        newCallHandler[GetInstalledCertificateIdsRequest, GetInstalledCertificateIdsResponse](o.GetInstalledCertificateIds),
		ActionGetLocalListVersion:               newCallHandler[GetLocalListVersionRequest, GetLocalListVersionResponse](o.GetLocalListVersion),
		ActionGetLog:                            newCallHandler[GetLogRequest, GetLogResponse](o.GetLog),
		ActionGetMonitoringReport:               newCallHandler[GetMonitoringReportRequest, GetMonitoringReportResponse](o.GetMonitoringReport),
		ActionGetPeriodicEventStream:            newCallHandler[GetPeriodicEventStreamRequest, GetPeriodicEventStreamResponse](o.GetPeriodicEventStream),
		ActionGetReport:                         newCallHandler[GetReportRequest, GetReportResponse](o.GetReport),
		ActionGetTariffs:                        newCallHandler[GetTariffsRequest, GetTariffsResponse](o.GetTariffs),
		ActionGetTransactionStatus:              newCallHandler[GetTransactionStatusRequest, GetTransactionStatusResponse](o.GetTransactionStatus),
		ActionGetVariables:                      newCallHandler[GetVariablesRequest, GetVariablesResponse](o.GetVariables),
		ActionHeartbeat:                         newCallHandler[HeartbeatRequest, HeartbeatResponse](o.Heartbeat),
		ActionInstallCertificate:                newCallHandler[InstallCertificateRequest, InstallCertificateResponse](o.InstallCertificate),
		ActionLogStatusNotification:             newCallHandler[LogStatusNotificationRequest, LogStatusNotificationResponse](o.LogStatusNotification),
		ActionMeterValues:                       newCallHandler[MeterValuesRequest, MeterValuesResponse](o.MeterValues),
		ActionNotifyAllowedEnergyTransfer:       newCallHandler[NotifyAllowedEnergyTransferRequest, NotifyAllowedEnergyTransferResponse](o.NotifyAllowedEnergyTransfer),
		ActionNotifyChargingLimit:               newCallHandler[NotifyChargingLimitRequest, NotifyChargingLimitResponse](o.NotifyChargingLimit),
		ActionNotifyCustomerInformation:         newCallHandler[NotifyCustomerInformationRequest, NotifyCustomerInformationResponse](o.NotifyCustomerInformation),
		ActionNotifyDERAlarm:                    newCallHandler[NotifyDERAlarmRequest, NotifyDERAlarmResponse](o.NotifyDERAlarm),
		ActionNotifyDERStartStop:                newCallHandler[NotifyDERStartStopRequest, NotifyDERStartStopResponse](o.NotifyDERStartStop),
		ActionNotifyDisplayMessages:             newCallHandler[NotifyDisplayMessagesRequest, NotifyDisplayMessagesResponse](o.NotifyDisplayMessages),
		ActionNotifyEVChargingNeeds:             newCallHandler[NotifyEVChargingNeedsRequest, NotifyEVChargingNeedsResponse](o.NotifyEVChargingNeeds),
		ActionNotifyEVChargingSchedule:          newCallHandler[NotifyEVChargingScheduleRequest, NotifyEVChargingScheduleResponse](o.NotifyEVChargingSchedule),
		ActionNotifyEvent:                       newCallHandler[NotifyEventRequest, NotifyEventResponse](o.NotifyEvent),
		ActionNotifyMonitoringReport:            newCallHandler[NotifyMonitoringReportRequest, NotifyMonitoringReportResponse](o.NotifyMonitoringReport),
		ActionNotifyPeriodicEventStream:         newCallHandler[NotifyPeriodicEventStreamRequest, NotifyPeriodicEventStreamResponse](o.NotifyPeriodicEventStream),
		ActionNotifyPriorityCharging:            newCallHandler[NotifyPriorityChargingRequest, NotifyPriorityChargingResponse](o.NotifyPriorityCharging),
		ActionNotifyReport:                      newCallHandler[NotifyReportRequest, NotifyReportResponse](o.NotifyReport),
		ActionNotifySettlement:                  newCallHandler[NotifySettlementRequest, NotifySettlementResponse](o.NotifySettlement),
		ActionNotifyWebPaymentStarted:           newCallHandler[NotifyWebPaymentStartedRequest, NotifyWebPaymentStartedResponse](o.NotifyWebPaymentStarted),
		ActionOpenPeriodicEventStream:           newCallHandler[OpenPeriodicEventStreamRequest, OpenPeriodicEventStreamResponse](o.OpenPeriodicEventStream),
		ActionPublishFirmware:                   newCallHandler[PublishFirmwareRequest, PublishFirmwareResponse](o.PublishFirmware),
		ActionPublishFirmwareStatusNotification: newCallHandler[PublishFirmwareStatusNotificationRequest, PublishFirmwareStatusNotificationResponse](o.PublishFirmwareStatusNotification),
		ActionPullDynamicScheduleUpdate:         newCallHandler[PullDynamicScheduleUpdateRequest, PullDynamicScheduleUpdateResponse](o.PullDynamicScheduleUpdate),
		ActionReportChargingProfiles:            newCallHandler[ReportChargingProfilesRequest, ReportChargingProfilesResponse](o.ReportChargingProfiles),
		ActionReportDERControl:                  newCallHandler[ReportDERControlRequest, ReportDERControlResponse](o.ReportDERControl),
		ActionRequestBatterySwap:                newCallHandler[RequestBatterySwapRequest, RequestBatterySwapResponse](o.RequestBatterySwap),
		ActionRequestStartTransaction:           newCallHandler[RequestStartTransactionRequest, RequestStartTransactionResponse](o.RequestStartTransaction),
		ActionRequestStopTransaction:            newCallHandler[RequestStopTransactionRequest, RequestStopTransactionResponse](o.RequestStopTransaction),
		ActionReservationStatusUpdate:           newCallHandler[ReservationStatusUpdateRequest, ReservationStatusUpdateResponse](o.ReservationStatusUpdate),
		ActionReserveNow:                        newCallHandler[ReserveNowRequest, ReserveNowResponse](o.ReserveNow),
		ActionReset:                             newCallHandler[ResetRequest, ResetResponse](o.Reset),
		ActionSecurityEventNotification:         newCallHandler[SecurityEventNotificationRequest, SecurityEventNotificationResponse](o.SecurityEventNotification),
		ActionSendLocalList:                     newCallHandler[SendLocalListRequest, SendLocalListResponse](o.SendLocalList),
		ActionSetChargingProfile:                newCallHandler[SetChargingProfileRequest, SetChargingProfileResponse](o.SetChargingProfile),
		ActionSetDERControl:                     newCallHandler[SetDERControlRequest, SetDERControlResponse](o.SetDERControl),
		ActionSetDefaultTariff:                  newCallHandler[SetDefaultTariffRequest, SetDefaultTariffResponse](o.SetDefaultTariff),
		ActionSetDisplayMessage:                 newCallHandler[SetDisplayMessageRequest, SetDisplayMessageResponse](o.SetDisplayMessage),
		ActionSetMonitoringBase:                 newCallHandler[SetMonitoringBaseRequest, SetMonitoringBaseResponse](o.SetMonitoringBase),
		ActionSetMonitoringLevel:                newCallHandler[SetMonitoringLevelRequest, SetMonitoringLevelResponse](o.SetMonitoringLevel),
		ActionSetNetworkProfile:                 newCallHandler[SetNetworkProfileRequest, SetNetworkProfileResponse](o.SetNetworkProfile),
		ActionSetVariableMonitoring:             newCallHandler[SetVariableMonitoringRequest, SetVariableMonitoringResponse](o.SetVariableMonitoring),
		ActionSetVariables:                      newCallHandler[SetVariablesRequest, SetVariablesResponse](o.SetVariables),
		ActionSignCertificate:                   newCallHandler[SignCertificateRequest, SignCertificateResponse](o.SignCertificate),
		ActionStatusNotification:                newCallHandler[StatusNotificationRequest, StatusNotificationResponse](o.StatusNotification),
		ActionTransactionEvent:                  newCallHandler[TransactionEventRequest, TransactionEventResponse](o.TransactionEvent),
		ActionTriggerMessage:                    newCallHandler[TriggerMessageRequest, TriggerMessageResponse](o.TriggerMessage),
		ActionUnlockConnector:                   newCallHandler[UnlockConnectorRequest, UnlockConnectorResponse](o.UnlockConnector),
		ActionUnpublishFirmware:                 newCallHandler[UnpublishFirmwareRequest, UnpublishFirmwareResponse](o.UnpublishFirmware),
		ActionUpdateDynamicSchedule:             newCallHandler[UpdateDynamicScheduleRequest, UpdateDynamicScheduleResponse](o.UpdateDynamicSchedule),
		ActionUpdateFirmware:                    newCallHandler[UpdateFirmwareRequest, UpdateFirmwareResponse](o.UpdateFirmware),
		ActionUsePriorityCharging:               newCallHandler[UsePriorityChargingRequest, UsePriorityChargingResponse](o.UsePriorityCharging),
		ActionVatNumberValidation:               newCallHandler[VatNumberValidationRequest, VatNumberValidationResponse](o.VatNumberValidation),
	}
}

func newCallHandler[Req any, Resp any](callback func(*OCPPContext, Req) (*Resp, *OCPPError)) callHandler {
	var req Req
	var resp Resp
	var cb any
	if callback != nil {
		cb = callback
	}
	return callHandler{
		requestType:  reflect.TypeOf(req),
		responseType: reflect.TypeOf(resp),
		callback:     cb,
	}
}

func decodePayload[T any](payload any) (T, error) {
	var result T

	if typed, ok := payload.(T); ok {
		return typed, nil
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return result, fmt.Errorf("failed to marshal payload: %w", err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("failed to unmarshal payload into %T: %w", result, err)
	}

	return result, nil
}

func decodePayloadByType(payload any, typ reflect.Type) (any, error) {
	value := reflect.New(typ)
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}
	if err := json.Unmarshal(data, value.Interface()); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload into %s: %w", typ, err)
	}
	return value.Elem().Interface(), nil
}
