// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// ContentRepository defines the interface for content metadata access.
type ContentRepository interface {
	Create(ctx context.Context, content *entity.Content) error
	GetByID(ctx context.Context, id string) (*entity.Content, error)
	Update(ctx context.Context, content *entity.Content) error
	Delete(ctx context.Context, id string) error
	ListByLesson(ctx context.Context, lessonID string) ([]*entity.Content, error)
}
