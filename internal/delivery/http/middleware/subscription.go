// Package middleware provides HTTP middleware.
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SubscriptionMiddleware handles premium content access.
type SubscriptionMiddleware struct {
	// TODO: Add subscription use case dependency
}

// NewSubscriptionMiddleware creates a new SubscriptionMiddleware.
func NewSubscriptionMiddleware() *SubscriptionMiddleware {
	return &SubscriptionMiddleware{}
}

// RequirePremium returns a middleware that requires an active subscription.
func (m *SubscriptionMiddleware) RequirePremium() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// TODO: Check if user has active subscription
		_ = userID

		c.Next()
	}
}
