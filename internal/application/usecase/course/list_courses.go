// Package course contains course management use cases.
package course

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// ListCoursesUseCase handles listing courses.
type ListCoursesUseCase struct {
	// TODO: Add dependencies
}

// NewListCoursesUseCase creates a new ListCoursesUseCase.
func NewListCoursesUseCase() *ListCoursesUseCase {
	return &ListCoursesUseCase{}
}

// Execute lists courses with pagination and filters.
func (uc *ListCoursesUseCase) Execute(ctx context.Context, req *dto.ListCoursesRequest) (*dto.ListCoursesResponse, error) {
	// TODO: List courses with filters
	// TODO: Map to response DTOs
	return nil, nil
}
