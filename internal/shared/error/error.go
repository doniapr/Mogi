package error

import "net/http"

// AppError represents application error
type AppError struct {
	Code    int
	Message string
	Err     error
}

// NewAppError creates new app error
func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// NewAppErrorWithError creates new app error with wrapped error
func NewAppErrorWithError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Error implements error interface
func (ae *AppError) Error() string {
	if ae.Err != nil {
		return ae.Message + ": " + ae.Err.Error()
	}
	return ae.Message
}

// NotFound returns 404 error
func NotFound(message string) *AppError {
	return NewAppError(http.StatusNotFound, message)
}

// BadRequest returns 400 error
func BadRequest(message string) *AppError {
	return NewAppError(http.StatusBadRequest, message)
}

// InternalServerError returns 500 error
func InternalServerError(message string) *AppError {
	return NewAppError(http.StatusInternalServerError, message)
}

// Unauthorized returns 401 error
func Unauthorized(message string) *AppError {
	return NewAppError(http.StatusUnauthorized, message)
}

// Forbidden returns 403 error
func Forbidden(message string) *AppError {
	return NewAppError(http.StatusForbidden, message)
}

