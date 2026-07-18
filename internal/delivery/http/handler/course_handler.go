// Package handler provides HTTP request handlers.
package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/course"
)

// CourseHandler handles course requests.
type CourseHandler struct {
	createUC *course.CreateCourseUseCase
	updateUC *course.UpdateCourseUseCase
	deleteUC *course.DeleteCourseUseCase
	getUC    *course.GetCourseUseCase
	listUC   *course.ListCoursesUseCase
}

// NewCourseHandler creates a new CourseHandler.
func NewCourseHandler(
	createUC *course.CreateCourseUseCase,
	updateUC *course.UpdateCourseUseCase,
	deleteUC *course.DeleteCourseUseCase,
	getUC *course.GetCourseUseCase,
	listUC *course.ListCoursesUseCase,
) *CourseHandler {
	return &CourseHandler{
		createUC: createUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
		getUC:    getUC,
		listUC:   listUC,
	}
}

// CreateCourse creates a course (admin).
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var req dto.CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	authorID := c.GetString("user_id")
	response, err := h.createUC.Execute(c.Request.Context(), authorID, &req)
	if err != nil {
		internalError(c, "Failed to create course", err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateCourse updates a course (admin).
func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	var req dto.UpdateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.updateUC.Execute(c.Request.Context(), c.Param("id"), &req)
	if err != nil {
		if errors.Is(err, course.ErrCourseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
			return
		}
		internalError(c, "Failed to update course", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteCourse deletes a course (admin).
func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	if err := h.deleteUC.Execute(c.Request.Context(), c.Param("id")); err != nil {
		if errors.Is(err, course.ErrCourseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
			return
		}
		internalError(c, "Failed to delete course", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}

// GetCourse returns a course with its modules.
func (h *CourseHandler) GetCourse(c *gin.Context) {
	response, err := h.getUC.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, course.ErrCourseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
			return
		}
		internalError(c, "Failed to get course", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// ListCourses lists courses.
func (h *CourseHandler) ListCourses(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	response, err := h.listUC.Execute(c.Request.Context(), &dto.ListCoursesRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		internalError(c, "Failed to list courses", err)
		return
	}
	c.JSON(http.StatusOK, response)
}
