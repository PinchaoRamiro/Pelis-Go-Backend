package controllers

import (
	"mi-proyecto/config"
	"mi-proyecto/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangePasswordAdmin(c *gin.Context) {
	var admin models.Admin
	var adminIndatabase models.Admin

	id := c.Param("id")

	if err := config.DB.First(&adminIndatabase, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found the Admin"})
		return
	}

	if err := c.ShouldBindBodyWithJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The Structure of the Admin doest is correct"})
		return
	}

	admin.HashPassword()

	config.DB.Save(&admin)
	c.JSON(http.StatusOK, gin.H{"msg": "Updated"})
}

func CreateAdmin(c *gin.Context) {
	var newAdmin models.Admin

	if err := c.ShouldBindJSON(&newAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := newAdmin.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := config.DB.Create(&newAdmin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Admin created"})
}

func GetUsers(c *gin.Context) {
	var users []models.User

	config.DB.Preload("Users").Find(&users)

	if len(users) > 0 {
		c.JSON(http.StatusOK, &users)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "No exist any users"})
	}
}
