// Package handler provides HTTP request handlers.
package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/user"
)

// UserHandler handles user requests.
type UserHandler struct {
	getProfileUC    *user.GetProfileUseCase
	updateProfileUC *user.UpdateProfileUseCase
	listUsersUC     *user.ListUsersUseCase
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(
	getProfileUC *user.GetProfileUseCase,
	updateProfileUC *user.UpdateProfileUseCase,
	listUsersUC *user.ListUsersUseCase,
) *UserHandler {
	return &UserHandler{
		getProfileUC:    getProfileUC,
		updateProfileUC: updateProfileUC,
		listUsersUC:     listUsersUC,
	}
}

// GetProfile returns the authenticated user's profile.
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	profile, err := h.getProfileUC.Execute(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		internalError(c, "Failed to get profile", err)
		return
	}

	c.JSON(http.StatusOK, profile)
}

// UpdateProfile updates the authenticated user's profile.
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	profile, err := h.updateProfileUC.Execute(c.Request.Context(), userID, &req)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		internalError(c, "Failed to update profile", err)
		return
	}

	c.JSON(http.StatusOK, profile)
}

// ListUsers lists all users (admin only).
func (h *UserHandler) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	response, err := h.listUsersUC.Execute(c.Request.Context(), &dto.ListUsersRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		internalError(c, "Failed to list users", err)
		return
	}

	c.JSON(http.StatusOK, response)
}
