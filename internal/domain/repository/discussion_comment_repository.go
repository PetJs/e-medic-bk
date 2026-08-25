// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// DiscussionCommentRepository defines the interface for discussion comment data access.
type DiscussionCommentRepository interface {
	Create(ctx context.Context, comment *entity.DiscussionComment) error
	GetByID(ctx context.Context, id string) (*entity.DiscussionComment, error)
	Delete(ctx context.Context, id string) error
	// ListByPost returns every comment on a post, unpaginated, ordered oldest
	// first so callers can build the reply tree client-side.
	ListByPost(ctx context.Context, postID string) ([]*entity.DiscussionComment, error)
}
