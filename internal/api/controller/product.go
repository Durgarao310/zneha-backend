package controller

import (
	"strconv"

	"github.com/Durgarao310/zneha-backend/internal/dto"
	"github.com/Durgarao310/zneha-backend/internal/model"
	"github.com/Durgarao310/zneha-backend/internal/service"
	"github.com/Durgarao310/zneha-backend/pkg/middleware"
	"github.com/Durgarao310/zneha-backend/pkg/validator"
	"github.com/Durgarao310/zneha-backend/utils"

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

	// Bind JSON request to struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(middleware.New(middleware.InvalidFieldFormat, "Invalid request format", err))
		return
	}

	// Additional business validation
	if err := validator.ValidateBusinessRules(&req.Name, &req.Description, &req.ShortDescription); err != nil {
		ctx.Error(err)
		return
	}

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
		ctx.Error(middleware.New(middleware.InternalServerError, "Failed to create product", err))
		return
	}

	utils.CreatedResponse(ctx, dto.ToProductResponse(&product))
}

// GetAll handles retrieving all products with optional pagination
func (c *productController) GetAll(ctx *gin.Context) {
	products, err := c.service.GetAll()
	if err != nil {
		ctx.Error(middleware.New(middleware.InternalServerError, "Failed to retrieve products", err))
		return
	}

	responses := dto.ToProductResponseList(products)

	meta := map[string]interface{}{
		"total": len(responses),
		"page":  1,
		"limit": len(responses),
	}

	utils.SuccessResponse(ctx, responses, meta)
}

// GetByID handles retrieving a product by its ID
func (c *productController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.Error(middleware.New(middleware.InvalidFieldFormat, "Invalid product ID format", err))
		return
	}

	product, err := c.service.GetByID(id)
	if err != nil {
		ctx.Error(middleware.New(middleware.UserNotFound, "Product not found", err))
		return
	}

	utils.SuccessResponse(ctx, dto.ToProductResponse(product))
}

// Update handles updating an existing product
func (c *productController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.Error(middleware.New(middleware.InvalidFieldFormat, "Invalid product ID format", err))
		return
	}

	var req dto.ProductUpdateRequest

	// Bind JSON request to struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(middleware.New(middleware.InvalidFieldFormat, "Invalid request format", err))
		return
	}

	// Additional business validation
	if err := validator.ValidateBusinessRules(&req.Name, &req.Description, &req.ShortDescription); err != nil {
		ctx.Error(err)
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
		ctx.Error(middleware.New(middleware.InternalServerError, "Failed to update product", err))
		return
	}

	utils.SuccessResponse(ctx, dto.ToProductResponse(&product))
}

// Delete handles deleting a product by its ID
func (c *productController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.Error(middleware.New(middleware.InvalidFieldFormat, "Invalid product ID format", err))
		return
	}

	if err := c.service.Delete(id); err != nil {
		ctx.Error(middleware.New(middleware.InternalServerError, "Failed to delete product", err))
		return
	}

	utils.NoContentResponse(ctx)
}
