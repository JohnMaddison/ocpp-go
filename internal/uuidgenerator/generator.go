package uuidgenerator

import "github.com/google/uuid"

// MessageIDGeneratorMethod creates OCPP message IDs.
type MessageIDGeneratorMethod func() string

// DefaultUUIDGenerator returns a random UUID string.
func DefaultUUIDGenerator() string {
	return uuid.New().String()
}
