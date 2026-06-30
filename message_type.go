// Package ocpp contains shared OCPP message envelope types and helpers.
package ocpp

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// MessageType constants define the type of the OCPP message.
const (
	MessageTypeCall       = 2
	MessageTypeCallResult = 3
	MessageTypeCallError  = 4
)

// Call represents a request from a client to a server.
type Call struct {
	MessageType int    `json:"-"`
	MessageID   string `json:"-"`
	Action      string `json:"-"`
	Payload     any    `json:"-"`
}

// SerializeOCPP converts a Call object to its JSON representation.
func (c *Call) SerializeOCPP() ([]byte, error) {
	message := []any{
		c.MessageType,
		c.MessageID,
		c.Action,
		c.Payload,
	}
	return json.Marshal(message)
}

func NewCall(action string, payload any) *Call {
	return &Call{MessageType: MessageTypeCall, MessageID: uuid.New().String(), Action: action, Payload: payload}
}

// CallResult represents a successful response to a Call.
type CallResult struct {
	MessageType int    `json:"-"`
	MessageID   string `json:"-"`
	Payload     any    `json:"-"`
}

// SerializeOCPP converts a CallResult object to its JSON representation.
func (cr *CallResult) SerializeOCPP() ([]byte, error) {
	message := []any{
		cr.MessageType,
		cr.MessageID,
		cr.Payload,
	}
	return json.Marshal(message)
}

// CallError represents an error response to a Call.
type CallError struct {
	MessageType      int    `json:"-"`
	MessageID        string `json:"-"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	ErrorDetails     any    `json:"errorDetails,omitempty"`
}

// SerializeOCPP converts a CallError object to its JSON representation.
func (ce *CallError) SerializeOCPP() ([]byte, error) {
	message := []any{
		ce.MessageType,
		ce.MessageID,
		ce.ErrorCode,
		ce.ErrorDescription,
		ce.ErrorDetails,
	}
	// The OCPP spec requires the details object to be present, even if empty.
	if ce.ErrorDetails == nil {
		message[4] = struct{}{}
	}
	return json.Marshal(message)
}

// Implement Stringer
func (c CallError) String() string {
	// Marshal only the exported JSON fields
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Sprintf("CallError{ErrorCode:%q, ErrorDescription:%q, ErrorDetails:%v}",
			c.ErrorCode, c.ErrorDescription, c.ErrorDetails)
	}
	return string(data)
}
