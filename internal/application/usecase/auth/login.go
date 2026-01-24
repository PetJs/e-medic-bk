// Package auth contains authentication use cases.
package auth

import (
	"context"
	"errors"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/repository"
)

// ErrInvalidCredentials is returned when email or password is incorrect.
var ErrInvalidCredentials = errors.New("invalid email or password")

// ErrUserInactive is returned when user account is deactivated.
var ErrUserInactive = errors.New("user account is inactive")

// LoginUseCase handles user login.
type LoginUseCase struct {
	userRepo repository.UserRepository
	hasher   port.Hasher
	tokenGen port.TokenGenerator
}

// NewLoginUseCase creates a new LoginUseCase.
func NewLoginUseCase(
	userRepo repository.UserRepository,
	hasher port.Hasher,
	tokenGen port.TokenGenerator,
) *LoginUseCase {
	return &LoginUseCase{
		userRepo: userRepo,
		hasher:   hasher,
		tokenGen: tokenGen,
	}
}

// Execute logs in a user and returns tokens.
func (uc *LoginUseCase) Execute(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error) {
	// Find user by email
	user, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	// Check if user is active
	if !user.IsActive {
		return nil, ErrUserInactive
	}

	// Verify password
	if err := uc.hasher.Compare(req.Password, user.PasswordHash); err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate tokens
	accessToken, expiresIn, err := uc.tokenGen.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := uc.tokenGen.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
		User: dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}
