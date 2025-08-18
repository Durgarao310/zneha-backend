package v1

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

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req dto.ProductCreateRequest

	// Additional business validation
	if err := validator.ValidateBusinessRules(&req.Name, &req.Description, &req.ShortDescription); err != nil {
		c.Error(err)
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

	if err := h.service.Create(&product); err != nil {
		c.Error(middleware.New(middleware.InternalServerError, "Failed to create product", err))
		return
	}

	utils.CreatedResponse(c, dto.ToProductResponse(&product))
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		c.Error(middleware.New(middleware.InternalServerError, "Failed to retrieve products", err))
		return
	}

	responses := dto.ToProductResponseList(products)

	meta := map[string]interface{}{
		"total": len(responses),
		"page":  1,
		"limit": len(responses),
	}

	utils.SuccessResponse(c, responses, meta)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.Error(middleware.New(middleware.InvalidFieldFormat, "Invalid product ID format", err))
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		c.Error(middleware.New(middleware.UserNotFound, "Product not found", err))
		return
	}

	utils.SuccessResponse(c, dto.ToProductResponse(product))
}

func (h *ProductHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.Error(middleware.New(middleware.InvalidFieldFormat, "Invalid product ID format", err))
		return
	}

	var req dto.ProductUpdateRequest
	// Additional business validation
	if err := validator.ValidateBusinessRules(&req.Name, &req.Description, &req.ShortDescription); err != nil {
		c.Error(err)
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

	if err := h.service.Update(&product); err != nil {
		c.Error(middleware.New(middleware.InternalServerError, "Failed to update product", err))
		return
	}

	utils.SuccessResponse(c, dto.ToProductResponse(&product))
}

func (h *ProductHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.Error(middleware.New(middleware.InvalidFieldFormat, "Invalid product ID format", err))
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.Error(middleware.New(middleware.InternalServerError, "Failed to delete product", err))
		return
	}

	utils.NoContentResponse(c)
}
