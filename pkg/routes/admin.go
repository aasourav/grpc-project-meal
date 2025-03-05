package routes

import (
	"aas.dev/pkg/handlers"
	"aas.dev/pkg/repository"
	"aas.dev/pkg/services"
	"aas.dev/pkg/utils"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(router *gin.Engine) {
	db := utils.MongoDatabase

	adminRepo := repository.NewAdminRepo(db)
	varificationRepo := repository.NewVerificationRepo(db, true)

	adminService := services.NewAdminService(adminRepo, varificationRepo)
	adminHandler := handlers.NewAdminHandler(adminService)

	adminRoutes := router.Group("/admins")
	{
		adminRoutes.POST("/register", adminHandler.RegisterUser)
		adminRoutes.POST("/login", adminHandler.Login)
		adminRoutes.GET("/verify", adminHandler.VerifyAccount)
	}
}
