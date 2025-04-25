package routes

import (
	"fiyathunter/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// base route
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Waiting for your request"})
	})

	server.GET("/game/:keyword", data.FetchData)
}
