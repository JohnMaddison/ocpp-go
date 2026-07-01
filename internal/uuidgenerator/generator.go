package uuidgenerator

import "github.com/google/uuid"

type MessageIDGeneratorMethod func() string

func DefaultUUIDGenerator() string {
	return uuid.New().String()
}
