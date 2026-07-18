// Package user contains user management use cases.
package user

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/shared/pagination"
)

// ListUsersUseCase handles listing users (admin only).
type ListUsersUseCase struct {
	userRepo repository.UserRepository
}

// NewListUsersUseCase creates a new ListUsersUseCase.
func NewListUsersUseCase(userRepo repository.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{userRepo: userRepo}
}

// Execute lists users with pagination.
func (uc *ListUsersUseCase) Execute(ctx context.Context, req *dto.ListUsersRequest) (*dto.ListUsersResponse, error) {
	p := pagination.Pagination{Page: req.Page, Limit: req.Limit}
	p.Normalize()

	users, err := uc.userRepo.List(ctx, p.Limit, p.Offset())
	if err != nil {
		return nil, err
	}

	count, err := uc.userRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.UserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, &dto.UserResponse{
			ID:        u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
		})
	}

	return &dto.ListUsersResponse{
		Users:      responses,
		TotalCount: count,
		Page:       p.Page,
		Limit:      p.Limit,
	}, nil
}
