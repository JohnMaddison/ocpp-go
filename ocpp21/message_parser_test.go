package ocpp21

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/johnmaddison/ocpp-go"
)

func TestParseMessage_BootNotificationCall(t *testing.T) {
	callbacks := Callbacks{
		BootNotification: func(ctx *Context, request BootNotificationRequest) (*BootNotificationResponse, *Error) {
			if ctx.ChargePointID != "CP_21" {
				t.Fatalf("unexpected charge point ID: %s", ctx.ChargePointID)
			}
			return &BootNotificationResponse{
				CurrentTime: time.Now(),
				Interval:    30,
				Status:      RegistrationStatusEnumAccepted,
			}, nil
		},
	}
	callbacks.InitHandlers()

	response, err := callbacks.ParseMessage([]byte(`[2,"msg-1","BootNotification",{}]`), NewContext("CP_21"))
	if err != nil {
		t.Fatalf("ParseMessage returned error: %v", err)
	}

	var message []any
	if err := json.Unmarshal(response, &message); err != nil {
		t.Fatalf("response is not JSON: %v", err)
	}
	if message[0] != float64(3) || message[1] != "msg-1" {
		t.Fatalf("unexpected response envelope: %#v", message)
	}
}

func TestParseMessage_UnsupportedCallReturnsCallError(t *testing.T) {
	callbacks := Callbacks{}
	callbacks.InitHandlers()

	response, err := callbacks.ParseMessage([]byte(`[2,"msg-1","BootNotification",{}]`), NewContext("CP_21"))
	if err != nil {
		t.Fatalf("ParseMessage returned error: %v", err)
	}

	var message []any
	if err := json.Unmarshal(response, &message); err != nil {
		t.Fatalf("response is not JSON: %v", err)
	}
	if message[0] != float64(4) || message[2] != "NotSupported" {
		t.Fatalf("unexpected call error: %#v", message)
	}
}

func TestParseMessage_MalformedPayloadReturnsFormationViolation(t *testing.T) {
	callbacks := Callbacks{
		BootNotification: func(ctx *Context, request BootNotificationRequest) (*BootNotificationResponse, *Error) {
			t.Fatal("callback should not be called")
			return nil, nil
		},
	}
	callbacks.InitHandlers()

	response, err := callbacks.ParseMessage([]byte(`[2,"msg-1","BootNotification","bad-payload"]`), NewContext("CP_21"))
	if err != nil {
		t.Fatalf("ParseMessage returned error: %v", err)
	}

	var message []any
	if err := json.Unmarshal(response, &message); err != nil {
		t.Fatalf("response is not JSON: %v", err)
	}
	if message[0] != float64(4) || message[2] != "FormationViolation" {
		t.Fatalf("unexpected call error: %#v", message)
	}
}

func TestParseMessage_DecodesCallResult(t *testing.T) {
	callbacks := Callbacks{}
	callbacks.InitHandlers()
	ctx := NewContext("CP_21")
	request := Request{
		Call:   ocpp.Call{MessageID: "msg-1", Action: string(ActionBootNotification)},
		result: make(chan ResultOrError, 1),
	}
	ctx.storage.Add(request)

	_, err := callbacks.ParseMessage([]byte(`[3,"msg-1",{"currentTime":"2026-06-26T20:22:36Z","interval":10,"status":"Accepted"}]`), ctx)
	if err != nil {
		t.Fatalf("ParseMessage returned error: %v", err)
	}

	result := <-request.result
	payload, ok := result.CallResult.Payload.(BootNotificationResponse)
	if !ok {
		t.Fatalf("unexpected payload type: %T", result.CallResult.Payload)
	}
	if payload.Status != RegistrationStatusEnumAccepted {
		t.Fatalf("unexpected status: %s", payload.Status)
	}
}

func TestParseMessage_DeliversCallError(t *testing.T) {
	callbacks := Callbacks{}
	callbacks.InitHandlers()
	ctx := NewContext("CP_21")
	request := Request{
		Call:   ocpp.Call{MessageID: "msg-1", Action: string(ActionBootNotification)},
		result: make(chan ResultOrError, 1),
	}
	ctx.storage.Add(request)

	_, err := callbacks.ParseMessage([]byte(`[4,"msg-1","InternalError","failed",{}]`), ctx)
	if err != nil {
		t.Fatalf("ParseMessage returned error: %v", err)
	}

	result := <-request.result
	if result.CallError == nil || result.CallError.ErrorCode != "InternalError" {
		t.Fatalf("unexpected call error: %#v", result.CallError)
	}
}
