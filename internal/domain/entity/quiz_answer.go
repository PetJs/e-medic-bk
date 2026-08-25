// Package entity contains the core domain entities.
package entity

import "time"

// QuizAnswer is a student's single, locked-in submission for one question.
type QuizAnswer struct {
	ID         string
	QuestionID string
	UserID     string

	SelectedOptionID *string // multiple_choice
	FreeTextBody     *string // free_text
	IsCorrect        *bool   // nil for free_text (ungraded)

	CreatedAt time.Time

	// Author is populated only by ListByQuestion (admin review) — the
	// caller of ListByUserAndLesson already knows who they are.
	Author *User
}
