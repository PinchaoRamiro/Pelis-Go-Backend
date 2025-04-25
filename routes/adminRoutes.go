package routes

import (
	"mi-proyecto/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		admin.POST("/movies", controllers.CreateMovie)
		admin.GET("/movies", controllers.GetMovies)
		admin.PUT("/movies/:id", controllers.UpdateMovie)
		admin.DELETE("/movies/:id", controllers.DeleteMovie)
		admin.POST("/series", controllers.CreateSerie)
	}
}
