package utils

import (
	"github.com/gin-gonic/gin"
)

func ErrorJSON(c *gin.Context, err error, statusCode int) {
	error := gin.H{
		"status":       statusCode,
		"errorMessage": err.Error(),
		"message":      nil,
		"body":         nil,
	}
	c.JSON(statusCode, error)
}

func SuccessJSON(c *gin.Context, message string, statusCode int, body any) {
	success := gin.H{
		"status":       statusCode,
		"errorMessage": nil,
		"message":      message,
		"body":         body,
	}
	c.JSON(statusCode, success)
}
