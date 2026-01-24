// Package port defines secondary port interfaces for the application layer.
package port

// Hasher defines the interface for password hashing.
type Hasher interface {
	Hash(password string) (string, error)
	Compare(password, hash string) error
}
