package service

import (
	"github.com/Durgarao310/zneha-backend/internal/model"
	"github.com/Durgarao310/zneha-backend/internal/repository"
)

type MediaService struct {
	mediaRepo repository.MediaRepository
}

func NewMediaService(mediaRepo repository.MediaRepository) *MediaService {
	return &MediaService{
		mediaRepo: mediaRepo,
	}
}

func (s *MediaService) CreateMedia(media *model.Media) error {
	return s.mediaRepo.Create(media)
}

func (s *MediaService) GetMediaByID(id uint64) (*model.Media, error) {
	return s.mediaRepo.GetByID(id)
}

func (s *MediaService) GetMediaByProductID(productID uint64) ([]model.Media, error) {
	return s.mediaRepo.GetByProductID(productID)
}

func (s *MediaService) GetMediaByVariantID(variantID uint64) ([]model.Media, error) {
	return s.mediaRepo.GetByVariantID(variantID)
}

func (s *MediaService) UpdateMedia(media *model.Media) error {
	return s.mediaRepo.Update(media)
}

func (s *MediaService) DeleteMedia(id uint64) error {
	return s.mediaRepo.Delete(id)
}

func (s *MediaService) SetPrimaryMedia(productID, mediaID uint64) error {
	return s.mediaRepo.SetPrimary(productID, mediaID)
}

func (s *MediaService) GetPrimaryMedia(productID uint64) (*model.Media, error) {
	return s.mediaRepo.GetPrimaryByProductID(productID)
}
