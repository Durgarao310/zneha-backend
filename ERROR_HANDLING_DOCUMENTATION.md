# Error Handling Documentation

## Overview

This document provides a comprehensive overview of the error handling implementation in the Zneha Backend API. The system implements a structured approach to error handling with custom error types, centralized error codes, and consistent response formats.

## Architecture

The error handling system is organized into several components:

```
internal/common/errors/
├── app_error.go      # Custom AppError type definition
├── error_codes.go    # Centralized error code constants
├── error_handler.go  # Gin middleware for error handling
└── mapper.go         # Error mapping utilities

utils/
└── response.go       # Response utilities and helper functions
```

## Core Components

### 1. AppError Structure (`app_error.go`)

The `AppError` is the foundation of the error handling system, providing a structured approach to error representation.

```go
type AppError struct {
    Code       string      `json:"code"`              // machine-readable error code
    Message    string      `json:"message"`           // human-readable error message
    StatusCode int         `json:"-"`                 // HTTP status code
    Details    interface{} `json:"details,omitempty"` // extra info for debugging/client
    Err        error       `json:"-"`                 // underlying error
}
```

#### Key Features:
- **Machine-readable codes**: Consistent error identification
- **Human-readable messages**: User-friendly error descriptions
- **HTTP status codes**: Proper REST API status mapping
- **Optional details**: Additional context for debugging
- **Error wrapping**: Preserves underlying Go errors

#### Predefined Errors:
```go
var (
    ErrInternalServer = NewAppError("INTERNAL_ERROR", "Internal server error", 500, nil)
    ErrBadRequest     = NewAppError("BAD_REQUEST", "Invalid request", 400, nil)
    ErrUnauthorized   = NewAppError("UNAUTHORIZED", "Unauthorized", 401, nil)
    ErrNotFound       = NewAppError("NOT_FOUND", "Resource not found", 404, nil)
)
```

### 2. Error Codes (`error_codes.go`)

Centralized error code definitions ensure consistency across the application.

```go
const (
    CodeValidationFailed = "VALIDATION_FAILED"
    CodeDatabaseError    = "DATABASE_ERROR"
    CodeConflict         = "CONFLICT"
    CodeForbidden        = "FORBIDDEN"
    CodeRateLimited      = "RATE_LIMITED"
)
```

### 3. Error Handler Middleware (`error_handler.go`)

Gin middleware that processes errors accumulated during request handling.

```go
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next() // process request

        // Check if errors were added during request
        if len(c.Errors) > 0 {
            lastErr := c.Errors.Last().Err
            handleError(c, lastErr)
        }
    }
}
```

#### Features:
- **Automatic error processing**: Handles errors after request completion
- **Structured logging**: Logs internal errors with stack traces
- **Consistent responses**: Returns standardized error JSON format
- **Stack trace logging**: For debugging internal server errors

### 4. Error Mapper (`mapper.go`)

Converts generic Go errors into structured `AppError` instances.

```go
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
```

#### Mapping Rules:
- **AppError passthrough**: Returns existing AppErrors unchanged
- **SQL no rows**: Maps to 404 Not Found
- **Default fallback**: Unknown errors become 500 Internal Server Error

## Response Utilities (`utils/response.go`)

### Standard Response Structure

All API responses follow a consistent format:

```go
type Response struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   *ErrorInfo  `json:"error,omitempty"`
    Meta    interface{} `json:"meta,omitempty"`
}

type ErrorInfo struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}
```

### Available Response Functions

#### Success Responses:
- `SuccessResponse(c, data, meta...)` - 200 OK with data
- `CreatedResponse(c, data)` - 201 Created
- `NoContentResponse(c)` - 204 No Content

#### Error Responses:
- `ErrorResponse(c, statusCode, code, message)` - Generic error
- `BadRequestResponse(c, message)` - 400 Bad Request
- `NotFoundResponse(c, message)` - 404 Not Found
- `ValidationErrorResponse(c, message)` - 400 Validation Error
- `InternalServerErrorResponse(c, message)` - 500 Internal Server Error

## Current Implementation Issues

### 1. **Middleware Not Registered**
❌ **Problem**: The error handling middleware is defined but not used in the main application.

**Current main.go:**
```go
func main() {
    db := database.InitPostgres()
    
    // Dependencies setup...
    
    r := gin.Default()
    routes.RegisterRoutes(r, productHandler)
    
    r.Run(":8080")
}
```

✅ **Solution**: Register the error middleware:
```go
func main() {
    db := database.InitPostgres()
    
    // Dependencies setup...
    
    r := gin.Default()
    r.Use(errors.ErrorHandler()) // Add this line
    routes.RegisterRoutes(r, productHandler)
    
    r.Run(":8080")
}
```

### 2. **Inconsistent Error Handling in Handlers**
❌ **Problem**: Handlers bypass the AppError system and use utils functions directly.

**Current handler code:**
```go
func (h *ProductHandler) Create(c *gin.Context) {
    var product model.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        utils.ValidationErrorResponse(c, err.Error())  // Direct response
        return
    }

    if err := h.service.Create(&product); err != nil {
        utils.InternalServerErrorResponse(c, err.Error())  // All errors become 500
        return
    }

    utils.CreatedResponse(c, product)
}
```

