// Package course contains course management use cases.
package course

import (
	"context"

	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// DeleteCourseUseCase handles course deletion.
type DeleteCourseUseCase struct {
	courseRepo  repository.CourseRepository
	moduleRepo  repository.ModuleRepository
	lessonRepo  repository.LessonRepository
	contentRepo repository.ContentRepository
	storage     service.StorageService
}

// NewDeleteCourseUseCase creates a new DeleteCourseUseCase.
func NewDeleteCourseUseCase(
	courseRepo repository.CourseRepository,
	moduleRepo repository.ModuleRepository,
	lessonRepo repository.LessonRepository,
	contentRepo repository.ContentRepository,
	storage service.StorageService,
) *DeleteCourseUseCase {
	return &DeleteCourseUseCase{
		courseRepo:  courseRepo,
		moduleRepo:  moduleRepo,
		lessonRepo:  lessonRepo,
		contentRepo: contentRepo,
		storage:     storage,
	}
}

// Execute deletes a course and everything under it. Module, lesson, and
// content rows go via DB cascade; stored objects are removed best effort first.
func (uc *DeleteCourseUseCase) Execute(ctx context.Context, id string) error {
	course, err := uc.courseRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if course == nil {
		return ErrCourseNotFound
	}

	if modules, err := uc.moduleRepo.ListByCourse(ctx, id); err == nil {
		for _, module := range modules {
			if lessons, err := uc.lessonRepo.ListByModule(ctx, module.ID); err == nil {
				for _, lesson := range lessons {
					if contents, err := uc.contentRepo.ListByLesson(ctx, lesson.ID); err == nil {
						for _, content := range contents {
							_ = uc.storage.Delete(ctx, content.StorageKey)
						}
					}
				}
			}
		}
	}
	return uc.courseRepo.Delete(ctx, id)
}
