package main

import (
	"github.com/Durgarao310/zneha-backend/internal/api/controller"
	"github.com/Durgarao310/zneha-backend/internal/database"
	"github.com/Durgarao310/zneha-backend/internal/repository"
	"github.com/Durgarao310/zneha-backend/internal/routes"
	"github.com/Durgarao310/zneha-backend/internal/service"
	"github.com/Durgarao310/zneha-backend/pkg/api"
	"github.com/Durgarao310/zneha-backend/pkg/logger"
	pkgMiddleware "github.com/Durgarao310/zneha-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	log := logger.New()

	// Initialize zap logger for error middleware
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	// Init DB
	db := database.InitPostgres() // implement in internal/database/postgres.go

	// Dependencies - Repositories
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	mediaRepo := repository.NewMediaRepository(db)
	variantRepo := repository.NewVariantRepository(db)

	// Dependencies - Services
	productService := service.NewProductService(productRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	mediaService := service.NewMediaService(mediaRepo)
	variantService := service.NewVariantService(variantRepo)

	// Dependencies - Controllers
	productController := controller.NewProductController(productService)
	categoryController := controller.NewCategoryController(categoryService)
	mediaController := controller.NewMediaController(mediaService)
	variantController := controller.NewVariantController(variantService)

	// Router
	r := gin.Default()

	// Apply request metadata middleware
	r.Use(api.RequestMetaMiddleware())

	// Apply global middleware from pkg/middleware
	r.Use(pkgMiddleware.JSONMiddleware())
	r.Use(pkgMiddleware.GlobalErrorHandler())

	log.Info("Starting server with middleware...")
	routes.RegisterRoutes(r, productController, categoryController, mediaController, variantController)

	log.Info("Server listening on :8080")
	r.Run(":8080") // Start server
}
