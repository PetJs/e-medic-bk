// Package content contains content management use cases.
package content

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// ErrContentNotFound is returned when the content does not exist.
var ErrContentNotFound = errors.New("content not found")

// ErrSubscriptionRequired is returned when premium content is accessed
// without an active subscription.
var ErrSubscriptionRequired = errors.New("active subscription required")

// URLExpiry is how long signed streaming URLs stay valid. Short-lived on
// purpose: links can't be shared meaningfully outside the player.
const URLExpiry = 15 * time.Minute

// GetContentURLUseCase generates signed URLs for content streaming.
type GetContentURLUseCase struct {
	contentRepo repository.ContentRepository
	lessonRepo  repository.LessonRepository
	moduleRepo  repository.ModuleRepository
	subRepo     repository.SubscriptionRepository
	storage     service.StorageService
}

// NewGetContentURLUseCase creates a new GetContentURLUseCase.
func NewGetContentURLUseCase(
	contentRepo repository.ContentRepository,
	lessonRepo repository.LessonRepository,
	moduleRepo repository.ModuleRepository,
	subRepo repository.SubscriptionRepository,
	storage service.StorageService,
) *GetContentURLUseCase {
	return &GetContentURLUseCase{
		contentRepo: contentRepo,
		lessonRepo:  lessonRepo,
		moduleRepo:  moduleRepo,
		subRepo:     subRepo,
		storage:     storage,
	}
}

// Execute returns a short-lived signed URL after enforcing the premium gate.
func (uc *GetContentURLUseCase) Execute(ctx context.Context, contentID, userID, role string) (*dto.ContentURLResponse, error) {
	content, err := uc.contentRepo.GetByID(ctx, contentID)
	if err != nil {
		return nil, err
	}
	if content == nil {
		return nil, ErrContentNotFound
	}

	lesson, err := uc.lessonRepo.GetByID(ctx, content.LessonID)
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, ErrContentNotFound
	}

	module, err := uc.moduleRepo.GetByID(ctx, lesson.ModuleID)
	if err != nil {
		return nil, err
	}
	if module == nil {
		return nil, ErrContentNotFound
	}

	if module.IsPremium && role != "admin" {
		sub, err := uc.subRepo.GetActiveByUser(ctx, userID)
		if err != nil {
			return nil, err
		}
		if sub == nil {
			return nil, ErrSubscriptionRequired
		}
	}

	url, err := uc.storage.GetSignedURL(ctx, content.StorageKey, URLExpiry)
	if err != nil {
		return nil, err
	}

	return &dto.ContentURLResponse{
		URL:       url,
		ExpiresAt: time.Now().Add(URLExpiry),
	}, nil
}
