package server

import (
	"github.com/Durgarao310/zneha-backend/internal/container"
	"github.com/Durgarao310/zneha-backend/internal/routes"
	"github.com/Durgarao310/zneha-backend/pkg/api"
	"github.com/Durgarao310/zneha-backend/pkg/logger"
	pkgMiddleware "github.com/Durgarao310/zneha-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server
type Server struct {
	router    *gin.Engine
	container *container.Container
	logger    *logger.Logger
}

// NewServer creates a new server instance
func NewServer(container *container.Container, log *logger.Logger) *Server {
	return &Server{
		container: container,
		logger:    log,
	}
}

// SetupRouter configures the gin router with middleware and routes
func (s *Server) SetupRouter() {
	s.router = gin.Default()

	// Configure trusted proxies for security
	if err := s.router.SetTrustedProxies([]string{"127.0.0.1", "::1"}); err != nil {
		s.logger.Error("Failed to set trusted proxies", "error", err)
	}

	// Apply middleware
	s.setupMiddleware()

	// Register routes
	s.registerRoutes()
}

// setupMiddleware applies global middleware
func (s *Server) setupMiddleware() {
	// CORS middleware (should be first to handle preflight requests)
	s.router.Use(pkgMiddleware.DevelopmentCORS())
	
	// Other middleware
	s.router.Use(api.RequestMetaMiddleware())
	s.router.Use(pkgMiddleware.JSONMiddleware())
	s.router.Use(pkgMiddleware.GlobalErrorHandler())
}

// registerRoutes registers all application routes
func (s *Server) registerRoutes() {
	s.logger.Info("Starting server with middleware...")
	routes.RegisterRoutes(
		s.router,
		s.container.ProductController,
		s.container.CategoryController,
		s.container.MediaController,
		s.container.VariantController,
	)
}

// Start starts the HTTP server
func (s *Server) Start(port string) error {
	s.logger.Info("Server listening on :" + port)
	return s.router.Run(":" + port)
}

// GetRouter returns the gin router (useful for testing)
func (s *Server) GetRouter() *gin.Engine {
	return s.router
}
