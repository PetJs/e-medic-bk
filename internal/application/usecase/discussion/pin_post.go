package discussion

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// PinPostUseCase handles pinning a discussion post (admin only — enforced by
// the route's role middleware, not this use case).
type PinPostUseCase struct {
	postRepo repository.DiscussionPostRepository
}

// NewPinPostUseCase creates a new PinPostUseCase.
func NewPinPostUseCase(postRepo repository.DiscussionPostRepository) *PinPostUseCase {
	return &PinPostUseCase{postRepo: postRepo}
}

// Execute pins a post so it sorts first on its module's board.
func (uc *PinPostUseCase) Execute(ctx context.Context, id string) (*dto.PostResponse, error) {
	return setPinned(ctx, uc.postRepo, id, true)
}

// UnpinPostUseCase handles unpinning a discussion post (admin only).
type UnpinPostUseCase struct {
	postRepo repository.DiscussionPostRepository
}

// NewUnpinPostUseCase creates a new UnpinPostUseCase.
func NewUnpinPostUseCase(postRepo repository.DiscussionPostRepository) *UnpinPostUseCase {
	return &UnpinPostUseCase{postRepo: postRepo}
}

// Execute unpins a post, returning it to normal newest-first ordering.
func (uc *UnpinPostUseCase) Execute(ctx context.Context, id string) (*dto.PostResponse, error) {
	return setPinned(ctx, uc.postRepo, id, false)
}

func setPinned(ctx context.Context, repo repository.DiscussionPostRepository, id string, pinned bool) (*dto.PostResponse, error) {
	post, err := repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, ErrPostNotFound
	}
	if err := repo.SetPinned(ctx, id, pinned); err != nil {
		return nil, err
	}
	post.IsPinned = pinned
	return ToPostResponse(post), nil
}
