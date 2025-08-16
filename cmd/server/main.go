package main

import (
	"your-app/internal/database"
	"your-app/internal/repository"
	"your-app/internal/routes"
	"your-app/internal/service"

	"github.com/gin-gonic/gin"
	v1 "github.com/yourname/go-gin-hello/internal/api/v1"
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
