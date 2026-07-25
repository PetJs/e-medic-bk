// Package course contains course management use cases.
package course

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/module"
	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// GetCourseUseCase handles fetching a course with its modules.
type GetCourseUseCase struct {
	courseRepo repository.CourseRepository
	moduleRepo repository.ModuleRepository
	storageSvc service.StorageService
}

// NewGetCourseUseCase creates a new GetCourseUseCase.
func NewGetCourseUseCase(courseRepo repository.CourseRepository, moduleRepo repository.ModuleRepository, storageSvc service.StorageService) *GetCourseUseCase {
	return &GetCourseUseCase{courseRepo: courseRepo, moduleRepo: moduleRepo, storageSvc: storageSvc}
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
		moduleResponses = append(moduleResponses, module.ToModuleResponse(ctx, m, uc.storageSvc))
	}

	return &dto.CourseDetailResponse{
		CourseResponse: *toCourseResponse(course),
		Modules:        moduleResponses,
	}, nil
}
