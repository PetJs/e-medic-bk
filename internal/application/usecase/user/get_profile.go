// Package user contains user management use cases.
package user

import (
	"context"
	"errors"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// ErrUserNotFound is returned when the user does not exist.
var ErrUserNotFound = errors.New("user not found")

// GetProfileUseCase handles getting user profile.
type GetProfileUseCase struct {
	userRepo repository.UserRepository
}

// NewGetProfileUseCase creates a new GetProfileUseCase.
func NewGetProfileUseCase(userRepo repository.UserRepository) *GetProfileUseCase {
	return &GetProfileUseCase{userRepo: userRepo}
}

// Execute retrieves a user's profile.
func (uc *GetProfileUseCase) Execute(ctx context.Context, userID string) (*dto.UserResponse, error) {
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}, nil
}
