package main

import (
	"log"

	"github.com/durgarao310/zneha-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDB()

	// Initialize Gin
	r := gin.Default()

	// Sample route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	// Start server on port 8080
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
