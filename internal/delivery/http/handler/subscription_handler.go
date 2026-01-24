// Package handler provides HTTP request handlers.
package handler

import "github.com/gin-gonic/gin"

// SubscriptionHandler handles subscription requests.
type SubscriptionHandler struct{}

func NewSubscriptionHandler() *SubscriptionHandler { return &SubscriptionHandler{} }

func (h *SubscriptionHandler) Create(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *SubscriptionHandler) Cancel(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *SubscriptionHandler) List(c *gin.Context)   { c.JSON(501, gin.H{"error": "not implemented"}) }
