// Package entity contains the core domain entities.
package entity

import "time"

// Module represents a module within a course.
type Module struct {
	ID               string
	CourseID         string
	Title            string
	Description      string
	Order            int
	IsPremium        bool // Requires subscription
	CoverImageKey    string
	CoverImageStatus string // "pending" | "ready" | "failed"
	CreatedAt        time.Time
	UpdatedAt        time.Time

	// Read-time aggregates populated by list queries (not stored columns).
	LessonCount   int
	TotalDuration int // sum of lesson durations, in seconds
}
