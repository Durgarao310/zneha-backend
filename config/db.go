package config

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the connection to the PostgreSQL database
func ConnectDB() {
    // Load environment variables from the .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Retrieve database credentials from environment variables
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

    // Validate required environment variables
    if host == "" || user == "" || password == "" || dbName == "" || port == "" {
        log.Fatal("Database environment variables are not properly set")
    }

    // Create the connection string for PostgreSQL
    connectionString := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, password, dbName, port,
    )

    // Open the connection using GORM
    var err error
    DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Verify the database connection
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get database instance: %v", err)
    }

    if err := sqlDB.Ping(); err != nil {
        log.Fatalf("Database connection failed: %v", err)
    }

    log.Println("Successfully connected to PostgreSQL")
}