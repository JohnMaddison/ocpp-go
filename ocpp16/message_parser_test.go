package ocpp16

import (
	"testing"
	"time"

	"github.com/JohnMaddison/ocpp-go"
)

func TestParseMessage_DecodesBootNotificationCallResult(t *testing.T) {
	callbacks := OcppCallbacks{}
	ctx := NewOcppContext("CP_TEST")
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
