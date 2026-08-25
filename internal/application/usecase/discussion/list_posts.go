package discussion

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/shared/pagination"
)

// ListPostsUseCase handles listing a module's discussion posts.
type ListPostsUseCase struct {
	postRepo   repository.DiscussionPostRepository
	moduleRepo repository.ModuleRepository
}

// NewListPostsUseCase creates a new ListPostsUseCase.
func NewListPostsUseCase(postRepo repository.DiscussionPostRepository, moduleRepo repository.ModuleRepository) *ListPostsUseCase {
	return &ListPostsUseCase{postRepo: postRepo, moduleRepo: moduleRepo}
}

// Execute lists a module's discussion posts, pinned first then newest first.
func (uc *ListPostsUseCase) Execute(ctx context.Context, moduleID string, req *dto.ListPostsRequest) (*dto.ListPostsResponse, error) {
	mod, err := uc.moduleRepo.GetByID(ctx, moduleID)
	if err != nil {
		return nil, err
	}
	if mod == nil {
		return nil, ErrModuleNotFound
	}

	p := pagination.Pagination{Page: req.Page, Limit: req.Limit}
	p.Normalize()

	posts, err := uc.postRepo.ListByModule(ctx, moduleID, p.Limit, p.Offset())
	if err != nil {
		return nil, err
	}
	count, err := uc.postRepo.CountByModule(ctx, moduleID)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.PostResponse, 0, len(posts))
	for _, post := range posts {
		responses = append(responses, ToPostResponse(post))
	}

	return &dto.ListPostsResponse{
		Posts:      responses,
		TotalCount: count,
		Page:       p.Page,
		Limit:      p.Limit,
	}, nil
}
