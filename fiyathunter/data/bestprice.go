package data

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func BestPrice(c *gin.Context) {
	keyword := c.Param("keyword")
	keyword = strings.ToLower(keyword)

	fmt.Println(keyword)
}
