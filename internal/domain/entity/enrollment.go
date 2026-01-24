// Package entity contains the core domain entities.
package entity

import "time"

// Enrollment represents a user's enrollment in a course.
type Enrollment struct {
	ID         string
	UserID     string
	CourseID   string
	EnrolledAt time.Time
	ExpiresAt  *time.Time // nil means no expiration
	IsActive   bool
}
