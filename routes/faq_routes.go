package routes

import (
	"app/config"
	"app/handlers"
	"app/repository"
	"app/services"

	"github.com/gin-gonic/gin"
)

func SetupFaqRoutes(r *gin.Engine) {
	// Get the global database connection
	db := config.DB

	// Create repository, service, and handler instances for FAQ
	faqRepo := repository.NewFaqRepository(db)
	faqService := services.NewFaqService(faqRepo)
	faqHandler := handlers.NewFaqHandler(faqService)

	// Define FAQ-related routes
	faqRoutes := r.Group("/faqs")
	{
		faqRoutes.GET("/", faqHandler.GetAll)        // Get all FAQs
		faqRoutes.POST("/", faqHandler.Create)       // Create a new FAQ
		faqRoutes.GET("/:id", faqHandler.GetByID)    // Get an FAQ by ID
		faqRoutes.PUT("/:id", faqHandler.Update)     // Update an FAQ
		faqRoutes.DELETE("/:id", faqHandler.Delete)  // Delete an FAQ
	}
}