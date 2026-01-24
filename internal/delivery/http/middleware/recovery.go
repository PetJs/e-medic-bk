// Package middleware provides HTTP middleware.
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"emedic-bk/pkg/logger"
)

// Recovery returns a middleware that recovers from panics.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic recovered", "error", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
				})
			}
		}()
		c.Next()
	}
}
