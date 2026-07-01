package ocpp16

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/johnmaddison/ocpp-go"
)

func TestParseMessage_BootNotificationCall(t *testing.T) {
	callbacks := OCPPCallbacks{
		BootNotification: func(ctx *OCPPContext, request BootNotificationRequest) (*BootNotificationResponse, *OCPPError) {
			if ctx.ChargePointID != "CP_TEST" {
				t.Fatalf("unexpected charge point ID: %s", ctx.ChargePointID)
			}
			if request.ChargePointModel != "Simulator" {
				t.Fatalf("unexpected charge point model: %s", request.ChargePointModel)
			}
			return &BootNotificationResponse{
				CurrentTime: time.Now(),
				Interval:    30,
				Status:      RegistrationStatusAccepted,
			}, nil
		},
	}

	response, err := callbacks.ParseMessage([]byte(`[2,"msg-1","BootNotification",{"chargePointModel":"Simulator","chargePointVendor":"Maddison"}]`), NewOCPPContext("CP_TEST"))
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
	callbacks := OCPPCallbacks{}

	response, err := callbacks.ParseMessage([]byte(`[2,"msg-1","BootNotification",{}]`), NewOCPPContext("CP_TEST"))
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
	callbacks := OCPPCallbacks{
		BootNotification: func(ctx *OCPPContext, request BootNotificationRequest) (*BootNotificationResponse, *OCPPError) {
			t.Fatal("callback should not be called")
			return nil, nil
		},
	}

	response, err := callbacks.ParseMessage([]byte(`[2,"msg-1","BootNotification","bad-payload"]`), NewOCPPContext("CP_TEST"))
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

func TestParseMessage_DecodesBootNotificationCallResult(t *testing.T) {
	callbacks := OCPPCallbacks{}
	ctx := NewOCPPContext("CP_TEST")
	result := make(chan ResultOrError, 1)

	ctx.storage.Add(Request{
		Call: ocpp.Call{
			MessageType: ocpp.MessageTypeCall,
			MessageID:   "msg-1",
			Action:      string(ActionBootNotification),
			Payload: BootNotificationRequest{
				ChargePointModel:  "Simulator",
				ChargePointVendor: "Maddison",
			},
		},
		result: result,
	})

	response, err := callbacks.ParseMessage([]byte(`[3,"msg-1",{"currentTime":"2026-06-26T20:22:36.360732585+02:00","interval":10,"status":"Accepted"}]`), ctx)
	if err != nil {
		t.Fatalf("ParseMessage returned error: %v", err)
	}
	if response != nil {
		t.Fatalf("expected no websocket response for CALLRESULT, got %s", string(response))
	}

	select {
	case got := <-result:
		if got.CallError != nil {
			t.Fatalf("unexpected call error: %v", got.CallError)
		}
		if got.CallResult == nil {
			t.Fatal("expected call result")
		}

		payload, ok := got.CallResult.Payload.(BootNotificationResponse)
		if !ok {
			t.Fatalf("expected BootNotificationResponse payload, got %T", got.CallResult.Payload)
		}
		if payload.Interval != 10 {
			t.Fatalf("expected interval 10, got %d", payload.Interval)
		}
		if payload.Status != RegistrationStatusAccepted {
			t.Fatalf("expected status Accepted, got %s", payload.Status)
		}
		if payload.CurrentTime.IsZero() {
			t.Fatal("expected current time to be decoded")
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for call result")
	}
}
