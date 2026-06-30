package ocpp16

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JohnMaddison/ocpp-go"
)

type OCPPMessage []any

func (o *OCPPCallbacks) ParseMessage(message []byte, ctx *OCPPContext) ([]byte, error) {
	// Parse the OCPP message using standard JSON
	var ocppMessage OCPPMessage
	if err := json.Unmarshal(message, &ocppMessage); err != nil {
		return nil, fmt.Errorf("failed to unmarshal OCPP message: %w", err)
	}

	// Validate message structure
	if len(ocppMessage) < 3 {
		return nil, fmt.Errorf("invalid OCPP message format")
	}

	messageType, ok := ocppMessage[0].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid message type")
	}

	// Handle CALL messages (type 2)
	if int(messageType) == ocpp.MessageTypeCall {
		if len(ocppMessage) != 4 {
			return nil, fmt.Errorf("invalid CALL message format")
		}

		messageID, ok := ocppMessage[1].(string)
		if !ok {
			return nil, fmt.Errorf("invalid message ID")
		}

		actionStr, ok := ocppMessage[2].(string)
		if !ok {
			return nil, fmt.Errorf("invalid action")
		}

		callResult, callError, err := o.processCall(messageID, Action(actionStr), ocppMessage[3], ctx)
		if err != nil {
			return nil, err
		}

		if callError != nil {
			return callError.SerializeOCPP()
		} else if callResult != nil {
			return callResult.SerializeOCPP()
		}
	}

	// Handle CALLRESULT messages (type 3)
	if int(messageType) == ocpp.MessageTypeCallResult {
		if len(ocppMessage) != 3 {
			return nil, fmt.Errorf("invalid CALLRESULT message format")
		}

		messageID, ok := ocppMessage[1].(string)
		if !ok {
			return nil, fmt.Errorf("invalid message ID")
		}

		item, found := ctx.storage.FindByMessageID(messageID)
		if found {
			payload, err := decodeCallResultPayload(Action(item.Call.Action), ocppMessage[2])
			if err != nil {
				return nil, err
			}

			callResult := ocpp.CallResult{
				MessageType: int(messageType),
				MessageID:   messageID,
				Payload:     payload,
			}

			item.result <- ResultOrError{CallResult: &callResult, CallError: nil}
		}

		return nil, nil
	}

	// Handle CALLERROR messages (type 4)
	if int(messageType) == ocpp.MessageTypeCallError {
		if len(ocppMessage) < 4 || len(ocppMessage) > 5 {
			return nil, fmt.Errorf("invalid CALLERROR message format")
		}

		messageID, ok := ocppMessage[1].(string)
		if !ok {
			return nil, fmt.Errorf("invalid message ID")
		}

		errorCode, ok := ocppMessage[2].(string)
		if !ok {
			return nil, fmt.Errorf("invalid error code")
		}

		errorDescription, ok := ocppMessage[3].(string)
		if !ok {
			return nil, fmt.Errorf("invalid error description")
		}

		var errorDetails any
		if len(ocppMessage) == 5 {
			errorDetails = ocppMessage[4]
		}

		callError := ocpp.CallError{
			MessageType:      int(messageType),
			MessageID:        messageID,
			ErrorCode:        errorCode,
			ErrorDescription: errorDescription,
			ErrorDetails:     errorDetails,
		}

		item, found := ctx.storage.FindByMessageID(messageID)

		if found {
			item.result <- ResultOrError{CallError: &callError, CallResult: nil}
		}

		return nil, nil
	}

	return nil, fmt.Errorf("message type not implemented: %d", int(messageType))
}

func (o *OCPPCallbacks) processCall(messageID string, action Action, payload any, ctx *OCPPContext) (*ocpp.CallResult, *ocpp.CallError, error) {

	handler, ok := o.handlers[action]
	if !ok || handler.callback == nil {
		log.Printf("messageID %s no handler found for action: %s", messageID, action)
		return nil, &ocpp.CallError{
			MessageType:      ocpp.MessageTypeCallError,
			MessageID:        messageID,
			ErrorCode:        "NotSupported",
			ErrorDescription: "Requested Action is recognized but not supported by the receiver",
			ErrorDetails:     nil,
		}, nil
	}

	req, err := handler.decode(payload)
	if err != nil {
		return nil, &ocpp.CallError{
			MessageType:      ocpp.MessageTypeCallError,
			MessageID:        messageID,
			ErrorCode:        "FormationViolation",
			ErrorDescription: "Payload for Action is syntactically incorrect or not conform the PDU structure for Action",
			ErrorDetails:     err.Error(),
		}, nil
	}

	reply, callErr := handler.callback(ctx, req)
	if callErr != nil {
		return nil, &ocpp.CallError{
			MessageType:      ocpp.MessageTypeCallError,
			MessageID:        messageID,
			ErrorCode:        callErr.ErrorCode,
			ErrorDescription: callErr.ErrorDescription,
			ErrorDetails:     callErr.ErrorDetails,
		}, nil
	}

	if reply != nil {
		return &ocpp.CallResult{
			MessageType: ocpp.MessageTypeCallResult,
			MessageID:   messageID,
			Payload:     reply,
		}, nil, nil
	}

	return nil, nil, nil

}

