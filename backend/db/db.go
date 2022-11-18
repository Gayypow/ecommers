package db

import (
	"fmt"
	"log"
	"os"

	"github.com/MilliGoshant/merci/backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}
	// connStr := "user=dvlt dbname=merci password=123456 host=localhost sslmode=disable"
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DBHost, DBUser, DBPass, DBName, DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Panic(err.Error())
	}

	db.AutoMigrate(&models.Category{}, &models.Product{}, &models.ProductImage{}, &models.User{}, &models.Admin{})
	DB = db
}
