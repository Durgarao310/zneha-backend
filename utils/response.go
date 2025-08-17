package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response standard structure for all API responses
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    interface{} `json:"meta,omitempty"` // optional for pagination, etc.
}

// ErrorInfo detailed error block
type ErrorInfo struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Fields  []FieldError `json:"fields,omitempty"`
}

// FieldError represents a single field validation error
type FieldError struct {
	Field   string      `json:"field"`
	Tag     string      `json:"tag"`
	Value   interface{} `json:"value,omitempty"`
	Param   string      `json:"param,omitempty"`
	Message string      `json:"message"`
}

// SuccessResponse returns a success JSON response
func SuccessResponse(c *gin.Context, data interface{}, meta ...interface{}) {
	resp := Response{
		Success: true,
		Data:    data,
	}

	if len(meta) > 0 {
		resp.Meta = meta[0]
	}

	c.JSON(http.StatusOK, resp)
}

// CreatedResponse returns a 201 created JSON response
func CreatedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    data,
	})
}

// ErrorResponse returns an error JSON response
func ErrorResponse(c *gin.Context, statusCode int, code string, message string) {
	c.JSON(statusCode, Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
		},
	})
}

// BadRequestResponse returns a 400 bad request error
func BadRequestResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, "BAD_REQUEST", message)
}

// NotFoundResponse returns a 404 not found error
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, "NOT_FOUND", message)
}

// ValidationErrorResponse returns a 400 validation error
func ValidationErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, "VALIDATION_ERROR", message)
}

// DetailedValidationErrorResponse returns a structured validation error response
func DetailedValidationErrorResponse(c *gin.Context, fieldErrors []FieldError) {
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    "VALIDATION_ERROR",
			Message: "Validation failed",
			Fields:  fieldErrors,
		},
	})
}

// InternalServerErrorResponse returns a 500 internal server error
func InternalServerErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, "INTERNAL_ERROR", message)
}

// NoContentResponse returns a 204 no content response
func NoContentResponse(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
