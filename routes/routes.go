package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes and sets up the routes for the API
func SetupRoutes(router *gin.Engine) {
	// Set up a group of routes under the /api path
	api := router.Group("/api")

	// Define a route for the /api/hello endpoint
	api.GET("/hello", handlers.HelloWorldHandler)
}
