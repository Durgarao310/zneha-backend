package errors

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"sync"

	"github.com/Durgarao310/zneha-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// Logger is an interface for logging, allowing dependency injection.
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

// ErrorCode represents a standardized error code for the application.
type ErrorCode string

// Error codes grouped by category.
const (
	// Authentication Errors
	AuthInvalidCredentials ErrorCode = "AUTH_INVALID_CREDENTIALS"
	AuthTokenExpired       ErrorCode = "AUTH_TOKEN_EXPIRED"
	AuthUnauthorized       ErrorCode = "AUTH_UNAUTHORIZED"
	AuthTokenInvalid       ErrorCode = "AUTH_TOKEN_INVALID"
	AuthAccountLocked      ErrorCode = "AUTH_ACCOUNT_LOCKED"

	// Authorization Errors
	Forbidden              ErrorCode = "FORBIDDEN"
	RoleNotAllowed         ErrorCode = "ROLE_NOT_ALLOWED"
	InsufficientPrivileges ErrorCode = "INSUFFICIENT_PRIVILEGES"

	// Validation Errors
	ValidationError      ErrorCode = "VALIDATION_ERROR"
	InvalidFieldFormat   ErrorCode = "INVALID_FIELD_FORMAT"
	MissingRequiredField ErrorCode = "MISSING_REQUIRED_FIELD"
	ValueOutOfRange      ErrorCode = "VALUE_OUT_OF_RANGE"

	// Business Logic Errors
	UserNotFound          ErrorCode = "USER_NOT_FOUND"
	OrderAlreadyCompleted ErrorCode = "ORDER_ALREADY_COMPLETED"
	StockNotAvailable     ErrorCode = "STOCK_NOT_AVAILABLE"
	PaymentFailed         ErrorCode = "PAYMENT_FAILED"
	DuplicateRequest      ErrorCode = "DUPLICATE_REQUEST"

	// Database Errors
	DBConnectionFailed ErrorCode = "DB_CONNECTION_FAILED"
	DBQueryFailed      ErrorCode = "DB_QUERY_FAILED"
	DBDuplicateKey     ErrorCode = "DB_DUPLICATE_KEY"
	DBTimeout          ErrorCode = "DB_TIMEOUT"

	// External Service Errors
	ThirdPartyAPIError  ErrorCode = "THIRD_PARTY_API_ERROR"
	ThirdPartyTimeout   ErrorCode = "THIRD_PARTY_TIMEOUT"
	WebhookFailed       ErrorCode = "WEBHOOK_FAILED"
	PaymentGatewayError ErrorCode = "PAYMENT_GATEWAY_ERROR"

	// Rate Limiting Errors
	RateLimitExceeded ErrorCode = "RATE_LIMIT_EXCEEDED"
	TooManyAttempts   ErrorCode = "TOO_MANY_ATTEMPTS"
	ServiceOverloaded ErrorCode = "SERVICE_OVERLOADED"

	// File & Media Errors
	FileTooLarge       ErrorCode = "FILE_TOO_LARGE"
	FileTypeNotAllowed ErrorCode = "FILE_TYPE_NOT_ALLOWED"
	FileUploadFailed   ErrorCode = "FILE_UPLOAD_FAILED"

	// System / Server Errors
	InternalServerError ErrorCode = "INTERNAL_SERVER_ERROR"
	ServiceUnavailable  ErrorCode = "SERVICE_UNAVAILABLE"
	Timeout             ErrorCode = "TIMEOUT"
	ConfigurationError  ErrorCode = "CONFIGURATION_ERROR"

	// Networking Errors
	DNSResolutionFailed ErrorCode = "DNS_RESOLUTION_FAILED"
	ConnectionReset     ErrorCode = "CONNECTION_RESET"
	SSLHandshakeFailed  ErrorCode = "SSL_HANDSHAKE_FAILED"

	// Concurrency Errors
	OptimisticLockFailed ErrorCode = "OPTIMISTIC_LOCK_FAILED"
	TransactionAborted   ErrorCode = "TRANSACTION_ABORTED"
	StaleData            ErrorCode = "STALE_DATA"

	// Security Errors
	CSRFTokenInvalid    ErrorCode = "CSRF_TOKEN_INVALID"
	XSSDetected         ErrorCode = "XSS_DETECTED"
	SQLInjectionAttempt ErrorCode = "SQL_INJECTION_ATTEMPT"
	InvalidSignature    ErrorCode = "INVALID_SIGNATURE"
)

// AppError represents a custom application error with context.
type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Err     error     `json:"-"` // Underlying error, not exposed to JSON
	Context map[string]interface{}
}

// Error returns the error message.
func (e *AppError) Error() string {
	return e.Message
}

// Unwrap returns the underlying error for errors.Is/As.
func (e *AppError) Unwrap() error {
	return e.Err
}

