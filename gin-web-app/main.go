package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("web/template/*")

	// Serve static files
	r.Static("/static", "./web/static")

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "chat.html", gin.H{
			"title": "Real-time Chat Application",
		})
	})

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
