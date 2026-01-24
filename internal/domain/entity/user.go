// Package entity contains the core domain entities.
package entity

import "time"

// User represents a user in the system.
type User struct {
	ID           string
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	Role         string // "student" or "admin"
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
