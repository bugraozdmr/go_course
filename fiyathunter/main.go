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

	// choose a port which can be opened 8888 / 80 / 8080 etc.
	server.Run(":8888")
}
