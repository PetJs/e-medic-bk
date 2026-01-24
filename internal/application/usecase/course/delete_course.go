// Package course contains course management use cases.
package course

import "context"

// DeleteCourseUseCase handles course deletion.
type DeleteCourseUseCase struct {
	// TODO: Add dependencies
}

// NewDeleteCourseUseCase creates a new DeleteCourseUseCase.
func NewDeleteCourseUseCase() *DeleteCourseUseCase {
	return &DeleteCourseUseCase{}
}

// Execute deletes a course.
func (uc *DeleteCourseUseCase) Execute(ctx context.Context, courseID string) error {
	// TODO: Check if course exists
	// TODO: Delete associated modules, lessons, content
	// TODO: Delete course
	return nil
}
