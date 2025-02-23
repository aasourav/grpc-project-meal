package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorJSON(c *gin.Context, err error, statusCode int) {
	errorMessage := gin.H{
		"status":       statusCode,
		"errorMessage": err.Error(),
		"message":      nil,
		"body":         nil,
	}
	c.JSON(http.StatusBadRequest, errorMessage)
}
