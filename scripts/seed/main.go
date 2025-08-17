package main

import (
	"log"

	"github.com/Durgarao310/zneha-backend/internal/database"
	"github.com/Durgarao310/zneha-backend/internal/model"
)

func main() {
	db := database.InitPostgres()

	products := []model.Product{
		{Name: "iPhone 15", Description: "Latest Apple iPhone", ShortDescription: "Apple flagship", Status: "active"},
		{Name: "Samsung Galaxy S24", Description: "Flagship Android phone", ShortDescription: "Samsung flagship", Status: "active"},
	}

	for _, p := range products {
		if err := db.Create(&p).Error; err != nil {
			log.Fatalf("❌ Failed to seed product: %v", err)
		}
	}
	log.Println("🌱 Database seeded successfully")
}
