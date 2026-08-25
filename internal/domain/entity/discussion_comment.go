// Package entity contains the core domain entities.
package entity

import "time"

// DiscussionComment represents a comment (or threaded reply) on a discussion post.
type DiscussionComment struct {
	ID              string
	PostID          string
	UserID          string
	ParentCommentID *string // nil for a top-level comment
	Body            string

	CreatedAt time.Time
	UpdatedAt time.Time

	// Author is populated by joined queries (not a stored column).
	Author *User
}
