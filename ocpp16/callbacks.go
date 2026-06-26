package ocpp16

import (
	"encoding/json"
	"fmt"
)

type callHandler struct {
	decode   func(any) (any, error)
	callback func(*OcppContext, any) (any, *OcppError)
}

type OcppError struct {
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	ErrorDetails     any    `json:"errorDetails,omitempty"`
}

type AuthorizeCallback func(ctx *OcppContext, request AuthorizeRequest) (*AuthorizeResponse, *OcppError)
type BootNotificationCallback func(ctx *OcppContext, request BootNotificationRequest) (*BootNotificationResponse, *OcppError)
type CancelReservationCallback func(ctx *OcppContext, request CancelReservationRequest) (*CancelReservationResponse, *OcppError)
type CertificateSignedCallback func(ctx *OcppContext, request CertificateSignedRequest) (*CertificateSignedResponse, *OcppError)
type ChangeAvailabilityCallback func(ctx *OcppContext, request ChangeAvailabilityRequest) (*ChangeAvailabilityResponse, *OcppError)
type ChangeConfigurationCallback func(ctx *OcppContext, request ChangeConfigurationRequest) (*ChangeConfigurationResponse, *OcppError)
type ClearCacheCallback func(ctx *OcppContext, request ClearCacheRequest) (*ClearCacheResponse, *OcppError)
type ClearChargingProfileCallback func(ctx *OcppContext, request ClearChargingProfileRequest) (*ClearChargingProfileResponse, *OcppError)
type ClearReservationCallback func(ctx *OcppContext, request CancelReservationRequest) (*CancelReservationResponse, *OcppError)
type DataTransferCallback func(ctx *OcppContext, request DataTransferRequest) (*DataTransferResponse, *OcppError)
type DeleteCertificateCallback func(ctx *OcppContext, request DeleteCertificateRequest) (*DeleteCertificateResponse, *OcppError)
type DiagnosticsStatusNotificationCallback func(ctx *OcppContext, request DiagnosticsStatusNotificationRequest) (*DiagnosticsStatusNotificationResponse, *OcppError)
type ExtendedTriggerMessageCallback func(ctx *OcppContext, request ExtendedTriggerMessageRequest) (*ExtendedTriggerMessageResponse, *OcppError)
type GetCompositeScheduleCallback func(ctx *OcppContext, request GetCompositeScheduleRequest) (*GetCompositeScheduleResponse, *OcppError)
type GetConfigurationCallback func(ctx *OcppContext, request GetConfigurationRequest) (*GetConfigurationResponse, *OcppError)
type GetDiagnosticsCallback func(ctx *OcppContext, request GetDiagnosticsRequest) (*GetDiagnosticsResponse, *OcppError)
type GetInstalledCertificatesCallback func(ctx *OcppContext, request GetInstalledCertificateIdsRequest) (*GetInstalledCertificateIdsResponse, *OcppError)
type GetLogCallback func(ctx *OcppContext, request GetLogRequest) (*GetLogResponse, *OcppError)
type HeartbeatCallback func(ctx *OcppContext, request HeartbeatRequest) (*HeartbeatResponse, *OcppError)
type InstallCertificateCallback func(ctx *OcppContext, request InstallCertificateRequest) (*InstallCertificateResponse, *OcppError)
type LogStatusNotificationCallback func(ctx *OcppContext, request LogStatusNotificationRequest) (*LogStatusNotificationResponse, *OcppError)
type MeterValuesCallback func(ctx *OcppContext, request MeterValuesRequest) (*MeterValuesResponse, *OcppError)
type RemoteStartTransactionCallback func(ctx *OcppContext, request RemoteStartTransactionRequest) (*RemoteStartTransactionResponse, *OcppError)
type RemoteStopTransactionCallback func(ctx *OcppContext, request RemoteStopTransactionRequest) (*RemoteStopTransactionResponse, *OcppError)
type ReserveNowCallback func(ctx *OcppContext, request ReserveNowRequest) (*ReserveNowResponse, *OcppError)
type ResetCallback func(ctx *OcppContext, request ResetRequest) (*ResetResponse, *OcppError)
type SendLocalListCallback func(ctx *OcppContext, request SendLocalListRequest) (*SendLocalListResponse, *OcppError)
type SetChargingProfileCallback func(ctx *OcppContext, request SetChargingProfileRequest) (*SetChargingProfileResponse, *OcppError)
type SignCertificateCallback func(ctx *OcppContext, request SignCertificateRequest) (*SignCertificateResponse, *OcppError)
type SignedFirmwareStatusNotificationCallback func(ctx *OcppContext, request SignedFirmwareStatusNotificationRequest) (*SignedFirmwareStatusNotificationResponse, *OcppError)
type SignedUpdateFirmwareCallback func(ctx *OcppContext, request SignedUpdateFirmwareRequest) (*SignedUpdateFirmwareResponse, *OcppError)
type StartTransactionCallback func(ctx *OcppContext, request StartTransactionRequest) (*StartTransactionResponse, *OcppError)
type StatusNotificationCallback func(ctx *OcppContext, request StatusNotificationRequest) (*StatusNotificationResponse, *OcppError)
type StopTransactionCallback func(ctx *OcppContext, request StopTransactionRequest) (*StopTransactionResponse, *OcppError)
type TriggerMessageCallback func(ctx *OcppContext, request TriggerMessageRequest) (*TriggerMessageResponse, *OcppError)
type UnlockConnectorCallback func(ctx *OcppContext, request UnlockConnectorRequest) (*UnlockConnectorResponse, *OcppError)
type UpdateFirmwareCallback func(ctx *OcppContext, request UpdateFirmwareRequest) (*UpdateFirmwareResponse, *OcppError)

