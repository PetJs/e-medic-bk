// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"

	"emedic-bk/internal/domain/entity"
)

// EnrollmentRepository defines the interface for enrollment data access.
type EnrollmentRepository interface {
	Create(ctx context.Context, enrollment *entity.Enrollment) error
	GetByID(ctx context.Context, id string) (*entity.Enrollment, error)
	GetByUserAndCourse(ctx context.Context, userID, courseID string) (*entity.Enrollment, error)
	Update(ctx context.Context, enrollment *entity.Enrollment) error
	Delete(ctx context.Context, id string) error
	ListByUser(ctx context.Context, userID string) ([]*entity.Enrollment, error)
	ListByCourse(ctx context.Context, courseID string) ([]*entity.Enrollment, error)
}
