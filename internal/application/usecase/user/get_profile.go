// Package user contains user management use cases.
package user

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// GetProfileUseCase handles getting user profile.
type GetProfileUseCase struct {
	// TODO: Add dependencies
}

// NewGetProfileUseCase creates a new GetProfileUseCase.
func NewGetProfileUseCase() *GetProfileUseCase {
	return &GetProfileUseCase{}
}

// Execute retrieves a user's profile.
func (uc *GetProfileUseCase) Execute(ctx context.Context, userID string) (*dto.UserResponse, error) {
	// TODO: Get user by ID
	// TODO: Map to response DTO
	return nil, nil
}
