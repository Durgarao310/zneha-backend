package errors

// AppError represents a custom application error.
type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Err     error     `json:"-"` // Underlying error, not exposed to the user
}

// Error returns the error message.
func (e *AppError) Error() string {
	return e.Message
}

// Unwrap returns the underlying error for errors.Is/As.
func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new AppError.
func New(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
