// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"context"

	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

type EnrollmentRepository struct{ db *DB }

func NewEnrollmentRepository(db *DB) repository.EnrollmentRepository {
	return &EnrollmentRepository{db: db}
}

func (r *EnrollmentRepository) Create(ctx context.Context, enrollment *entity.Enrollment) error {
	return nil
}
func (r *EnrollmentRepository) GetByID(ctx context.Context, id string) (*entity.Enrollment, error) {
	return nil, nil
}
func (r *EnrollmentRepository) GetByUserAndCourse(ctx context.Context, userID, courseID string) (*entity.Enrollment, error) {
	return nil, nil
}
func (r *EnrollmentRepository) Update(ctx context.Context, enrollment *entity.Enrollment) error {
	return nil
}
func (r *EnrollmentRepository) Delete(ctx context.Context, id string) error { return nil }
func (r *EnrollmentRepository) ListByUser(ctx context.Context, userID string) ([]*entity.Enrollment, error) {
	return nil, nil
}
func (r *EnrollmentRepository) ListByCourse(ctx context.Context, courseID string) ([]*entity.Enrollment, error) {
	return nil, nil
}