✅ **Solution**: Use the error middleware system:
```go
func (h *ProductHandler) Create(c *gin.Context) {
    var product model.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.Error(errors.NewAppError("VALIDATION_FAILED", "Invalid request data", 400, err))
        return
    }

    if err := h.service.Create(&product); err != nil {
        c.Error(err)  // Let middleware handle error mapping
        return
    }

    utils.CreatedResponse(c, product)
}
```

### 3. **Repository Layer Error Handling**
❌ **Problem**: Repository methods return raw GORM errors without proper typing.

**Current repository code:**
```go
func (r *productRepository) FindByID(id uint64) (*model.Product, error) {
    var product model.Product
    err := r.db.First(&product, id).Error
    if err != nil {
        return nil, err  // Raw GORM error
    }
    return &product, nil
}
```

✅ **Solution**: Map database errors to AppErrors:
```go
func (r *productRepository) FindByID(id uint64) (*model.Product, error) {
    var product model.Product
    err := r.db.First(&product, id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.NewAppError("PRODUCT_NOT_FOUND", "Product not found", 404, err)
        }
        return nil, errors.NewAppError("DATABASE_ERROR", "Database operation failed", 500, err)
    }
    return &product, nil
}
```

## Example Error Responses

### Validation Error (400)
```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_FAILED",
    "message": "Invalid request data"
  }
}
```

### Not Found Error (404)
```json
{
  "success": false,
  "error": {
    "code": "PRODUCT_NOT_FOUND",
    "message": "Product not found"
  }
}
```

### Internal Server Error (500)
```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "Internal server error"
  }
}
```

## Best Practices

### 1. Use AppError for Business Logic Errors
```go
// Service layer
func (s *productService) Create(product *model.Product) error {
    if product.Price < 0 {
        return errors.NewAppError("INVALID_PRICE", "Product price cannot be negative", 400, nil)
    }
    return s.repo.Create(product)
}
```

### 2. Let Middleware Handle Error Responses
```go
// Handler layer
func (h *ProductHandler) Create(c *gin.Context) {
    // ... validation ...
    
    if err := h.service.Create(&product); err != nil {
        c.Error(err)  // Don't handle response here
        return
    }
    
    utils.CreatedResponse(c, product)
}
```

### 3. Add Error Context in Repository Layer
```go
func (r *productRepository) Create(product *model.Product) error {
    err := r.db.Create(product).Error
    if err != nil {
        if isDuplicateKeyError(err) {
            return errors.NewAppError("DUPLICATE_PRODUCT", "Product already exists", 409, err)
        }
        return errors.NewAppError("DATABASE_ERROR", "Failed to create product", 500, err)
    }
    return nil
}
```

## Recommended Improvements

### 1. Expand Error Codes
```go
const (
    // Validation errors
    CodeValidationFailed     = "VALIDATION_FAILED"
    CodeMissingField        = "MISSING_FIELD"
    CodeInvalidFormat       = "INVALID_FORMAT"
    
    // Business logic errors
    CodeProductNotFound     = "PRODUCT_NOT_FOUND"
    CodeDuplicateProduct    = "DUPLICATE_PRODUCT"
    CodeInsufficientStock   = "INSUFFICIENT_STOCK"
    
    // Authentication/Authorization
    CodeUnauthorized        = "UNAUTHORIZED"
    CodeForbidden          = "FORBIDDEN"
    CodeTokenExpired       = "TOKEN_EXPIRED"
    
    // System errors
    CodeDatabaseError      = "DATABASE_ERROR"
    CodeExternalService    = "EXTERNAL_SERVICE_ERROR"
    CodeRateLimited       = "RATE_LIMITED"
)
```

### 2. Add Structured Logging
```go
func handleError(c *gin.Context, err error) {
    appErr := MapError(err)
    
    // Add request context to logs
    logger := log.WithFields(log.Fields{
        "request_id": c.GetString("request_id"),
        "method":     c.Request.Method,
        "path":       c.Request.URL.Path,
        "error_code": appErr.Code,
        "user_id":    c.GetString("user_id"),
    })
    
    if appErr.StatusCode == http.StatusInternalServerError {
        logger.WithField("stack_trace", string(debug.Stack())).Error(appErr.Error())
    } else {
        logger.Warn(appErr.Error())
    }
    
    // Respond with structured error
    c.AbortWithStatusJSON(appErr.StatusCode, gin.H{
        "error": gin.H{
            "code":    appErr.Code,
            "message": appErr.Message,
            "details": appErr.Details,
        },
    })
}
```

### 3. Add Request ID Middleware
```go
func RequestIDMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestID := uuid.New().String()
        c.Set("request_id", requestID)
        c.Header("X-Request-ID", requestID)
        c.Next()
    }
}
```

## Integration Checklist

- [ ] Register error handling middleware in main.go
- [ ] Update all handlers to use c.Error() instead of direct responses
- [ ] Implement proper error mapping in repository layer
- [ ] Add comprehensive error codes
- [ ] Implement structured logging with request context
- [ ] Add request ID tracking
- [ ] Create error handling tests
- [ ] Document error codes for API consumers
- [ ] Add error monitoring/alerting

## Conclusion

The current error handling system has a solid foundation with custom error types and structured responses. However, it requires integration work to be fully functional. The main issues are:

1. **Missing middleware registration**
2. **Inconsistent error handling patterns**
3. **Lack of proper error mapping in lower layers**

Once these issues are addressed, the system will provide robust, consistent error handling across the entire application.
