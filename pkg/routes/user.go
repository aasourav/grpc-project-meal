package routes

import (
	"aas.dev/pkg/handlers"
	"aas.dev/pkg/repository"
	"aas.dev/pkg/services"
	"aas.dev/pkg/utils"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	db := utils.MongoDatabase
	// user
	userRepo := repository.NewUserRepo(db)
	varificationRepo := repository.NewVerificationRepo(db, true)
	userService := services.NewUserService(userRepo, varificationRepo)
	userHandler := handlers.NewUserHandler(userService)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", userHandler.RegisterUser)
		userRoutes.POST("/login", userHandler.RegisterUser)
		// userRoutes.POST("/change-password", pendingUserRepoHandler.RegisterUser)
	}
}
