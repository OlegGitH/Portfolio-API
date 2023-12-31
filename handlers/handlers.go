package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorldHandler handles the /hello endpoint
func HelloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
