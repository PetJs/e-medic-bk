// Package handler provides HTTP request handlers.
package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// internalError logs the underlying error server-side and responds with a
// generic 500 message so internals are never leaked to clients.
func internalError(c *gin.Context, msg string, err error) {
	slog.Error(msg, "error", err, "method", c.Request.Method, "path", c.FullPath())
	c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
}
