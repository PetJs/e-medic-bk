// Package lesson contains lesson management use cases.
package lesson

import "context"

// DeleteLessonUseCase handles lesson deletion.
type DeleteLessonUseCase struct{}

func NewDeleteLessonUseCase() *DeleteLessonUseCase { return &DeleteLessonUseCase{} }

func (uc *DeleteLessonUseCase) Execute(ctx context.Context, lessonID string) error {
	return nil
}
