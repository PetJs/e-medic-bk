// Package user contains user management use cases.
package user

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// ListUsersUseCase handles listing users (admin only).
type ListUsersUseCase struct {
	// TODO: Add dependencies
}

// NewListUsersUseCase creates a new ListUsersUseCase.
func NewListUsersUseCase() *ListUsersUseCase {
	return &ListUsersUseCase{}
}

// Execute lists users with pagination.
func (uc *ListUsersUseCase) Execute(ctx context.Context, req *dto.ListUsersRequest) (*dto.ListUsersResponse, error) {
	// TODO: List users with pagination
	// TODO: Map to response DTOs
	return nil, nil
}
