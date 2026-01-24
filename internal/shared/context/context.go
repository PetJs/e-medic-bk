// Package context provides context key helpers.
package context

type contextKey string

const (
	// UserIDKey is the context key for user ID.
	UserIDKey contextKey = "user_id"

	// RoleKey is the context key for user role.
	RoleKey contextKey = "role"

	// RequestIDKey is the context key for request ID.
	RequestIDKey contextKey = "request_id"
)
