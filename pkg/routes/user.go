package routes

import (
	"aas.dev/pkg/handlers"
	"aas.dev/pkg/repository"
	"aas.dev/pkg/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupUserRoutes(router *gin.Engine, db *mongo.Database) {
	// user
	userRepo := repository.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// pending user
	pendingUserRepo := repository.NewUserRepo(db)
	pendingUserRepoService := services.NewUserService(pendingUserRepo)
	pendingUserRepoHandler := handlers.NewUserHandler(pendingUserRepoService)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", pendingUserRepoHandler.RegisterUser)
		userRoutes.POST("/login", userHandler.RegisterUser)
		// userRoutes.POST("/change-password", pendingUserRepoHandler.RegisterUser)
	}
}
