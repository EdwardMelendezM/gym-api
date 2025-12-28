package main

import (
	"net/http"
	"os"

	"gym-api/internal/database"
	"gym-api/internal/middlware"
	users2 "gym-api/internal/modules/users"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin engine (modern default: logger + recovery)
	r := gin.Default()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	// Database
	client := database.NewEntClient(dsn)

	// middlewares base
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger())

	// versioning
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// dependencies
	userRepo := users2.NewEntRepository(client)
	userService := users2.NewService(userRepo)

	// routes
	users2.RegisterRoutes(v1, userService)

	// Simple test endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Start server
	r.Run(":8080")
}
