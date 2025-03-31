package routes

import (
	"rest/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // GET POST DELETE PATCH PUT HEAD
	server.GET("/events/:id", getEvent)

	// grouplayarak seslemek once middleware/ler secilir
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register",cancelRegistration)

	// same package lower case

	// ** ROUTES **
	server.POST("/signup", signup)
	server.POST("/login", login)
}
