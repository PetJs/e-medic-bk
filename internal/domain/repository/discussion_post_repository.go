// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// DiscussionPostRepository defines the interface for discussion post data access.
type DiscussionPostRepository interface {
	Create(ctx context.Context, post *entity.DiscussionPost) error
	GetByID(ctx context.Context, id string) (*entity.DiscussionPost, error)
	Delete(ctx context.Context, id string) error
	ListByModule(ctx context.Context, moduleID string, limit, offset int) ([]*entity.DiscussionPost, error)
	CountByModule(ctx context.Context, moduleID string) (int64, error)
	SetPinned(ctx context.Context, id string, pinned bool) error
}
