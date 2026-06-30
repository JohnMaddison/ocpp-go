package client

import (
	"encoding/json"
	"testing"

	"github.com/JohnMaddison/ocpp-go/ocpp16"
)

func TestNewClientInitializesOCPPHandlers(t *testing.T) {
	c := NewClient("CP_1", "ws://127.0.0.1:9001/ocpp")

	response, err := c.ocppCallbacks.ParseMessage([]byte(`[2,"msg-1","GetConfiguration",{"key":[]}]`), ocpp16.NewOCPPContext("CP_1"))
	if err != nil {
		t.Fatalf("ParseMessage returned error: %v", err)
	}

	frame := decodeFrame(t, response)
	if frame[0].(float64) != 4 {
		t.Fatalf("expected CALLERROR message type, got %v", frame[0])
	}
	if frame[2].(string) != ocpp16.NotSupported {
		t.Fatalf("expected NotSupported error, got %v", frame[2])
	}
}

func TestWithConfigurationHandlerRoutesIncomingCall(t *testing.T) {
	const value = "Simulator"
	called := false

	c := NewClient("CP_1", "ws://127.0.0.1:9001/ocpp").
		WithConfigurationHandler(func(ctx *ocpp16.OCPPContext, request ocpp16.GetConfigurationRequest) (*ocpp16.GetConfigurationResponse, *ocpp16.OCPPError) {
			called = true
			if ctx.ChargePointID != "CP_1" {
				t.Fatalf("expected charge point ID CP_1, got %s", ctx.ChargePointID)
			}
			if len(request.Key) != 1 || request.Key[0] != "ChargePointModel" {
				t.Fatalf("unexpected request keys: %#v", request.Key)
			}
			return &ocpp16.GetConfigurationResponse{
				ConfigurationKey: []ocpp16.KeyValue{
					{Key: "ChargePointModel", Readonly: true, Value: ocpp16Value(value)},
				},
			}, nil
		})

	response, err := c.ocppCallbacks.ParseMessage([]byte(`[2,"msg-1","GetConfiguration",{"key":["ChargePointModel"]}]`), ocpp16.NewOCPPContext("CP_1"))
	if err != nil {
		t.Fatalf("ParseMessage returned error: %v", err)
	}
	if !called {
		t.Fatal("expected configuration handler to be called")
	}

	frame := decodeFrame(t, response)
	if frame[0].(float64) != 3 {
		t.Fatalf("expected CALLRESULT message type, got %v", frame[0])
	}
	payload := frame[2].(map[string]any)
	keys := payload["configurationKey"].([]any)
	first := keys[0].(map[string]any)
	if first["value"].(string) != value {
		t.Fatalf("expected configuration value %q, got %v", value, first["value"])
	}
}

func decodeFrame(t *testing.T, response []byte) []any {
	t.Helper()

	var frame []any
	if err := json.Unmarshal(response, &frame); err != nil {
		t.Fatalf("failed to decode response frame %s: %v", response, err)
	}
	return frame
}

func ocpp16Value(value string) *string {
	return &value
}