// WithContext adds additional context to the error.
func (e *AppError) WithContext(key string, value interface{}) *AppError {
	if e.Context == nil {
		e.Context = make(map[string]interface{})
	}
	e.Context[key] = value
	return e
}

// New creates a new AppError.
func New(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
		Context: make(map[string]interface{}),
	}
}

// ErrorConfig defines mappings for error codes to HTTP status codes and messages.
type ErrorConfig struct {
	StatusCode     int
	DefaultMessage string
}

// errorConfigMap holds the mapping of ErrorCode to HTTP status and default message.
var errorConfigMap = map[ErrorCode]ErrorConfig{
	AuthInvalidCredentials: {http.StatusUnauthorized, "Invalid credentials provided."},
	AuthTokenExpired:       {http.StatusUnauthorized, "Authentication token has expired."},
	AuthUnauthorized:       {http.StatusUnauthorized, "Unauthorized access."},
	AuthTokenInvalid:       {http.StatusUnauthorized, "Invalid authentication token."},
	AuthAccountLocked:      {http.StatusUnauthorized, "Account is locked."},
	Forbidden:              {http.StatusForbidden, "Access forbidden."},
	RoleNotAllowed:         {http.StatusForbidden, "Role not allowed for this action."},
	InsufficientPrivileges: {http.StatusForbidden, "Insufficient privileges."},
	ValidationError:        {http.StatusBadRequest, "Validation failed."},
	InvalidFieldFormat:     {http.StatusBadRequest, "Invalid field format."},
	MissingRequiredField:   {http.StatusBadRequest, "Missing required field."},
	ValueOutOfRange:        {http.StatusBadRequest, "Value out of range."},
	UserNotFound:           {http.StatusNotFound, "User not found."},
	OrderAlreadyCompleted:  {http.StatusConflict, "Order already completed."},
	StockNotAvailable:      {http.StatusBadRequest, "Stock not available."},
	PaymentFailed:          {http.StatusBadRequest, "Payment failed."},
	DuplicateRequest:       {http.StatusConflict, "Duplicate request detected."},
	DBConnectionFailed:     {http.StatusInternalServerError, "Database connection failed."},
	DBQueryFailed:          {http.StatusInternalServerError, "Database query failed."},
	DBDuplicateKey:         {http.StatusConflict, "Duplicate key in database."},
	DBTimeout:              {http.StatusInternalServerError, "Database operation timed out."},
	ThirdPartyAPIError:     {http.StatusBadGateway, "Third-party API error."},
	ThirdPartyTimeout:      {http.StatusGatewayTimeout, "Third-party service timed out."},
	WebhookFailed:          {http.StatusBadGateway, "Webhook delivery failed."},
	PaymentGatewayError:    {http.StatusGatewayTimeout, "Payment gateway error."},
	RateLimitExceeded:      {http.StatusTooManyRequests, "Rate limit exceeded."},
	TooManyAttempts:        {http.StatusTooManyRequests, "Too many attempts."},
	ServiceOverloaded:      {http.StatusTooManyRequests, "Service overloaded."},
	FileTooLarge:           {http.StatusRequestEntityTooLarge, "File too large."},
	FileTypeNotAllowed:     {http.StatusBadRequest, "File type not allowed."},
	FileUploadFailed:       {http.StatusBadRequest, "File upload failed."},
	InternalServerError:    {http.StatusInternalServerError, "Internal server error."},
	ServiceUnavailable:     {http.StatusServiceUnavailable, "Service unavailable."},
	Timeout:                {http.StatusGatewayTimeout, "Operation timed out."},
	ConfigurationError:     {http.StatusInternalServerError, "Configuration error."},
	DNSResolutionFailed:    {http.StatusInternalServerError, "DNS resolution failed."},
	ConnectionReset:        {http.StatusInternalServerError, "Connection reset."},
	SSLHandshakeFailed:     {http.StatusInternalServerError, "SSL handshake failed."},
	OptimisticLockFailed:   {http.StatusConflict, "Optimistic lock failed."},
	TransactionAborted:     {http.StatusConflict, "Transaction aborted."},
	StaleData:              {http.StatusConflict, "Stale data detected."},
	CSRFTokenInvalid:       {http.StatusBadRequest, "Invalid CSRF token."},
	XSSDetected:            {http.StatusBadRequest, "XSS attack detected."},
	SQLInjectionAttempt:    {http.StatusBadRequest, "SQL injection attempt detected."},
	InvalidSignature:       {http.StatusBadRequest, "Invalid signature."},
}

// validationMessages defines user-friendly messages for validation tags.
var validationMessages = map[string]func(fe validator.FieldError) string{
	"required": func(fe validator.FieldError) string { return "This field is required." },
	"email":    func(fe validator.FieldError) string { return "Invalid email format." },
	"min":      func(fe validator.FieldError) string { return fmt.Sprintf("Value must be at least %s.", fe.Param()) },
	"max":      func(fe validator.FieldError) string { return fmt.Sprintf("Value must not exceed %s.", fe.Param()) },
	"len":      func(fe validator.FieldError) string { return fmt.Sprintf("Length must be exactly %s.", fe.Param()) },
}

