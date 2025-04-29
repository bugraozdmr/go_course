package handlers

import (
	"fiyathunter/sites"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BestPrice(c *gin.Context) {
	keyword := c.Param("keyword")
	keyword = strings.ToLower(keyword)

	response := sites.BestPrice(keyword)

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
