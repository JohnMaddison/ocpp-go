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

type OCPPError struct {
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	ErrorDetails     any    `json:"errorDetails,omitempty"`
}

type AuthorizeCallback func(ctx *OCPPContext, request AuthorizeRequest) (*AuthorizeResponse, *OCPPError)
type BootNotificationCallback func(ctx *OCPPContext, request BootNotificationRequest) (*BootNotificationResponse, *OCPPError)
type CancelReservationCallback func(ctx *OCPPContext, request CancelReservationRequest) (*CancelReservationResponse, *OCPPError)
type CertificateSignedCallback func(ctx *OCPPContext, request CertificateSignedRequest) (*CertificateSignedResponse, *OCPPError)
type ChangeAvailabilityCallback func(ctx *OCPPContext, request ChangeAvailabilityRequest) (*ChangeAvailabilityResponse, *OCPPError)
type ChangeConfigurationCallback func(ctx *OCPPContext, request ChangeConfigurationRequest) (*ChangeConfigurationResponse, *OCPPError)
type ClearCacheCallback func(ctx *OCPPContext, request ClearCacheRequest) (*ClearCacheResponse, *OCPPError)
type ClearChargingProfileCallback func(ctx *OCPPContext, request ClearChargingProfileRequest) (*ClearChargingProfileResponse, *OCPPError)
type ClearReservationCallback func(ctx *OCPPContext, request CancelReservationRequest) (*CancelReservationResponse, *OCPPError)
type DataTransferCallback func(ctx *OCPPContext, request DataTransferRequest) (*DataTransferResponse, *OCPPError)
type DeleteCertificateCallback func(ctx *OCPPContext, request DeleteCertificateRequest) (*DeleteCertificateResponse, *OCPPError)
type DiagnosticsStatusNotificationCallback func(ctx *OCPPContext, request DiagnosticsStatusNotificationRequest) (*DiagnosticsStatusNotificationResponse, *OCPPError)
type ExtendedTriggerMessageCallback func(ctx *OCPPContext, request ExtendedTriggerMessageRequest) (*ExtendedTriggerMessageResponse, *OCPPError)
type GetCompositeScheduleCallback func(ctx *OCPPContext, request GetCompositeScheduleRequest) (*GetCompositeScheduleResponse, *OCPPError)
type GetConfigurationCallback func(ctx *OCPPContext, request GetConfigurationRequest) (*GetConfigurationResponse, *OCPPError)
type GetDiagnosticsCallback func(ctx *OCPPContext, request GetDiagnosticsRequest) (*GetDiagnosticsResponse, *OCPPError)
type GetInstalledCertificatesCallback func(ctx *OCPPContext, request GetInstalledCertificateIdsRequest) (*GetInstalledCertificateIdsResponse, *OCPPError)
type GetLogCallback func(ctx *OCPPContext, request GetLogRequest) (*GetLogResponse, *OCPPError)
type HeartbeatCallback func(ctx *OCPPContext, request HeartbeatRequest) (*HeartbeatResponse, *OCPPError)
type InstallCertificateCallback func(ctx *OCPPContext, request InstallCertificateRequest) (*InstallCertificateResponse, *OCPPError)
type LogStatusNotificationCallback func(ctx *OCPPContext, request LogStatusNotificationRequest) (*LogStatusNotificationResponse, *OCPPError)
type MeterValuesCallback func(ctx *OCPPContext, request MeterValuesRequest) (*MeterValuesResponse, *OCPPError)
type RemoteStartTransactionCallback func(ctx *OCPPContext, request RemoteStartTransactionRequest) (*RemoteStartTransactionResponse, *OCPPError)
type RemoteStopTransactionCallback func(ctx *OCPPContext, request RemoteStopTransactionRequest) (*RemoteStopTransactionResponse, *OCPPError)
type ReserveNowCallback func(ctx *OCPPContext, request ReserveNowRequest) (*ReserveNowResponse, *OCPPError)
type ResetCallback func(ctx *OCPPContext, request ResetRequest) (*ResetResponse, *OCPPError)
type SendLocalListCallback func(ctx *OCPPContext, request SendLocalListRequest) (*SendLocalListResponse, *OCPPError)
type SetChargingProfileCallback func(ctx *OCPPContext, request SetChargingProfileRequest) (*SetChargingProfileResponse, *OCPPError)
type SignCertificateCallback func(ctx *OCPPContext, request SignCertificateRequest) (*SignCertificateResponse, *OCPPError)
type SignedFirmwareStatusNotificationCallback func(ctx *OCPPContext, request SignedFirmwareStatusNotificationRequest) (*SignedFirmwareStatusNotificationResponse, *OCPPError)
type SignedUpdateFirmwareCallback func(ctx *OCPPContext, request SignedUpdateFirmwareRequest) (*SignedUpdateFirmwareResponse, *OCPPError)
type StartTransactionCallback func(ctx *OCPPContext, request StartTransactionRequest) (*StartTransactionResponse, *OCPPError)
type StatusNotificationCallback func(ctx *OCPPContext, request StatusNotificationRequest) (*StatusNotificationResponse, *OCPPError)
type StopTransactionCallback func(ctx *OCPPContext, request StopTransactionRequest) (*StopTransactionResponse, *OCPPError)
type TriggerMessageCallback func(ctx *OCPPContext, request TriggerMessageRequest) (*TriggerMessageResponse, *OCPPError)
type UnlockConnectorCallback func(ctx *OCPPContext, request UnlockConnectorRequest) (*UnlockConnectorResponse, *OCPPError)
type UpdateFirmwareCallback func(ctx *OCPPContext, request UpdateFirmwareRequest) (*UpdateFirmwareResponse, *OCPPError)

type OCPPCallbacks struct {
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

func (o *OCPPCallbacks) InitHandlers() {
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
