package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error sends a JSON error response
func Error(c *gin.Context, code int, message string, err error) {
	c.JSON(code, gin.H{
		"error":   message,
		"details": err.Error(),
	})
}

// GlobalErrorHandler middleware
func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				Error(c, http.StatusInternalServerError, "Internal Server Error", r.(error))
				c.Abort()
			}
		}()
		c.Next()
	}
}
