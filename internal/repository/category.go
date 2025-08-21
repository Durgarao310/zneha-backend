package repository

import (
	"github.com/Durgarao310/zneha-backend/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *model.Category) error
	GetByID(id uint64) (*model.Category, error)
	GetAll() ([]model.Category, error)
	GetByParentID(parentID *uint64) ([]model.Category, error)
	GetMainCategories() ([]model.Category, error)
	Update(category *model.Category) error
	Delete(id uint64) error
	GetAllWithPagination(page, limit int) ([]model.Category, int64, error)
	GetByParentIDWithPagination(parentID *uint64, page, limit int) ([]model.Category, int64, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetByID(id uint64) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *categoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetByParentID(parentID *uint64) ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Where("parent_id = ?", parentID).Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetMainCategories() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Where("parent_id IS NULL").Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Category{}, id).Error
}

func (r *categoryRepository) GetAllWithPagination(page, limit int) ([]model.Category, int64, error) {
	var categories []model.Category
	var total int64

	// Get total count
	if err := r.db.Model(&model.Category{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	err := r.db.Offset(offset).Limit(limit).Find(&categories).Error
	return categories, total, err
}

func (r *categoryRepository) GetByParentIDWithPagination(parentID *uint64, page, limit int) ([]model.Category, int64, error) {
	var categories []model.Category
	var total int64

	query := r.db.Model(&model.Category{}).Where("parent_id = ?", parentID)

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	err := query.Offset(offset).Limit(limit).Find(&categories).Error
	return categories, total, err
}
