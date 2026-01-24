// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type LessonRepository struct{ db *DB }

func NewLessonRepository(db *DB) repository.LessonRepository { return &LessonRepository{db: db} }

func (r *LessonRepository) Create(ctx context.Context, lesson *entity.Lesson) error { return nil }
func (r *LessonRepository) GetByID(ctx context.Context, id string) (*entity.Lesson, error) {
	return nil, nil
}
func (r *LessonRepository) Update(ctx context.Context, lesson *entity.Lesson) error { return nil }
func (r *LessonRepository) Delete(ctx context.Context, id string) error             { return nil }
func (r *LessonRepository) ListByModule(ctx context.Context, moduleID string) ([]*entity.Lesson, error) {
	return nil, nil
}
