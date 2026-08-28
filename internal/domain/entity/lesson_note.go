// Package entity contains the core domain entities.
package entity

import "time"

// LessonNote is a student's own private, timestamped note on a lesson —
// tied to a moment in the video (VideoPosition, in seconds).
type LessonNote struct {
	ID            string
	UserID        string
	LessonID      string
	Body          string
	VideoPosition int

	CreatedAt time.Time
	UpdatedAt time.Time
}
