// Package entity contains the core domain entities.
package entity

import "time"

// Answer represents an answer to a question.
type Answer struct {
	ID         string
	QuestionID string
	UserID     string
	Body       string
	IsBest     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
