// Package entity contains the core domain entities.
package entity

import "time"

// QuestionType is the shape of a lesson quiz question.
type QuestionType string

const (
	QuestionTypeMultipleChoice QuestionType = "multiple_choice"
	QuestionTypeFreeText       QuestionType = "free_text"
)

// QuizQuestion represents one admin-authored question attached to a lesson.
type QuizQuestion struct {
	ID       string
	LessonID string
	Type     QuestionType
	Prompt   string
	Order    int

	CreatedAt time.Time
	UpdatedAt time.Time

	// Options is populated by ListByLesson/GetByID for multiple_choice
	// questions (not a stored column on this table).
	Options []*QuizOption
}
