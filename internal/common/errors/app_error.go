// common/errors/app_error.go
// Custom AppError type (with code, message, details)
package errors

import "net/http"

// AppError is a structured error for clean API responses
type AppError struct {
	Code       string      `json:"code"`              // machine-readable error code
	Message    string      `json:"message"`           // human-readable error message
	StatusCode int         `json:"-"`                 // HTTP status code
	Details    interface{} `json:"details,omitempty"` // extra info for debugging/client
	Err        error       `json:"-"`                 // underlying error
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

// NewAppError creates a new AppError
func NewAppError(code, message string, status int, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: status,
		Err:        err,
	}
}

// Predefined errors (you can expand this list)
var (
	ErrInternalServer = NewAppError("INTERNAL_ERROR", "Internal server error", http.StatusInternalServerError, nil)
	ErrBadRequest     = NewAppError("BAD_REQUEST", "Invalid request", http.StatusBadRequest, nil)
	ErrUnauthorized   = NewAppError("UNAUTHORIZED", "Unauthorized", http.StatusUnauthorized, nil)
	ErrNotFound       = NewAppError("NOT_FOUND", "Resource not found", http.StatusNotFound, nil)
)
