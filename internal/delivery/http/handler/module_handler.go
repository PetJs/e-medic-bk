// Package handler provides HTTP request handlers.
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/module"
)

// ModuleHandler handles module requests.
type ModuleHandler struct {
	createUC *module.CreateModuleUseCase
	updateUC *module.UpdateModuleUseCase
	deleteUC *module.DeleteModuleUseCase
	listUC   *module.ListModulesUseCase
	getUC    *module.GetModuleUseCase
}

// NewModuleHandler creates a new ModuleHandler.
func NewModuleHandler(
	createUC *module.CreateModuleUseCase,
	updateUC *module.UpdateModuleUseCase,
	deleteUC *module.DeleteModuleUseCase,
	listUC *module.ListModulesUseCase,
	getUC *module.GetModuleUseCase,
) *ModuleHandler {
	return &ModuleHandler{
		createUC: createUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
		listUC:   listUC,
		getUC:    getUC,
	}
}

// CreateModule creates a module (admin).
func (h *ModuleHandler) CreateModule(c *gin.Context) {
	var req dto.CreateModuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.createUC.Execute(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, module.ErrCourseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
			return
		}
		internalError(c, "Failed to create module", err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateModule updates a module (admin).
func (h *ModuleHandler) UpdateModule(c *gin.Context) {
	var req dto.UpdateModuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.updateUC.Execute(c.Request.Context(), c.Param("id"), &req)
	if err != nil {
		if errors.Is(err, module.ErrModuleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
			return
		}
		internalError(c, "Failed to update module", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteModule deletes a module (admin).
func (h *ModuleHandler) DeleteModule(c *gin.Context) {
	if err := h.deleteUC.Execute(c.Request.Context(), c.Param("id")); err != nil {
		if errors.Is(err, module.ErrModuleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
			return
		}
		internalError(c, "Failed to delete module", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Module deleted"})
}

// ListModules lists modules, optionally filtered by ?course_id.
func (h *ModuleHandler) ListModules(c *gin.Context) {
	response, err := h.listUC.Execute(c.Request.Context(), c.Query("course_id"))
	if err != nil {
		internalError(c, "Failed to list modules", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetModule returns a single module.
func (h *ModuleHandler) GetModule(c *gin.Context) {
	response, err := h.getUC.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, module.ErrModuleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
			return
		}
		internalError(c, "Failed to get module", err)
		return
	}
	c.JSON(http.StatusOK, response)
}
