package routes

import (
	"aas.dev/pkg/handlers"
	middleware "aas.dev/pkg/middleware"
	"aas.dev/pkg/models/types"
	models "aas.dev/pkg/models/user"
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
		userRoutes.POST("/register", middleware.RequestValidatorMiddleware(&models.User{}), userHandler.RegisterUser)
		userRoutes.POST("/login", middleware.RequestValidatorMiddleware(&models.UserLogin{}), userHandler.Login)
		userRoutes.GET("/verify", userHandler.VerifyAccount)
		userRoutes.POST("/password-reset", middleware.RequestValidatorMiddleware(&types.ResetPassword{}), userHandler.PassowordChange)
		userRoutes.POST("/update-meal-plan", middleware.RequestValidatorMiddleware(&types.UpdateWeeklyMealPlan{}), middleware.UserValidatorMiddleware(), userHandler.MealPlanUpdate)
	}
}
