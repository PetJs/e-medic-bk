package discussion

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// ListCommentsUseCase handles listing every comment on a post.
type ListCommentsUseCase struct {
	commentRepo repository.DiscussionCommentRepository
	postRepo    repository.DiscussionPostRepository
}

// NewListCommentsUseCase creates a new ListCommentsUseCase.
func NewListCommentsUseCase(commentRepo repository.DiscussionCommentRepository, postRepo repository.DiscussionPostRepository) *ListCommentsUseCase {
	return &ListCommentsUseCase{commentRepo: commentRepo, postRepo: postRepo}
}

// Execute returns every comment on a post, unpaginated, so the client can
// build the reply tree from the flat parent_comment_id list.
func (uc *ListCommentsUseCase) Execute(ctx context.Context, postID string) (*dto.ListCommentsResponse, error) {
	post, err := uc.postRepo.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, ErrPostNotFound
	}

	comments, err := uc.commentRepo.ListByPost(ctx, postID)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.CommentResponse, 0, len(comments))
	for _, c := range comments {
		responses = append(responses, ToCommentResponse(c))
	}
	return &dto.ListCommentsResponse{Comments: responses}, nil
}
