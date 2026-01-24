// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type ModuleRepository struct{ db *DB }

func NewModuleRepository(db *DB) repository.ModuleRepository { return &ModuleRepository{db: db} }

func (r *ModuleRepository) Create(ctx context.Context, module *entity.Module) error { return nil }
func (r *ModuleRepository) GetByID(ctx context.Context, id string) (*entity.Module, error) {
	return nil, nil
}
func (r *ModuleRepository) Update(ctx context.Context, module *entity.Module) error { return nil }
func (r *ModuleRepository) Delete(ctx context.Context, id string) error             { return nil }
func (r *ModuleRepository) ListByCourse(ctx context.Context, courseID string) ([]*entity.Module, error) {
	return nil, nil
}
