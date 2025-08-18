package main

import (
	"github.com/Durgarao310/zneha-backend/internal/database"
	"github.com/Durgarao310/zneha-backend/internal/repository"
	"github.com/Durgarao310/zneha-backend/internal/routes"
	"github.com/Durgarao310/zneha-backend/internal/service"
	"github.com/Durgarao310/zneha-backend/pkg/logger"
	pkgMiddleware "github.com/Durgarao310/zneha-backend/pkg/middleware"

	v1 "github.com/Durgarao310/zneha-backend/internal/api/v1"
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

	// Dependencies
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := v1.NewProductHandler(productService)

	// Router
	r := gin.Default()

	// Apply global middleware from pkg/middleware
	r.Use(pkgMiddleware.JSONMiddleware())
	r.Use(pkgMiddleware.GlobalErrorHandler(pkgMiddleware.ErrorHandlerConfig{
		Logger: zapLogger,
	}))

	log.Info("Starting server with middleware...")
	routes.RegisterRoutes(r, productHandler)

	log.Info("Server listening on :8080")
	r.Run(":8080") // Start server
}
