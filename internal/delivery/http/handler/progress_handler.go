// Package handler provides HTTP request handlers.
package handler

import "github.com/gin-gonic/gin"

// ProgressHandler handles progress requests.
type ProgressHandler struct{}

func NewProgressHandler() *ProgressHandler { return &ProgressHandler{} }

func (h *ProgressHandler) UpdateProgress(c *gin.Context) {
	c.JSON(501, gin.H{"error": "not implemented"})
}
func (h *ProgressHandler) GetProgress(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *ProgressHandler) GetCourseProgress(c *gin.Context) {
	c.JSON(501, gin.H{"error": "not implemented"})
}
