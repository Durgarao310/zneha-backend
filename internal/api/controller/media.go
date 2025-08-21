package controller

import (
	"net/http"
	"strconv"

	"github.com/Durgarao310/zneha-backend/internal/model"
	"github.com/Durgarao310/zneha-backend/internal/service"
	"github.com/Durgarao310/zneha-backend/pkg/api"
	"github.com/gin-gonic/gin"
)

type MediaController struct {
	mediaService *service.MediaService
}

func NewMediaController(mediaService *service.MediaService) *MediaController {
	return &MediaController{
		mediaService: mediaService,
	}
}

func (c *MediaController) CreateMedia(ctx *gin.Context) {
	var media model.Media
	if err := ctx.ShouldBindJSON(&media); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.mediaService.CreateMedia(&media); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusCreated, media)
}

func (c *MediaController) GetMedia(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid media ID"})
		return
	}

	media, err := c.mediaService.GetMediaByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, media)
}

func (c *MediaController) GetMediaByProduct(ctx *gin.Context) {
	idParam := ctx.Param("productId")
	productID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	media, err := c.mediaService.GetMediaByProductID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, media)
}

func (c *MediaController) GetMediaByVariant(ctx *gin.Context) {
	idParam := ctx.Param("variantId")
	variantID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid variant ID"})
		return
	}

	media, err := c.mediaService.GetMediaByVariantID(variantID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, media)
}

func (c *MediaController) UpdateMedia(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid media ID"})
		return
	}

	var media model.Media
	if err := ctx.ShouldBindJSON(&media); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	media.ID = id
	if err := c.mediaService.UpdateMedia(&media); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, media)
}

func (c *MediaController) DeleteMedia(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid media ID"})
		return
	}

	if err := c.mediaService.DeleteMedia(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusNoContent, struct{}{})
}

func (c *MediaController) SetPrimaryMedia(ctx *gin.Context) {
	productIDParam := ctx.Param("productId")
	productID, err := strconv.ParseUint(productIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	mediaIDParam := ctx.Param("mediaId")
	mediaID, err := strconv.ParseUint(mediaIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid media ID"})
		return
	}

	if err := c.mediaService.SetPrimaryMedia(productID, mediaID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.SendSuccess(ctx, http.StatusOK, gin.H{"message": "Primary media set successfully"})
}
