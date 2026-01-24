// Package entity contains the core domain entities.
package entity

import "time"

// Question represents a question in the Q&A system.
type Question struct {
	ID          string
	LessonID    string
	UserID      string
	Title       string
	Body        string
	AnswerCount int
	IsResolved  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
