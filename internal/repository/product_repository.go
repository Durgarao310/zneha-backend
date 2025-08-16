package repository

import (
	"github.com/Durgarao310/zneha-backend/internal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
	FindAll() ([]model.Product, error)
	FindByID(id uint64) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id uint64) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint64) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Product{}, id).Error
}
