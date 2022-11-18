package controllers

import (
	"net/http"

	"github.com/MilliGoshant/merci/backend/db"

	"github.com/MilliGoshant/merci/backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// CreateProduct creates a product
func CreateProduct(c *gin.Context) {
	var product models.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := db.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, product)
}

// GetProducts finds products and responses it
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := db.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

// GetProduct collects informations about product and responses it
func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product
	if err := db.DB.Where("id = ?", id).Preload(clause.Associations).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, product)
}

// UpdateProduct updates a product by id
func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	err := db.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := db.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, product)
}

// DeleteProduct finds a product by id, deletes and responses it
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	if err := db.DB.Where("id = ?", id).Delete(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, product)
}
