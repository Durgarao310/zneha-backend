package service

import (
	"backend/internal/model"
	"backend/internal/repository"
)

type ProductService interface {
	Create(product *model.Product) error
	GetAll() ([]model.Product, error)
	GetByID(id uint64) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id uint64) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) Create(product *model.Product) error {
	return s.repo.Create(product)
}

func (s *productService) GetAll() ([]model.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetByID(id uint64) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) Update(product *model.Product) error {
	return s.repo.Update(product)
}

func (s *productService) Delete(id uint64) error {
	return s.repo.Delete(id)
}
