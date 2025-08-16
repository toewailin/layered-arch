package routes

import (
	"app/handlers"
	"app/repository"
	"app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupUserRoutes sets up all user-related routes with dependency injection
func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	// Create repository, service, and handler instances
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Define routes
	userRoutes := r.Group("/api/users")
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
