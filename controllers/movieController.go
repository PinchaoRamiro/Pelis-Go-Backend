package controllers

import (
	"mi-proyecto/config"
	"mi-proyecto/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	var movies []models.Movie
	config.DB.Preload("Movies").Find(&movies)
	if len(movies) > 0 {
		c.JSON(http.StatusOK, &movies)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "No exist any movies"})
	}
}

func CreateMovie(c *gin.Context) {
	var movie models.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&movie)
	c.JSON(http.StatusCreated, movie)
}

func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie

	if err := config.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&movie)
	c.JSON(http.StatusOK, movie)
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie

	if err := config.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Film not found"})
		return
	}

	config.DB.Delete(&movie)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted movie"})
}

func SearchMovieByName(c *gin.Context) {
	tittlemovie := c.Query("tittlemovie")

	if tittlemovie == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tittlemovie parameter is required"})
		return
	}

	var movies []models.Movie
	if err := config.DB.Where("Name =?", tittlemovie).Find(&movies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching for movies", "movies": movies})
		return
	}
	if len(movies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No movie with this name were found"})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func ShowBestMovies(c *gin.Context) {
	var bestmovies []models.Movie
	config.DB.Order("rating DESC").Limit(10).Find(&bestmovies)
	if len(bestmovies) > 0 {
		c.JSON(http.StatusOK, &bestmovies)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "No exist list of best movies"})
	}
}
