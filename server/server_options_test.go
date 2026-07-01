package server

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/johnmaddison/ocpp-go/ocpp21"
)

func TestServerUsesDefaultResourceLimits(t *testing.T) {
	s := NewServer(":0")

	if s.readHeaderTimeout != 15*time.Second {
		t.Fatalf("readHeaderTimeout = %v, want 15s", s.readHeaderTimeout)
	}
	if s.readTimeout != 60*time.Second {
		t.Fatalf("readTimeout = %v, want 60s", s.readTimeout)
	}
	if s.idleTimeout != 120*time.Second {
		t.Fatalf("idleTimeout = %v, want 120s", s.idleTimeout)
	}
	if s.maxHeaderBytes != 1<<20 {
		t.Fatalf("maxHeaderBytes = %d, want 1 MiB", s.maxHeaderBytes)
	}
	if s.websocketReadLimit != 4<<20 {
		t.Fatalf("websocketReadLimit = %d, want 4 MiB", s.websocketReadLimit)
	}
	if !s.websocketCompression {
		t.Fatal("expected websocket compression to be enabled by default")
	}
	if s.pingInterval != 30*time.Second {
		t.Fatalf("pingInterval = %v, want 30s", s.pingInterval)
	}
	if s.pongTimeout != 45*time.Second {
		t.Fatalf("pongTimeout = %v, want 45s", s.pongTimeout)
	}
}

func TestServerResourceLimitOptionsOverrideDefaults(t *testing.T) {
	s := NewServer(":0",
		WithHTTPTimeouts(time.Second, 2*time.Second, 3*time.Second),
		WithMaxHeaderBytes(4096),
		WithWebsocketReadLimit(1024),
		WithWebsocketCompression(false),
		WithWebsocketKeepalive(4*time.Second, 5*time.Second),
	)

	if s.readHeaderTimeout != time.Second {
		t.Fatalf("readHeaderTimeout = %v, want 1s", s.readHeaderTimeout)
	}
	if s.readTimeout != 2*time.Second {
		t.Fatalf("readTimeout = %v, want 2s", s.readTimeout)
	}
	if s.idleTimeout != 3*time.Second {
		t.Fatalf("idleTimeout = %v, want 3s", s.idleTimeout)
	}
	if s.maxHeaderBytes != 4096 {
		t.Fatalf("maxHeaderBytes = %d, want 4096", s.maxHeaderBytes)
	}
	if s.websocketReadLimit != 1024 {
		t.Fatalf("websocketReadLimit = %d, want 1024", s.websocketReadLimit)
	}
	if s.websocketCompression {
		t.Fatal("expected websocket compression to be disabled")
	}
	if s.pingInterval != 4*time.Second {
		t.Fatalf("pingInterval = %v, want 4s", s.pingInterval)
	}
	if s.pongTimeout != 5*time.Second {
		t.Fatalf("pongTimeout = %v, want 5s", s.pongTimeout)
	}
}

func TestServerWebsocketKeepaliveCanBeDisabled(t *testing.T) {
	s := NewServer(":0", WithWebsocketKeepalive(0, 0))

	if s.pingInterval != 0 {
		t.Fatalf("pingInterval = %v, want disabled", s.pingInterval)
	}
	if s.pongTimeout != 0 {
		t.Fatalf("pongTimeout = %v, want disabled", s.pongTimeout)
	}
}

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

func TestWith16BootNotificationHandlerRoutesIncomingCall(t *testing.T) {
	called := false
	s := NewServer(":0",
		With16BootNotificationHandler(func(ctx *ocpp16.Context, request ocpp16.BootNotificationRequest) (*ocpp16.BootNotificationResponse, *ocpp16.Error) {
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

	response, err := s.parser([]byte(`[2,"msg-1","BootNotification",{"chargePointModel":"Simulator","chargePointVendor":"OCPP-GO"}]`), ocpp16.NewContext("CP_1"))
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
	customParser := func(message []byte, ctx *ocpp16.Context) ([]byte, error) {
		return []byte(`[3,"custom",{}]`), nil
	}

	s := NewServer(":0",
		WithParser(customParser),
		With16BootNotificationHandler(func(ctx *ocpp16.Context, request ocpp16.BootNotificationRequest) (*ocpp16.BootNotificationResponse, *ocpp16.Error) {
			t.Fatal("custom parser should have been used")
			return nil, nil
		}),
	)

	response, err := s.parser([]byte(`[2,"msg-1","BootNotification",{}]`), ocpp16.NewContext("CP_1"))
	if err != nil {
		t.Fatalf("parser returned error: %v", err)
	}

	frame := decodeServerFrame(t, response)
	if frame[1].(string) != "custom" {
		t.Fatalf("expected custom parser response, got %v", frame[1])
	}
}

func TestWith21BootNotificationHandlerRoutesIncomingCall(t *testing.T) {
	called := false
	s := NewServer(":0",
		With21BootNotificationHandler(func(ctx *ocpp21.Context, request ocpp21.BootNotificationRequest) (*ocpp21.BootNotificationResponse, *ocpp21.Error) {
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

	response, err := s.ocpp21Callbacks.ParseMessage([]byte(`[2,"msg-1","BootNotification",{}]`), ocpp21.NewContext("CP_1"))
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
		With16BootNotificationHandler(func(ctx *ocpp16.Context, request ocpp16.BootNotificationRequest) (*ocpp16.BootNotificationResponse, *ocpp16.Error) {
			return &ocpp16.BootNotificationResponse{CurrentTime: time.Now(), Interval: 30, Status: ocpp16.RegistrationStatusAccepted}, nil
		}),
		With21BootNotificationHandler(func(ctx *ocpp21.Context, request ocpp21.BootNotificationRequest) (*ocpp21.BootNotificationResponse, *ocpp21.Error) {
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
