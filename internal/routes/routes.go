package routes

import (
	"github.com/Durgarao310/zneha-backend/internal/api/controller"
	"github.com/Durgarao310/zneha-backend/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, productController controller.ProductController) {
	api := router.Group("/api/v1")
	api.Use(middleware.ValidationMiddleware())
	{
		products := api.Group("/products")
		{
			products.POST("/", productController.Create)
			products.GET("/", productController.GetAll)
			products.GET("/:id", productController.GetByID)
			products.PUT("/:id", productController.Update)
			products.DELETE("/:id", productController.Delete)
		}
	}
}
