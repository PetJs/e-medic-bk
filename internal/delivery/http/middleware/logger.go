// Package middleware provides HTTP middleware.
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"emedic-bk/pkg/logger"
)

// Logger returns a middleware that logs HTTP requests.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Info("HTTP Request",
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"status", status,
			"latency", latency.String(),
			"client_ip", c.ClientIP(),
		)
	}
}
