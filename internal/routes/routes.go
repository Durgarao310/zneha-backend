package routes

import (
	"github.com/Durgarao310/zneha-backend/internal/api/controller"
	"github.com/Durgarao310/zneha-backend/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, 
	productController controller.ProductController,
	categoryController *controller.CategoryController,
	mediaController *controller.MediaController,
	variantController *controller.VariantController) {
	api := router.Group("/api/v1")
	api.Use(middleware.ValidationMiddleware())
	{
		// Products routes
		products := api.Group("/products")
		{
			products.POST("/", productController.Create)
			products.GET("/", productController.GetAll)
			products.GET("/:id", productController.GetByID)
			products.PUT("/:id", productController.Update)
			products.DELETE("/:id", productController.Delete)
		}

		// Categories routes
		categories := api.Group("/categories")
		{
			categories.POST("/", categoryController.CreateCategory)
			categories.GET("/", categoryController.GetAllCategories)
			categories.GET("/root", categoryController.GetRootCategories)
			categories.GET("/:id", categoryController.GetCategory)
			categories.GET("/:id/subcategories", categoryController.GetSubcategories)
			categories.PUT("/:id", categoryController.UpdateCategory)
			categories.DELETE("/:id", categoryController.DeleteCategory)
		}

		// Media routes
		media := api.Group("/media")
		{
			media.POST("/", mediaController.CreateMedia)
			media.GET("/:id", mediaController.GetMedia)
			media.PUT("/:id", mediaController.UpdateMedia)
			media.DELETE("/:id", mediaController.DeleteMedia)
			media.GET("/product/:productId", mediaController.GetMediaByProduct)
			media.GET("/variant/:variantId", mediaController.GetMediaByVariant)
			media.PUT("/product/:productId/primary/:mediaId", mediaController.SetPrimaryMedia)
		}

		// Variants routes
		variants := api.Group("/variants")
		{
			variants.POST("/", variantController.CreateVariant)
			variants.GET("/:id", variantController.GetVariant)
			variants.GET("/sku/:sku", variantController.GetVariantBySKU)
			variants.GET("/product/:productId", variantController.GetVariantsByProduct)
			variants.GET("/product/:productId/active", variantController.GetActiveVariantsByProduct)
			variants.PUT("/:id", variantController.UpdateVariant)
			variants.PUT("/:id/stock", variantController.UpdateStock)
			variants.PUT("/:id/activate", variantController.ActivateVariant)
			variants.PUT("/:id/deactivate", variantController.DeactivateVariant)
			variants.DELETE("/:id", variantController.DeleteVariant)
		}
	}
}
