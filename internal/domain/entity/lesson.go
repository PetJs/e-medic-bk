// Package entity contains the core domain entities.
package entity

import "time"

// Lesson represents a lesson within a module.
type Lesson struct {
	ID          string
	ModuleID    string
	Title       string
	Description string
	Order       int
	Duration    int // in minutes
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
