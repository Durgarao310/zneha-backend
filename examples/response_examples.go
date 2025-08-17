package examples

import (
	"strconv"

	"github.com/Durgarao310/zneha-backend/utils"
	"github.com/gin-gonic/gin"
)

// Example handler demonstrating various response utilities
func ExampleProductHandler(c *gin.Context) {
	// Example 1: Success response with data
	product := map[string]interface{}{
		"id":              1,
		"name":            "iPhone 15 Pro",
		"description":     "Latest Apple iPhone",
		"shortDescription": "Flagship phone",
		"status":          "active",
	}
	utils.SuccessResponse(c, product)
}

// Example 2: Success response with metadata (pagination)
func ExampleProductListHandler(c *gin.Context) {
	products := []map[string]interface{}{
		{"id": 1, "name": "iPhone 15 Pro"},
		{"id": 2, "name": "Samsung Galaxy S24"},
	}

	meta := map[string]interface{}{
		"page":       1,
		"limit":      10,
		"total":      2,
		"totalPages": 1,
	}

	utils.SuccessResponse(c, products, meta)
}

// Example 3: Created response (201)
func ExampleCreateProductHandler(c *gin.Context) {
	// Validation
	var body struct {
		Name            string `json:"name" binding:"required"`
		Description     string `json:"description"`
		ShortDescription string `json:"shortDescription"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	// Simulate creating product
	newProduct := map[string]interface{}{
		"id":              3,
		"name":            body.Name,
		"description":     body.Description,
		"shortDescription": body.ShortDescription,
		"status":          "active",
	}

	utils.CreatedResponse(c, newProduct)
}

// Example 4: Various error responses
func ExampleErrorHandler(c *gin.Context) {
	errorType := c.Query("type")

	switch errorType {
	case "validation":
		utils.ValidationErrorResponse(c, "Name field is required")

	case "notfound":
		utils.NotFoundResponse(c, "Product not found")

	case "badrequest":
		utils.BadRequestResponse(c, "Invalid product ID format")

	case "internal":
		utils.InternalServerErrorResponse(c, "Database connection failed")

	case "custom":
		// Custom error with specific status code
		utils.ErrorResponse(c, 422, "UNPROCESSABLE_ENTITY", "Product name already exists")

	default:
		utils.BadRequestResponse(c, "Unknown error type")
	}
}

// Example 5: Delete operation with no content response
func ExampleDeleteHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID format")
		return
	}

	// Simulate deletion
	if id == 999 {
		utils.NotFoundResponse(c, "Product not found")
		return
	}

	// Successful deletion - returns 204 No Content
	utils.NoContentResponse(c)
}

// Example 6: Input validation with multiple errors
func ExampleValidationHandler(c *gin.Context) {
	var body struct {
		Name  string  `json:"name" binding:"required,min=3,max=100"`
		Price float64 `json:"price" binding:"required,gt=0"`
		Email string  `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		// You can customize the validation error message
		utils.ValidationErrorResponse(c, "Validation failed: "+err.Error())
		return
	}

	utils.SuccessResponse(c, map[string]string{
		"message": "Validation passed successfully",
	})
}

// Example 7: Conditional response based on business logic
func ExampleConditionalHandler(c *gin.Context) {
	userRole := c.GetHeader("X-User-Role") // Simulate getting user role

	if userRole != "admin" {
		utils.ErrorResponse(c, 403, "FORBIDDEN", "Insufficient permissions")
		return
	}

	// Admin-only data
	adminData := map[string]interface{}{
		"users":    100,
		"products": 50,
		"orders":   200,
	}

	utils.SuccessResponse(c, adminData)
}
