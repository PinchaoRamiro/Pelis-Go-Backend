package main

import (
	"mi-proyecto/config"

	"mi-proyecto/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	server := gin.Default()
	routes.AdminRoutes(server)
	routes.SeriesRoutes(server)

	server.Run(":8080")
}
