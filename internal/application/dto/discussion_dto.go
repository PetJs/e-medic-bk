// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// CreatePostRequest represents a discussion post creation request.
// The module is identified by the URL path, not the body.
type CreatePostRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

// PostResponse represents a discussion post in API responses.
type PostResponse struct {
	ID           string        `json:"id"`
	ModuleID     string        `json:"module_id"`
	Title        string        `json:"title"`
	Body         string        `json:"body"`
	IsPinned     bool          `json:"is_pinned"`
	CommentCount int           `json:"comment_count"`
	Author       *UserResponse `json:"author"`
	CreatedAt    time.Time     `json:"created_at"`
}

// ListPostsRequest represents a list discussion posts request.
type ListPostsRequest struct {
	Page  int `json:"page" validate:"min=1"`
	Limit int `json:"limit" validate:"min=1,max=100"`
}

// ListPostsResponse represents a list discussion posts response.
type ListPostsResponse struct {
	Posts      []*PostResponse `json:"posts"`
	TotalCount int64           `json:"total_count"`
	Page       int             `json:"page"`
	Limit      int             `json:"limit"`
}

// CreateCommentRequest represents a discussion comment creation request.
// The post is identified by the URL path, not the body.
type CreateCommentRequest struct {
	Body            string  `json:"body" binding:"required"`
	ParentCommentID *string `json:"parent_comment_id,omitempty"`
}

// CommentResponse represents a discussion comment in API responses.
type CommentResponse struct {
	ID              string        `json:"id"`
	PostID          string        `json:"post_id"`
	ParentCommentID *string       `json:"parent_comment_id,omitempty"`
	Body            string        `json:"body"`
	Author          *UserResponse `json:"author"`
	CreatedAt       time.Time     `json:"created_at"`
}

// ListCommentsResponse represents a list discussion comments response.
// Unpaginated by design — the client builds the reply tree from the flat list.
type ListCommentsResponse struct {
	Comments []*CommentResponse `json:"comments"`
}