func decodeCallResultPayload(action Action, payload any) (any, error) {
	switch action {
	case ActionAuthorize:
		return decodePayload[AuthorizeResponse](payload)
	case ActionBootNotification:
		return decodePayload[BootNotificationResponse](payload)
	case ActionCancelReservation:
		return decodePayload[CancelReservationResponse](payload)
	case ActionCertificateSigned:
		return decodePayload[CertificateSignedResponse](payload)
	case ActionChangeAvailability:
		return decodePayload[ChangeAvailabilityResponse](payload)
	case ActionChangeConfiguration:
		return decodePayload[ChangeConfigurationResponse](payload)
	case ActionClearCache:
		return decodePayload[ClearCacheResponse](payload)
	case ActionClearChargingProfile:
		return decodePayload[ClearChargingProfileResponse](payload)
	case ActionDataTransfer:
		return decodePayload[DataTransferResponse](payload)
	case ActionDeleteCertificate:
		return decodePayload[DeleteCertificateResponse](payload)
	case ActionDiagnosticsStatusNotification:
		return decodePayload[DiagnosticsStatusNotificationResponse](payload)
	case ActionExtendedTriggerMessage:
		return decodePayload[ExtendedTriggerMessageResponse](payload)
	case ActionFirmwareStatusNotification:
		return decodePayload[FirmwareStatusNotificationResponse](payload)
	case ActionGetCompositeSchedule:
		return decodePayload[GetCompositeScheduleResponse](payload)
	case ActionGetConfiguration:
		return decodePayload[GetConfigurationResponse](payload)
	case ActionGetDiagnostics:
		return decodePayload[GetDiagnosticsResponse](payload)
	case ActionGetInstalledCertificateIds:
		return decodePayload[GetInstalledCertificateIdsResponse](payload)
	case ActionGetLocalListVersion:
		return decodePayload[GetLocalListVersionResponse](payload)
	case ActionGetLog:
		return decodePayload[GetLogResponse](payload)
	case ActionHeartbeat:
		return decodePayload[HeartbeatResponse](payload)
	case ActionInstallCertificate:
		return decodePayload[InstallCertificateResponse](payload)
	case ActionLogStatusNotification:
		return decodePayload[LogStatusNotificationResponse](payload)
	case ActionMeterValues:
		return decodePayload[MeterValuesResponse](payload)
	case ActionRemoteStartTransaction:
		return decodePayload[RemoteStartTransactionResponse](payload)
	case ActionRemoteStopTransaction:
		return decodePayload[RemoteStopTransactionResponse](payload)
	case ActionReserveNow:
		return decodePayload[ReserveNowResponse](payload)
	case ActionReset:
		return decodePayload[ResetResponse](payload)
	case ActionSecurityEventNotification:
		return decodePayload[SecurityEventNotificationResponse](payload)
	case ActionSendLocalList:
		return decodePayload[SendLocalListResponse](payload)
	case ActionSetChargingProfile:
		return decodePayload[SetChargingProfileResponse](payload)
	case ActionSignCertificate:
		return decodePayload[SignCertificateResponse](payload)
	case ActionSignedFirmwareStatusNotification:
		return decodePayload[SignedFirmwareStatusNotificationResponse](payload)
	case ActionSignedUpdateFirmware:
		return decodePayload[SignedUpdateFirmwareResponse](payload)
	case ActionStartTransaction:
		return decodePayload[StartTransactionResponse](payload)
	case ActionStatusNotification:
		return decodePayload[StatusNotificationResponse](payload)
	case ActionStopTransaction:
		return decodePayload[StopTransactionResponse](payload)
	case ActionTriggerMessage:
		return decodePayload[TriggerMessageResponse](payload)
	case ActionUnlockConnector:
		return decodePayload[UnlockConnectorResponse](payload)
	case ActionUpdateFirmware:
		return decodePayload[UpdateFirmwareResponse](payload)
	default:
		return payload, nil
	}
}
