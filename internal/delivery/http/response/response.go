// Package response provides HTTP response helpers.
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success sends a successful JSON response.
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

// Created sends a created JSON response.
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
	})
}

// NoContent sends a no content response.
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// Paginated sends a paginated JSON response.
func Paginated(c *gin.Context, data interface{}, total int64, page, limit int) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}
