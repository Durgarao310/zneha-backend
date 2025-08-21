package service

import (
	"github.com/Durgarao310/zneha-backend/internal/model"
	"github.com/Durgarao310/zneha-backend/internal/repository"
)

type VariantService struct {
	variantRepo repository.VariantRepository
}

func NewVariantService(variantRepo repository.VariantRepository) *VariantService {
	return &VariantService{
		variantRepo: variantRepo,
	}
}

func (s *VariantService) CreateVariant(variant *model.Variant) error {
	return s.variantRepo.Create(variant)
}

func (s *VariantService) GetVariantByID(id uint64) (*model.Variant, error) {
	return s.variantRepo.GetByID(id)
}

func (s *VariantService) GetVariantBySKU(sku string) (*model.Variant, error) {
	return s.variantRepo.GetBySKU(sku)
}

func (s *VariantService) GetVariantsByProductID(productID uint64) ([]model.Variant, error) {
	return s.variantRepo.GetByProductID(productID)
}

func (s *VariantService) GetActiveVariantsByProductID(productID uint64) ([]model.Variant, error) {
	return s.variantRepo.GetActiveByProductID(productID)
}

func (s *VariantService) UpdateVariant(variant *model.Variant) error {
	return s.variantRepo.Update(variant)
}

func (s *VariantService) UpdateStock(id uint64, quantity int) error {
	return s.variantRepo.UpdateStock(id, quantity)
}

func (s *VariantService) DeleteVariant(id uint64) error {
	return s.variantRepo.Delete(id)
}

func (s *VariantService) DeactivateVariant(id uint64) error {
	variant, err := s.variantRepo.GetByID(id)
	if err != nil {
		return err
	}
	variant.IsActive = false
	return s.variantRepo.Update(variant)
}

func (s *VariantService) ActivateVariant(id uint64) error {
	variant, err := s.variantRepo.GetByID(id)
	if err != nil {
		return err
	}
	variant.IsActive = true
	return s.variantRepo.Update(variant)
}
