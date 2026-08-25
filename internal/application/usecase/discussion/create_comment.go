package discussion

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// ErrParentCommentNotFound is returned when a reply targets a parent comment
// that doesn't exist, or that belongs to a different post.
var ErrParentCommentNotFound = errors.New("parent comment not found")

// ToCommentResponse maps a discussion comment entity to its API response.
func ToCommentResponse(c *entity.DiscussionComment) *dto.CommentResponse {
	resp := &dto.CommentResponse{
		ID:              c.ID,
		PostID:          c.PostID,
		ParentCommentID: c.ParentCommentID,
		Body:            c.Body,
		CreatedAt:       c.CreatedAt,
	}
	if c.Author != nil {
		resp.Author = &dto.UserResponse{
			ID:        c.Author.ID,
			Email:     c.Author.Email,
			FirstName: c.Author.FirstName,
			LastName:  c.Author.LastName,
			Role:      c.Author.Role,
			CreatedAt: c.Author.CreatedAt,
		}
	}
	return resp
}

// CreateCommentUseCase handles discussion comment creation.
type CreateCommentUseCase struct {
	commentRepo repository.DiscussionCommentRepository
	postRepo    repository.DiscussionPostRepository
	idGen       port.IDGenerator
}

// NewCreateCommentUseCase creates a new CreateCommentUseCase.
func NewCreateCommentUseCase(
	commentRepo repository.DiscussionCommentRepository,
	postRepo repository.DiscussionPostRepository,
	idGen port.IDGenerator,
) *CreateCommentUseCase {
	return &CreateCommentUseCase{commentRepo: commentRepo, postRepo: postRepo, idGen: idGen}
}

// Execute creates a new comment (or threaded reply) on a post.
func (uc *CreateCommentUseCase) Execute(ctx context.Context, postID, userID string, req *dto.CreateCommentRequest) (*dto.CommentResponse, error) {
	post, err := uc.postRepo.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, ErrPostNotFound
	}

	if req.ParentCommentID != nil {
		parent, err := uc.commentRepo.GetByID(ctx, *req.ParentCommentID)
		if err != nil {
			return nil, err
		}
		if parent == nil || parent.PostID != postID {
			return nil, ErrParentCommentNotFound
		}
	}

	now := time.Now()
	comment := &entity.DiscussionComment{
		ID:              uc.idGen.Generate(),
		PostID:          postID,
		UserID:          userID,
		ParentCommentID: req.ParentCommentID,
		Body:            req.Body,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	if err := uc.commentRepo.Create(ctx, comment); err != nil {
		return nil, err
	}

	created, err := uc.commentRepo.GetByID(ctx, comment.ID)
	if err != nil {
		return nil, err
	}
	return ToCommentResponse(created), nil
}
