package server

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/johnmaddison/ocpp-go/ocpp21"
	"github.com/google/uuid"
)

func TestServerUsesDefaultMessageIDGenerator(t *testing.T) {
	s := NewServer(":0")
	if s.messageIDGenerator == nil {
		t.Fatal("expected default message ID generator")
	}
	if _, err := uuid.Parse(s.messageIDGenerator()); err != nil {
		t.Fatalf("expected UUID message ID, got error: %v", err)
	}
}

func TestServerMessageIDGeneratorCanBeOverridden(t *testing.T) {
	s := NewServer(":0", WithMessageIDGenerator(func() string { return "custom-id" }))
	if got := s.messageIDGenerator(); got != "custom-id" {
		t.Fatalf("expected custom-id, got %q", got)
	}
}

func TestWithOCPP16BootNotificationHandlerRoutesIncomingCall(t *testing.T) {
	called := false
	s := NewServer(":0",
		WithOCPP16BootNotificationHandler(func(ctx *ocpp16.OCPPContext, request ocpp16.BootNotificationRequest) (*ocpp16.BootNotificationResponse, *ocpp16.OCPPError) {
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
		WithOCPP16BootNotificationHandler(func(ctx *ocpp16.OCPPContext, request ocpp16.BootNotificationRequest) (*ocpp16.BootNotificationResponse, *ocpp16.OCPPError) {
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

func TestWithOCPP21BootNotificationHandlerRoutesIncomingCall(t *testing.T) {
	called := false
	s := NewServer(":0",
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
		}),
	)

	if !s.ocpp21Enabled {
		t.Fatal("expected ocpp2.1 to be enabled")
	}

	response, err := s.ocpp21Callbacks.ParseMessage([]byte(`[2,"msg-1","BootNotification",{}]`), ocpp21.NewOCPPContext("CP_1"))
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
}

func TestVersionedBootNotificationHandlersDoNotConflict(t *testing.T) {
	s := NewServer(":0",
		WithOCPP16BootNotificationHandler(func(ctx *ocpp16.OCPPContext, request ocpp16.BootNotificationRequest) (*ocpp16.BootNotificationResponse, *ocpp16.OCPPError) {
			return &ocpp16.BootNotificationResponse{CurrentTime: time.Now(), Interval: 30, Status: ocpp16.RegistrationStatusAccepted}, nil
		}),
		WithOCPP21BootNotificationHandler(func(ctx *ocpp21.OCPPContext, request ocpp21.BootNotificationRequest) (*ocpp21.BootNotificationResponse, *ocpp21.OCPPError) {
			return &ocpp21.BootNotificationResponse{CurrentTime: time.Now(), Interval: 30, Status: ocpp21.RegistrationStatusEnumAccepted}, nil
		}),
	)

	if got := s.protocols(); len(got) != 2 || got[0] != "ocpp1.6" || got[1] != "ocpp2.1" {
		t.Fatalf("expected both protocols, got %#v", got)
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
