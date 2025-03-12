package routes

import (
	"aas.dev/pkg/handlers"
	middleware "aas.dev/pkg/middleware"
	models "aas.dev/pkg/models/admin"
	"aas.dev/pkg/models/types"
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
		adminRoutes.POST("/register", middleware.RequestValidatorMiddleware(&models.Admin{}), adminHandler.RegisterUser)
		adminRoutes.POST("/login", middleware.RequestValidatorMiddleware(&models.AdminLogin{}), adminHandler.Login)
		adminRoutes.POST("/password-reset", middleware.RequestValidatorMiddleware(&types.ResetPassword{}), adminHandler.PassowordChange)
		adminRoutes.GET("/verify", adminHandler.VerifyAccount)
	}
}
