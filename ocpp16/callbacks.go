package ocpp16

import (
	"encoding/json"
	"fmt"
	"reflect"
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

type AuthorizeCallback func(ctx *Context, request AuthorizeRequest) (*AuthorizeResponse, *Error)
type BootNotificationCallback func(ctx *Context, request BootNotificationRequest) (*BootNotificationResponse, *Error)
type CancelReservationCallback func(ctx *Context, request CancelReservationRequest) (*CancelReservationResponse, *Error)
type CertificateSignedCallback func(ctx *Context, request CertificateSignedRequest) (*CertificateSignedResponse, *Error)
type ChangeAvailabilityCallback func(ctx *Context, request ChangeAvailabilityRequest) (*ChangeAvailabilityResponse, *Error)
type ChangeConfigurationCallback func(ctx *Context, request ChangeConfigurationRequest) (*ChangeConfigurationResponse, *Error)
type ClearCacheCallback func(ctx *Context, request ClearCacheRequest) (*ClearCacheResponse, *Error)
type ClearChargingProfileCallback func(ctx *Context, request ClearChargingProfileRequest) (*ClearChargingProfileResponse, *Error)
type ClearReservationCallback func(ctx *Context, request CancelReservationRequest) (*CancelReservationResponse, *Error)
type DataTransferCallback func(ctx *Context, request DataTransferRequest) (*DataTransferResponse, *Error)
type DeleteCertificateCallback func(ctx *Context, request DeleteCertificateRequest) (*DeleteCertificateResponse, *Error)
type DiagnosticsStatusNotificationCallback func(ctx *Context, request DiagnosticsStatusNotificationRequest) (*DiagnosticsStatusNotificationResponse, *Error)
type ExtendedTriggerMessageCallback func(ctx *Context, request ExtendedTriggerMessageRequest) (*ExtendedTriggerMessageResponse, *Error)
type GetCompositeScheduleCallback func(ctx *Context, request GetCompositeScheduleRequest) (*GetCompositeScheduleResponse, *Error)
type GetConfigurationCallback func(ctx *Context, request GetConfigurationRequest) (*GetConfigurationResponse, *Error)
type GetDiagnosticsCallback func(ctx *Context, request GetDiagnosticsRequest) (*GetDiagnosticsResponse, *Error)
type GetInstalledCertificatesCallback func(ctx *Context, request GetInstalledCertificateIdsRequest) (*GetInstalledCertificateIdsResponse, *Error)
type GetLogCallback func(ctx *Context, request GetLogRequest) (*GetLogResponse, *Error)
type HeartbeatCallback func(ctx *Context, request HeartbeatRequest) (*HeartbeatResponse, *Error)
type InstallCertificateCallback func(ctx *Context, request InstallCertificateRequest) (*InstallCertificateResponse, *Error)
type LogStatusNotificationCallback func(ctx *Context, request LogStatusNotificationRequest) (*LogStatusNotificationResponse, *Error)
type MeterValuesCallback func(ctx *Context, request MeterValuesRequest) (*MeterValuesResponse, *Error)
type RemoteStartTransactionCallback func(ctx *Context, request RemoteStartTransactionRequest) (*RemoteStartTransactionResponse, *Error)
type RemoteStopTransactionCallback func(ctx *Context, request RemoteStopTransactionRequest) (*RemoteStopTransactionResponse, *Error)
type ReserveNowCallback func(ctx *Context, request ReserveNowRequest) (*ReserveNowResponse, *Error)
type ResetCallback func(ctx *Context, request ResetRequest) (*ResetResponse, *Error)
type SendLocalListCallback func(ctx *Context, request SendLocalListRequest) (*SendLocalListResponse, *Error)
type SetChargingProfileCallback func(ctx *Context, request SetChargingProfileRequest) (*SetChargingProfileResponse, *Error)
type SignCertificateCallback func(ctx *Context, request SignCertificateRequest) (*SignCertificateResponse, *Error)
type SignedFirmwareStatusNotificationCallback func(ctx *Context, request SignedFirmwareStatusNotificationRequest) (*SignedFirmwareStatusNotificationResponse, *Error)
type SignedUpdateFirmwareCallback func(ctx *Context, request SignedUpdateFirmwareRequest) (*SignedUpdateFirmwareResponse, *Error)
type StartTransactionCallback func(ctx *Context, request StartTransactionRequest) (*StartTransactionResponse, *Error)
type StatusNotificationCallback func(ctx *Context, request StatusNotificationRequest) (*StatusNotificationResponse, *Error)
type StopTransactionCallback func(ctx *Context, request StopTransactionRequest) (*StopTransactionResponse, *Error)
type TriggerMessageCallback func(ctx *Context, request TriggerMessageRequest) (*TriggerMessageResponse, *Error)
type UnlockConnectorCallback func(ctx *Context, request UnlockConnectorRequest) (*UnlockConnectorResponse, *Error)
type UpdateFirmwareCallback func(ctx *Context, request UpdateFirmwareRequest) (*UpdateFirmwareResponse, *Error)

type Callbacks struct {
	handlers map[Action]callHandler

	Authorize                        AuthorizeCallback
	BootNotification                 BootNotificationCallback
	CancelReservation                CancelReservationCallback
	CertificateSigned                CertificateSignedCallback
	ChangeAvailability               ChangeAvailabilityCallback
	ChangeConfiguration              ChangeConfigurationCallback
	ClearCache                       ClearCacheCallback
	ClearChargingProfile             ClearChargingProfileCallback
	ClearReservation                 ClearReservationCallback
	DataTransfer                     DataTransferCallback
	DeleteCertificate                DeleteCertificateCallback
	DiagnosticsStatusNotification    DiagnosticsStatusNotificationCallback
	ExtendedTriggerMessage           ExtendedTriggerMessageCallback
	GetCompositeSchedule             GetCompositeScheduleCallback
	GetConfiguration                 GetConfigurationCallback
	GetDiagnostics                   GetDiagnosticsCallback
	GetInstalledCertificates         GetInstalledCertificatesCallback
	GetLog                           GetLogCallback
	Heartbeat                        HeartbeatCallback
	InstallCertificate               InstallCertificateCallback
	LogStatusNotification            LogStatusNotificationCallback
	MeterValues                      MeterValuesCallback
	RemoteStartTransaction           RemoteStartTransactionCallback
	RemoteStopTransaction            RemoteStopTransactionCallback
	ReserveNow                       ReserveNowCallback
	Reset                            ResetCallback
	SendLocalList                    SendLocalListCallback
	SetChargingProfile               SetChargingProfileCallback
	SignCertificate                  SignCertificateCallback
	SignedFirmwareStatusNotification SignedFirmwareStatusNotificationCallback
	SignedUpdateFirmware             SignedUpdateFirmwareCallback
	StartTransaction                 StartTransactionCallback
	StatusNotification               StatusNotificationCallback
	StopTransaction                  StopTransactionCallback
	TriggerMessage                   TriggerMessageCallback
	UnlockConnector                  UnlockConnectorCallback
	UpdateFirmware                   UpdateFirmwareCallback
}

func (o *Callbacks) InitHandlers() {
	o.handlers = map[Action]callHandler{
		ActionAuthorize:                        newCallHandler[AuthorizeRequest, AuthorizeResponse](o.Authorize),
		ActionBootNotification:                 newCallHandler[BootNotificationRequest, BootNotificationResponse](o.BootNotification),
		ActionCancelReservation:                newCallHandler[CancelReservationRequest, CancelReservationResponse](o.CancelReservation),
		ActionCertificateSigned:                newCallHandler[CertificateSignedRequest, CertificateSignedResponse](o.CertificateSigned),
		ActionChangeAvailability:               newCallHandler[ChangeAvailabilityRequest, ChangeAvailabilityResponse](o.ChangeAvailability),
		ActionChangeConfiguration:              newCallHandler[ChangeConfigurationRequest, ChangeConfigurationResponse](o.ChangeConfiguration),
		ActionClearCache:                       newCallHandler[ClearCacheRequest, ClearCacheResponse](o.ClearCache),
		ActionClearChargingProfile:             newCallHandler[ClearChargingProfileRequest, ClearChargingProfileResponse](o.ClearChargingProfile),
		ActionDataTransfer:                     newCallHandler[DataTransferRequest, DataTransferResponse](o.DataTransfer),
		ActionDeleteCertificate:                newCallHandler[DeleteCertificateRequest, DeleteCertificateResponse](o.DeleteCertificate),
		ActionDiagnosticsStatusNotification:    newCallHandler[DiagnosticsStatusNotificationRequest, DiagnosticsStatusNotificationResponse](o.DiagnosticsStatusNotification),
		ActionExtendedTriggerMessage:           newCallHandler[ExtendedTriggerMessageRequest, ExtendedTriggerMessageResponse](o.ExtendedTriggerMessage),
		ActionFirmwareStatusNotification:       newCallHandler[FirmwareStatusNotificationRequest, FirmwareStatusNotificationResponse](nil),
		ActionGetCompositeSchedule:             newCallHandler[GetCompositeScheduleRequest, GetCompositeScheduleResponse](o.GetCompositeSchedule),
		ActionGetConfiguration:                 newCallHandler[GetConfigurationRequest, GetConfigurationResponse](o.GetConfiguration),
		ActionGetDiagnostics:                   newCallHandler[GetDiagnosticsRequest, GetDiagnosticsResponse](o.GetDiagnostics),
		ActionGetInstalledCertificateIds:       newCallHandler[GetInstalledCertificateIdsRequest, GetInstalledCertificateIdsResponse](o.GetInstalledCertificates),
		ActionGetLocalListVersion:              newCallHandler[GetLocalListVersionRequest, GetLocalListVersionResponse](nil),
		ActionGetLog:                           newCallHandler[GetLogRequest, GetLogResponse](o.GetLog),
		ActionHeartbeat:                        newCallHandler[HeartbeatRequest, HeartbeatResponse](o.Heartbeat),
		ActionInstallCertificate:               newCallHandler[InstallCertificateRequest, InstallCertificateResponse](o.InstallCertificate),
		ActionLogStatusNotification:            newCallHandler[LogStatusNotificationRequest, LogStatusNotificationResponse](o.LogStatusNotification),
		ActionMeterValues:                      newCallHandler[MeterValuesRequest, MeterValuesResponse](o.MeterValues),
		ActionRemoteStartTransaction:           newCallHandler[RemoteStartTransactionRequest, RemoteStartTransactionResponse](o.RemoteStartTransaction),
		ActionRemoteStopTransaction:            newCallHandler[RemoteStopTransactionRequest, RemoteStopTransactionResponse](o.RemoteStopTransaction),
		ActionReserveNow:                       newCallHandler[ReserveNowRequest, ReserveNowResponse](o.ReserveNow),
		ActionReset:                            newCallHandler[ResetRequest, ResetResponse](o.Reset),
		ActionSecurityEventNotification:        newCallHandler[SecurityEventNotificationRequest, SecurityEventNotificationResponse](nil),
		ActionSendLocalList:                    newCallHandler[SendLocalListRequest, SendLocalListResponse](o.SendLocalList),
		ActionSetChargingProfile:               newCallHandler[SetChargingProfileRequest, SetChargingProfileResponse](o.SetChargingProfile),
		ActionSignCertificate:                  newCallHandler[SignCertificateRequest, SignCertificateResponse](o.SignCertificate),
		ActionSignedFirmwareStatusNotification: newCallHandler[SignedFirmwareStatusNotificationRequest, SignedFirmwareStatusNotificationResponse](o.SignedFirmwareStatusNotification),
		ActionSignedUpdateFirmware:             newCallHandler[SignedUpdateFirmwareRequest, SignedUpdateFirmwareResponse](o.SignedUpdateFirmware),
		ActionStartTransaction:                 newCallHandler[StartTransactionRequest, StartTransactionResponse](o.StartTransaction),
		ActionStatusNotification:               newCallHandler[StatusNotificationRequest, StatusNotificationResponse](o.StatusNotification),
		ActionStopTransaction:                  newCallHandler[StopTransactionRequest, StopTransactionResponse](o.StopTransaction),
		ActionTriggerMessage:                   newCallHandler[TriggerMessageRequest, TriggerMessageResponse](o.TriggerMessage),
		ActionUnlockConnector:                  newCallHandler[UnlockConnectorRequest, UnlockConnectorResponse](o.UnlockConnector),
		ActionUpdateFirmware:                   newCallHandler[UpdateFirmwareRequest, UpdateFirmwareResponse](o.UpdateFirmware),
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