type OcppCallbacks struct {
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

func (o *OcppCallbacks) InitHandlers() {
	o.handlers = map[Action]callHandler{
		ActionAuthorize: {
			decode: func(p any) (any, error) { return decodePayload[AuthorizeRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.Authorize == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.Authorize(ctx, req.(AuthorizeRequest))
			},
		},
		ActionBootNotification: {
			decode: func(p any) (any, error) { return decodePayload[BootNotificationRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.BootNotification == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.BootNotification(ctx, req.(BootNotificationRequest))
			},
		},
		ActionCancelReservation: {
			decode: func(p any) (any, error) { return decodePayload[CancelReservationRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.CancelReservation == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.CancelReservation(ctx, req.(CancelReservationRequest))
			},
		},
		ActionCertificateSigned: {
			decode: func(p any) (any, error) { return decodePayload[CertificateSignedRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.CertificateSigned == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.CertificateSigned(ctx, req.(CertificateSignedRequest))
			},
		},
		ActionChangeAvailability: {
			decode: func(p any) (any, error) { return decodePayload[ChangeAvailabilityRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.ChangeAvailability == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ChangeAvailability(ctx, req.(ChangeAvailabilityRequest))
			},
		},
		ActionChangeConfiguration: {
			decode: func(p any) (any, error) { return decodePayload[ChangeConfigurationRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.ChangeConfiguration == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ChangeConfiguration(ctx, req.(ChangeConfigurationRequest))
			},
		},
		ActionClearCache: {
			decode: func(p any) (any, error) { return decodePayload[ClearCacheRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.ClearCache == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ClearCache(ctx, req.(ClearCacheRequest))
			},
		},
		ActionClearChargingProfile: {
			decode: func(p any) (any, error) { return decodePayload[ClearChargingProfileRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.ClearChargingProfile == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ClearChargingProfile(ctx, req.(ClearChargingProfileRequest))
			},
		},
		ActionDataTransfer: {
			decode: func(p any) (any, error) { return decodePayload[DataTransferRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.DataTransfer == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.DataTransfer(ctx, req.(DataTransferRequest))
			},
		},
		ActionDeleteCertificate: {
			decode: func(p any) (any, error) { return decodePayload[DeleteCertificateRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.DeleteCertificate == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.DeleteCertificate(ctx, req.(DeleteCertificateRequest))
			},
		},
		ActionDiagnosticsStatusNotification: {
			decode: func(p any) (any, error) { return decodePayload[DiagnosticsStatusNotificationRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.DiagnosticsStatusNotification == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.DiagnosticsStatusNotification(ctx, req.(DiagnosticsStatusNotificationRequest))
			},
		},
		ActionExtendedTriggerMessage: {
			decode: func(p any) (any, error) { return decodePayload[ExtendedTriggerMessageRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.ExtendedTriggerMessage == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ExtendedTriggerMessage(ctx, req.(ExtendedTriggerMessageRequest))
			},
		},
		ActionGetCompositeSchedule: {
			decode: func(p any) (any, error) { return decodePayload[GetCompositeScheduleRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.GetCompositeSchedule == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.GetCompositeSchedule(ctx, req.(GetCompositeScheduleRequest))
			},
		},
		ActionGetConfiguration: {
			decode: func(p any) (any, error) { return decodePayload[GetConfigurationRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.GetConfiguration == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.GetConfiguration(ctx, req.(GetConfigurationRequest))
			},
		},
		ActionGetDiagnostics: {
			decode: func(p any) (any, error) { return decodePayload[GetDiagnosticsRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.GetDiagnostics == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.GetDiagnostics(ctx, req.(GetDiagnosticsRequest))
			},
		},
		ActionGetInstalledCertificateIds: {
			decode: func(p any) (any, error) { return decodePayload[GetInstalledCertificateIdsRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.GetInstalledCertificates == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.GetInstalledCertificates(ctx, req.(GetInstalledCertificateIdsRequest))
			},
		},
		ActionGetLog: {
			decode: func(p any) (any, error) { return decodePayload[GetLogRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.GetLog == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.GetLog(ctx, req.(GetLogRequest))
			},
		},
		ActionHeartbeat: {
			decode: func(p any) (any, error) { return decodePayload[HeartbeatRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.Heartbeat == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.Heartbeat(ctx, req.(HeartbeatRequest))
			},
		},
		ActionInstallCertificate: {
			decode: func(p any) (any, error) { return decodePayload[InstallCertificateRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.InstallCertificate == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.InstallCertificate(ctx, req.(InstallCertificateRequest))
			},
		},
		ActionLogStatusNotification: {
			decode: func(p any) (any, error) { return decodePayload[LogStatusNotificationRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.LogStatusNotification == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.LogStatusNotification(ctx, req.(LogStatusNotificationRequest))
			},
		},
		ActionMeterValues: {
			decode: func(p any) (any, error) { return decodePayload[MeterValuesRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.MeterValues == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.MeterValues(ctx, req.(MeterValuesRequest))
			},
		},
		ActionRemoteStartTransaction: {
			decode: func(p any) (any, error) { return decodePayload[RemoteStartTransactionRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.RemoteStartTransaction == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.RemoteStartTransaction(ctx, req.(RemoteStartTransactionRequest))
			},
		},
		ActionRemoteStopTransaction: {
			decode: func(p any) (any, error) { return decodePayload[RemoteStopTransactionRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.RemoteStopTransaction == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.RemoteStopTransaction(ctx, req.(RemoteStopTransactionRequest))
			},
		},
		ActionReserveNow: {
			decode: func(p any) (any, error) { return decodePayload[ReserveNowRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.ReserveNow == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.ReserveNow(ctx, req.(ReserveNowRequest))
			},
		},
		ActionReset: {
			decode: func(p any) (any, error) { return decodePayload[ResetRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.Reset == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.Reset(ctx, req.(ResetRequest))
			},
		},
		ActionSendLocalList: {
			decode: func(p any) (any, error) { return decodePayload[SendLocalListRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.SendLocalList == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.SendLocalList(ctx, req.(SendLocalListRequest))
			},
		},
		ActionSetChargingProfile: {
			decode: func(p any) (any, error) { return decodePayload[SetChargingProfileRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.SetChargingProfile == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.SetChargingProfile(ctx, req.(SetChargingProfileRequest))
			},
		},
		ActionSignCertificate: {
			decode: func(p any) (any, error) { return decodePayload[SignCertificateRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.SignCertificate == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.SignCertificate(ctx, req.(SignCertificateRequest))
			},
		},
		ActionSignedFirmwareStatusNotification: {
			decode: func(p any) (any, error) { return decodePayload[SignedFirmwareStatusNotificationRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.SignedFirmwareStatusNotification == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.SignedFirmwareStatusNotification(ctx, req.(SignedFirmwareStatusNotificationRequest))
			},
		},
		ActionSignedUpdateFirmware: {
			decode: func(p any) (any, error) { return decodePayload[SignedUpdateFirmwareRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.SignedUpdateFirmware == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.SignedUpdateFirmware(ctx, req.(SignedUpdateFirmwareRequest))
			},
		},
		ActionStartTransaction: {
			decode: func(p any) (any, error) { return decodePayload[StartTransactionRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.StartTransaction == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.StartTransaction(ctx, req.(StartTransactionRequest))
			},
		},
		ActionStatusNotification: {
			decode: func(p any) (any, error) { return decodePayload[StatusNotificationRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.StatusNotification == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.StatusNotification(ctx, req.(StatusNotificationRequest))
			},
		},
		ActionStopTransaction: {
			decode: func(p any) (any, error) { return decodePayload[StopTransactionRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.StopTransaction == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.StopTransaction(ctx, req.(StopTransactionRequest))
			},
		},
		ActionTriggerMessage: {
			decode: func(p any) (any, error) { return decodePayload[TriggerMessageRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.TriggerMessage == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.TriggerMessage(ctx, req.(TriggerMessageRequest))
			},
		},
		ActionUnlockConnector: {
			decode: func(p any) (any, error) { return decodePayload[UnlockConnectorRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.UnlockConnector == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
				}
				return o.UnlockConnector(ctx, req.(UnlockConnectorRequest))
			},
		},
		ActionUpdateFirmware: {
			decode: func(p any) (any, error) { return decodePayload[UpdateFirmwareRequest](p) },
			callback: func(ctx *OcppContext, req any) (any, *OcppError) {
				if o.UpdateFirmware == nil {
					return nil, &OcppError{ErrorCode: NotSupported, ErrorDescription: ActionNotRecognized, ErrorDetails: nil}
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
