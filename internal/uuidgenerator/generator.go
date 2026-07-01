package uuidgenerator

import "github.com/google/uuid"

type MessageIdGeneratorMethod func() string

func DefaultUUIDGenerator() string {
	return uuid.New().String()
}
