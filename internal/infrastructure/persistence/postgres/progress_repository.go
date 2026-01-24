// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type ProgressRepository struct{ db *DB }

func NewProgressRepository(db *DB) repository.ProgressRepository { return &ProgressRepository{db: db} }

func (r *ProgressRepository) Upsert(ctx context.Context, progress *entity.Progress) error { return nil }
func (r *ProgressRepository) GetByUserAndLesson(ctx context.Context, userID, lessonID string) (*entity.Progress, error) {
	return nil, nil
}
func (r *ProgressRepository) ListByUser(ctx context.Context, userID string) ([]*entity.Progress, error) {
	return nil, nil
}
func (r *ProgressRepository) ListByUserAndCourse(ctx context.Context, userID, courseID string) ([]*entity.Progress, error) {
	return nil, nil
}
func (r *ProgressRepository) GetCourseCompletionStats(ctx context.Context, userID, courseID string) (completed, total int, err error) {
	return 0, 0, nil
}
