// common/errors/mapper.go
// Maps Go errors → AppError (business rules)
package errors

import (
	"database/sql"
	"errors"
	"net/http"
)

// MapError maps any generic error into an AppError
func MapError(err error) *AppError {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}

	// Example: DB not found → 404
	if errors.Is(err, sql.ErrNoRows) {
		return NewAppError(ErrNotFound.Code, ErrNotFound.Message, http.StatusNotFound, err)
	}

	// Default fallback
	return NewAppError(ErrInternalServer.Code, ErrInternalServer.Message, http.StatusInternalServerError, err)
}
