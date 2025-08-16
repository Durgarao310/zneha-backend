package model

import "time"

type Product struct {
	ID               uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name             string    `json:"name" gorm:"size:255;not null"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"shortDescription"`
	Status           string    `json:"status" gorm:"default:'active'"` // active, inactive
	CreatedAt        time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
