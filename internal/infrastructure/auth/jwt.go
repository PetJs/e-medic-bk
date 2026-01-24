// Package auth provides authentication infrastructure implementations.
package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"emedic-bk/internal/application/port"
)

// JWTGenerator implements port.TokenGenerator using JWT.
type JWTGenerator struct {
	accessSecret  string
	refreshSecret string
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

// NewJWTGenerator creates a new JWT token generator.
func NewJWTGenerator(accessSecret, refreshSecret string, accessTTL, refreshTTL time.Duration) port.TokenGenerator {
	return &JWTGenerator{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessTTL:     accessTTL,
		refreshTTL:    refreshTTL,
	}
}

type accessClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type refreshClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *JWTGenerator) GenerateAccessToken(userID, role string) (string, int64, error) {
	expiresAt := time.Now().Add(j.accessTTL)
	claims := accessClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.accessSecret))
	if err != nil {
		return "", 0, err
	}

	return signedToken, int64(j.accessTTL.Seconds()), nil
}

func (j *JWTGenerator) GenerateRefreshToken(userID string) (string, error) {
	claims := refreshClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.refreshSecret))
}

func (j *JWTGenerator) ValidateAccessToken(tokenString string) (userID, role string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &accessClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.accessSecret), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*accessClaims)
	if !ok || !token.Valid {
		return "", "", errors.New("invalid token")
	}

	return claims.UserID, claims.Role, nil
}

func (j *JWTGenerator) ValidateRefreshToken(tokenString string) (userID string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &refreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.refreshSecret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*refreshClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.UserID, nil
}
