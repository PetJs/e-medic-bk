// Package handler provides HTTP request handlers.
package handler

import "github.com/gin-gonic/gin"

// LessonHandler handles lesson requests.
type LessonHandler struct{}

func NewLessonHandler() *LessonHandler { return &LessonHandler{} }

func (h *LessonHandler) CreateLesson(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *LessonHandler) UpdateLesson(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *LessonHandler) DeleteLesson(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *LessonHandler) GetLesson(c *gin.Context)    { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *LessonHandler) ListLessons(c *gin.Context)  { c.JSON(501, gin.H{"error": "not implemented"}) }
