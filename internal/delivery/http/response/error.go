// Package response provides HTTP response helpers.
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error sends an error JSON response.
func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": false,
		"error":   message,
	})
}

// BadRequest sends a bad request error response.
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}

// Unauthorized sends an unauthorized error response.
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message)
}

// Forbidden sends a forbidden error response.
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, message)
}

// NotFound sends a not found error response.
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message)
}

// InternalServerError sends an internal server error response.
func InternalServerError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message)
}

// ValidationError sends a validation error response.
func ValidationError(c *gin.Context, errors interface{}) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"success": false,
		"error":   "validation failed",
		"details": errors,
	})
}
