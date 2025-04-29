package routes

import (
	"fiyathunter/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// base route
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Waiting for your request",
			"hint":    "usage : /game/game_name",
		})
	})

	// epic games api
	server.GET("/epic/game/:keyword", data.FetchDataEpic)

	// gog games api
	server.GET("/gog/game/:keyword", data.FetchDataGog)


	
	// find best price
	server.GET("/bestprice/:keyword", data.BestPrice)
}
