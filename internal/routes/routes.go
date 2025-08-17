package routes

import (
	"github.com/Durgarao310/zneha-backend/internal/api/middleware"
	v1 "github.com/Durgarao310/zneha-backend/internal/api/v1"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, productHandler *v1.ProductHandler) {
	api := router.Group("/api/v1")
	api.Use(middleware.ValidationMiddleware())
	{
		products := api.Group("/products")
		{
			products.POST("/", productHandler.Create)
			products.GET("/", productHandler.GetAll)
			products.GET("/:id", productHandler.GetByID)
			products.PUT("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)
		}
	}
}
