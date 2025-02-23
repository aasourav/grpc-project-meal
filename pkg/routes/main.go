package routes

import (
	"aas.dev/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func SetupMainRoutes(router *gin.Engine) {
	mainHandler := handlers.NewMainHandler()

	router.GET("/health", mainHandler.HealthCheck)
	router.GET("/", mainHandler.AboutUs)
}
