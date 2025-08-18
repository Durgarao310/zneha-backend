package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONMiddleware enforces application/json for requests
func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Always set response header to JSON
		c.Writer.Header().Set("Content-Type", "application/json")

		// Enforce JSON for write operations
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			if c.GetHeader("Content-Type") != "application/json" {
				c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{
					"success": false,
					"error": gin.H{
						"code":        "UNSUPPORTED_MEDIA_TYPE",
						"http_status": http.StatusUnsupportedMediaType,
						"message":     "Content-Type must be application/json",
					},
				})
				return
			}
		}

		c.Next()
	}
}
