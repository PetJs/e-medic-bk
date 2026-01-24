// Package entity contains the core domain entities.
package entity

import "time"

// Module represents a module within a course.
type Module struct {
	ID          string
	CourseID    string
	Title       string
	Description string
	Order       int
	IsPremium   bool // Requires subscription
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
