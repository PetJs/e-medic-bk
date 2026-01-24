// Package errors provides custom application errors.
package errors

import "errors"

// Application errors
var (
	ErrNotFound             = errors.New("resource not found")
	ErrAlreadyExists        = errors.New("resource already exists")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrForbidden            = errors.New("forbidden")
	ErrInvalidInput         = errors.New("invalid input")
	ErrInternalServer       = errors.New("internal server error")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrTokenExpired         = errors.New("token expired")
	ErrSubscriptionRequired = errors.New("subscription required")
)

// AppError represents an application error with context.
type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new application error.
func NewAppError(code, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
