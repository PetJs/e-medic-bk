// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// ProgressRepository defines the interface for progress data access.
type ProgressRepository interface {
	Upsert(ctx context.Context, progress *entity.Progress) error
	GetByUserAndLesson(ctx context.Context, userID, lessonID string) (*entity.Progress, error)
	ListByUser(ctx context.Context, userID string) ([]*entity.Progress, error)
	ListByUserAndCourse(ctx context.Context, userID, courseID string) ([]*entity.Progress, error)
	GetCourseCompletionStats(ctx context.Context, userID, courseID string) (completed, total int, err error)
}
