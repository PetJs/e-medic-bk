// Package course contains course management use cases.
package course

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// GetCourseUseCase handles fetching a course with its modules.
type GetCourseUseCase struct {
	courseRepo repository.CourseRepository
	moduleRepo repository.ModuleRepository
}

// NewGetCourseUseCase creates a new GetCourseUseCase.
func NewGetCourseUseCase(courseRepo repository.CourseRepository, moduleRepo repository.ModuleRepository) *GetCourseUseCase {
	return &GetCourseUseCase{courseRepo: courseRepo, moduleRepo: moduleRepo}
}

// Execute fetches a course and its modules.
func (uc *GetCourseUseCase) Execute(ctx context.Context, id string) (*dto.CourseDetailResponse, error) {
	course, err := uc.courseRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if course == nil {
		return nil, ErrCourseNotFound
	}

	modules, err := uc.moduleRepo.ListByCourse(ctx, id)
	if err != nil {
		return nil, err
	}

	moduleResponses := make([]*dto.ModuleResponse, 0, len(modules))
	for _, m := range modules {
		moduleResponses = append(moduleResponses, &dto.ModuleResponse{
			ID:            m.ID,
			CourseID:      m.CourseID,
			Title:         m.Title,
			Description:   m.Description,
			Order:         m.Order,
			IsPremium:     m.IsPremium,
			LessonCount:   m.LessonCount,
			TotalDuration: m.TotalDuration,
			CreatedAt:     m.CreatedAt,
		})
	}

	return &dto.CourseDetailResponse{
		CourseResponse: *toCourseResponse(course),
		Modules:        moduleResponses,
	}, nil
}
