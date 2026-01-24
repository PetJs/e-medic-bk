// Package port defines secondary port interfaces for the application layer.
package port

// TokenGenerator defines the interface for JWT token generation.
type TokenGenerator interface {
	GenerateAccessToken(userID, role string) (string, int64, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateAccessToken(token string) (userID, role string, err error)
	ValidateRefreshToken(token string) (userID string, err error)
}
