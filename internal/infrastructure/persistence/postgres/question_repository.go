// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type QuestionRepository struct{ db *DB }

func NewQuestionRepository(db *DB) repository.QuestionRepository { return &QuestionRepository{db: db} }

func (r *QuestionRepository) Create(ctx context.Context, question *entity.Question) error { return nil }
func (r *QuestionRepository) GetByID(ctx context.Context, id string) (*entity.Question, error) {
	return nil, nil
}
func (r *QuestionRepository) Update(ctx context.Context, question *entity.Question) error { return nil }
func (r *QuestionRepository) Delete(ctx context.Context, id string) error                 { return nil }
func (r *QuestionRepository) ListByLesson(ctx context.Context, lessonID string, limit, offset int) ([]*entity.Question, error) {
	return nil, nil
}
func (r *QuestionRepository) ListByUser(ctx context.Context, userID string, limit, offset int) ([]*entity.Question, error) {
	return nil, nil
}
