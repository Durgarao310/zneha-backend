// common/errors/error_handler.go
// Global error handler for Gin
package errors

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// ErrorHandler middleware for Gin
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

func handleError(c *gin.Context, err error) {
	appErr := MapError(err)

	// Log internal errors with stacktrace
	if appErr.StatusCode == http.StatusInternalServerError {
		log.Printf("[ERROR] %v\n%s", err, debug.Stack())
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
