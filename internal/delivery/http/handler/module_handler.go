// Package handler provides HTTP request handlers.
package handler

import "github.com/gin-gonic/gin"

// ModuleHandler handles module requests.
type ModuleHandler struct{}

func NewModuleHandler() *ModuleHandler { return &ModuleHandler{} }

func (h *ModuleHandler) CreateModule(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *ModuleHandler) UpdateModule(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *ModuleHandler) DeleteModule(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *ModuleHandler) ListModules(c *gin.Context)  { c.JSON(501, gin.H{"error": "not implemented"}) }
