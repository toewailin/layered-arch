package routes

import (
	"app/handlers"
	"app/repository"
	"app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupAuthRoutes sets up authentication-related routes (login, register)
func SetupAuthRoutes(r *gin.Engine, db *gorm.DB) {
	// Create repository, service, and handler instances for authentication
	authRepo := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Define routes for authentication
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}
}
