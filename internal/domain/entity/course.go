// Package entity contains the core domain entities.
package entity

import "time"

// Course represents a course in the platform.
type Course struct {
	ID          string
	Title       string
	Description string
	Slug        string
	CoverImage  string
	AuthorID    string
	IsPublished bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
