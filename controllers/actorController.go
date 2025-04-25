package controllers

import (
	"mi-proyecto/config"
	"mi-proyecto/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActors(c *gin.Context) {
	var actors []models.Actor
	config.DB.Preload("Actors").Find(&actors)
	if len(actors) > 0 {
		c.JSON(http.StatusOK, &actors)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "No exist any actors"})
	}
}
func CreateActor(c *gin.Context) {
	var actor models.Actor
	if err := c.ShouldBindJSON(&actor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&actor)
	c.JSON(http.StatusCreated, actor)
}

func UpdateActor(c *gin.Context) {
	id := c.Param("id")
	var actor models.Actor
	if err := config.DB.First(&actor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actor not found"})
		return
	}
	if err := c.ShouldBindJSON(&actor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&actor)
	c.JSON(http.StatusOK, actor)
}

func DeleteActor(c *gin.Context) {
	id := c.Param("id")
	var actor models.Actor
	if err := config.DB.Delete(&actor, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actor not found"})
		return
	}
	config.DB.Delete(&actor)
	c.JSON(http.StatusOK, gin.H{"message": "Actor deleted"})
}
