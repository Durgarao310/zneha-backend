package middleware

import (
	"github.com/Durgarao310/zneha-backend/internal/common/errors"
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
		if _, ok := err.Err.(*errors.AppError); ok {
			// Let the global error handler deal with it
			return
		}

		// Handle JSON parsing errors
		if err.Error() == "invalid JSON" || err.Type == gin.ErrorTypeBind {
			c.Error(errors.New(errors.ValidationError, "Invalid JSON format", err.Err))
			c.Abort()
			return
		}
	}
}

// BindJSON is a helper function that binds JSON and adds errors to gin context
func BindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.Error(err)
		return false
	}
	return true
}
