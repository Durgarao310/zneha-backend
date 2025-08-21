package middleware

import (
	"net/http"

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

		// Handle JSON parsing errors
		if err.Error() == "invalid JSON" || err.Type == gin.ErrorTypeBind {
			pkgMiddleware.Error(c, http.StatusBadRequest, "Invalid JSON format", err.Err)
			c.Abort()
			return
		}
	}
}
