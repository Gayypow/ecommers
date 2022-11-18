package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Name        Lang           `json:"name" gorm:"embedded;embeddedPrefix:name_"`
	Description Lang           `json:"description" gorm:"embedded;embeddedPrefix:description_"`
	Image       string         `json:"image"`
}
