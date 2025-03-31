package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // GET POST DELETE PATCH PUT HEAD
	server.GET("/events/:id", getEvent)

	// same package lower case
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id",deleteEvent)
}
