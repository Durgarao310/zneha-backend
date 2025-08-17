package v1

import (
	"strconv"

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
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	if err := h.service.Create(&product); err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.CreatedResponse(c, product)
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}
	
	// Example of adding metadata for pagination (can be enhanced later)
	meta := map[string]interface{}{
		"total": len(products),
		"page":  1,
		"limit": len(products),
	}
	
	utils.SuccessResponse(c, products, meta)
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
	
	utils.SuccessResponse(c, product)
}

func (h *ProductHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID format")
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}
	product.ID = id

	if err := h.service.Update(&product); err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}
	
	utils.SuccessResponse(c, product)
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
