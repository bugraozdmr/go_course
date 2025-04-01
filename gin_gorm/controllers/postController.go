package controllers

import (
	"gin_gorm/helpers"
	"gin_gorm/initializers"
	"gin_gorm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostsCreate(context *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	if err := context.ShouldBindJSON(&body); err != nil {
		helpers.HandleValidationError(context, err) // Validation hatası
		return
	}

	// Create a post
	// post := models.Post{Title: "AI is crazy mate", Body: "AI is going to make our jobs automated"}
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		context.Status(400)
		return
	}

	// return it

	context.JSON(201, gin.H{
		"post": &post,
	})
}

func PostsIndex(c *gin.Context) {
	// Query params: page, limit, query
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	squery := c.DefaultQuery("query", "")

	// Convert to integers
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		helpers.HandleBadRequestError(c, err)
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		helpers.HandleBadRequestError(c, err)
		return
	}

	// Offset calculate
	offset := (pageInt - 1) * limitInt

	// get count
	var postCount int64
	query := initializers.DB.Model(&models.Post{})

	//
	if squery != "" {
		query = query.Where("title LIKE ?", "%"+squery+"%")
	}

	// get filtered post count
	if err := query.Count(&postCount).Error; err != nil {
		helpers.HandleDBError(c, err)
		return
	}

	// get posts
	var posts []models.Post
	if err := query.Offset(offset).Limit(limitInt).Find(&posts).Error; err != nil {
		helpers.HandleDBError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post_count": postCount,
		"posts":      posts,
		"page":       pageInt,
		"limit":      limitInt,
	})
}

// rest of the CRUD methods are homework for you !!!
