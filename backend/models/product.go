package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
	Name        Lang           `json:"name" gorm:"embedded;embeddedPrefix:name_"`
	Description Lang           `json:"description" gorm:"embedded;embeddedPrefix:description_"`
	Price       int            `json:"price"`
	Discount    int            `json:"discount"`
	CategoryID  uint           `json:"category_id"`
	Category    *Category      `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	Images      []ProductImage `json:"images"`
}
