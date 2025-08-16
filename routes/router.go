package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes sets up all routes including authentication, user, product, and other routes
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	
	// Setup Authentication Routes under "/api/auth"
	SetupAuthRoutes(r, db)

	// Setup User Routes under "/api/users"
	SetupUserRoutes(r, db)

	// Setup Product Routes under "/api/products"
	SetupProductRoutes(r, db)

	// Setup FAQ Routes under "/api/faqs"
	SetupFaqRoutes(r, db)

}
