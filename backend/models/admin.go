package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name      string         `json:"name" binding:"required"`
	Email     string         `json:"email" binding:"required" gorm:"unique"`
	Password  string         `json:"password" binding:"required"`
}

// HashPassword hashes admin's password and replaces it with given password
func (admin *Admin) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	admin.Password = string(bytes)
	return nil
}

// CheckPassword hashes the given password and compares it with admin's password, responses error if it isn't equal
func (admin *Admin) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
