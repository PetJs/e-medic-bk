// Package middleware provides HTTP middleware.
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware handles role-based authorization.
type RoleMiddleware struct{}

// NewRoleMiddleware creates a new RoleMiddleware.
func NewRoleMiddleware() *RoleMiddleware {
	return &RoleMiddleware{}
}

// RequireRole returns a middleware that requires a specific role.
func (m *RoleMiddleware) RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		if role != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			return
		}

		c.Next()
	}
}
