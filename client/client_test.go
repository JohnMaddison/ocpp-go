package client

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/johnmaddison/ocpp-go/ocpp21"
	"github.com/google/uuid"
)

func TestNewClientInitializesOCPPHandlers(t *testing.T) {
	c := NewOCPP16Client("CP_1", "ws://127.0.0.1:9001/ocpp")

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

func TestVersionedConstructorsSelectSingleSubprotocol(t *testing.T) {
	client16 := NewOCPP16Client("CP_1", "ws://127.0.0.1:9001/ocpp")
	if client16.subprotocol != "ocpp1.6" {
		t.Fatalf("OCPP 1.6 client subprotocol = %q", client16.subprotocol)
	}

	client21 := NewOCPP21Client("CP_1", "ws://127.0.0.1:9001/ocpp")
	if client21.subprotocol != "ocpp2.1" {
		t.Fatalf("OCPP 2.1 client subprotocol = %q", client21.subprotocol)
	}
}

func TestClientConstructorsUseDefaultMessageIDGenerator(t *testing.T) {
	for _, c := range []*Client{
		NewOCPP16Client("CP_1", "ws://127.0.0.1:9001/ocpp"),
		NewOCPP21Client("CP_1", "ws://127.0.0.1:9001/ocpp"),
	} {
		if c.messageIDGenerator == nil {
			t.Fatal("expected default message ID generator")
		}
		if _, err := uuid.Parse(c.messageIDGenerator()); err != nil {
			t.Fatalf("expected UUID message ID, got error: %v", err)
		}
	}
}

func TestClientMessageIDGeneratorCanBeOverridden(t *testing.T) {
	c := NewOCPP16Client("CP_1", "ws://127.0.0.1:9001/ocpp").
		WithMessageIDGenerator(func() string { return "custom-id" })

	if got := c.messageIDGenerator(); got != "custom-id" {
		t.Fatalf("expected custom-id, got %q", got)
	}
}

func TestOCPP21ClientCallbacksUseOCPP21Types(t *testing.T) {
	called := false
	c := NewOCPP21Client("CP_1", "ws://127.0.0.1:9001/ocpp").WithOCPP21Callbacks(ocpp21.OCPPCallbacks{
		BootNotification: func(ctx *ocpp21.OCPPContext, request ocpp21.BootNotificationRequest) (*ocpp21.BootNotificationResponse, *ocpp21.OCPPError) {
			called = true
			if ctx.ChargePointID != "CP_1" {
				t.Fatalf("expected charge point ID CP_1, got %s", ctx.ChargePointID)
			}
			return &ocpp21.BootNotificationResponse{
				CurrentTime: time.Now(),
				Interval:    30,
				Status:      ocpp21.RegistrationStatusEnumAccepted,
			}, nil
		},
	})

	response, err := c.ocpp21Callbacks.ParseMessage([]byte(`[2,"msg-1","BootNotification",{}]`), ocpp21.NewOCPPContext("CP_1"))
	if err != nil {
		t.Fatalf("ParseMessage returned error: %v", err)
	}
	if !called {
		t.Fatal("expected OCPP 2.1 callback to be called")
	}

	frame := decodeFrame(t, response)
	if frame[0].(float64) != 3 {
		t.Fatalf("expected CALLRESULT message type, got %v", frame[0])
	}
}

func TestWithOCPP16GetConfigurationHandlerRoutesIncomingCall(t *testing.T) {
	const value = "Simulator"
	called := false

	c := NewOCPP16Client("CP_1", "ws://127.0.0.1:9001/ocpp").
		WithOCPP16GetConfigurationHandler(func(ctx *ocpp16.OCPPContext, request ocpp16.GetConfigurationRequest) (*ocpp16.GetConfigurationResponse, *ocpp16.OCPPError) {
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

func TestWithOCPP21BootNotificationHandlerRoutesIncomingCall(t *testing.T) {
	called := false
	c := NewOCPP21Client("CP_1", "ws://127.0.0.1:9001/ocpp").
		WithOCPP21BootNotificationHandler(func(ctx *ocpp21.OCPPContext, request ocpp21.BootNotificationRequest) (*ocpp21.BootNotificationResponse, *ocpp21.OCPPError) {
			called = true
			if ctx.ChargePointID != "CP_1" {
				t.Fatalf("expected charge point ID CP_1, got %s", ctx.ChargePointID)
			}
			return &ocpp21.BootNotificationResponse{
				CurrentTime: time.Now(),
				Interval:    30,
				Status:      ocpp21.RegistrationStatusEnumAccepted,
			}, nil
		})

	response, err := c.ocpp21Callbacks.ParseMessage([]byte(`[2,"msg-1","BootNotification",{}]`), ocpp21.NewOCPPContext("CP_1"))
	if err != nil {
		t.Fatalf("ParseMessage returned error: %v", err)
	}
	if !called {
		t.Fatal("expected OCPP 2.1 boot notification handler to be called")
	}

	frame := decodeFrame(t, response)
	if frame[0].(float64) != 3 {
		t.Fatalf("expected CALLRESULT message type, got %v", frame[0])
	}
	if c.subprotocol != "ocpp2.1" {
		t.Fatalf("expected ocpp2.1 subprotocol, got %s", c.subprotocol)
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
