// Package service defines domain service interfaces for external services.
package service

import "context"

// TokenPair represents an access and refresh token pair.
type TokenPair struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64 // seconds until access token expires
}

// AuthService defines the interface for authentication operations.
type AuthService interface {
	GenerateTokenPair(ctx context.Context, userID, role string) (*TokenPair, error)
	ValidateAccessToken(ctx context.Context, token string) (userID, role string, err error)
	ValidateRefreshToken(ctx context.Context, token string) (userID string, err error)
	RevokeToken(ctx context.Context, token string) error
}
