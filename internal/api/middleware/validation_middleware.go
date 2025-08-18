package middleware

import (
	pkgMiddleware "github.com/Durgarao310/zneha-backend/pkg/middleware"
	pkgValidator "github.com/Durgarao310/zneha-backend/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ValidationMiddleware handles JSON binding and validation errors automatically
func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors that occurred during request processing
		err := c.Errors.Last()
		if err == nil {
			return
		}

		// Handle validation errors from ShouldBindJSON
		if validationErrors, ok := err.Err.(validator.ValidationErrors); ok {
			pkgValidator.HandleValidationErrors(c, validationErrors)
			c.Abort()
			return
		}

		// Handle custom AppError with specific error codes
		if _, ok := err.Err.(*pkgMiddleware.AppError); ok {
			// Let the global error handler deal with it
			return
		}

		// Handle JSON parsing errors
		if err.Error() == "invalid JSON" || err.Type == gin.ErrorTypeBind {
			c.Error(pkgMiddleware.New(pkgMiddleware.ValidationError, "Invalid JSON format", err.Err))
			c.Abort()
			return
		}
	}
}
