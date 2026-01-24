// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// CourseRepository implements repository.CourseRepository.
type CourseRepository struct {
	db *DB
}

func NewCourseRepository(db *DB) repository.CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) Create(ctx context.Context, course *entity.Course) error { return nil }
func (r *CourseRepository) GetByID(ctx context.Context, id string) (*entity.Course, error) {
	return nil, nil
}
func (r *CourseRepository) GetBySlug(ctx context.Context, slug string) (*entity.Course, error) {
	return nil, nil
}
func (r *CourseRepository) Update(ctx context.Context, course *entity.Course) error { return nil }
func (r *CourseRepository) Delete(ctx context.Context, id string) error             { return nil }
func (r *CourseRepository) List(ctx context.Context, limit, offset int) ([]*entity.Course, error) {
	return nil, nil
}
func (r *CourseRepository) ListByAuthor(ctx context.Context, authorID string, limit, offset int) ([]*entity.Course, error) {
	return nil, nil
}
