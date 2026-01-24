// Package validator provides input validation utilities.
package validator

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// IsValidEmail checks if a string is a valid email.
func IsValidEmail(email string) bool {
	return emailRegex.MatchString(strings.TrimSpace(email))
}

// IsValidPassword checks if a password meets requirements.
func IsValidPassword(password string) bool {
	return len(password) >= 8
}

// Sanitize removes leading/trailing whitespace and normalizes strings.
func Sanitize(s string) string {
	return strings.TrimSpace(s)
}

// IsNotEmpty checks if a string is not empty after trimming.
func IsNotEmpty(s string) bool {
	return strings.TrimSpace(s) != ""
}
