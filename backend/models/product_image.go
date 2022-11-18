package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	ProductID uint           `json:"product_id"`
	URL       string         `json:"url"`
}
