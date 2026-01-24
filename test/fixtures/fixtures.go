// Package fixtures provides test data fixtures.
package fixtures

import (
	"time"

	"emedic-bk/internal/domain/entity"
)

// User fixtures
var (
	StudentUser = &entity.User{
		ID:           "user-student-1",
		Email:        "student@example.com",
		PasswordHash: "$2a$10$hash",
		FirstName:    "John",
		LastName:     "Student",
		Role:         "student",
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	AdminUser = &entity.User{
		ID:           "user-admin-1",
		Email:        "admin@example.com",
		PasswordHash: "$2a$10$hash",
		FirstName:    "Jane",
		LastName:     "Admin",
		Role:         "admin",
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
)

// Course fixtures
var (
	SampleCourse = &entity.Course{
		ID:          "course-1",
		Title:       "Introduction to Go",
		Description: "Learn Go programming",
		Slug:        "intro-to-go",
		AuthorID:    AdminUser.ID,
		IsPublished: true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
)
