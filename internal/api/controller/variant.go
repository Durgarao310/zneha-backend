package controller

import (
	"net/http"
	"strconv"

	"github.com/Durgarao310/zneha-backend/internal/model"
	"github.com/Durgarao310/zneha-backend/internal/service"
	"github.com/Durgarao310/zneha-backend/pkg/api"
	"github.com/gin-gonic/gin"
)

type VariantController struct {
	variantService *service.VariantService
}

func NewVariantController(variantService *service.VariantService) *VariantController {
	return &VariantController{
		variantService: variantService,
	}
}

func (c *VariantController) CreateVariant(ctx *gin.Context) {
	var variant model.Variant
	if err := ctx.ShouldBindJSON(&variant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.variantService.CreateVariant(&variant); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusCreated, variant)
}

func (c *VariantController) GetVariant(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid variant ID"})
		return
	}

	variant, err := c.variantService.GetVariantByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Variant not found"})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, variant)
}

func (c *VariantController) GetVariantBySKU(ctx *gin.Context) {
	sku := ctx.Param("sku")

	variant, err := c.variantService.GetVariantBySKU(sku)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Variant not found"})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, variant)
}

func (c *VariantController) GetVariantsByProduct(ctx *gin.Context) {
	idParam := ctx.Param("productId")
	productID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	variants, err := c.variantService.GetVariantsByProductID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, variants)
}

func (c *VariantController) GetActiveVariantsByProduct(ctx *gin.Context) {
	idParam := ctx.Param("productId")
	productID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	variants, err := c.variantService.GetActiveVariantsByProductID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, variants)
}

func (c *VariantController) UpdateVariant(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid variant ID"})
		return
	}

	var variant model.Variant
	if err := ctx.ShouldBindJSON(&variant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	variant.ID = id
	if err := c.variantService.UpdateVariant(&variant); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, variant)
}

func (c *VariantController) UpdateStock(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid variant ID"})
		return
	}

	var stockUpdate struct {
		Quantity int `json:"quantity" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&stockUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.variantService.UpdateStock(id, stockUpdate.Quantity); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, gin.H{"message": "Stock updated successfully"})
}

func (c *VariantController) DeleteVariant(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid variant ID"})
		return
	}

	if err := c.variantService.DeleteVariant(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusNoContent, struct{}{})
}

func (c *VariantController) DeactivateVariant(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid variant ID"})
		return
	}

	if err := c.variantService.DeactivateVariant(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, gin.H{"message": "Variant deactivated successfully"})
}

func (c *VariantController) ActivateVariant(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid variant ID"})
		return
	}

	if err := c.variantService.ActivateVariant(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, gin.H{"message": "Variant activated successfully"})
}
