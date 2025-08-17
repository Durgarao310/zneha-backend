package errors

import (
	"net/http"

	"github.com/Durgarao310/zneha-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GlobalErrorHandler is a middleware that recovers from panics and handles errors.
func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.ErrorResponse(c, http.StatusInternalServerError, string(InternalServerError), "An unexpected error occurred.")
				c.Abort()
			}
		}()

		c.Next()

		// Handle errors that occurred during the request
		err := c.Errors.Last()
		if err == nil {
			return
		}

		// Handle custom application errors
		if appErr, ok := err.Err.(*AppError); ok {
			statusCode := getStatusCode(appErr.Code)
			utils.ErrorResponse(c, statusCode, string(appErr.Code), appErr.Message)
			c.Abort()
			return
		}

		// Handle validation errors
		if validationErrs, ok := err.Err.(validator.ValidationErrors); ok {
			fieldErrors := make([]utils.FieldError, len(validationErrs))
			for i, fe := range validationErrs {
				fieldErrors[i] = utils.FieldError{
					Field:   fe.Field(),
					Tag:     fe.Tag(),
					Value:   fe.Value(),
					Message: validationMessage(fe),
				}
			}
			utils.DetailedValidationErrorResponse(c, fieldErrors)
			c.Abort()
			return
		}

		// Fallback for other unhandled errors
		utils.ErrorResponse(c, http.StatusInternalServerError, string(InternalServerError), "An internal server error occurred.")
		c.Abort()
	}
}

// getStatusCode maps an ErrorCode to an HTTP status code.
func getStatusCode(code ErrorCode) int {
	switch code {
	// Authentication
	case AuthInvalidCredentials, AuthTokenExpired, AuthUnauthorized, AuthTokenInvalid, AuthAccountLocked:
		return http.StatusUnauthorized
	// Authorization
	case Forbidden, RoleNotAllowed, InsufficientPrivileges:
		return http.StatusForbidden
	// Validation
	case ValidationError, InvalidFieldFormat, MissingRequiredField, ValueOutOfRange:
		return http.StatusBadRequest
	// Business Logic
	case UserNotFound:
		return http.StatusNotFound
	case DuplicateRequest, OrderAlreadyCompleted:
		return http.StatusConflict
	case StockNotAvailable, PaymentFailed:
		return http.StatusBadRequest
	// Database
	case DBDuplicateKey:
		return http.StatusConflict
	case DBConnectionFailed, DBQueryFailed, DBTimeout:
		return http.StatusInternalServerError
	// External Services
	case ThirdPartyAPIError, WebhookFailed:
		return http.StatusBadGateway
	case ThirdPartyTimeout, PaymentGatewayError:
		return http.StatusGatewayTimeout
	// Rate Limiting
	case RateLimitExceeded, TooManyAttempts, ServiceOverloaded:
		return http.StatusTooManyRequests
	// File & Media
	case FileTooLarge:
		return http.StatusRequestEntityTooLarge
	case FileTypeNotAllowed, FileUploadFailed:
		return http.StatusBadRequest
	// Concurrency
	case OptimisticLockFailed, TransactionAborted, StaleData:
		return http.StatusConflict
	// Security
	case CSRFTokenInvalid, XSSDetected, SQLInjectionAttempt, InvalidSignature:
		return http.StatusBadRequest
	// Default to Internal Server Error
	default:
		return http.StatusInternalServerError
	}
}

// validationMessage creates a user-friendly message for a validation error.
func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required."
	case "email":
		return "Invalid email format."
	case "min":
		return "Value must be at least " + fe.Param() + "."
	case "max":
		return "Value must not exceed " + fe.Param() + "."
	case "len":
		return "Length must be exactly " + fe.Param() + "."
	default:
		return "Invalid value for field " + fe.Field() + "."
	}
}
