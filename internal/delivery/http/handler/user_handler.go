// Package handler provides HTTP request handlers.
package handler

import "github.com/gin-gonic/gin"

// UserHandler handles user requests.
type UserHandler struct{}

func NewUserHandler() *UserHandler { return &UserHandler{} }

func (h *UserHandler) GetProfile(c *gin.Context)    { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *UserHandler) UpdateProfile(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *UserHandler) ListUsers(c *gin.Context)     { c.JSON(501, gin.H{"error": "not implemented"}) }
