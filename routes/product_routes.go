package routes

import (
	"app/config"
	"app/handlers"
	"app/repository"
	"app/services"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(r *gin.Engine) {
	db := config.DB

	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	productRoutes := r.Group("/products")
	{
		productRoutes.GET("/", productHandler.GetAll)
		productRoutes.POST("/", productHandler.Create)
		productRoutes.GET("/:id", productHandler.GetByID)
		productRoutes.PUT("/:id", productHandler.Update)
		productRoutes.DELETE("/:id", productHandler.Delete)
	}
}