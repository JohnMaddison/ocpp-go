package ocpp21

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/JohnMaddison/ocpp-go"
)

type OCPPMessage []any

func (o *OCPPCallbacks) ParseMessage(message []byte, ctx *OCPPContext) ([]byte, error) {
	var ocppMessage OCPPMessage
	if err := json.Unmarshal(message, &ocppMessage); err != nil {
		return nil, fmt.Errorf("failed to unmarshal OCPP message: %w", err)
	}

	if len(ocppMessage) < 3 {
		return nil, fmt.Errorf("invalid OCPP message format")
	}

	messageType, ok := ocppMessage[0].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid message type")
	}

	switch int(messageType) {
	case ocpp.MessageTypeCall:
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
		}
		if callResult != nil {
			return callResult.SerializeOCPP()
		}
		return nil, nil
	case ocpp.MessageTypeCallResult:
		if len(ocppMessage) != 3 {
			return nil, fmt.Errorf("invalid CALLRESULT message format")
		}
		messageID, ok := ocppMessage[1].(string)
		if !ok {
			return nil, fmt.Errorf("invalid message ID")
		}
		item, found := ctx.storage.FindByMessageID(messageID)
		if found {
			payload, err := o.decodeCallResultPayload(Action(item.Call.Action), ocppMessage[2])
			if err != nil {
				return nil, err
			}
			callResult := ocpp.CallResult{MessageType: int(messageType), MessageID: messageID, Payload: payload}
			item.result <- ResultOrError{CallResult: &callResult}
		}
		return nil, nil
	case ocpp.MessageTypeCallError:
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
		callError := ocpp.CallError{MessageType: int(messageType), MessageID: messageID, ErrorCode: errorCode, ErrorDescription: errorDescription, ErrorDetails: errorDetails}
		item, found := ctx.storage.FindByMessageID(messageID)
		if found {
			item.result <- ResultOrError{CallError: &callError}
		}
		return nil, nil
	default:
		return nil, fmt.Errorf("message type not implemented: %d", int(messageType))
	}
}

func (o *OCPPCallbacks) processCall(messageID string, action Action, payload any, ctx *OCPPContext) (*ocpp.CallResult, *ocpp.CallError, error) {
	handler, ok := o.handlers[action]
	if !ok || handler.callback == nil {
		log.Printf("messageID %s no handler found for action: %s", messageID, action)
		return nil, &ocpp.CallError{MessageType: ocpp.MessageTypeCallError, MessageID: messageID, ErrorCode: "NotSupported", ErrorDescription: "Requested Action is recognized but not supported by the receiver"}, nil
	}

	req, err := decodePayloadByType(payload, handler.requestType)
	if err != nil {
		return nil, &ocpp.CallError{MessageType: ocpp.MessageTypeCallError, MessageID: messageID, ErrorCode: "FormationViolation", ErrorDescription: "Payload for Action is syntactically incorrect or not conform the PDU structure for Action", ErrorDetails: err.Error()}, nil
	}

	results := reflect.ValueOf(handler.callback).Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(req)})
	if !results[1].IsNil() {
		callErr := results[1].Interface().(*OCPPError)
		return nil, &ocpp.CallError{MessageType: ocpp.MessageTypeCallError, MessageID: messageID, ErrorCode: callErr.ErrorCode, ErrorDescription: callErr.ErrorDescription, ErrorDetails: callErr.ErrorDetails}, nil
	}
	if !results[0].IsNil() {
		return &ocpp.CallResult{MessageType: ocpp.MessageTypeCallResult, MessageID: messageID, Payload: results[0].Interface()}, nil, nil
	}
	return nil, nil, nil
}

func (o *OCPPCallbacks) decodeCallResultPayload(action Action, payload any) (any, error) {
	handler, ok := o.handlers[action]
	if !ok {
		return payload, nil
	}
	return decodePayloadByType(payload, handler.responseType)
}
