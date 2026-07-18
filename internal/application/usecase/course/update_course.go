// Package course contains course management use cases.
package course

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// UpdateCourseUseCase handles course updates.
type UpdateCourseUseCase struct {
	courseRepo repository.CourseRepository
}

// NewUpdateCourseUseCase creates a new UpdateCourseUseCase.
func NewUpdateCourseUseCase(courseRepo repository.CourseRepository) *UpdateCourseUseCase {
	return &UpdateCourseUseCase{courseRepo: courseRepo}
}

// Execute updates a course.
func (uc *UpdateCourseUseCase) Execute(ctx context.Context, id string, req *dto.UpdateCourseRequest) (*dto.CourseResponse, error) {
	course, err := uc.courseRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if course == nil {
		return nil, ErrCourseNotFound
	}

	if req.Title != nil {
		course.Title = *req.Title
	}
	if req.Description != nil {
		course.Description = *req.Description
	}
	if req.CoverImage != nil {
		course.CoverImage = *req.CoverImage
	}
	if req.IsPublished != nil {
		course.IsPublished = *req.IsPublished
	}

	if err := uc.courseRepo.Update(ctx, course); err != nil {
		return nil, err
	}
	return toCourseResponse(course), nil
}
