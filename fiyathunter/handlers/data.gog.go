package handlers

import (
	"net/http"
	"strings"
	gog "fiyathunter/sites/gog"
	
	"github.com/gin-gonic/gin"
)

func FetchDataGog(c *gin.Context) {
	keyword := c.Param("keyword")
	//! No replacing
	keyword = strings.ToLower(keyword)

	response := gog.FetchData(keyword)

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
