package main

import (
	"github.com/Durgarao310/zneha-backend/internal/database"
	"github.com/Durgarao310/zneha-backend/internal/repository"
	"github.com/Durgarao310/zneha-backend/internal/routes"
	"github.com/Durgarao310/zneha-backend/internal/service"

	"github.com/gin-gonic/gin"
	v1 "github.com/Durgarao310/zneha-backend/internal/api/v1"
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
	routes.RegisterRoutes(r, productHandler)

	r.Run(":8080") // Start server
}
