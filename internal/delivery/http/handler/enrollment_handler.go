// Package handler provides HTTP request handlers.
package handler

import "github.com/gin-gonic/gin"

// EnrollmentHandler handles enrollment requests.
type EnrollmentHandler struct{}

func NewEnrollmentHandler() *EnrollmentHandler { return &EnrollmentHandler{} }

func (h *EnrollmentHandler) Enroll(c *gin.Context)   { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *EnrollmentHandler) Unenroll(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *EnrollmentHandler) ListEnrollments(c *gin.Context) {
	c.JSON(501, gin.H{"error": "not implemented"})
}
