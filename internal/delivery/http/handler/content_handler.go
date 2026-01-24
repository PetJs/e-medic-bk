// Package handler provides HTTP request handlers.
package handler

import "github.com/gin-gonic/gin"

// ContentHandler handles content requests.
type ContentHandler struct{}

func NewContentHandler() *ContentHandler { return &ContentHandler{} }

func (h *ContentHandler) UploadContent(c *gin.Context) {
	c.JSON(501, gin.H{"error": "not implemented"})
}
func (h *ContentHandler) GetContentURL(c *gin.Context) {
	c.JSON(501, gin.H{"error": "not implemented"})
}
func (h *ContentHandler) DeleteContent(c *gin.Context) {
	c.JSON(501, gin.H{"error": "not implemented"})
}
