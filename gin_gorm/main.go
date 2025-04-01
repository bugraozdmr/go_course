package main

import (
	"gin_gorm/controllers"
	"gin_gorm/initializers"

	"github.com/gin-gonic/gin"
)

// initialize
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)

	r.Run()
}
