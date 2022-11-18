package controllers

import (
	"net/http"

	"github.com/MilliGoshant/merci/backend/db"
	"github.com/MilliGoshant/merci/backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// Create a category
func CreateCategory(c *gin.Context) {
	var category models.Category

	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, category)
}

// GetCategories finds all categories and responses it
func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := db.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

// GetCategory collects informations about category and responses it
func GetCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	var category models.Category
	if err := db.DB.Where("id = ?", id).Preload(clause.Associations).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, category)
}

// UpdateCategory updates a category by id
func UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")
	err := db.DB.Where("id = ?", id).First(&category).Error
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := db.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, category)
}

// DeleteCategory finds a category by id, deletes and responses it
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := db.DB.Where("id = ?", id).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	if err := db.DB.Where("id = ?", id).Delete(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, category)
}
