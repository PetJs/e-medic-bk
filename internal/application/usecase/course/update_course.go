// Package course contains course management use cases.
package course

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// UpdateCourseUseCase handles course updates.
type UpdateCourseUseCase struct {
	// TODO: Add dependencies
}

// NewUpdateCourseUseCase creates a new UpdateCourseUseCase.
func NewUpdateCourseUseCase() *UpdateCourseUseCase {
	return &UpdateCourseUseCase{}
}

// Execute updates an existing course.
func (uc *UpdateCourseUseCase) Execute(ctx context.Context, courseID string, req *dto.UpdateCourseRequest) (*dto.CourseResponse, error) {
	// TODO: Validate input
	// TODO: Get course
	// TODO: Update fields
	// TODO: Save course
	return nil, nil
}
