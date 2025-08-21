package service

import (
	"errors"

	"github.com/Durgarao310/zneha-backend/internal/model"
	"github.com/Durgarao310/zneha-backend/internal/repository"
)

type CategoryService interface {
	CreateCategory(category *model.Category) error
	GetCategoryByID(id uint64) (*model.Category, error)
	GetAllCategories() ([]model.Category, error)
	GetMainCategories() ([]model.Category, error)
	GetSubCategories(parentID uint64) ([]model.Category, error)
	UpdateCategory(category *model.Category) error
	DeleteCategory(id uint64) error
	GetAllCategoriesWithPagination(page, limit int) ([]model.Category, int64, error)
	GetSubCategoriesWithPagination(parentID uint64, page, limit int) ([]model.Category, int64, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) CreateCategory(category *model.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}

	// Validate parent category exists if parentID is provided
	if category.ParentID != nil {
		_, err := s.categoryRepo.GetByID(*category.ParentID)
		if err != nil {
			return errors.New("parent category does not exist")
		}
		category.Depth = 1 // Set as subcategory
	} else {
		category.Depth = 0 // Set as main category
	}

	return s.categoryRepo.Create(category)
}

func (s *categoryService) GetCategoryByID(id uint64) (*model.Category, error) {
	return s.categoryRepo.GetByID(id)
}

func (s *categoryService) GetAllCategories() ([]model.Category, error) {
	return s.categoryRepo.GetAll()
}

func (s *categoryService) GetMainCategories() ([]model.Category, error) {
	return s.categoryRepo.GetMainCategories()
}

func (s *categoryService) GetSubCategories(parentID uint64) ([]model.Category, error) {
	return s.categoryRepo.GetByParentID(&parentID)
}

func (s *categoryService) UpdateCategory(category *model.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}

	// Check if category exists
	_, err := s.categoryRepo.GetByID(category.ID)
	if err != nil {
		return errors.New("category not found")
	}

	return s.categoryRepo.Update(category)
}

func (s *categoryService) DeleteCategory(id uint64) error {
	// Check if category has subcategories
	subCategories, err := s.categoryRepo.GetByParentID(&id)
	if err == nil && len(subCategories) > 0 {
		return errors.New("cannot delete category with subcategories")
	}

	return s.categoryRepo.Delete(id)
}

func (s *categoryService) GetAllCategoriesWithPagination(page, limit int) ([]model.Category, int64, error) {
	return s.categoryRepo.GetAllWithPagination(page, limit)
}

func (s *categoryService) GetSubCategoriesWithPagination(parentID uint64, page, limit int) ([]model.Category, int64, error) {
	return s.categoryRepo.GetByParentIDWithPagination(&parentID, page, limit)
}
