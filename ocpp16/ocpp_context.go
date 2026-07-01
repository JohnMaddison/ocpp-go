package ocpp16

import (
	"context"
	"fmt"
	"time"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/internal/uuidgenerator"
)

type OCPPContext struct {
	ChargePointID      string
	Queue              chan Request
	storage            *CircularBuffer
	messageIdGenerator uuidgenerator.MessageIdGeneratorMethod
}

func NewOCPPContext(ChargePointID string) *OCPPContext {
	return NewOCPPContextWithMessageIdGenerator(ChargePointID, nil)
}

func NewOCPPContextWithMessageIdGenerator(ChargePointID string, generator func() string) *OCPPContext {
	capacity := 10
	if generator == nil {
		generator = uuidgenerator.DefaultUUIDGenerator
	}
	return &OCPPContext{
		ChargePointID:      ChargePointID,
		Queue:              make(chan Request, capacity),
		storage:            NewCircularBuffer(capacity),
		messageIdGenerator: generator,
	}
}

type ResultOrError struct {
	CallResult *ocpp.CallResult
	CallError  *ocpp.CallError
}

func (s ResultOrError) IsCallError() bool {
	return s.CallError != nil
}

func (s ResultOrError) IsCallResult() bool {
	return s.CallResult != nil
}

func (s ResultOrError) GetPayload() any {
	return s.CallResult.Payload
}

type Request struct {
	Call   ocpp.Call
	result chan ResultOrError
}

// Send sends a call with a default 10-second timeout
func (s *OCPPContext) Send(call ocpp.Call) (*ResultOrError, error) {
	return s.SendWithTimeout(call, 10*time.Second)
}

// SendCall sends a typed OCPP 1.6 CALL with a default 10-second timeout.
func (s *OCPPContext) SendCall(action Action, payload any) (*ResultOrError, error) {
	return s.Send(ocpp.Call{MessageType: ocpp.MessageTypeCall, Action: string(action), Payload: payload})
}

// SendWithTimeout sends a call with a custom timeout
func (s *OCPPContext) SendWithTimeout(call ocpp.Call, timeout time.Duration) (*ResultOrError, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return s.SendWithContext(ctx, call)
}

// SendCallWithContext sends a typed OCPP 1.6 CALL with context cancellation.
func (s *OCPPContext) SendCallWithContext(ctx context.Context, action Action, payload any) (*ResultOrError, error) {
	return s.SendWithContext(ctx, ocpp.Call{MessageType: ocpp.MessageTypeCall, Action: string(action), Payload: payload})
}

// SendWithContext sends a call with context for cancellation
func (s *OCPPContext) SendWithContext(ctx context.Context, call ocpp.Call) (*ResultOrError, error) {
	if call.MessageID == "" {
		call.MessageID = s.messageIdGenerator()
	}
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

func (s *OCPPContext) SendCallAndExpectResult(action string, payload any) (*ocpp.CallResult, error) {
	call := &ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: s.messageIdGenerator(), Action: action, Payload: payload}
	request, err := s.Send(*call)

	if err != nil {
		return nil, fmt.Errorf("Failed to send %s call, error: %s", action, err)
	} else if request.CallError != nil {
		return nil, fmt.Errorf("Received CallError when sending %s, error: %s", action, err)
	}
	return request.CallResult, nil
}
