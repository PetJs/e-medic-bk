// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// QuestionRepository defines the interface for question data access.
type QuestionRepository interface {
	Create(ctx context.Context, question *entity.Question) error
	GetByID(ctx context.Context, id string) (*entity.Question, error)
	Update(ctx context.Context, question *entity.Question) error
	Delete(ctx context.Context, id string) error
	ListByLesson(ctx context.Context, lessonID string, limit, offset int) ([]*entity.Question, error)
	ListByUser(ctx context.Context, userID string, limit, offset int) ([]*entity.Question, error)
}
