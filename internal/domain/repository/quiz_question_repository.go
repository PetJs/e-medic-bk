// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// QuizQuestionRepository defines the interface for lesson quiz question data access.
type QuizQuestionRepository interface {
	// CreateWithOptions inserts the question and (for multiple_choice) its
	// options atomically, so a question is never persisted half-written.
	CreateWithOptions(ctx context.Context, q *entity.QuizQuestion, options []*entity.QuizOption) error
	GetByID(ctx context.Context, id string) (*entity.QuizQuestion, error) // includes Options
	Update(ctx context.Context, q *entity.QuizQuestion) error
	Delete(ctx context.Context, id string) error
	ListByLesson(ctx context.Context, lessonID string) ([]*entity.QuizQuestion, error) // includes Options
	CountByLesson(ctx context.Context, lessonID string) (int, error)
}
