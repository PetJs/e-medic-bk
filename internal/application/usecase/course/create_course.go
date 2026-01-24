// Package course contains course management use cases.
package course

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// CreateCourseUseCase handles course creation.
type CreateCourseUseCase struct {
	// TODO: Add dependencies
}

// NewCreateCourseUseCase creates a new CreateCourseUseCase.
func NewCreateCourseUseCase() *CreateCourseUseCase {
	return &CreateCourseUseCase{}
}

// Execute creates a new course.
func (uc *CreateCourseUseCase) Execute(ctx context.Context, authorID string, req *dto.CreateCourseRequest) (*dto.CourseResponse, error) {
	// TODO: Validate input
	// TODO: Generate slug
	// TODO: Create course entity
	// TODO: Save course
	return nil, nil
}
