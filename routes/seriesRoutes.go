package routes

import (
	"mi-proyecto/controllers"

	"github.com/gin-gonic/gin"
)

func SeriesRoutes(r *gin.Engine) {
	serie := r.Group("/serie")
	{
		serie.GET("/search", controllers.SearchSeriesByName)
	}
}