// ErrorHandlerConfig holds configuration for the global error handler.
type ErrorHandlerConfig struct {
	Logger          Logger
	I18nProvider    func(ctx context.Context, key string) string          // Optional i18n provider
	MetricsRecorder func(ctx context.Context, code ErrorCode, status int) // Optional metrics
}

// GlobalErrorHandler creates a Gin middleware for error handling with advanced features.
func GlobalErrorHandler(config ErrorHandlerConfig) gin.HandlerFunc {
	if config.Logger == nil {
		config.Logger = zap.NewNop() // Default to no-op logger
	}

	// Pre-allocate buffer for validation errors to reduce allocations
	var fieldErrorsPool = sync.Pool{
		New: func() interface{} {
			return make([]utils.FieldError, 0, 10) // Adjust size based on typical validation error count
		},
	}

	return func(c *gin.Context) {
		// Add request ID to context for logging
		requestID := c.GetString("request_id")
		if requestID == "" {
			requestID = "unknown"
		}
		logFields := []zap.Field{zap.String("request_id", requestID)}

		// Recover from panics
		defer func() {
			if err := recover(); err != nil {
				config.Logger.Error(
					"Panic occurred",
					append(logFields,
						zap.Any("error", err),
						zap.String("stack", string(debug.Stack())),
					)...,
				)
				if config.MetricsRecorder != nil {
					config.MetricsRecorder(c.Request.Context(), InternalServerError, http.StatusInternalServerError)
				}
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
			cfg, exists := errorConfigMap[appErr.Code]
			if !exists {
				cfg = ErrorConfig{StatusCode: http.StatusInternalServerError, DefaultMessage: "Unknown error."}
			}

			// Use i18n provider if available
			message := appErr.Message
			if config.I18nProvider != nil && message == "" {
				message = config.I18nProvider(c.Request.Context(), string(appErr.Code))
			}
			if message == "" {
				message = cfg.DefaultMessage
			}

			config.Logger.Error(
				"Application error",
				append(logFields,
					zap.String("code", string(appErr.Code)),
					zap.String("message", message),
					zap.Any("context", appErr.Context),
				)...,
			)

			if config.MetricsRecorder != nil {
				config.MetricsRecorder(c.Request.Context(), appErr.Code, cfg.StatusCode)
			}

			// Add rate-limiting headers if applicable
			if appErr.Code == RateLimitExceeded || appErr.Code == TooManyAttempts {
				if retryAfter, ok := appErr.Context["retry_after"]; ok {
					c.Header("Retry-After", fmt.Sprintf("%v", retryAfter))
				}
			}

			utils.ErrorResponse(c, cfg.StatusCode, string(appErr.Code), message)
			c.Abort()
			return
		}

		// Handle validation errors
		if validationErrs, ok := err.Err.(validator.ValidationErrors); ok {
			fieldErrors := fieldErrorsPool.Get().([]utils.FieldError)
			defer fieldErrorsPool.Put(fieldErrors[:0]) // Reset and return to pool

			for _, fe := range validationErrs {
				message := validationMessage(fe, c.Request.Context(), config.I18nProvider)
				fieldErrors = append(fieldErrors, utils.FieldError{
					Field:   fe.Field(),
					Tag:     fe.Tag(),
					Value:   fe.Value(),
					Message: message,
				})
			}

			config.Logger.Info(
				"Validation error",
				append(logFields, zap.Any("errors", fieldErrors))...,
			)

			if config.MetricsRecorder != nil {
				config.MetricsRecorder(c.Request.Context(), ValidationError, http.StatusBadRequest)
			}

			utils.DetailedValidationErrorResponse(c, fieldErrors)
			c.Abort()
			return
		}

		// Fallback for unhandled errors
		config.Logger.Error(
			"Unhandled error",
			append(logFields, zap.Error(err.Err))...,
		)

		if config.MetricsRecorder != nil {
			config.MetricsRecorder(c.Request.Context(), InternalServerError, http.StatusInternalServerError)
		}

		utils.ErrorResponse(c, http.StatusInternalServerError, string(InternalServerError), "An internal server error occurred.")
		c.Abort()
	}
}

// validationMessage creates a user-friendly message for a validation error with i18n support.
func validationMessage(fe validator.FieldError, ctx context.Context, i18nProvider func(ctx context.Context, key string) string) string {
	if i18nProvider != nil {
		key := fmt.Sprintf("validation.%s.%s", fe.Field(), fe.Tag())
		if msg := i18nProvider(ctx, key); msg != "" {
			return msg
		}
	}

	if fn, exists := validationMessages[fe.Tag()]; exists {
		return fn(fe)
	}
	return fmt.Sprintf("Invalid value for field %s.", fe.Field())
}
