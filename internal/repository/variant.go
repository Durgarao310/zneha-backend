package repository

import (
	"github.com/Durgarao310/zneha-backend/internal/model"
	"gorm.io/gorm"
)

type VariantRepository interface {
	Create(variant *model.Variant) error
	GetByID(id uint64) (*model.Variant, error)
	GetByProductID(productID uint64) ([]model.Variant, error)
	GetByProductIDWithPagination(productID uint64, page, limit int) ([]model.Variant, int64, error)
	GetBySKU(sku string) (*model.Variant, error)
	GetActiveByProductID(productID uint64) ([]model.Variant, error)
	GetActiveByProductIDWithPagination(productID uint64, page, limit int) ([]model.Variant, int64, error)
	Update(variant *model.Variant) error
	Delete(id uint64) error
	UpdateStock(id uint64, quantity int) error
}

type variantRepository struct {
	db *gorm.DB
}

func NewVariantRepository(db *gorm.DB) VariantRepository {
	return &variantRepository{db: db}
}

func (r *variantRepository) Create(variant *model.Variant) error {
	return r.db.Create(variant).Error
}

func (r *variantRepository) GetByID(id uint64) (*model.Variant, error) {
	var variant model.Variant
	err := r.db.Preload("Product").Preload("Media").First(&variant, id).Error
	return &variant, err
}

func (r *variantRepository) GetByProductID(productID uint64) ([]model.Variant, error) {
	var variants []model.Variant
	err := r.db.Where("product_id = ?", productID).Find(&variants).Error
	return variants, err
}

func (r *variantRepository) GetByProductIDWithPagination(productID uint64, page, limit int) ([]model.Variant, int64, error) {
	var variants []model.Variant
	var total int64

	// Get total count for this product
	if err := r.db.Model(&model.Variant{}).Where("product_id = ?", productID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results using database-level pagination
	offset := (page - 1) * limit
	err := r.db.Where("product_id = ?", productID).
		Limit(limit).
		Offset(offset).
		Find(&variants).Error

	return variants, total, err
}

func (r *variantRepository) GetBySKU(sku string) (*model.Variant, error) {
	var variant model.Variant
	err := r.db.Where("sku = ?", sku).First(&variant).Error
	return &variant, err
}

func (r *variantRepository) GetActiveByProductID(productID uint64) ([]model.Variant, error) {
	var variants []model.Variant
	err := r.db.Where("product_id = ? AND is_active = ?", productID, true).Find(&variants).Error
	return variants, err
}

func (r *variantRepository) GetActiveByProductIDWithPagination(productID uint64, page, limit int) ([]model.Variant, int64, error) {
	var variants []model.Variant
	var total int64

	// Get total count for active variants of this product
	if err := r.db.Model(&model.Variant{}).Where("product_id = ? AND is_active = ?", productID, true).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results using database-level pagination
	offset := (page - 1) * limit
	err := r.db.Where("product_id = ? AND is_active = ?", productID, true).
		Limit(limit).
		Offset(offset).
		Find(&variants).Error

	return variants, total, err
}

func (r *variantRepository) Update(variant *model.Variant) error {
	return r.db.Save(variant).Error
}

func (r *variantRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Variant{}, id).Error
}

func (r *variantRepository) UpdateStock(id uint64, quantity int) error {
	return r.db.Model(&model.Variant{}).Where("id = ?", id).Update("stock_quantity", quantity).Error
}
