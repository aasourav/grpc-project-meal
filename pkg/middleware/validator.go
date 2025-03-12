package middleware

import (
	"fmt"
	"net/http"

	"aas.dev/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RequestValidatorMiddleware(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(obj); err != nil {
			utils.ErrorJSON(c, err, http.StatusBadRequest)
			c.Abort()
			return
		}

		validate := validator.New()
		if err := validate.Struct(obj); err != nil {
			fmt.Println("Error validating request: ", err)
			utils.ErrorJSON(c, err, http.StatusBadRequest)
			c.Abort()
			return
		}

		c.Set("req", obj) // Store validated data in context
		c.Next()
	}
}
