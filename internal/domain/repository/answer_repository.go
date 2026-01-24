// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// AnswerRepository defines the interface for answer data access.
type AnswerRepository interface {
	Create(ctx context.Context, answer *entity.Answer) error
	GetByID(ctx context.Context, id string) (*entity.Answer, error)
	Update(ctx context.Context, answer *entity.Answer) error
	Delete(ctx context.Context, id string) error
	ListByQuestion(ctx context.Context, questionID string) ([]*entity.Answer, error)
	ClearBestAnswer(ctx context.Context, questionID string) error
}
