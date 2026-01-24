// Package uid provides unique ID generation utilities.
package uid

import (
	"github.com/google/uuid"

	"emedic-bk/internal/application/port"
)

// Generate generates a new UUID string.
func Generate() string {
	return uuid.New().String()
}

// IsValid checks if a string is a valid UUID.
func IsValid(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// MustParse parses a UUID string or panics.
func MustParse(id string) uuid.UUID {
	return uuid.MustParse(id)
}

// Generator implements port.IDGenerator using UUID.
type Generator struct{}

// NewGenerator creates a new UUID generator.
func NewGenerator() port.IDGenerator {
	return &Generator{}
}

// Generate generates a new UUID string.
func (g *Generator) Generate() string {
	return uuid.New().String()
}
