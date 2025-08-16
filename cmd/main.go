package main

import (
	"app/config"
	"app/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	r := gin.Default()

	// Load configuration and MySQL connection
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load config:", err)
	}

	// Set up routes
	routes.SetupRoutes(r, config.DB)

	// Start the server
	err = r.Run(":8081")
	if err != nil {
		log.Fatal("Could not start server:", err)
	}
}
