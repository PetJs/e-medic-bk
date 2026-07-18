// Package course contains course management use cases.
package course

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/shared/pagination"
)

// ListCoursesUseCase handles listing courses.
type ListCoursesUseCase struct {
	courseRepo repository.CourseRepository
}

// NewListCoursesUseCase creates a new ListCoursesUseCase.
func NewListCoursesUseCase(courseRepo repository.CourseRepository) *ListCoursesUseCase {
	return &ListCoursesUseCase{courseRepo: courseRepo}
}

// Execute lists courses with pagination.
func (uc *ListCoursesUseCase) Execute(ctx context.Context, req *dto.ListCoursesRequest) (*dto.ListCoursesResponse, error) {
	p := pagination.Pagination{Page: req.Page, Limit: req.Limit}
	p.Normalize()

	courses, err := uc.courseRepo.List(ctx, p.Limit, p.Offset())
	if err != nil {
		return nil, err
	}

	count, err := uc.courseRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.CourseResponse, 0, len(courses))
	for _, c := range courses {
		responses = append(responses, toCourseResponse(c))
	}

	return &dto.ListCoursesResponse{
		Courses:    responses,
		TotalCount: count,
		Page:       p.Page,
		Limit:      p.Limit,
	}, nil
}
