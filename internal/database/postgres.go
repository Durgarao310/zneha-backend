package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Durgarao310/zneha-backend/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB // global db instance

// InitPostgres initializes the PostgreSQL database connection
func InitPostgres() *gorm.DB {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: .env file not found, falling back to system env")
	}

	// Get database configuration from environment variables with defaults
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "zneha")
	dbname := getEnv("DB_NAME", "zneha_backend")
	sslmode := getEnv("DB_SSLMODE", "disable")
	timezone := getEnv("DB_TIMEZONE", "Asia/Kolkata")

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)

	// Configure GORM logger (only show slow queries in production)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // log queries slower than 1s
			LogLevel:      logger.Info, // change to logger.Warn in prod
			Colorful:      true,
		},
	)

	// Open DB connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // "product" instead of "products"
		},
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Setup connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to configure connection pool: %v", err)
	}
	sqlDB.SetMaxOpenConns(25)                 // max open connections
	sqlDB.SetMaxIdleConns(5)                  // max idle connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // recycle connections

	// Auto migrate models
	if err := db.AutoMigrate(
		&model.Product{}, // add more models here as your app grows
	); err != nil {
		log.Fatalf("❌ Failed to auto-migrate models: %v", err)
	}

	log.Printf("✅ Connected to Postgres: %s@%s:%s/%s", user, host, port, dbname)

	DB = db
	return db
}

// getEnv returns environment variable or fallback
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
