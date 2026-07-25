// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// ModuleRepository defines the interface for module data access.
type ModuleRepository interface {
	Create(ctx context.Context, module *entity.Module) error
	GetByID(ctx context.Context, id string) (*entity.Module, error)
	Update(ctx context.Context, module *entity.Module) error
	Delete(ctx context.Context, id string) error
	ListByCourse(ctx context.Context, courseID string) ([]*entity.Module, error)
	// ListAll returns modules across all published courses.
	ListAll(ctx context.Context) ([]*entity.Module, error)
	// UpdateCoverImage sets the generated cover image's storage key and status.
	UpdateCoverImage(ctx context.Context, id, coverImageKey, coverImageStatus string) error
}
