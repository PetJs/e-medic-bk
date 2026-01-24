// Package valueobject contains immutable value objects with validation.
package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

// ErrInvalidEmail is returned when an email address is invalid.
var ErrInvalidEmail = errors.New("invalid email address")

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Email represents a validated email address.
type Email struct {
	value string
}

// NewEmail creates a new Email value object with validation.
func NewEmail(email string) (Email, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	if !emailRegex.MatchString(email) {
		return Email{}, ErrInvalidEmail
	}
	return Email{value: email}, nil
}

// String returns the email as a string.
func (e Email) String() string {
	return e.value
}
