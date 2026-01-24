// Package validator provides request validation helpers.
package validator

import (
	"github.com/go-playground/validator/v10"
)

// Validator wraps the go-playground validator.
var validate = validator.New()

// Validate validates a struct using tags.
func Validate(s interface{}) error {
	return validate.Struct(s)
}

// FormatValidationErrors formats validation errors for API response.
func FormatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors[e.Field()] = formatErrorMessage(e)
		}
	}

	return errors
}

func formatErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	default:
		return "Invalid value"
	}
}
