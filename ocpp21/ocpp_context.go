package ocpp21

import (
	"context"
	"fmt"
	"time"

	"github.com/johnmaddison/ocpp-go"
	"github.com/johnmaddison/ocpp-go/internal/uuidgenerator"
)

type Context struct {
	ChargePointID      string
	Queue              chan Request
	storage            *CircularBuffer
	messageIDGenerator uuidgenerator.MessageIDGeneratorMethod
}

// NewContext creates an OCPP 2.1 context using the default message ID generator.
func NewContext(chargePointID string) *Context {
	return NewContextWithMessageIDGenerator(chargePointID, nil)
}

// NewContextWithMessageIDGenerator creates an OCPP 2.1 context with a custom message ID generator.
func NewContextWithMessageIDGenerator(chargePointID string, generator func() string) *Context {
	capacity := 10
	if generator == nil {
		generator = uuidgenerator.DefaultUUIDGenerator
	}
	return &Context{
		ChargePointID:      chargePointID,
		Queue:              make(chan Request, capacity),
		storage:            NewCircularBuffer(capacity),
		messageIDGenerator: generator,
	}
}

// ResultOrError contains either a CALLRESULT or CALLERROR response.
type ResultOrError struct {
	CallResult *ocpp.CallResult
	CallError  *ocpp.CallError
}

// IsCallError reports whether the response is a CALLERROR.
func (s ResultOrError) IsCallError() bool {
	return s.CallError != nil
}

// IsCallResult reports whether the response is a CALLRESULT.
func (s ResultOrError) IsCallResult() bool {
	return s.CallResult != nil
}

// GetPayload returns the CALLRESULT payload.
func (s ResultOrError) GetPayload() any {
	return s.CallResult.Payload
}

// Request tracks an outbound CALL and its eventual response.
type Request struct {
	Call   ocpp.Call
	result chan ResultOrError
}

// Send sends a CALL with a default 10-second timeout.
func (s *Context) Send(call ocpp.Call) (*ResultOrError, error) {
	return s.SendWithTimeout(call, 10*time.Second)
}

// SendCall sends a typed OCPP 2.1 CALL with a default 10-second timeout.
func (s *Context) SendCall(action Action, payload any) (*ResultOrError, error) {
	return s.Send(ocpp.Call{MessageType: ocpp.MessageTypeCall, Action: string(action), Payload: payload})
}

// SendWithTimeout sends a CALL with a custom timeout.
func (s *Context) SendWithTimeout(call ocpp.Call, timeout time.Duration) (*ResultOrError, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return s.SendWithContext(ctx, call)
}

// SendCallWithContext sends a typed OCPP 2.1 CALL with context cancellation.
func (s *Context) SendCallWithContext(ctx context.Context, action Action, payload any) (*ResultOrError, error) {
	return s.SendWithContext(ctx, ocpp.Call{MessageType: ocpp.MessageTypeCall, Action: string(action), Payload: payload})
}

// SendWithContext sends a CALL and waits until a response arrives or ctx is canceled.
func (s *Context) SendWithContext(ctx context.Context, call ocpp.Call) (*ResultOrError, error) {
	if call.MessageID == "" {
		call.MessageID = s.messageIDGenerator()
	}
	msg := Request{
		Call:   call,
		result: make(chan ResultOrError, 1),
	}

	s.storage.Add(msg)
	select {
	case s.Queue <- msg:
	case <-ctx.Done():
		return nil, fmt.Errorf("failed to queue call %s: %w", call.MessageID, ctx.Err())
	}

	select {
	case result := <-msg.result:
		return &result, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("call %s timed out: %w", call.MessageID, ctx.Err())
	}
}

// SendCallAndExpectResult sends a CALL and returns only a successful CALLRESULT.
func (s *Context) SendCallAndExpectResult(action string, payload any) (*ocpp.CallResult, error) {
	call := &ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: s.messageIDGenerator(), Action: action, Payload: payload}
	request, err := s.Send(*call)

	if err != nil {
		return nil, fmt.Errorf("Failed to send %s call, error: %s", action, err)
	} else if request.CallError != nil {
		return nil, fmt.Errorf("Received CallError when sending %s, error: %s", action, err)
	}
	return request.CallResult, nil
}
