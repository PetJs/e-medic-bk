// Package repository defines the repository interfaces for data access.
package repository

import (
	"context"
	"time"

	"emedic-bk/internal/domain/entity"
)

// UserRepository defines the interface for user data access.
type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*entity.User, error)
	Count(ctx context.Context) (int64, error)
	CountByRole(ctx context.Context, role string) (int64, error)
	// SignupsByDay returns new user signups grouped by day, since a point in time.
	SignupsByDay(ctx context.Context, since time.Time) ([]entity.DailyMetric, error)
}
