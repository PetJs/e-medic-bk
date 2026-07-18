// Package handler provides HTTP request handlers.
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/lesson"
)

// LessonHandler handles lesson requests.
type LessonHandler struct {
	createUC *lesson.CreateLessonUseCase
	updateUC *lesson.UpdateLessonUseCase
	deleteUC *lesson.DeleteLessonUseCase
	getUC    *lesson.GetLessonUseCase
	listUC   *lesson.ListLessonsUseCase
}

// NewLessonHandler creates a new LessonHandler.
func NewLessonHandler(
	createUC *lesson.CreateLessonUseCase,
	updateUC *lesson.UpdateLessonUseCase,
	deleteUC *lesson.DeleteLessonUseCase,
	getUC *lesson.GetLessonUseCase,
	listUC *lesson.ListLessonsUseCase,
) *LessonHandler {
	return &LessonHandler{
		createUC: createUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
		getUC:    getUC,
		listUC:   listUC,
	}
}

// CreateLesson creates a lesson (admin).
func (h *LessonHandler) CreateLesson(c *gin.Context) {
	var req dto.CreateLessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.createUC.Execute(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, lesson.ErrModuleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
			return
		}
		internalError(c, "Failed to create lesson", err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateLesson updates a lesson (admin).
func (h *LessonHandler) UpdateLesson(c *gin.Context) {
	var req dto.UpdateLessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.updateUC.Execute(c.Request.Context(), c.Param("id"), &req)
	if err != nil {
		if errors.Is(err, lesson.ErrLessonNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
			return
		}
		internalError(c, "Failed to update lesson", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteLesson deletes a lesson (admin).
func (h *LessonHandler) DeleteLesson(c *gin.Context) {
	if err := h.deleteUC.Execute(c.Request.Context(), c.Param("id")); err != nil {
		if errors.Is(err, lesson.ErrLessonNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
			return
		}
		internalError(c, "Failed to delete lesson", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted"})
}

// GetLesson returns a lesson with contents; premium lessons need a subscription.
func (h *LessonHandler) GetLesson(c *gin.Context) {
	userID := c.GetString("user_id")
	role := c.GetString("role")

	response, err := h.getUC.Execute(c.Request.Context(), c.Param("id"), userID, role)
	if err != nil {
		switch {
		case errors.Is(err, lesson.ErrLessonNotFound), errors.Is(err, lesson.ErrModuleNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		case errors.Is(err, lesson.ErrSubscriptionRequired):
			c.JSON(http.StatusPaymentRequired, gin.H{"error": "Active subscription required"})
		default:
			internalError(c, "Failed to get lesson", err)
		}
		return
	}
	c.JSON(http.StatusOK, response)
}

// ListLessons lists a module's lessons.
func (h *LessonHandler) ListLessons(c *gin.Context) {
	response, err := h.listUC.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, lesson.ErrModuleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
			return
		}
		internalError(c, "Failed to list lessons", err)
		return
	}
	c.JSON(http.StatusOK, response)
}
