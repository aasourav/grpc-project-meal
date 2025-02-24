package routes

import (
	"aas.dev/pkg/handlers"
	"aas.dev/pkg/repository"
	"aas.dev/pkg/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAdminRoutes(router *gin.Engine, db *mongo.Database) {
	adminRepo := repository.NewAdminRepo(db)
	adminService := services.NewAdminService(adminRepo)
	adminHandler := handlers.NewAdminHandler(adminService)

	adminRoutes := router.Group("/admins")
	{
		adminRoutes.POST("/register", adminHandler.RegisterUser)
		adminRoutes.POST("/login", adminHandler.Login, adminHandler.Login)
	}
}
