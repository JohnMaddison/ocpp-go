package ocpp16

import (
	"context"
	"fmt"
	"time"

	"github.com/JohnMaddison/ocpp-go"
)

type OcppContext struct {
	ChargePointID string
	Queue         chan Request
	storage       *CircularBuffer
}

func NewOcppContext(ChargePointID string) *OcppContext {
	capacity := 10
	return &OcppContext{
		ChargePointID: ChargePointID,
		Queue:         make(chan Request, capacity),
		storage:       NewCircularBuffer(capacity),
	}
}

type ResultOrError struct {
	CallResult *ocpp.CallResult
	CallError  *ocpp.CallError
}

type Request struct {
	Call   ocpp.Call
	result chan ResultOrError
}

// Send sends a call with a default 10-second timeout
func (s *OcppContext) Send(call ocpp.Call) (*ResultOrError, error) {
	return s.SendWithTimeout(call, 10*time.Second)
}

// SendWithTimeout sends a call with a custom timeout
func (s *OcppContext) SendWithTimeout(call ocpp.Call, timeout time.Duration) (*ResultOrError, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return s.SendWithContext(ctx, call)
}

// SendWithContext sends a call with context for cancellation
func (s *OcppContext) SendWithContext(ctx context.Context, call ocpp.Call) (*ResultOrError, error) {
	msg := Request{
		Call:   call,
		result: make(chan ResultOrError, 1), // Buffered to prevent goroutine leaks
	}

	s.storage.Add(msg)
	// Non-blocking send to queue with context
	select {
	case s.Queue <- msg:
		// Successfully queued
	case <-ctx.Done():
		return nil, fmt.Errorf("failed to queue call %s: %w", call.MessageID, ctx.Err())
	}

	// Wait for result
	select {
	case result := <-msg.result:
		return &result, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("call %s timed out: %w", call.MessageID, ctx.Err())
	}
}

func (s *OcppContext) SendCallAndExpectResult(action string, payload any) (*ocpp.CallResult, error) {
	call := ocpp.NewCall(action, payload)
	request, err := s.Send(*call)

	if err != nil {
		return nil, fmt.Errorf("Failed to send %s call, error: %s", action, err)
	} else if request.CallError != nil {
		return nil, fmt.Errorf("Recieved Callerror when sending %s, error: %s", action, err)
	}
	return request.CallResult, nil
}
