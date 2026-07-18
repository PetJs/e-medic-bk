// Package lesson contains lesson management use cases.
package lesson

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// GetLessonUseCase handles fetching a lesson with its contents,
// enforcing the premium-access rule.
type GetLessonUseCase struct {
	lessonRepo  repository.LessonRepository
	moduleRepo  repository.ModuleRepository
	contentRepo repository.ContentRepository
	subRepo     repository.SubscriptionRepository
}

// NewGetLessonUseCase creates a new GetLessonUseCase.
func NewGetLessonUseCase(
	lessonRepo repository.LessonRepository,
	moduleRepo repository.ModuleRepository,
	contentRepo repository.ContentRepository,
	subRepo repository.SubscriptionRepository,
) *GetLessonUseCase {
	return &GetLessonUseCase{
		lessonRepo:  lessonRepo,
		moduleRepo:  moduleRepo,
		contentRepo: contentRepo,
		subRepo:     subRepo,
	}
}

// Execute fetches a lesson and its contents. Premium lessons require an
// active subscription unless the caller is an admin.
func (uc *GetLessonUseCase) Execute(ctx context.Context, lessonID, userID, role string) (*dto.LessonDetailResponse, error) {
	lesson, err := uc.lessonRepo.GetByID(ctx, lessonID)
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, ErrLessonNotFound
	}

	module, err := uc.moduleRepo.GetByID(ctx, lesson.ModuleID)
	if err != nil {
		return nil, err
	}
	if module == nil {
		return nil, ErrModuleNotFound
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

	contents, err := uc.contentRepo.ListByLesson(ctx, lessonID)
	if err != nil {
		return nil, err
	}

	contentResponses := make([]*dto.ContentResponse, 0, len(contents))
	for _, c := range contents {
		contentResponses = append(contentResponses, &dto.ContentResponse{
			ID:        c.ID,
			LessonID:  c.LessonID,
			Type:      string(c.Type),
			Title:     c.Title,
			MimeType:  c.MimeType,
			Size:      c.Size,
			Duration:  c.Duration,
			Order:     c.Order,
			CreatedAt: c.CreatedAt,
		})
	}

	return &dto.LessonDetailResponse{
		LessonResponse: *toLessonResponse(lesson),
		Contents:       contentResponses,
	}, nil
}
