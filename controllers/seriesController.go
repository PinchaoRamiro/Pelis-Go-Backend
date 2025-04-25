package controllers

import (
	"fmt"
	"mi-proyecto/config"
	"mi-proyecto/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSeries(c *gin.Context) {
	var series []models.Serie
	config.DB.Preload("Series").Find(&series)
	if len(series) > 0 {
		c.JSON(http.StatusOK, &series)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "No exist any series"})
	}
}

func CreateSerie(c *gin.Context) {
	var serie models.Serie
	if err := c.ShouldBindJSON(&serie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&serie)
	c.JSON(http.StatusCreated, serie)
}

func UpdateSerie(c *gin.Context) {
	id := c.Param("id")
	var serie models.Serie
	if err := config.DB.First(&serie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serie not found"})
		return
	}
	if err := c.ShouldBindJSON(&serie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&serie)
	c.JSON(http.StatusOK, serie)
}

func DeleteSerie(c *gin.Context) {
	id := c.Param("id")
	var serie models.Serie
	if err := config.DB.Delete(&serie, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serie not found"})
		return
	}
	config.DB.Delete(&serie)
	c.JSON(http.StatusOK, gin.H{"message": "Serie deleted"})
}

func SearchSeriesByName(c *gin.Context) {
	tittleserie := c.Query("tittleserie")

	if tittleserie == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tittleserie parameter is required"})
		return
	}

	var series []models.Serie
	if err := config.DB.Where("Name = ?", tittleserie).Find(&series).Error; err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching for series", "series": series})
		return
	}
	if len(series) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No series with this name were found"})
		return
	}
	relatedSeries, _ := GetRelatedByGender(series[0].Genre, series)
	c.JSON(http.StatusOK, gin.H{"series": series, "relatedSeries": relatedSeries})
}

func GetRelatedByGender(gender string, excludeSeries []models.Serie) ([]models.Serie, error) {
	var relatedSeries []models.Serie
	var excludeIDs []uint
	for _, serie := range excludeSeries {
		excludeIDs = append(excludeIDs, serie.ID)
	}

	if err := config.DB.Where("Genre = ? AND ID NOT IN (?)", gender, excludeIDs).Find(&relatedSeries).Error; err != nil {
		return []models.Serie{}, err
	}

	return relatedSeries, nil
}
func ShowBestSeries(c *gin.Context) {
	var bestseries []models.Serie
	config.DB.Order("rating DESC").Limit(10).Find(&bestseries)
	if len(bestseries) > 0 {
		c.JSON(http.StatusOK, &bestseries)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "No exist list of best series"})
	}
}
