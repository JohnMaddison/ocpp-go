package ocpp16

import (
	"context"
	"testing"
	"time"

	"github.com/JohnMaddison/ocpp-go"
)

func TestSendWithContextUsesMessageIdGeneratorWhenMissing(t *testing.T) {
	ctx := NewOCPPContextWithMessageIdGenerator("CP_1", func() string { return "custom-id" })

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
