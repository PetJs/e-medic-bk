// Package middleware provides HTTP middleware.
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/port"
)

// AuthMiddleware handles JWT authentication.
type AuthMiddleware struct {
	tokenGenerator port.TokenGenerator
}

// NewAuthMiddleware creates a new AuthMiddleware.
func NewAuthMiddleware(tokenGenerator port.TokenGenerator) *AuthMiddleware {
	return &AuthMiddleware{tokenGenerator: tokenGenerator}
}

// Authenticate returns a middleware that validates JWT tokens.
func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		userID, role, err := m.tokenGenerator.ValidateAccessToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set("user_id", userID)
		c.Set("role", role)
		c.Next()
	}
}
