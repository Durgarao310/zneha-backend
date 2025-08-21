package model

import "time"

type Media struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	VariantID *uint64   `json:"variantId" gorm:"index"`            // nullable, for product variant specific media
	ProductID uint64    `json:"productId" gorm:"not null;index"`   // required, media belongs to a product
	MediaType string    `json:"mediaType" gorm:"size:50;not null"` // image, video, etc.
	URL       string    `json:"url" gorm:"size:500;not null"`      // media file URL
	Alt       string    `json:"alt" gorm:"size:255"`               // alt text for accessibility
	Position  int       `json:"position" gorm:"default:0"`         // order of media in product gallery
	IsPrimary bool      `json:"isPrimary" gorm:"default:false"`    // main product image
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relationships
	Product *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}
