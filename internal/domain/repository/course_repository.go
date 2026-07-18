// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// CourseRepository defines the interface for course data access.
type CourseRepository interface {
	Create(ctx context.Context, course *entity.Course) error
	GetByID(ctx context.Context, id string) (*entity.Course, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Course, error)
	Update(ctx context.Context, course *entity.Course) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*entity.Course, error)
	ListByAuthor(ctx context.Context, authorID string, limit, offset int) ([]*entity.Course, error)
	Count(ctx context.Context) (int64, error)
}
