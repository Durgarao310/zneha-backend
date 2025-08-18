package errors

import (
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
