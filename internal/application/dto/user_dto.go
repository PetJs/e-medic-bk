// Package dto contains Data Transfer Objects for use cases.
package dto

// UpdateProfileRequest represents a profile update request.
type UpdateProfileRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
}

// ListUsersRequest represents a list users request.
type ListUsersRequest struct {
	Page  int `json:"page" validate:"min=1"`
	Limit int `json:"limit" validate:"min=1,max=100"`
}

// ListUsersResponse represents a list users response.
type ListUsersResponse struct {
	Users      []*UserResponse `json:"users"`
	TotalCount int64           `json:"total_count"`
	Page       int             `json:"page"`
	Limit      int             `json:"limit"`
}
