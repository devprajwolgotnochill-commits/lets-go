package main

import (
	"net/http"

	"github.com/devprajwalgotnochill/gin-frmwork/controller"
	"github.com/devprajwalgotnochill/gin-frmwork/entity"
	"github.com/devprajwalgotnochill/gin-frmwork/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize service and controller
	videoService := services.New()
	videoController := controller.New(videoService)

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "hello, world",
		})
	})

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Video API endpoints
	r.GET("/videos", func(c *gin.Context) {
		videos := videoController.FindAll()
		c.JSON(http.StatusOK, videos)
	})

	r.POST("/videos", func(c *gin.Context) {
		var video entity.Videos
		if err := c.ShouldBindJSON(&video); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		savedVideo := videoController.Save(video)
		c.JSON(http.StatusCreated, savedVideo)
	})

	// Start server on port 8000
	// Server will listen on 0.0.0.0:8000 (localhost:8000 on Windows)
	r.Run(":8000")
}
