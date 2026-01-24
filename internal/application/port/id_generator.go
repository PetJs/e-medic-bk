// Package port defines secondary port interfaces for the application layer.
package port

// IDGenerator defines the interface for generating unique IDs.
type IDGenerator interface {
	Generate() string
}
