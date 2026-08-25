// Package discussion contains discussion board (posts + comments) use cases.
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

// ErrModuleNotFound is returned when the parent module does not exist.
var ErrModuleNotFound = errors.New("module not found")

// ErrPostNotFound is returned when a post does not exist, or the caller is
// not allowed to act on it (kept indistinguishable from "not found" so
// ownership checks don't leak existence of other users' posts).
var ErrPostNotFound = errors.New("post not found")

// ToPostResponse maps a discussion post entity to its API response.
func ToPostResponse(p *entity.DiscussionPost) *dto.PostResponse {
	resp := &dto.PostResponse{
		ID:           p.ID,
		ModuleID:     p.ModuleID,
		Title:        p.Title,
		Body:         p.Body,
		IsPinned:     p.IsPinned,
		CommentCount: p.CommentCount,
		CreatedAt:    p.CreatedAt,
	}
	if p.Author != nil {
		resp.Author = &dto.UserResponse{
			ID:        p.Author.ID,
			Email:     p.Author.Email,
			FirstName: p.Author.FirstName,
			LastName:  p.Author.LastName,
			Role:      p.Author.Role,
			CreatedAt: p.Author.CreatedAt,
		}
	}
	return resp
}

// CreatePostUseCase handles discussion post creation.
type CreatePostUseCase struct {
	postRepo   repository.DiscussionPostRepository
	moduleRepo repository.ModuleRepository
	idGen      port.IDGenerator
}

// NewCreatePostUseCase creates a new CreatePostUseCase.
func NewCreatePostUseCase(
	postRepo repository.DiscussionPostRepository,
	moduleRepo repository.ModuleRepository,
	idGen port.IDGenerator,
) *CreatePostUseCase {
	return &CreatePostUseCase{postRepo: postRepo, moduleRepo: moduleRepo, idGen: idGen}
}

// Execute creates a new discussion post under a module.
func (uc *CreatePostUseCase) Execute(ctx context.Context, moduleID, userID string, req *dto.CreatePostRequest) (*dto.PostResponse, error) {
	mod, err := uc.moduleRepo.GetByID(ctx, moduleID)
	if err != nil {
		return nil, err
	}
	if mod == nil {
		return nil, ErrModuleNotFound
	}

	now := time.Now()
	post := &entity.DiscussionPost{
		ID:        uc.idGen.Generate(),
		ModuleID:  moduleID,
		UserID:    userID,
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := uc.postRepo.Create(ctx, post); err != nil {
		return nil, err
	}

	created, err := uc.postRepo.GetByID(ctx, post.ID)
	if err != nil {
		return nil, err
	}
	return ToPostResponse(created), nil
}
