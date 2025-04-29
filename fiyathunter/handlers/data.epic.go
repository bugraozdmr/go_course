package handlers

import (
	epic "fiyathunter/sites/epic"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func FetchDataEpic(c *gin.Context) {
	keyword := c.Param("keyword")
	keyword = strings.ToLower(strings.ReplaceAll(keyword, " ", "+"))

	response := epic.FetchData(keyword)

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
