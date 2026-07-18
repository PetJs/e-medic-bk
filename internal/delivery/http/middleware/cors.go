// Package middleware provides HTTP middleware.
package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORS returns a middleware that handles CORS.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Auth uses Bearer headers, not cookies, so a wildcard origin is safe
		// (and Allow-Credentials must not be combined with it per the spec).
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
