package ocpp16

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JohnMaddison/ocpp-go"
)

type OCPPMessage []any

func (o *OcppCallbacks) ParseMessage(message []byte, ctx *OcppContext) ([]byte, error) {
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
			return callError.SerializeToOcpp()
		} else if callResult != nil {
			return callResult.SerializeToOcpp()
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

		callResult := ocpp.CallResult{
			MessageType: int(messageType),
			MessageID:   messageID,
			Payload:     ocppMessage[2],
		}

		item, found := ctx.storage.FindByMessageID(messageID)

		if found {
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

		var errorDetails interface{}
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

func (o *OcppCallbacks) processCall(MessageID string, Action Action, payload any, ctx *OcppContext) (*ocpp.CallResult, *ocpp.CallError, error) {

	handler, ok := o.handlers[Action]
	if !ok || handler.callback == nil {
		log.Printf("MessageId %s no handler found for Action: %s", MessageID, Action)
		return nil, &ocpp.CallError{
			MessageType:      ocpp.MessageTypeCallError,
			MessageID:        MessageID,
			ErrorCode:        "NotSupported",
			ErrorDescription: "Requested Action is recognized but not supported by the receiver",
			ErrorDetails:     nil,
		}, nil
	}

	req, err := handler.decode(payload)
	if err != nil {
		return nil, &ocpp.CallError{
			MessageType:      ocpp.MessageTypeCallError,
			MessageID:        MessageID,
			ErrorCode:        "FormationViolation",
			ErrorDescription: "Payload for Action is syntactically incorrect or not conform the PDU structure for Action",
			ErrorDetails:     err.Error(),
		}, nil
	}

	reply, callErr := handler.callback(ctx, req)
	if callErr != nil {
		return nil, &ocpp.CallError{
			MessageType:      ocpp.MessageTypeCallError,
			MessageID:        MessageID,
			ErrorCode:        callErr.ErrorCode,
			ErrorDescription: callErr.ErrorDescription,
			ErrorDetails:     callErr.ErrorDetails,
		}, nil
	}

	if reply != nil {
		return &ocpp.CallResult{
			MessageType: ocpp.MessageTypeCallResult,
			MessageID:   MessageID,
			Payload:     reply,
		}, nil, nil
	}

	return nil, nil, nil

}
