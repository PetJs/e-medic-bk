// Package enrollment contains enrollment use cases.
package enrollment

import "context"

// UnenrollCourseUseCase handles unenrolling a user from a course.
type UnenrollCourseUseCase struct{}

func NewUnenrollCourseUseCase() *UnenrollCourseUseCase { return &UnenrollCourseUseCase{} }

func (uc *UnenrollCourseUseCase) Execute(ctx context.Context, userID, courseID string) error {
	return nil
}
