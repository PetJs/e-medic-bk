// Package lesson contains lesson management use cases.
package lesson

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// ErrLessonNotFound is returned when a lesson does not exist.
var ErrLessonNotFound = errors.New("lesson not found")

// ErrModuleNotFound is returned when the parent module does not exist.
var ErrModuleNotFound = errors.New("module not found")

// ErrSubscriptionRequired is returned when premium content is accessed without an active subscription.
var ErrSubscriptionRequired = errors.New("active subscription required")

func toLessonResponse(l *entity.Lesson) *dto.LessonResponse {
	return &dto.LessonResponse{
		ID:          l.ID,
		ModuleID:    l.ModuleID,
		Title:       l.Title,
		Description: l.Description,
		Order:       l.Order,
		Duration:    l.Duration,
		CreatedAt:   l.CreatedAt,
	}
}

// CreateLessonUseCase handles lesson creation.
type CreateLessonUseCase struct {
	lessonRepo repository.LessonRepository
	moduleRepo repository.ModuleRepository
	idGen      port.IDGenerator
}

// NewCreateLessonUseCase creates a new CreateLessonUseCase.
func NewCreateLessonUseCase(
	lessonRepo repository.LessonRepository,
	moduleRepo repository.ModuleRepository,
	idGen port.IDGenerator,
) *CreateLessonUseCase {
	return &CreateLessonUseCase{lessonRepo: lessonRepo, moduleRepo: moduleRepo, idGen: idGen}
}

// Execute creates a new lesson.
func (uc *CreateLessonUseCase) Execute(ctx context.Context, req *dto.CreateLessonRequest) (*dto.LessonResponse, error) {
	module, err := uc.moduleRepo.GetByID(ctx, req.ModuleID)
	if err != nil {
		return nil, err
	}
	if module == nil {
		return nil, ErrModuleNotFound
	}

	now := time.Now()
	lesson := &entity.Lesson{
		ID:          uc.idGen.Generate(),
		ModuleID:    req.ModuleID,
		Title:       req.Title,
		Description: req.Description,
		Order:       req.Order,
		Duration:    req.Duration,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := uc.lessonRepo.Create(ctx, lesson); err != nil {
		return nil, err
	}
	return toLessonResponse(lesson), nil
}
