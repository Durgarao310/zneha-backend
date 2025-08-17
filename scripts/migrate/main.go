package main

import (
	"log"

	"github.com/Durgarao310/zneha-backend/internal/database"
	"github.com/Durgarao310/zneha-backend/internal/model"
)

func main() {
	db := database.InitPostgres()
	if err := db.AutoMigrate(&model.Product{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("✅ Migrations completed")
}
