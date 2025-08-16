package routes

import (
	"app/config"
	"app/handlers"
	"app/repository"
	"app/services"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes sets up all user-related routes
func SetupUserRoutes(r *gin.Engine) {
	// Assuming the database connection is available
	db := config.DB

	// Create repository, service, and handler instances
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Define routes
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
