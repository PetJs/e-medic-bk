// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// LessonRepository defines the interface for lesson data access.
type LessonRepository interface {
	Create(ctx context.Context, lesson *entity.Lesson) error
	GetByID(ctx context.Context, id string) (*entity.Lesson, error)
	Update(ctx context.Context, lesson *entity.Lesson) error
	Delete(ctx context.Context, id string) error
	ListByModule(ctx context.Context, moduleID string) ([]*entity.Lesson, error)
}
