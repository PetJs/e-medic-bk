// Package user contains user management use cases.
package user

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// UpdateProfileUseCase handles updating user profile.
type UpdateProfileUseCase struct {
	// TODO: Add dependencies
}

// NewUpdateProfileUseCase creates a new UpdateProfileUseCase.
func NewUpdateProfileUseCase() *UpdateProfileUseCase {
	return &UpdateProfileUseCase{}
}

// Execute updates a user's profile.
func (uc *UpdateProfileUseCase) Execute(ctx context.Context, userID string, req *dto.UpdateProfileRequest) (*dto.UserResponse, error) {
	// TODO: Validate input
	// TODO: Get user by ID
	// TODO: Update user fields
	// TODO: Save user
	return nil, nil
}
