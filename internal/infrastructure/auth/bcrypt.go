// Package auth provides authentication infrastructure implementations.
package auth

import (
	"golang.org/x/crypto/bcrypt"

	"emedic-bk/internal/application/port"
)

// BcryptHasher implements port.Hasher using bcrypt.
type BcryptHasher struct {
	cost int
}

// NewBcryptHasher creates a new bcrypt hasher.
func NewBcryptHasher(cost int) port.Hasher {
	if cost < bcrypt.MinCost {
		cost = bcrypt.DefaultCost
	}
	return &BcryptHasher{cost: cost}
}

// Hash hashes a password using bcrypt.
func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Compare compares a password with a hash.
func (h *BcryptHasher) Compare(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
