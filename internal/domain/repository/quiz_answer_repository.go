// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"
	"errors"

	"emedic-bk/internal/domain/entity"
)

// ErrAlreadyAnswered is returned by Create when the student has already
// submitted an answer for this question (a concurrent submission won the
// race, or they're attempting a disallowed second attempt).
var ErrAlreadyAnswered = errors.New("question already answered")

// QuizAnswerRepository defines the interface for lesson quiz answer data access.
type QuizAnswerRepository interface {
	Create(ctx context.Context, a *entity.QuizAnswer) error
	ListByUserAndLesson(ctx context.Context, userID, lessonID string) ([]*entity.QuizAnswer, error)
	ListByQuestion(ctx context.Context, questionID string) ([]*entity.QuizAnswer, error)
	CountAnsweredByUser(ctx context.Context, lessonID, userID string) (int, error)
}
