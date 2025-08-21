package model

import (
	"time"
)

type Variant struct {
	ID            uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID     uint64    `json:"productId" gorm:"not null;index"`
	SKU           string    `json:"sku" gorm:"size:100;uniqueIndex;not null"`
	Price         float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	StockQuantity int       `json:"stock_quantity" gorm:"default:0"`
	IsActive      bool      `json:"isActive" gorm:"default:true"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relationships
	Product *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	Media   []Media  `json:"media,omitempty" gorm:"foreignKey:VariantID"`
}
