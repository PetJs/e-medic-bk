// Package errors provides custom application errors.
package errors

// Error codes
const (
	CodeNotFound             = "NOT_FOUND"
	CodeAlreadyExists        = "ALREADY_EXISTS"
	CodeUnauthorized         = "UNAUTHORIZED"
	CodeForbidden            = "FORBIDDEN"
	CodeInvalidInput         = "INVALID_INPUT"
	CodeInternalServer       = "INTERNAL_SERVER_ERROR"
	CodeInvalidCredentials   = "INVALID_CREDENTIALS"
	CodeTokenExpired         = "TOKEN_EXPIRED"
	CodeSubscriptionRequired = "SUBSCRIPTION_REQUIRED"
	CodeValidationFailed     = "VALIDATION_FAILED"
	CodeRateLimitExceeded    = "RATE_LIMIT_EXCEEDED"
)
