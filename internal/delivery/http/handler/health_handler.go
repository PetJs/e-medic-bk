// Package handler provides HTTP request handlers.
package handler

import "github.com/gin-gonic/gin"

// HealthHandler handles health check requests.
type HealthHandler struct{}

func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
