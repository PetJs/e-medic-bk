// Package enrollment contains enrollment use cases.
package enrollment

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// EnrollCourseUseCase handles enrolling a user in a course.
type EnrollCourseUseCase struct{}

func NewEnrollCourseUseCase() *EnrollCourseUseCase { return &EnrollCourseUseCase{} }

func (uc *EnrollCourseUseCase) Execute(ctx context.Context, userID, courseID string) (*dto.EnrollmentResponse, error) {
	// TODO: Check if already enrolled
	// TODO: Create enrollment
	return nil, nil
}
