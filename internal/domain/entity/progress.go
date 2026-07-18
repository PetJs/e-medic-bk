// Package entity contains the core domain entities.
package entity

import "time"

// Progress represents a user's progress on a lesson.
type Progress struct {
	ID           string
	UserID       string
	LessonID     string
	IsCompleted  bool
	ProgressPct  int // 0-100
	LastPosition int // for videos, in seconds
	CompletedAt  *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// ModuleID is a read-time aggregate joined from the lesson (list queries only).
	ModuleID string
}
