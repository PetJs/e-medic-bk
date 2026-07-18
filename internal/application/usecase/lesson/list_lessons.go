// Package lesson contains lesson management use cases.
package lesson

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// ListLessonsUseCase handles listing lessons within a module.
type ListLessonsUseCase struct {
	lessonRepo repository.LessonRepository
	moduleRepo repository.ModuleRepository
}

// NewListLessonsUseCase creates a new ListLessonsUseCase.
func NewListLessonsUseCase(
	lessonRepo repository.LessonRepository,
	moduleRepo repository.ModuleRepository,
) *ListLessonsUseCase {
	return &ListLessonsUseCase{lessonRepo: lessonRepo, moduleRepo: moduleRepo}
}

// Execute lists a module's lessons. Titles are visible to everyone —
// content access is enforced when fetching an individual lesson.
func (uc *ListLessonsUseCase) Execute(ctx context.Context, moduleID string) ([]*dto.LessonResponse, error) {
	module, err := uc.moduleRepo.GetByID(ctx, moduleID)
	if err != nil {
		return nil, err
	}
	if module == nil {
		return nil, ErrModuleNotFound
	}

	lessons, err := uc.lessonRepo.ListByModule(ctx, moduleID)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.LessonResponse, 0, len(lessons))
	for _, l := range lessons {
		responses = append(responses, toLessonResponse(l))
	}
	return responses, nil
}
