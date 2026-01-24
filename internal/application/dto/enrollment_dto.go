// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// EnrollmentResponse represents an enrollment in API responses.
type EnrollmentResponse struct {
	ID         string          `json:"id"`
	UserID     string          `json:"user_id"`
	CourseID   string          `json:"course_id"`
	Course     *CourseResponse `json:"course,omitempty"`
	EnrolledAt time.Time       `json:"enrolled_at"`
	ExpiresAt  *time.Time      `json:"expires_at,omitempty"`
	IsActive   bool            `json:"is_active"`
}
