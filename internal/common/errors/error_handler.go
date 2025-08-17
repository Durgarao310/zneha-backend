package errors

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	appErrors "github.com/Durgarao310/zneha-backend/internal/errors"
	"github.com/Durgarao310/zneha-backend/utils"
)

// GlobalErrorHandler catches everything
func GlobalErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Recover from panic
		defer func() {
			if r := recover(); r != nil {
				utils.InternalServerErrorResponse(ctx, "Internal Server Error")
				ctx.Abort()
			}
		}()

		ctx.Next()

		// Grab the first error
		err := ctx.Errors.ByType(gin.ErrorTypeAny).Last()
		if err == nil {
			return
		}

		// Handle known AppError
		if appErr, ok := err.Err.(*appErrors.AppError); ok {
			switch appErr.Code {
			case http.StatusBadRequest:
				utils.BadRequestResponse(ctx, appErr.Message)
			case http.StatusNotFound:
				utils.NotFoundResponse(ctx, appErr.Message)
			case http.StatusUnauthorized:
				utils.ErrorResponse(ctx, http.StatusUnauthorized, "UNAUTHORIZED", appErr.Message)
			default:
				utils.ErrorResponse(ctx, appErr.Code, "ERROR", appErr.Message)
			}
			ctx.Abort()
			return
		}

		// Handle validator errors
		if verrs, ok := err.Err.(validator.ValidationErrors); ok {
			fieldErrors := make([]utils.FieldError, 0, len(verrs))
			for _, fe := range verrs {
				fieldName := fe.Field()
				// Try to derive json tag if possible
				if fe.StructField() != "" && fe.Kind() != reflect.Invalid {
					// Attempt to get JSON tag via reflection on the struct type present in context (not trivial here)
					// Leaving as provided field name for simplicity
				}
				msg := validationMessage(fe)
				fieldErrors = append(fieldErrors, utils.FieldError{
					Field:   fieldName,
					Tag:     fe.Tag(),
					Value:   fe.Value(),
					Param:   fe.Param(),
					Message: msg,
				})
			}
			utils.DetailedValidationErrorResponse(ctx, fieldErrors)
			ctx.Abort()
			return
		}

		// Fallback: unknown error
		utils.InternalServerErrorResponse(ctx, err.Error())
		ctx.Abort()
	}
}

// validationMessage builds a human readable message for a validator FieldError
func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email"
	case "min":
		return "must be at least " + fe.Param()
	case "max":
		return "must be at most " + fe.Param()
	case "len":
		return "must have length " + fe.Param()
	}
	return fe.Tag()
}
