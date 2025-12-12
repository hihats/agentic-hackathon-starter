package main

import (
	"app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Shuffle Lunch API"})
	})

	// API routes
	shuffleHandler := handlers.NewShuffleHandler()
	api := r.Group("/api")
	{
		api.POST("/shuffle", shuffleHandler.Shuffle)
	}

	r.Run(":8080")
}
