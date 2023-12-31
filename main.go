package main

import (
	"backend/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Start the server on port 8080
	err := router.Run("localhost:8000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
