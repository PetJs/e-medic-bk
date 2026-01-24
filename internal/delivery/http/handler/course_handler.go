// Package handler provides HTTP request handlers.
package handler

import "github.com/gin-gonic/gin"

// CourseHandler handles course requests.
type CourseHandler struct{}

func NewCourseHandler() *CourseHandler { return &CourseHandler{} }

func (h *CourseHandler) CreateCourse(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *CourseHandler) UpdateCourse(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *CourseHandler) DeleteCourse(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *CourseHandler) GetCourse(c *gin.Context)    { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *CourseHandler) ListCourses(c *gin.Context)  { c.JSON(501, gin.H{"error": "not implemented"}) }
