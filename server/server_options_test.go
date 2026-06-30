package server

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/JohnMaddison/ocpp-go/ocpp16"
)

func TestWithBootNotificationHandlerRoutesIncomingCall(t *testing.T) {
	called := false
	s := NewServer(":0",
		WithBootNotificationHandler(func(ctx *ocpp16.OCPPContext, request ocpp16.BootNotificationRequest) (*ocpp16.BootNotificationResponse, *ocpp16.OCPPError) {
			called = true
			if ctx.ChargePointID != "CP_1" {
				t.Fatalf("expected charge point ID CP_1, got %s", ctx.ChargePointID)
			}
			if request.ChargePointModel != "Simulator" {
				t.Fatalf("expected model Simulator, got %s", request.ChargePointModel)
			}
			return &ocpp16.BootNotificationResponse{
				CurrentTime: time.Now(),
				Interval:    30,
				Status:      ocpp16.RegistrationStatusAccepted,
			}, nil
		}),
	)

	if s.parser == nil {
		t.Fatal("expected parser to be configured")
	}

	response, err := s.parser([]byte(`[2,"msg-1","BootNotification",{"chargePointModel":"Simulator","chargePointVendor":"OCPP-GO"}]`), ocpp16.NewOCPPContext("CP_1"))
	if err != nil {
		t.Fatalf("parser returned error: %v", err)
	}
	if !called {
		t.Fatal("expected boot notification handler to be called")
	}

	frame := decodeServerFrame(t, response)
	if frame[0].(float64) != 3 {
		t.Fatalf("expected CALLRESULT message type, got %v", frame[0])
	}
	payload := frame[2].(map[string]any)
	if payload["status"].(string) != string(ocpp16.RegistrationStatusAccepted) {
		t.Fatalf("expected Accepted status, got %v", payload["status"])
	}
}

func TestHandlerOptionDoesNotOverrideCustomParser(t *testing.T) {
	customParser := func(message []byte, ctx *ocpp16.OCPPContext) ([]byte, error) {
		return []byte(`[3,"custom",{}]`), nil
	}

	s := NewServer(":0",
		WithParser(customParser),
		WithBootNotificationHandler(func(ctx *ocpp16.OCPPContext, request ocpp16.BootNotificationRequest) (*ocpp16.BootNotificationResponse, *ocpp16.OCPPError) {
			t.Fatal("custom parser should have been used")
			return nil, nil
		}),
	)

	response, err := s.parser([]byte(`[2,"msg-1","BootNotification",{}]`), ocpp16.NewOCPPContext("CP_1"))
	if err != nil {
		t.Fatalf("parser returned error: %v", err)
	}

	frame := decodeServerFrame(t, response)
	if frame[1].(string) != "custom" {
		t.Fatalf("expected custom parser response, got %v", frame[1])
	}
}

func decodeServerFrame(t *testing.T, response []byte) []any {
	t.Helper()

	var frame []any
	if err := json.Unmarshal(response, &frame); err != nil {
		t.Fatalf("failed to decode response frame %s: %v", response, err)
	}
	return frame
}
