package middleware

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	models "aas.dev/pkg/models/user"
	"aas.dev/pkg/utils"
	"github.com/gin-gonic/gin"
)

func UserValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken, err := c.Cookie("user-token")
		if err != nil || userToken == "" {
			log.Println("user cookies read failed: ", err.Error())
			utils.ErrorJSON(c, errors.New("req forbidden"), http.StatusForbidden)
			c.Abort()
			return
		}

		parsedToken, err := utils.VerifyJWT(userToken, "user")
		if err != nil {
			log.Println("verify token failed: ", err.Error())
			utils.ErrorJSON(c, errors.New("req forbidden"), http.StatusForbidden)
			c.Abort()
			return
		}

		var userData models.User
		err = json.Unmarshal(parsedToken, &userData)
		if err != nil {
			log.Println("type cast error in user validator middleware: ", err.Error())
			utils.ErrorJSON(c, errors.New("req forbidden"), http.StatusForbidden)
			c.Abort()
			return
		}

		c.Set("userData", userData)
		c.Next()
	}
}
