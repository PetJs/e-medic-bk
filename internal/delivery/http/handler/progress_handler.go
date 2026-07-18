// Package handler provides HTTP request handlers.
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/progress"
)

// ProgressHandler handles progress requests.
type ProgressHandler struct {
	updateUC *progress.UpdateProgressUseCase
	getUC    *progress.GetProgressUseCase
	listUC   *progress.ListProgressUseCase
	courseUC *progress.GetCourseProgressUseCase
}

// NewProgressHandler creates a new ProgressHandler.
func NewProgressHandler(
	updateUC *progress.UpdateProgressUseCase,
	getUC *progress.GetProgressUseCase,
	listUC *progress.ListProgressUseCase,
	courseUC *progress.GetCourseProgressUseCase,
) *ProgressHandler {
	return &ProgressHandler{updateUC: updateUC, getUC: getUC, listUC: listUC, courseUC: courseUC}
}

// UpdateProgress upserts the current user's progress on a lesson.
func (h *ProgressHandler) UpdateProgress(c *gin.Context) {
	var req dto.UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.LessonID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lesson_id is required"})
		return
	}

	response, err := h.updateUC.Execute(c.Request.Context(), c.GetString("user_id"), &req)
	if err != nil {
		if errors.Is(err, progress.ErrLessonNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
			return
		}
		internalError(c, "Failed to update progress", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetProgress returns the current user's progress on one lesson (:id = lesson id).
func (h *ProgressHandler) GetProgress(c *gin.Context) {
	response, err := h.getUC.Execute(c.Request.Context(), c.GetString("user_id"), c.Param("id"))
	if err != nil {
		internalError(c, "Failed to get progress", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// ListProgress returns all of the current user's lesson progress.
func (h *ProgressHandler) ListProgress(c *gin.Context) {
	response, err := h.listUC.Execute(c.Request.Context(), c.GetString("user_id"))
	if err != nil {
		internalError(c, "Failed to list progress", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetCourseProgress returns completion stats for a course (:id = course id).
func (h *ProgressHandler) GetCourseProgress(c *gin.Context) {
	response, err := h.courseUC.Execute(c.Request.Context(), c.GetString("user_id"), c.Param("id"))
	if err != nil {
		internalError(c, "Failed to get course progress", err)
		return
	}
	c.JSON(http.StatusOK, response)
}
