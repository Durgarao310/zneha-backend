package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	JWT      JWTConfig      `json:"jwt"`
	App      AppConfig      `json:"app"`
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port        string `json:"port"`
	Host        string `json:"host"`
	Environment string `json:"environment"` // development, staging, production
}

// DatabaseConfig holds database connection configuration
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	SSLMode  string `json:"ssl_mode"`
	TimeZone string `json:"time_zone"`
}

// JWTConfig holds JWT-related configuration
type JWTConfig struct {
	Secret     string `json:"secret"`
	ExpiryHour int    `json:"expiry_hour"`
}

// AppConfig holds general application configuration
type AppConfig struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Debug   bool   `json:"debug"`
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		Server: ServerConfig{
			Port:        getEnv("PORT", "8080"),
			Host:        getEnv("HOST", "localhost"),
			Environment: getEnv("ENVIRONMENT", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "zneha"),
			DBName:   getEnv("DB_NAME", "zneha_backend"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			TimeZone: getEnv("DB_TIMEZONE", "Asia/Kolkata"),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
			ExpiryHour: getEnvAsInt("JWT_EXPIRY_HOUR", 24),
		},
		App: AppConfig{
			Name:    getEnv("APP_NAME", "zneha-backend"),
			Version: getEnv("APP_VERSION", "1.0.0"),
			Debug:   getEnvAsBool("DEBUG", true),
		},
	}

	// Validate required configurations
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return config, nil
}

// Validate checks if required configuration values are present
func (c *Config) Validate() error {
	if c.Database.Host == "" {
		return fmt.Errorf("database host is required")
	}
	if c.Database.User == "" {
		return fmt.Errorf("database user is required")
	}
	if c.Database.DBName == "" {
		return fmt.Errorf("database name is required")
	}
	if c.JWT.Secret == "" {
		return fmt.Errorf("JWT secret is required")
	}
	if c.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}
	return nil
}

// GetDatabaseDSN returns the database connection string
func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.Database.Host,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.Port,
		c.Database.SSLMode,
		c.Database.TimeZone,
	)
}

// IsProduction returns true if the environment is production
func (c *Config) IsProduction() bool {
	return c.Server.Environment == "production"
}

// IsDevelopment returns true if the environment is development
func (c *Config) IsDevelopment() bool {
	return c.Server.Environment == "development"
}

// Helper functions

// getEnv gets an environment variable with a fallback value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets an environment variable as integer with a fallback value
func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

// getEnvAsBool gets an environment variable as boolean with a fallback value
func getEnvAsBool(name string, defaultValue bool) bool {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}
