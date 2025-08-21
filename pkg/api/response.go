package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// APIResponse is the top-level response structure.
type APIResponse[T any] struct {
	Data T    `json:"data"`
	Meta Meta `json:"meta"`
}

// Meta contains metadata about the request.
type Meta struct {
	RequestID        string      `json:"requestId"`
	Timestamp        time.Time   `json:"timestamp"`
	APIVersion       string      `json:"apiVersion"`
	ProcessingTimeMs int64       `json:"processingTimeMs"`
	Pagination       *Pagination `json:"pagination,omitempty"`
}

// Pagination contains pagination information for list responses.
type Pagination struct {
	Page       int  `json:"page"`
	Limit      int  `json:"limit"`
	TotalPages int  `json:"totalPages"`
	TotalItems int  `json:"totalItems"`
	HasNext    bool `json:"hasNext"`
	HasPrev    bool `json:"hasPrev"`
}

// Links supports HATEOAS for resource navigation, universally applicable.
type Links map[string]any

// NewSuccessResponse creates a success response with the given data.
func NewSuccessResponse[T any](c *gin.Context, data T) APIResponse[T] {
	return APIResponse[T]{
		Data: data,
		Meta: Meta{
			RequestID:        getRequestID(c),
			Timestamp:        time.Now().UTC(),
			APIVersion:       "1.0.0",
			ProcessingTimeMs: getProcessingTime(c),
		},
	}
}

// NewPaginatedResponse creates a paginated success response with the given data and pagination info.
func NewPaginatedResponse[T any](c *gin.Context, data T, page, limit, totalItems int) APIResponse[T] {
	totalPages := (totalItems + limit - 1) / limit
	if totalPages == 0 {
		totalPages = 0
	}

	pagination := &Pagination{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalItems: totalItems,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}

	return APIResponse[T]{
		Data: data,
		Meta: Meta{
			RequestID:        getRequestID(c),
			Timestamp:        time.Now().UTC(),
			APIVersion:       "1.0.0",
			ProcessingTimeMs: getProcessingTime(c),
			Pagination:       pagination,
		},
	}
}

// SendSuccess sends a success response with the specified HTTP status code.
func SendSuccess[T any](c *gin.Context, status int, data T) {
	// Validate that the status code is in the success range (2xx)
	if status < 200 || status > 299 {
		// Fallback to 200 if an invalid status is provided
		c.JSON(200, NewSuccessResponse(c, data))
		return
	}
	c.JSON(status, NewSuccessResponse(c, data))
}

// SendPaginatedSuccess sends a paginated success response with the specified HTTP status code.
func SendPaginatedSuccess[T any](c *gin.Context, status int, data T, page, limit, totalItems int) {
	// Validate that the status code is in the success range (2xx)
	if status < 200 || status > 299 {
		// Fallback to 200 if an invalid status is provided
		c.JSON(200, NewPaginatedResponse(c, data, page, limit, totalItems))
		return
	}
	c.JSON(status, NewPaginatedResponse(c, data, page, limit, totalItems))
}

// Helper to get request ID from context (set by middleware).
func getRequestID(c *gin.Context) string {
	if id, exists := c.Get("requestID"); exists {
		return id.(string)
	}
	return uuid.New().String() // Fallback
}

// Helper to calculate processing time from context (set by middleware).
func getProcessingTime(c *gin.Context) int64 {
	if start, exists := c.Get("startTime"); exists {
		return time.Since(start.(time.Time)).Milliseconds()
	}
	return 0
}
