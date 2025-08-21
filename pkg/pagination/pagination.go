package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// PaginationParams represents pagination parameters
type PaginationParams struct {
	Page  int
	Limit int
}

// DefaultPage is the default page number
const DefaultPage = 1

// DefaultLimit is the default number of items per page
const DefaultLimit = 10

// MaxLimit is the maximum allowed limit per page
const MaxLimit = 100

// GetPaginationParams extracts and validates pagination parameters from query string
func GetPaginationParams(ctx *gin.Context) PaginationParams {
	page := DefaultPage
	limit := DefaultLimit

	// Parse page parameter
	if pageStr := ctx.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	// Parse limit parameter
	if limitStr := ctx.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= MaxLimit {
			limit = l
		}
	}

	return PaginationParams{
		Page:  page,
		Limit: limit,
	}
}

// GetPaginationParamsWithCustomLimits allows custom default and max limits
func GetPaginationParamsWithCustomLimits(ctx *gin.Context, defaultLimit, maxLimit int) PaginationParams {
	page := DefaultPage
	limit := defaultLimit

	// Parse page parameter
	if pageStr := ctx.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	// Parse limit parameter
	if limitStr := ctx.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= maxLimit {
			limit = l
		}
	}

	return PaginationParams{
		Page:  page,
		Limit: limit,
	}
}
