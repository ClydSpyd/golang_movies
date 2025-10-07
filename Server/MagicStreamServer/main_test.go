package main_test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main_test() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping/*name", func(c *gin.Context) {
		name := c.Param("name")
		fullPath := c.FullPath()

		// Return JSON response
		response := gin.H{
			"message":  "pong",
			"bing":     "bong",
			"fullPath": fullPath,
		}
		if name != "/" {
			response["name"] = name
		}
		c.JSON(http.StatusOK, response)
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
