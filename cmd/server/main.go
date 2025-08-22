package main

import (
	"log"

	"github.com/Durgarao310/zneha-backend/internal/container"
	"github.com/Durgarao310/zneha-backend/internal/database"
	"github.com/Durgarao310/zneha-backend/internal/server"
	"github.com/Durgarao310/zneha-backend/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	appLogger := logger.New()

	// Initialize zap logger for error middleware
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to initialize zap logger:", err)
	}
	defer zapLogger.Sync()

	// Initialize database
	db := database.InitPostgres()

	// Initialize dependency container
	appContainer := container.NewContainer(db)

	// Initialize server
	appServer := server.NewServer(appContainer, appLogger)

	// Setup router with middleware and routes
	appServer.SetupRouter()

	// Start server
	if err := appServer.Start("8080"); err != nil {
		appLogger.Error("Failed to start server", "error", err)
	}
}
