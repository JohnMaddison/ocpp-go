package ocpp16

import (
	"encoding/json"
	"fmt"
)

type callHandler struct {
	decode   func(any) (any, error)
	callback func(*OCPPContext, any) (any, *OCPPError)
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
		ActionAuthorize: {
			decode: func(p any) (any, error) { return decodePayload[AuthorizeRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.Authorize == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.Authorize(ctx, req.(AuthorizeRequest))
			},
		},
		ActionBootNotification: {
			decode: func(p any) (any, error) { return decodePayload[BootNotificationRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.BootNotification == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.BootNotification(ctx, req.(BootNotificationRequest))
			},
		},
		ActionCancelReservation: {
			decode: func(p any) (any, error) { return decodePayload[CancelReservationRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.CancelReservation == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.CancelReservation(ctx, req.(CancelReservationRequest))
			},
		},
		ActionCertificateSigned: {
			decode: func(p any) (any, error) { return decodePayload[CertificateSignedRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.CertificateSigned == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.CertificateSigned(ctx, req.(CertificateSignedRequest))
			},
		},
		ActionChangeAvailability: {
			decode: func(p any) (any, error) { return decodePayload[ChangeAvailabilityRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.ChangeAvailability == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ChangeAvailability(ctx, req.(ChangeAvailabilityRequest))
			},
		},
		ActionChangeConfiguration: {
			decode: func(p any) (any, error) { return decodePayload[ChangeConfigurationRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.ChangeConfiguration == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ChangeConfiguration(ctx, req.(ChangeConfigurationRequest))
			},
		},
		ActionClearCache: {
			decode: func(p any) (any, error) { return decodePayload[ClearCacheRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.ClearCache == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ClearCache(ctx, req.(ClearCacheRequest))
			},
		},
		ActionClearChargingProfile: {
			decode: func(p any) (any, error) { return decodePayload[ClearChargingProfileRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.ClearChargingProfile == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ClearChargingProfile(ctx, req.(ClearChargingProfileRequest))
			},
		},
		ActionDataTransfer: {
			decode: func(p any) (any, error) { return decodePayload[DataTransferRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.DataTransfer == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.DataTransfer(ctx, req.(DataTransferRequest))
			},
		},
		ActionDeleteCertificate: {
			decode: func(p any) (any, error) { return decodePayload[DeleteCertificateRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.DeleteCertificate == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.DeleteCertificate(ctx, req.(DeleteCertificateRequest))
			},
		},
		ActionDiagnosticsStatusNotification: {
			decode: func(p any) (any, error) { return decodePayload[DiagnosticsStatusNotificationRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.DiagnosticsStatusNotification == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.DiagnosticsStatusNotification(ctx, req.(DiagnosticsStatusNotificationRequest))
			},
		},
		ActionExtendedTriggerMessage: {
			decode: func(p any) (any, error) { return decodePayload[ExtendedTriggerMessageRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.ExtendedTriggerMessage == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ExtendedTriggerMessage(ctx, req.(ExtendedTriggerMessageRequest))
			},
		},
		ActionGetCompositeSchedule: {
			decode: func(p any) (any, error) { return decodePayload[GetCompositeScheduleRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.GetCompositeSchedule == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.GetCompositeSchedule(ctx, req.(GetCompositeScheduleRequest))
			},
		},
		ActionGetConfiguration: {
			decode: func(p any) (any, error) { return decodePayload[GetConfigurationRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.GetConfiguration == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.GetConfiguration(ctx, req.(GetConfigurationRequest))
			},
		},
		ActionGetDiagnostics: {
			decode: func(p any) (any, error) { return decodePayload[GetDiagnosticsRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.GetDiagnostics == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.GetDiagnostics(ctx, req.(GetDiagnosticsRequest))
			},
		},
		ActionGetInstalledCertificateIds: {
			decode: func(p any) (any, error) { return decodePayload[GetInstalledCertificateIdsRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.GetInstalledCertificates == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.GetInstalledCertificates(ctx, req.(GetInstalledCertificateIdsRequest))
			},
		},
		ActionGetLog: {
			decode: func(p any) (any, error) { return decodePayload[GetLogRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.GetLog == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.GetLog(ctx, req.(GetLogRequest))
			},
		},
		ActionHeartbeat: {
			decode: func(p any) (any, error) { return decodePayload[HeartbeatRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.Heartbeat == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.Heartbeat(ctx, req.(HeartbeatRequest))
			},
		},
		ActionInstallCertificate: {
			decode: func(p any) (any, error) { return decodePayload[InstallCertificateRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.InstallCertificate == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.InstallCertificate(ctx, req.(InstallCertificateRequest))
			},
		},
		ActionLogStatusNotification: {
			decode: func(p any) (any, error) { return decodePayload[LogStatusNotificationRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.LogStatusNotification == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.LogStatusNotification(ctx, req.(LogStatusNotificationRequest))
			},
		},
		ActionMeterValues: {
			decode: func(p any) (any, error) { return decodePayload[MeterValuesRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.MeterValues == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.MeterValues(ctx, req.(MeterValuesRequest))
			},
		},
		ActionRemoteStartTransaction: {
			decode: func(p any) (any, error) { return decodePayload[RemoteStartTransactionRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.RemoteStartTransaction == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.RemoteStartTransaction(ctx, req.(RemoteStartTransactionRequest))
			},
		},
		ActionRemoteStopTransaction: {
			decode: func(p any) (any, error) { return decodePayload[RemoteStopTransactionRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.RemoteStopTransaction == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.RemoteStopTransaction(ctx, req.(RemoteStopTransactionRequest))
			},
		},
		ActionReserveNow: {
			decode: func(p any) (any, error) { return decodePayload[ReserveNowRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.ReserveNow == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ReserveNow(ctx, req.(ReserveNowRequest))
			},
		},
		ActionReset: {
			decode: func(p any) (any, error) { return decodePayload[ResetRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.Reset == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.Reset(ctx, req.(ResetRequest))
			},
		},
		ActionSendLocalList: {
			decode: func(p any) (any, error) { return decodePayload[SendLocalListRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.SendLocalList == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.SendLocalList(ctx, req.(SendLocalListRequest))
			},
		},
		ActionSetChargingProfile: {
			decode: func(p any) (any, error) { return decodePayload[SetChargingProfileRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.SetChargingProfile == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.SetChargingProfile(ctx, req.(SetChargingProfileRequest))
			},
		},
		ActionSignCertificate: {
			decode: func(p any) (any, error) { return decodePayload[SignCertificateRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.SignCertificate == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.SignCertificate(ctx, req.(SignCertificateRequest))
			},
		},
		ActionSignedFirmwareStatusNotification: {
			decode: func(p any) (any, error) { return decodePayload[SignedFirmwareStatusNotificationRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.SignedFirmwareStatusNotification == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.SignedFirmwareStatusNotification(ctx, req.(SignedFirmwareStatusNotificationRequest))
			},
		},
		ActionSignedUpdateFirmware: {
			decode: func(p any) (any, error) { return decodePayload[SignedUpdateFirmwareRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.SignedUpdateFirmware == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.SignedUpdateFirmware(ctx, req.(SignedUpdateFirmwareRequest))
			},
		},
		ActionStartTransaction: {
			decode: func(p any) (any, error) { return decodePayload[StartTransactionRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.StartTransaction == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.StartTransaction(ctx, req.(StartTransactionRequest))
			},
		},
		ActionStatusNotification: {
			decode: func(p any) (any, error) { return decodePayload[StatusNotificationRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.StatusNotification == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.StatusNotification(ctx, req.(StatusNotificationRequest))
			},
		},
		ActionStopTransaction: {
			decode: func(p any) (any, error) { return decodePayload[StopTransactionRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.StopTransaction == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.StopTransaction(ctx, req.(StopTransactionRequest))
			},
		},
		ActionTriggerMessage: {
			decode: func(p any) (any, error) { return decodePayload[TriggerMessageRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.TriggerMessage == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.TriggerMessage(ctx, req.(TriggerMessageRequest))
			},
		},
		ActionUnlockConnector: {
			decode: func(p any) (any, error) { return decodePayload[UnlockConnectorRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.UnlockConnector == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.UnlockConnector(ctx, req.(UnlockConnectorRequest))
			},
		},
		ActionUpdateFirmware: {
			decode: func(p any) (any, error) { return decodePayload[UpdateFirmwareRequest](p) },
			callback: func(ctx *OCPPContext, req any) (any, *OCPPError) {
				if o.UpdateFirmware == nil {
					return nil, &OCPPError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.UpdateFirmware(ctx, req.(UpdateFirmwareRequest))
			},
		},
	}
}

func decodePayload[T any](payload any) (T, error) {
	var result T

	// Fast path: if it's already the right type
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
