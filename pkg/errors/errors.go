package errors

import "fmt"

// Common error types
var (
	ErrInvalidInput      = NewError("invalid_input", "The provided input is invalid")
	ErrNotFound          = NewError("not_found", "The requested resource was not found")
	ErrUnauthorized      = NewError("unauthorized", "You are not authorized to perform this action")
	ErrForbidden         = NewError("forbidden", "You don't have permission to access this resource")
	ErrInternal          = NewError("internal_error", "An internal error occurred")
	ErrDatabase          = NewError("database_error", "A database error occurred")
	ErrValidation        = NewError("validation_error", "The provided data is invalid")
	ErrDuplicate         = NewError("duplicate_error", "A resource with this information already exists")
	ErrInvalidToken      = NewError("invalid_token", "The provided token is invalid")
	ErrExpiredToken      = NewError("expired_token", "The provided token has expired")
	ErrRateLimit         = NewError("rate_limit", "Too many requests, please try again later")
	ErrPaymentFailed     = NewError("payment_failed", "The payment process failed")
	ErrInsufficientFunds = NewError("insufficient_funds", "Insufficient funds for this operation")
)

// AppError represents an application error
type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewError creates a new AppError
func NewError(code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// IsAppError checks if an error is an AppError
func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}

// GetAppError returns the AppError if the error is an AppError
func GetAppError(err error) *AppError {
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}
	return ErrInternal
}
