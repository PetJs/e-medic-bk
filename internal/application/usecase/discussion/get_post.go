package discussion

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// GetPostUseCase handles fetching a single discussion post.
type GetPostUseCase struct {
	postRepo repository.DiscussionPostRepository
}

// NewGetPostUseCase creates a new GetPostUseCase.
func NewGetPostUseCase(postRepo repository.DiscussionPostRepository) *GetPostUseCase {
	return &GetPostUseCase{postRepo: postRepo}
}

// Execute returns a single discussion post.
func (uc *GetPostUseCase) Execute(ctx context.Context, id string) (*dto.PostResponse, error) {
	post, err := uc.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, ErrPostNotFound
	}
	return ToPostResponse(post), nil
}
