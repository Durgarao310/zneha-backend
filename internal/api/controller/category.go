package controller

import (
	"net/http"
	"strconv"

	"github.com/Durgarao310/zneha-backend/internal/model"
	"github.com/Durgarao310/zneha-backend/internal/service"
	"github.com/Durgarao310/zneha-backend/pkg/api"
	"github.com/Durgarao310/zneha-backend/pkg/pagination"
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
	// Use common pagination utility
	params := pagination.GetPaginationParams(ctx)

	// Use efficient database pagination
	categories, totalItems, err := c.categoryService.GetAllCategoriesWithPagination(params.Page, params.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendPaginatedSuccess(ctx, http.StatusOK, categories, params.Page, params.Limit, int(totalItems))
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

	// Use common pagination utility
	params := pagination.GetPaginationParams(ctx)

	subcategories, totalItems, err := c.categoryService.GetSubCategoriesWithPagination(parentID, params.Page, params.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendPaginatedSuccess(ctx, http.StatusOK, subcategories, params.Page, params.Limit, int(totalItems))
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
