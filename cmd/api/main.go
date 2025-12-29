package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"gym-api/internal/database"
	"gym-api/internal/middleware"
	"gym-api/internal/modules/auth"
	"gym-api/internal/modules/users"
)

func main() {
	// Create Gin engine (modern default: logger + recovery)
	r := gin.Default()

	// Database connection
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL environment variable is not set")
	}
	//secret := os.Getenv("SECRET_KEY")
	//secret := "supersecretkey"
	// Database client
	client := database.NewEntClient(dsn)

	// middlewares base
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger())
	//r.Use(middleware.AuthMiddleware())

	/// versioning
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// Public
	authGroup := v1.Group("/auth")
	auth.SetupRoutes(authGroup, client)

	// Protected
	protected := v1.Group("/")
	protected.Use(middleware.SetupAuthMiddleware(client))

	users.SetupRoutes(protected, client)

	// Simple test endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Start server
	r.Run(":8080")
}
