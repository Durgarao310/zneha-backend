package database

import (
	"log"
	"os"
	"time"

	"github.com/Durgarao310/zneha-backend/internal/config"
	"github.com/Durgarao310/zneha-backend/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB // global db instance

// InitPostgres initializes the PostgreSQL database connection
func InitPostgres() *gorm.DB {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	// Get DSN from config
	dsn := cfg.GetDatabaseDSN()

	// Configure GORM logger based on environment
	logLevel := logger.Info
	if cfg.IsProduction() {
		logLevel = logger.Warn // Less verbose in production
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // log queries slower than 1s
			LogLevel:      logLevel,
			Colorful:      cfg.IsDevelopment(), // Only colorful in development
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

	log.Printf("✅ Connected to Postgres: %s@%s:%s/%s",
		cfg.Database.User, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	DB = db
	return db
}
