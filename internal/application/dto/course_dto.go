// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// CreateCourseRequest represents a course creation request.
type CreateCourseRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	CoverImage  string `json:"cover_image"`
}

// UpdateCourseRequest represents a course update request.
type UpdateCourseRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	CoverImage  *string `json:"cover_image,omitempty"`
	IsPublished *bool   `json:"is_published,omitempty"`
}

// CourseResponse represents a course in API responses.
type CourseResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	CoverImage  string    `json:"cover_image"`
	AuthorID    string    `json:"author_id"`
	IsPublished bool      `json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
}

// CourseDetailResponse represents detailed course info with modules.
type CourseDetailResponse struct {
	CourseResponse
	Modules []*ModuleResponse `json:"modules"`
}

// ListCoursesRequest represents a list courses request.
type ListCoursesRequest struct {
	Page   int    `json:"page" validate:"min=1"`
	Limit  int    `json:"limit" validate:"min=1,max=100"`
	Search string `json:"search,omitempty"`
}

// ListCoursesResponse represents a list courses response.
type ListCoursesResponse struct {
	Courses    []*CourseResponse `json:"courses"`
	TotalCount int64             `json:"total_count"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
}
