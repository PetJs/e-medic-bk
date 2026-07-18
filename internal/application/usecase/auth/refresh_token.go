// Package auth contains authentication use cases.
package auth

import (
	"context"
	"errors"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/repository"
)

// ErrInvalidRefreshToken is returned when the refresh token is invalid or expired.
var ErrInvalidRefreshToken = errors.New("invalid or expired refresh token")

// RefreshTokenUseCase handles JWT token refresh.
type RefreshTokenUseCase struct {
	userRepo repository.UserRepository
	tokenGen port.TokenGenerator
}

// NewRefreshTokenUseCase creates a new RefreshTokenUseCase.
func NewRefreshTokenUseCase(userRepo repository.UserRepository, tokenGen port.TokenGenerator) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		userRepo: userRepo,
		tokenGen: tokenGen,
	}
}

// Execute refreshes the access token using a refresh token.
func (uc *RefreshTokenUseCase) Execute(ctx context.Context, refreshToken string) (*dto.AuthResponse, error) {
	userID, err := uc.tokenGen.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidRefreshToken
	}
	if !user.IsActive {
		return nil, ErrUserInactive
	}

	accessToken, expiresIn, err := uc.tokenGen.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := uc.tokenGen.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
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
