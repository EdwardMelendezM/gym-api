package main

import (
	middleware "gym-api/middlware"
	"gym-api/modules/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin engine (modern default: logger + recovery)
	r := gin.Default()

	// middlewares base
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger())

	// versioning
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// dependencies
	repo := users.NewRepository()
	service := users.NewService(repo)

	// routes
	users.RegisterRoutes(v1, service)

	// Simple test endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Start server
	r.Run(":8080")
}
