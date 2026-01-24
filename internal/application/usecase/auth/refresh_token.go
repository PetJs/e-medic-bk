// Package auth contains authentication use cases.
package auth

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// RefreshTokenUseCase handles JWT token refresh.
type RefreshTokenUseCase struct {
	// TODO: Add dependencies
}

// NewRefreshTokenUseCase creates a new RefreshTokenUseCase.
func NewRefreshTokenUseCase() *RefreshTokenUseCase {
	return &RefreshTokenUseCase{}
}

// Execute refreshes the access token using a refresh token.
func (uc *RefreshTokenUseCase) Execute(ctx context.Context, refreshToken string) (*dto.AuthResponse, error) {
	// TODO: Validate refresh token
	// TODO: Generate new token pair
	return nil, nil
}
