// Package handler provides HTTP request handlers.
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/usecase/admin"
)

// AdminHandler handles admin dashboard requests.
type AdminHandler struct {
	statsUC *admin.GetStatsUseCase
}

// NewAdminHandler creates a new AdminHandler.
func NewAdminHandler(statsUC *admin.GetStatsUseCase) *AdminHandler {
	return &AdminHandler{statsUC: statsUC}
}

// Stats returns platform totals for the admin dashboard.
func (h *AdminHandler) Stats(c *gin.Context) {
	response, err := h.statsUC.Execute(c.Request.Context())
	if err != nil {
		internalError(c, "Failed to load stats", err)
		return
	}
	c.JSON(http.StatusOK, response)
}
