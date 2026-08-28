// Package handler provides HTTP request handlers.
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/usecase/admin"
)

// AdminHandler handles admin dashboard requests.
type AdminHandler struct {
	statsUC     *admin.GetStatsUseCase
	analyticsUC *admin.GetAnalyticsUseCase
}

// NewAdminHandler creates a new AdminHandler.
func NewAdminHandler(statsUC *admin.GetStatsUseCase, analyticsUC *admin.GetAnalyticsUseCase) *AdminHandler {
	return &AdminHandler{statsUC: statsUC, analyticsUC: analyticsUC}
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

// Analytics returns revenue/signup/subscription trends and per-module
// completion rates for the admin analytics dashboard (?days=7|30|90).
func (h *AdminHandler) Analytics(c *gin.Context) {
	days, _ := strconv.Atoi(c.DefaultQuery("days", "30"))
	response, err := h.analyticsUC.Execute(c.Request.Context(), days)
	if err != nil {
		internalError(c, "Failed to load analytics", err)
		return
	}
	c.JSON(http.StatusOK, response)
}
