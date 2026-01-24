// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type ContentRepository struct{ db *DB }

func NewContentRepository(db *DB) repository.ContentRepository { return &ContentRepository{db: db} }

func (r *ContentRepository) Create(ctx context.Context, content *entity.Content) error { return nil }
func (r *ContentRepository) GetByID(ctx context.Context, id string) (*entity.Content, error) {
	return nil, nil
}
func (r *ContentRepository) Update(ctx context.Context, content *entity.Content) error { return nil }
func (r *ContentRepository) Delete(ctx context.Context, id string) error               { return nil }
func (r *ContentRepository) ListByLesson(ctx context.Context, lessonID string) ([]*entity.Content, error) {
	return nil, nil
}
