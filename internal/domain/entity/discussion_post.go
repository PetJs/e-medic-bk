// Package entity contains the core domain entities.
package entity

import "time"

// DiscussionPost represents a top-level post on a module's discussion board.
type DiscussionPost struct {
	ID       string
	ModuleID string
	UserID   string
	Title    string
	Body     string
	IsPinned bool

	CreatedAt time.Time
	UpdatedAt time.Time

	// Read-time fields populated by joined queries (not stored columns).
	Author       *User
	CommentCount int
}
