package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() *gorm.DB {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Get database configuration from environment variables with defaults
	host := getEnvWithDefault("DB_HOST", "localhost")
	port := getEnvWithDefault("DB_PORT", "5432")
	user := getEnvWithDefault("DB_USER", "postgres")
	password := getEnvWithDefault("DB_PASSWORD", "zneha")
	dbname := getEnvWithDefault("DB_NAME", "zneha_backend")
	sslmode := getEnvWithDefault("DB_SSLMODE", "disable")
	timezone := getEnvWithDefault("DB_TIMEZONE", "Asia/Kolkata")

	// Build connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database with DSN: host=%s user=%s dbname=%s port=%s", host, user, dbname, port)
		log.Fatal("Failed to connect to database:", err)
	}

	log.Printf("Successfully connected to database: %s@%s:%s/%s", user, host, port, dbname)
	return db
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
