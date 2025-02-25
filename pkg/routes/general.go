package routes

import (
	"aas.dev/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func SetupMainRoutes(router *gin.Engine) {
	generalHandler := handlers.NewGeneralHandler()

	router.GET("/", generalHandler.AboutUs)
	router.GET("/health", generalHandler.HealthCheck)
	router.POST("/email-verify", generalHandler.EmailVerify)
}
