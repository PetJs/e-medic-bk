// Package handler provides HTTP request handlers.
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/usecase/search"
)

// SearchHandler handles site-search requests.
type SearchHandler struct {
	searchUC *search.SearchUseCase
}

// NewSearchHandler creates a new SearchHandler.
func NewSearchHandler(searchUC *search.SearchUseCase) *SearchHandler {
	return &SearchHandler{searchUC: searchUC}
}

// Search searches courses, modules, and lessons by title/description (?q=).
func (h *SearchHandler) Search(c *gin.Context) {
	response, err := h.searchUC.Execute(c.Request.Context(), c.Query("q"))
	if err != nil {
		internalError(c, "Search failed", err)
		return
	}
	c.JSON(http.StatusOK, response)
}
