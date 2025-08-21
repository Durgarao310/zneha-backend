package model

import (
	"time"
)

// JSONB represents a JSONB field for PostgreSQL

type Category struct {
	ID          uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Depth       int       `json:"depth" gorm:"default:0;not null"` // 0 = main category, 1 = subcategory
	ParentID    *uint64   `json:"parentId" gorm:"index"`           // nullable for root categories
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
