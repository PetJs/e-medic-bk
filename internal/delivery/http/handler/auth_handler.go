// Package handler provides HTTP request handlers.
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/auth"
)

// AuthHandler handles authentication requests.
type AuthHandler struct {
	registerUC *auth.RegisterUseCase
	loginUC    *auth.LoginUseCase
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(registerUC *auth.RegisterUseCase, loginUC *auth.LoginUseCase) *AuthHandler {
	return &AuthHandler{
		registerUC: registerUC,
		loginUC:    loginUC,
	}
}

// Register handles user registration.
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.registerUC.Execute(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, auth.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// Login handles user login.
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.loginUC.Execute(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		if errors.Is(err, auth.ErrUserInactive) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Account is inactive"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Logout handles user logout (clears token on client side).
func (h *AuthHandler) Logout(c *gin.Context) {
	// For JWT, logout is typically handled client-side by discarding the token
	// Optionally, you can implement a token blacklist here
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// RefreshToken handles token refresh.
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// TODO: Implement token refresh
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

// RequestPasswordReset handles password reset requests.
func (h *AuthHandler) RequestPasswordReset(c *gin.Context) {
	// TODO: Implement password reset request
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

// ConfirmPasswordReset handles password reset confirmation.
func (h *AuthHandler) ConfirmPasswordReset(c *gin.Context) {
	// TODO: Implement password reset confirmation
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}
