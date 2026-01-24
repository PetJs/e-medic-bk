// Package valueobject contains immutable value objects with validation.
package valueobject

import "errors"

// ErrInvalidRole is returned when a role is invalid.
var ErrInvalidRole = errors.New("invalid role")

// Role represents a user role in the system.
type Role string

const (
	RoleStudent Role = "student"
	RoleAdmin   Role = "admin"
)

// NewRole creates a new Role value object with validation.
func NewRole(role string) (Role, error) {
	switch Role(role) {
	case RoleStudent, RoleAdmin:
		return Role(role), nil
	default:
		return "", ErrInvalidRole
	}
}

// String returns the role as a string.
func (r Role) String() string {
	return string(r)
}

// IsAdmin returns true if the role is admin.
func (r Role) IsAdmin() bool {
	return r == RoleAdmin
}
