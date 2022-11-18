package controllers

import (
	"net/http"

	"github.com/MilliGoshant/merci/backend/db"
	"github.com/MilliGoshant/merci/backend/models"
	"gorm.io/gorm/clause"

	"github.com/gin-gonic/gin"
)

// CreateAdmin creates an admin
func CreateAdmin(c *gin.Context) {
	var admin models.Admin

	err := c.BindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := db.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, admin)
}

// GetAdmins finds all admins and responses it
func GetAdmins(c *gin.Context) {
	var admins []models.Admin
	if err := db.DB.Find(&admins).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"admins": admins,
	})
}

// GetAdmin collects informations about admin and responses it
func GetAdmin(c *gin.Context) {
	id := c.Params.ByName("id")
	var admin models.Admin
	if err := db.DB.Where("id = ?", id).Preload(clause.Associations).First(&admin).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, admin)
}

// UpdateAdmin updates an admin by id
func UpdateAdmin(c *gin.Context) {
	var admin models.Admin
	id := c.Param("id")
	err := db.DB.Where("id = ?", id).First(&admin).Error
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := db.DB.Save(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, admin)
}

// DeleteAdmin finds an admin by id, deletes and responses it
func DeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	var admin models.Admin
	if err := db.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	if err := db.DB.Where("id = ?", id).Delete(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, admin)
}
