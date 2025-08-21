package repository

import (
	"github.com/Durgarao310/zneha-backend/internal/model"
	"gorm.io/gorm"
)

type MediaRepository interface {
	Create(media *model.Media) error
	GetByID(id uint64) (*model.Media, error)
	GetByProductID(productID uint64) ([]model.Media, error)
	GetByVariantID(variantID uint64) ([]model.Media, error)
	GetPrimaryByProductID(productID uint64) (*model.Media, error)
	Update(media *model.Media) error
	Delete(id uint64) error
	SetPrimary(productID uint64, mediaID uint64) error
}

type mediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepository{db: db}
}

func (r *mediaRepository) Create(media *model.Media) error {
	return r.db.Create(media).Error
}

func (r *mediaRepository) GetByID(id uint64) (*model.Media, error) {
	var media model.Media
	err := r.db.First(&media, id).Error
	return &media, err
}

func (r *mediaRepository) GetByProductID(productID uint64) ([]model.Media, error) {
	var media []model.Media
	err := r.db.Where("product_id = ?", productID).Order("position ASC").Find(&media).Error
	return media, err
}

func (r *mediaRepository) GetByVariantID(variantID uint64) ([]model.Media, error) {
	var media []model.Media
	err := r.db.Where("variant_id = ?", variantID).Order("position ASC").Find(&media).Error
	return media, err
}

func (r *mediaRepository) GetPrimaryByProductID(productID uint64) (*model.Media, error) {
	var media model.Media
	err := r.db.Where("product_id = ? AND is_primary = ?", productID, true).First(&media).Error
	return &media, err
}

func (r *mediaRepository) Update(media *model.Media) error {
	return r.db.Save(media).Error
}

func (r *mediaRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Media{}, id).Error
}

func (r *mediaRepository) SetPrimary(productID uint64, mediaID uint64) error {
	// First, set all media for this product to not primary
	if err := r.db.Model(&model.Media{}).Where("product_id = ?", productID).Update("is_primary", false).Error; err != nil {
		return err
	}
	// Then set the specified media as primary
	return r.db.Model(&model.Media{}).Where("id = ? AND product_id = ?", mediaID, productID).Update("is_primary", true).Error
}
