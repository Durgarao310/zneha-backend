package main

import (
	"log"

	"github.com/Durgarao310/zneha-backend/internal/database"
	"github.com/Durgarao310/zneha-backend/internal/model"
)

func main() {
	db := database.InitPostgres()

	// Run migrations for all models in the correct order
	// Categories first (no dependencies)
	if err := db.AutoMigrate(&model.Category{}); err != nil {
		log.Fatalf("Category migration failed: %v", err)
	}
	log.Println("✅ Category table migrated")

	// Products next (no dependencies on other new models)
	if err := db.AutoMigrate(&model.Product{}); err != nil {
		log.Fatalf("Product migration failed: %v", err)
	}
	log.Println("✅ Product table migrated")

	// Variants (depends on products)
	if err := db.AutoMigrate(&model.Variant{}); err != nil {
		log.Fatalf("Variant migration failed: %v", err)
	}
	log.Println("✅ Variant table migrated")

	// Media last (depends on both products and variants)
	if err := db.AutoMigrate(&model.Media{}); err != nil {
		log.Fatalf("Media migration failed: %v", err)
	}
	log.Println("✅ Media table migrated")

	log.Println("🎉 All migrations completed successfully!")
}
