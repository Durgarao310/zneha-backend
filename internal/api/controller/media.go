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

	media, err := c.mediaService.GetMediaByProductID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalItems := len(media)

	// Calculate pagination
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	// Handle pagination bounds
	if startIndex >= totalItems {
		// If page is beyond available data, return empty results
		media = []model.Media{}
		api.SendPaginatedSuccess(ctx, http.StatusOK, media, page, limit, totalItems)
		return
	}

	if endIndex > totalItems {
		endIndex = totalItems
	}

	// Get the paginated slice
	if startIndex < totalItems {
		media = media[startIndex:endIndex]
	}

	api.SendPaginatedSuccess(ctx, http.StatusOK, media, page, limit, totalItems)
}

func (c *MediaController) GetMediaByVariant(ctx *gin.Context) {
	idParam := ctx.Param("variantId")
	variantID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid variant ID"})
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

	media, err := c.mediaService.GetMediaByVariantID(variantID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalItems := len(media)

	// Calculate pagination
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	// Handle pagination bounds
	if startIndex >= totalItems {
		// If page is beyond available data, return empty results
		media = []model.Media{}
		api.SendPaginatedSuccess(ctx, http.StatusOK, media, page, limit, totalItems)
		return
	}

	if endIndex > totalItems {
		endIndex = totalItems
	}

	// Get the paginated slice
	if startIndex < totalItems {
		media = media[startIndex:endIndex]
	}

	api.SendPaginatedSuccess(ctx, http.StatusOK, media, page, limit, totalItems)
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
