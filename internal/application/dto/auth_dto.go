// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// RegisterRequest represents a user registration request.
type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

// LoginRequest represents a user login request.
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse represents an authentication response.
type AuthResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresIn    int64        `json:"expires_in"`
	User         UserResponse `json:"user"`
}

// ResetPasswordRequest represents a password reset request.
type ResetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// ConfirmResetRequest represents a password reset confirmation.
type ConfirmResetRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

// UserResponse represents a user in API responses.
type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
