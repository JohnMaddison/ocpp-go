package ocpp16

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/johnmaddison/ocpp-go"
)

func TestSendWithContextUsesMessageIDGeneratorWhenMissing(t *testing.T) {
	ctx := NewContextWithMessageIDGenerator("CP_1", func() string { return "custom-id" })

	go func() {
		request := <-ctx.Queue
		if request.Call.MessageID != "custom-id" {
			t.Errorf("expected custom-id, got %q", request.Call.MessageID)
		}
		request.result <- ResultOrError{CallResult: &ocpp.CallResult{MessageID: request.Call.MessageID}}
	}()

	callCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := ctx.SendWithContext(callCtx, ocpp.Call{MessageType: ocpp.MessageTypeCall, Action: string(ActionHeartbeat), Payload: EmptyPayload{}}); err != nil {
		t.Fatalf("SendWithContext returned error: %v", err)
	}
}

func TestSendCallBuildsTypedCall(t *testing.T) {
	ctx := NewContextWithMessageIDGenerator("CP_1", func() string { return "custom-id" })
	payload := DataTransferRequest{VendorID: "vendor"}

	go func() {
		request := <-ctx.Queue
		if request.Call.MessageType != ocpp.MessageTypeCall {
			t.Errorf("MessageType = %d, want %d", request.Call.MessageType, ocpp.MessageTypeCall)
		}
		if request.Call.MessageID != "custom-id" {
			t.Errorf("MessageID = %q, want custom-id", request.Call.MessageID)
		}
		if request.Call.Action != string(ActionDataTransfer) {
			t.Errorf("Action = %q, want %q", request.Call.Action, ActionDataTransfer)
		}
		if !reflect.DeepEqual(request.Call.Payload, payload) {
			t.Errorf("Payload = %#v, want %#v", request.Call.Payload, payload)
		}
		request.result <- ResultOrError{CallResult: &ocpp.CallResult{MessageID: request.Call.MessageID}}
	}()

	if _, err := ctx.SendCall(ActionDataTransfer, payload); err != nil {
		t.Fatalf("SendCall returned error: %v", err)
	}
}
