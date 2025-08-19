package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Always set response header to JSON
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Enforce JSON for write operations
		if c.Request.Method == http.MethodPost ||
			c.Request.Method == http.MethodPut ||
			c.Request.Method == http.MethodPatch {

			ct := c.GetHeader("Content-Type")
			if !strings.HasPrefix(ct, "application/json") {
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
