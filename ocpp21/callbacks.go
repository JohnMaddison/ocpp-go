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

type Error struct {
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	ErrorDetails     any    `json:"errorDetails,omitempty"`
}

type AFRRSignalCallback func(ctx *Context, request AFRRSignalRequest) (*AFRRSignalResponse, *Error)
type AdjustPeriodicEventStreamCallback func(ctx *Context, request AdjustPeriodicEventStreamRequest) (*AdjustPeriodicEventStreamResponse, *Error)
type AuthorizeCallback func(ctx *Context, request AuthorizeRequest) (*AuthorizeResponse, *Error)
type BatterySwapCallback func(ctx *Context, request BatterySwapRequest) (*BatterySwapResponse, *Error)
type BootNotificationCallback func(ctx *Context, request BootNotificationRequest) (*BootNotificationResponse, *Error)
type CancelReservationCallback func(ctx *Context, request CancelReservationRequest) (*CancelReservationResponse, *Error)
type CertificateSignedCallback func(ctx *Context, request CertificateSignedRequest) (*CertificateSignedResponse, *Error)
type ChangeAvailabilityCallback func(ctx *Context, request ChangeAvailabilityRequest) (*ChangeAvailabilityResponse, *Error)
type ChangeTransactionTariffCallback func(ctx *Context, request ChangeTransactionTariffRequest) (*ChangeTransactionTariffResponse, *Error)
type ClearCacheCallback func(ctx *Context, request ClearCacheRequest) (*ClearCacheResponse, *Error)
type ClearChargingProfileCallback func(ctx *Context, request ClearChargingProfileRequest) (*ClearChargingProfileResponse, *Error)
type ClearDERControlCallback func(ctx *Context, request ClearDERControlRequest) (*ClearDERControlResponse, *Error)
type ClearDisplayMessageCallback func(ctx *Context, request ClearDisplayMessageRequest) (*ClearDisplayMessageResponse, *Error)
type ClearTariffsCallback func(ctx *Context, request ClearTariffsRequest) (*ClearTariffsResponse, *Error)
type ClearVariableMonitoringCallback func(ctx *Context, request ClearVariableMonitoringRequest) (*ClearVariableMonitoringResponse, *Error)
type ClearedChargingLimitCallback func(ctx *Context, request ClearedChargingLimitRequest) (*ClearedChargingLimitResponse, *Error)
type ClosePeriodicEventStreamCallback func(ctx *Context, request ClosePeriodicEventStreamRequest) (*ClosePeriodicEventStreamResponse, *Error)
type CostUpdatedCallback func(ctx *Context, request CostUpdatedRequest) (*CostUpdatedResponse, *Error)
type CustomerInformationCallback func(ctx *Context, request CustomerInformationRequest) (*CustomerInformationResponse, *Error)
type DataTransferCallback func(ctx *Context, request DataTransferRequest) (*DataTransferResponse, *Error)
type DeleteCertificateCallback func(ctx *Context, request DeleteCertificateRequest) (*DeleteCertificateResponse, *Error)
type FirmwareStatusNotificationCallback func(ctx *Context, request FirmwareStatusNotificationRequest) (*FirmwareStatusNotificationResponse, *Error)
type Get15118EVCertificateCallback func(ctx *Context, request Get15118EVCertificateRequest) (*Get15118EVCertificateResponse, *Error)
type GetBaseReportCallback func(ctx *Context, request GetBaseReportRequest) (*GetBaseReportResponse, *Error)
type GetCertificateChainStatusCallback func(ctx *Context, request GetCertificateChainStatusRequest) (*GetCertificateChainStatusResponse, *Error)
type GetCertificateStatusCallback func(ctx *Context, request GetCertificateStatusRequest) (*GetCertificateStatusResponse, *Error)
type GetChargingProfilesCallback func(ctx *Context, request GetChargingProfilesRequest) (*GetChargingProfilesResponse, *Error)
type GetCompositeScheduleCallback func(ctx *Context, request GetCompositeScheduleRequest) (*GetCompositeScheduleResponse, *Error)
type GetDERControlCallback func(ctx *Context, request GetDERControlRequest) (*GetDERControlResponse, *Error)
type GetDisplayMessagesCallback func(ctx *Context, request GetDisplayMessagesRequest) (*GetDisplayMessagesResponse, *Error)
type GetInstalledCertificateIdsCallback func(ctx *Context, request GetInstalledCertificateIdsRequest) (*GetInstalledCertificateIdsResponse, *Error)
type GetLocalListVersionCallback func(ctx *Context, request GetLocalListVersionRequest) (*GetLocalListVersionResponse, *Error)
type GetLogCallback func(ctx *Context, request GetLogRequest) (*GetLogResponse, *Error)
type GetMonitoringReportCallback func(ctx *Context, request GetMonitoringReportRequest) (*GetMonitoringReportResponse, *Error)
type GetPeriodicEventStreamCallback func(ctx *Context, request GetPeriodicEventStreamRequest) (*GetPeriodicEventStreamResponse, *Error)
type GetReportCallback func(ctx *Context, request GetReportRequest) (*GetReportResponse, *Error)
type GetTariffsCallback func(ctx *Context, request GetTariffsRequest) (*GetTariffsResponse, *Error)
type GetTransactionStatusCallback func(ctx *Context, request GetTransactionStatusRequest) (*GetTransactionStatusResponse, *Error)
type GetVariablesCallback func(ctx *Context, request GetVariablesRequest) (*GetVariablesResponse, *Error)
type HeartbeatCallback func(ctx *Context, request HeartbeatRequest) (*HeartbeatResponse, *Error)
type InstallCertificateCallback func(ctx *Context, request InstallCertificateRequest) (*InstallCertificateResponse, *Error)
type LogStatusNotificationCallback func(ctx *Context, request LogStatusNotificationRequest) (*LogStatusNotificationResponse, *Error)
type MeterValuesCallback func(ctx *Context, request MeterValuesRequest) (*MeterValuesResponse, *Error)
type NotifyAllowedEnergyTransferCallback func(ctx *Context, request NotifyAllowedEnergyTransferRequest) (*NotifyAllowedEnergyTransferResponse, *Error)
type NotifyChargingLimitCallback func(ctx *Context, request NotifyChargingLimitRequest) (*NotifyChargingLimitResponse, *Error)
type NotifyCustomerInformationCallback func(ctx *Context, request NotifyCustomerInformationRequest) (*NotifyCustomerInformationResponse, *Error)
type NotifyDERAlarmCallback func(ctx *Context, request NotifyDERAlarmRequest) (*NotifyDERAlarmResponse, *Error)
type NotifyDERStartStopCallback func(ctx *Context, request NotifyDERStartStopRequest) (*NotifyDERStartStopResponse, *Error)
type NotifyDisplayMessagesCallback func(ctx *Context, request NotifyDisplayMessagesRequest) (*NotifyDisplayMessagesResponse, *Error)
type NotifyEVChargingNeedsCallback func(ctx *Context, request NotifyEVChargingNeedsRequest) (*NotifyEVChargingNeedsResponse, *Error)
type NotifyEVChargingScheduleCallback func(ctx *Context, request NotifyEVChargingScheduleRequest) (*NotifyEVChargingScheduleResponse, *Error)
type NotifyEventCallback func(ctx *Context, request NotifyEventRequest) (*NotifyEventResponse, *Error)
type NotifyMonitoringReportCallback func(ctx *Context, request NotifyMonitoringReportRequest) (*NotifyMonitoringReportResponse, *Error)
type NotifyPeriodicEventStreamCallback func(ctx *Context, request NotifyPeriodicEventStreamRequest) (*NotifyPeriodicEventStreamResponse, *Error)
type NotifyPriorityChargingCallback func(ctx *Context, request NotifyPriorityChargingRequest) (*NotifyPriorityChargingResponse, *Error)
type NotifyReportCallback func(ctx *Context, request NotifyReportRequest) (*NotifyReportResponse, *Error)
type NotifySettlementCallback func(ctx *Context, request NotifySettlementRequest) (*NotifySettlementResponse, *Error)
type NotifyWebPaymentStartedCallback func(ctx *Context, request NotifyWebPaymentStartedRequest) (*NotifyWebPaymentStartedResponse, *Error)
type OpenPeriodicEventStreamCallback func(ctx *Context, request OpenPeriodicEventStreamRequest) (*OpenPeriodicEventStreamResponse, *Error)
type PublishFirmwareCallback func(ctx *Context, request PublishFirmwareRequest) (*PublishFirmwareResponse, *Error)
type PublishFirmwareStatusNotificationCallback func(ctx *Context, request PublishFirmwareStatusNotificationRequest) (*PublishFirmwareStatusNotificationResponse, *Error)
type PullDynamicScheduleUpdateCallback func(ctx *Context, request PullDynamicScheduleUpdateRequest) (*PullDynamicScheduleUpdateResponse, *Error)
type ReportChargingProfilesCallback func(ctx *Context, request ReportChargingProfilesRequest) (*ReportChargingProfilesResponse, *Error)
type ReportDERControlCallback func(ctx *Context, request ReportDERControlRequest) (*ReportDERControlResponse, *Error)
type RequestBatterySwapCallback func(ctx *Context, request RequestBatterySwapRequest) (*RequestBatterySwapResponse, *Error)
type RequestStartTransactionCallback func(ctx *Context, request RequestStartTransactionRequest) (*RequestStartTransactionResponse, *Error)
type RequestStopTransactionCallback func(ctx *Context, request RequestStopTransactionRequest) (*RequestStopTransactionResponse, *Error)
type ReservationStatusUpdateCallback func(ctx *Context, request ReservationStatusUpdateRequest) (*ReservationStatusUpdateResponse, *Error)
type ReserveNowCallback func(ctx *Context, request ReserveNowRequest) (*ReserveNowResponse, *Error)
type ResetCallback func(ctx *Context, request ResetRequest) (*ResetResponse, *Error)
type SecurityEventNotificationCallback func(ctx *Context, request SecurityEventNotificationRequest) (*SecurityEventNotificationResponse, *Error)
type SendLocalListCallback func(ctx *Context, request SendLocalListRequest) (*SendLocalListResponse, *Error)
type SetChargingProfileCallback func(ctx *Context, request SetChargingProfileRequest) (*SetChargingProfileResponse, *Error)
type SetDERControlCallback func(ctx *Context, request SetDERControlRequest) (*SetDERControlResponse, *Error)
type SetDefaultTariffCallback func(ctx *Context, request SetDefaultTariffRequest) (*SetDefaultTariffResponse, *Error)
type SetDisplayMessageCallback func(ctx *Context, request SetDisplayMessageRequest) (*SetDisplayMessageResponse, *Error)
type SetMonitoringBaseCallback func(ctx *Context, request SetMonitoringBaseRequest) (*SetMonitoringBaseResponse, *Error)
type SetMonitoringLevelCallback func(ctx *Context, request SetMonitoringLevelRequest) (*SetMonitoringLevelResponse, *Error)
type SetNetworkProfileCallback func(ctx *Context, request SetNetworkProfileRequest) (*SetNetworkProfileResponse, *Error)
type SetVariableMonitoringCallback func(ctx *Context, request SetVariableMonitoringRequest) (*SetVariableMonitoringResponse, *Error)
type SetVariablesCallback func(ctx *Context, request SetVariablesRequest) (*SetVariablesResponse, *Error)
type SignCertificateCallback func(ctx *Context, request SignCertificateRequest) (*SignCertificateResponse, *Error)
type StatusNotificationCallback func(ctx *Context, request StatusNotificationRequest) (*StatusNotificationResponse, *Error)
type TransactionEventCallback func(ctx *Context, request TransactionEventRequest) (*TransactionEventResponse, *Error)
type TriggerMessageCallback func(ctx *Context, request TriggerMessageRequest) (*TriggerMessageResponse, *Error)
type UnlockConnectorCallback func(ctx *Context, request UnlockConnectorRequest) (*UnlockConnectorResponse, *Error)
type UnpublishFirmwareCallback func(ctx *Context, request UnpublishFirmwareRequest) (*UnpublishFirmwareResponse, *Error)
type UpdateDynamicScheduleCallback func(ctx *Context, request UpdateDynamicScheduleRequest) (*UpdateDynamicScheduleResponse, *Error)
type UpdateFirmwareCallback func(ctx *Context, request UpdateFirmwareRequest) (*UpdateFirmwareResponse, *Error)
type UsePriorityChargingCallback func(ctx *Context, request UsePriorityChargingRequest) (*UsePriorityChargingResponse, *Error)
type VatNumberValidationCallback func(ctx *Context, request VatNumberValidationRequest) (*VatNumberValidationResponse, *Error)

type Callbacks struct {
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

func (o *Callbacks) InitHandlers() {
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

func newCallHandler[Req any, Resp any](callback func(*Context, Req) (*Resp, *Error)) callHandler {
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
