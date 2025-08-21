package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestMetaMiddleware adds request ID and start time to the context.
func RequestMetaMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set request ID
		requestID := uuid.New().String()
		c.Set("requestID", requestID)

		// Set start time for processing time calculation
		c.Set("startTime", time.Now())

		// Add request ID to response header for traceability
		c.Writer.Header().Set("X-Request-ID", requestID)

		c.Next()
	}
}
