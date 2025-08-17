package v1

import (
	"strconv"

	"github.com/Durgarao310/zneha-backend/internal/dto"
	"github.com/Durgarao310/zneha-backend/internal/model"
	"github.com/Durgarao310/zneha-backend/internal/service"
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
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
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
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.CreatedResponse(c, dto.ToProductResponse(&product))
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
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
		utils.BadRequestResponse(c, "Invalid product ID format")
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		utils.NotFoundResponse(c, "Product not found")
		return
	}

	utils.SuccessResponse(c, dto.ToProductResponse(product))
}

func (h *ProductHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID format")
		return
	}

	var req dto.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
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
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, dto.ToProductResponse(&product))
}

func (h *ProductHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID format")
		return
	}

	if err := h.service.Delete(id); err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.NoContentResponse(c)
}
