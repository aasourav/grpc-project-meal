package middleware

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	models "aas.dev/pkg/models/admin"
	"aas.dev/pkg/models/types"
	"aas.dev/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AdminValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminToken, err := c.Cookie("admin-token")
		if err != nil || adminToken == "" {
			log.Println("admin cookies read failed: ", err.Error())
			utils.ErrorJSON(c, errors.New("req forbidden"), http.StatusForbidden)
			c.Abort()
			return
		}

		parsedToken, err := utils.VerifyJWT(adminToken, "admin")
		if err != nil {
			log.Println("verify token failed: ", err.Error())
			utils.ErrorJSON(c, errors.New("req forbidden"), http.StatusForbidden)
			c.Abort()
			return
		}

		var adminData models.Admin
		err = json.Unmarshal(parsedToken, &adminData)
		if err != nil {
			log.Println("type cast error in admin validator middleware: ", err.Error())
			utils.ErrorJSON(c, errors.New("req forbidden"), http.StatusForbidden)
			c.Abort()
			return
		}

		if adminData.Role != types.ADMIN && adminData.Role != types.SUPERADMIN {
			utils.ErrorJSON(c, errors.New("req forbidden"), http.StatusForbidden)
			c.Abort()
			return
		}
		c.Set("adminData", adminData)
		c.Next()
	}
}

func SuperAdminValidatorMiddleware(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminToken, err := c.Cookie("admin-token")
		if err != nil || adminToken == "" {
			utils.ErrorJSON(c, errors.New("req forbidden"), http.StatusForbidden)
			return
		}

		parsedToken, err := utils.VerifyJWT(adminToken, "admin-token")
		if err != nil {
			utils.ErrorJSON(c, errors.New("req forbidden"), http.StatusForbidden)
			return
		}

		var adminData models.Admin
		err = json.Unmarshal(parsedToken, &adminData)
		if err != nil {
			log.Println("type cast error in admin validator middleware: ", err.Error())
			utils.ErrorJSON(c, errors.New("req forbidden"), http.StatusForbidden)
			c.Abort()
			return
		}

		if adminData.Role != types.SUPERADMIN {
			utils.ErrorJSON(c, errors.New("req forbidden"), http.StatusForbidden)
			return
		}
		c.Set("adminData", adminData)
		c.Next()
	}
}
