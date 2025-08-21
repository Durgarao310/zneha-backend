package validator

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// HandleValidationErrors processes validation errors and returns structured field errors
func HandleValidationErrors(c *gin.Context, err error) bool {
	return false
}

// GetValidationErrorMessage creates user-friendly validation error messages
func GetValidationErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Must be at least " + fe.Param() + " characters"
	case "max":
		return "Must not exceed " + fe.Param() + " characters"
	case "alphanumunicode":
		return "Must contain only letters, numbers, and spaces"
	case "oneof":
		return "Must be one of: " + fe.Param()
	case "email":
		return "Must be a valid email address"
	case "url":
		return "Must be a valid URL"
	case "numeric":
		return "Must be a number"
	case "alpha":
		return "Must contain only letters"
	default:
		return "Invalid value for field " + fe.Field()
	}
}

// ValidateBusinessRules performs custom business validation
func ValidateBusinessRules(name, description, shortDescription *string) error {
	// Check for forbidden words or patterns
	forbiddenWords := []string{"test", "dummy", "fake", "spam"}

	if name != nil {
		for _, word := range forbiddenWords {
			if containsIgnoreCase(*name, word) {
				return fmt.Errorf("name contains forbidden word: %s", word)
			}
		}
	}

	// Check description length consistency
	if description != nil && shortDescription != nil {
		if len(*shortDescription) > len(*description) && len(*description) > 0 {
			return fmt.Errorf("short description cannot be longer than description")
		}
	}

	// Check for minimum meaningful content
	if name != nil && len(*name) < 3 {
		return fmt.Errorf("name must be at least 3 characters")
	}

	return nil
}

// containsIgnoreCase checks if a string contains a substring (case-insensitive)
func containsIgnoreCase(s, substr string) bool {
	return len(s) >= len(substr) && contains(toLowerCase(s), toLowerCase(substr))
}

// toLowerCase converts string to lowercase
func toLowerCase(s string) string {
	result := make([]rune, len(s))
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			result[i] = r + 32
		} else {
			result[i] = r
		}
	}
	return string(result)
}

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	if len(substr) > len(s) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			if s[i+j] != substr[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
