package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(c *gin.Context, err error, statusCode int, message string) {
	if err != nil {
		// log.Println(err)

		c.JSON(statusCode, gin.H{
			"error":   message,
			"details": err.Error(),
		})
	}
}

func HandleDBError(c *gin.Context, err error) {
	HandleError(c, err, http.StatusInternalServerError, "Database error occurred")
}

// HandleValidationError, doğrulama hataları için özel hata yöneticisi
func HandleValidationError(c *gin.Context, err error) {
	HandleError(c, err, http.StatusBadRequest, "Validation error occurred")
}

// HandleNotFoundError, bulunamadığı durumlar için özel hata yöneticisi
func HandleNotFoundError(c *gin.Context, err error) {
	HandleError(c, err, http.StatusNotFound, "Resource not found")
}

// HandleUnauthorizedError, yetkisiz erişim hataları için özel yöneticisi
func HandleUnauthorizedError(c *gin.Context, err error) {
	HandleError(c, err, http.StatusUnauthorized, "Unauthorized access")
}

// HandleBadRequestError, kötü istek hataları için özel yöneticisi
func HandleBadRequestError(c *gin.Context, err error) {
	HandleError(c, err, http.StatusBadRequest, "Bad request")
}
