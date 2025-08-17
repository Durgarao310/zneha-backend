package errors

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

// Unwrap returns the underlying error for errors.Is / errors.As support
func (e *AppError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

// Constructor for AppError
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Common error helpers
func BadRequest(msg string, err error) *AppError {
	return NewAppError(http.StatusBadRequest, msg, err)
}

func NotFound(msg string, err error) *AppError {
	return NewAppError(http.StatusNotFound, msg, err)
}

func InternalServerError(msg string, err error) *AppError {
	return NewAppError(http.StatusInternalServerError, msg, err)
}

func Unauthorized(msg string, err error) *AppError {
	return NewAppError(http.StatusUnauthorized, msg, err)
}
