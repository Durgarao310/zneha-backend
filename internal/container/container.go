package container

import (
	"github.com/Durgarao310/zneha-backend/internal/api/controller"
	"github.com/Durgarao310/zneha-backend/internal/repository"
	"github.com/Durgarao310/zneha-backend/internal/service"
	"gorm.io/gorm"
)

// Container holds all application dependencies
type Container struct {
	// Repositories
	ProductRepo  repository.ProductRepository
	CategoryRepo repository.CategoryRepository
	MediaRepo    repository.MediaRepository
	VariantRepo  repository.VariantRepository

	// Services
	ProductService  service.ProductService
	CategoryService service.CategoryService
	MediaService    *service.MediaService
	VariantService  *service.VariantService

	// Controllers
	ProductController  controller.ProductController
	CategoryController *controller.CategoryController
	MediaController    *controller.MediaController
	VariantController  *controller.VariantController
}

// NewContainer creates and initializes all dependencies
func NewContainer(db *gorm.DB) *Container {
	c := &Container{}

	// Initialize repositories
	c.initRepositories(db)

	// Initialize services
	c.initServices()

	// Initialize controllers
	c.initControllers()

	return c
}

// initRepositories initializes all repository dependencies
func (c *Container) initRepositories(db *gorm.DB) {
	c.ProductRepo = repository.NewProductRepository(db)
	c.CategoryRepo = repository.NewCategoryRepository(db)
	c.MediaRepo = repository.NewMediaRepository(db)
	c.VariantRepo = repository.NewVariantRepository(db)
}

// initServices initializes all service dependencies
func (c *Container) initServices() {
	c.ProductService = service.NewProductService(c.ProductRepo)
	c.CategoryService = service.NewCategoryService(c.CategoryRepo)
	c.MediaService = service.NewMediaService(c.MediaRepo)
	c.VariantService = service.NewVariantService(c.VariantRepo)
}

// initControllers initializes all controller dependencies
func (c *Container) initControllers() {
	c.ProductController = controller.NewProductController(c.ProductService)
	c.CategoryController = controller.NewCategoryController(c.CategoryService)
	c.MediaController = controller.NewMediaController(c.MediaService)
	c.VariantController = controller.NewVariantController(c.VariantService)
}
