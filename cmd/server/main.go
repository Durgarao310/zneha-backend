package main

import (
	commonErrors "github.com/Durgarao310/zneha-backend/internal/common/errors"
	"github.com/Durgarao310/zneha-backend/internal/database"
	"github.com/Durgarao310/zneha-backend/internal/repository"
	"github.com/Durgarao310/zneha-backend/internal/routes"
	"github.com/Durgarao310/zneha-backend/internal/service"
	"github.com/Durgarao310/zneha-backend/pkg/middleware" // <- new pkg

	v1 "github.com/Durgarao310/zneha-backend/internal/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init DB
	db := database.InitPostgres() // implement in internal/database/postgres.go

	// Dependencies
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := v1.NewProductHandler(productService)

	// Router
	r := gin.Default()
	r.Use(commonErrors.GlobalErrorHandler())
	r.Use(middleware.JSONMiddleware()) // force application/json\
	routes.RegisterRoutes(r, productHandler)

	r.Run(":8080") // Start server
}
