package controller

import (
	"net/http"
	"strconv"

	"github.com/Durgarao310/zneha-backend/internal/dto"
	"github.com/Durgarao310/zneha-backend/internal/model"
	"github.com/Durgarao310/zneha-backend/internal/service"
	"github.com/Durgarao310/zneha-backend/pkg/api"
	"github.com/Durgarao310/zneha-backend/pkg/pagination"
	"github.com/gin-gonic/gin"
)

// ProductController defines the interface for product controller operations
type ProductController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// productController implements ProductController interface
type productController struct {
	service service.ProductService
}

// NewProductController creates a new instance of ProductController
func NewProductController(service service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

// Create handles product creation
func (c *productController) Create(ctx *gin.Context) {
	var req dto.ProductCreateRequest

	status := req.Status
	if status == "" {
		status = "active"
	}

	product := model.Product{
		Name:             req.Name,
		Description:      req.Description,
		ShortDescription: req.ShortDescription,
		Status:           status,
	}

	if err := c.service.Create(&product); err != nil {
		return
	}

	api.SendSuccess(ctx, http.StatusCreated, dto.ToProductResponse(&product))
}

// GetAll handles retrieving all products with optional pagination
func (c *productController) GetAll(ctx *gin.Context) {
	// Use common pagination utility
	params := pagination.GetPaginationParams(ctx)

	// Use efficient database pagination
	products, totalItems, err := c.service.GetWithPagination(params.Page, params.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := dto.ToProductResponseList(products)

	api.SendPaginatedSuccess(ctx, http.StatusOK, responses, params.Page, params.Limit, int(totalItems))
}

// GetByID handles retrieving a product by its ID
func (c *productController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return
	}

	product, err := c.service.GetByID(id)
	if err != nil {
		return
	}

	api.SendSuccess(ctx, http.StatusOK, dto.ToProductResponse(product))
}

// Update handles updating an existing product
func (c *productController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return
	}

	var req dto.ProductUpdateRequest

	// Bind JSON request to struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	status := req.Status
	if status == "" {
		status = "active"
	}

	product := model.Product{
		ID:               id,
		Name:             req.Name,
		Description:      req.Description,
		ShortDescription: req.ShortDescription,
		Status:           status,
	}

	if err := c.service.Update(&product); err != nil {
		return
	}

	api.SendSuccess(ctx, http.StatusOK, dto.ToProductResponse(&product))
}

// Delete handles deleting a product by its ID
func (c *productController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return
	}

	if err := c.service.Delete(id); err != nil {
		return
	}

	api.SendSuccess(ctx, http.StatusNoContent, struct{}{})
}
