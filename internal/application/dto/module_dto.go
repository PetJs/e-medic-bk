// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// CreateModuleRequest represents a module creation request.
type CreateModuleRequest struct {
	CourseID    string `json:"course_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	IsPremium   bool   `json:"is_premium"`
}

// UpdateModuleRequest represents a module update request.
type UpdateModuleRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Order       *int    `json:"order,omitempty"`
	IsPremium   *bool   `json:"is_premium,omitempty"`
}

// ModuleResponse represents a module in API responses.
type ModuleResponse struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Order       int       `json:"order"`
	IsPremium   bool      `json:"is_premium"`
	CreatedAt   time.Time `json:"created_at"`
}
