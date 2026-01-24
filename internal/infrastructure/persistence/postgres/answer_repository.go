// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type AnswerRepository struct{ db *DB }

func NewAnswerRepository(db *DB) repository.AnswerRepository { return &AnswerRepository{db: db} }

func (r *AnswerRepository) Create(ctx context.Context, answer *entity.Answer) error { return nil }
func (r *AnswerRepository) GetByID(ctx context.Context, id string) (*entity.Answer, error) {
	return nil, nil
}
func (r *AnswerRepository) Update(ctx context.Context, answer *entity.Answer) error { return nil }
func (r *AnswerRepository) Delete(ctx context.Context, id string) error             { return nil }
func (r *AnswerRepository) ListByQuestion(ctx context.Context, questionID string) ([]*entity.Answer, error) {
	return nil, nil
}
func (r *AnswerRepository) ClearBestAnswer(ctx context.Context, questionID string) error { return nil }
