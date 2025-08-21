package controller

import (
	"net/http"
	"strconv"

	"github.com/Durgarao310/zneha-backend/internal/model"
	"github.com/Durgarao310/zneha-backend/internal/service"
	"github.com/Durgarao310/zneha-backend/pkg/api"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.categoryService.CreateCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusCreated, category)
}

func (c *CategoryController) GetCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := c.categoryService.GetCategoryByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, category)
}

func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	// Get pagination parameters from query string
	page := 1
	limit := 10

	if pageStr := ctx.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr := ctx.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalItems := len(categories)

	// Calculate pagination
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	// Handle pagination bounds
	if startIndex >= totalItems {
		// If page is beyond available data, return empty results
		categories = []model.Category{}
		api.SendPaginatedSuccess(ctx, http.StatusOK, categories, page, limit, totalItems)
		return
	}

	if endIndex > totalItems {
		endIndex = totalItems
	}

	// Get the paginated slice
	if startIndex < totalItems {
		categories = categories[startIndex:endIndex]
	}

	api.SendPaginatedSuccess(ctx, http.StatusOK, categories, page, limit, totalItems)
}

func (c *CategoryController) GetRootCategories(ctx *gin.Context) {
	categories, err := c.categoryService.GetMainCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, categories)
}

func (c *CategoryController) GetSubcategories(ctx *gin.Context) {
	idParam := ctx.Param("id")
	parentID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parent category ID"})
		return
	}

	// Get pagination parameters from query string
	page := 1
	limit := 10

	if pageStr := ctx.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr := ctx.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	subcategories, err := c.categoryService.GetSubCategories(parentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalItems := len(subcategories)

	// Calculate pagination
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	// Handle pagination bounds
	if startIndex >= totalItems {
		// If page is beyond available data, return empty results
		subcategories = []model.Category{}
		api.SendPaginatedSuccess(ctx, http.StatusOK, subcategories, page, limit, totalItems)
		return
	}

	if endIndex > totalItems {
		endIndex = totalItems
	}

	// Get the paginated slice
	if startIndex < totalItems {
		subcategories = subcategories[startIndex:endIndex]
	}

	api.SendPaginatedSuccess(ctx, http.StatusOK, subcategories, page, limit, totalItems)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = id
	if err := c.categoryService.UpdateCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, category)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := c.categoryService.DeleteCategory(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusNoContent, struct{}{})
}
