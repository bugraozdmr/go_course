package main

import (
	"fiyathunter/middleware"
	"fiyathunter/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Use(middleware.CORSMiddleware())

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
