// Package course contains course management use cases.
package course

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// GetCourseUseCase handles retrieving course details.
type GetCourseUseCase struct {
	// TODO: Add dependencies
}

// NewGetCourseUseCase creates a new GetCourseUseCase.
func NewGetCourseUseCase() *GetCourseUseCase {
	return &GetCourseUseCase{}
}

// Execute retrieves course details by ID or slug.
func (uc *GetCourseUseCase) Execute(ctx context.Context, idOrSlug string) (*dto.CourseDetailResponse, error) {
	// TODO: Get course by ID or slug
	// TODO: Get modules
	// TODO: Map to response DTO
	return nil, nil
}
