package main

import (
	"gym-api/internal/modules/shared/database"
	middleware2 "gym-api/internal/modules/shared/middleware"
	"net/http"
	"os"

	_ "gym-api/docs"
	"gym-api/internal/modules/auth"
	"gym-api/internal/modules/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gym API
// @version 1.0
// @description API de ejemplo para gestión de usuarios en un gym.
// @termsOfService http://gym-api.local/terms/

// @contact.name Edward
// @contact.url http://gym-api.local
// @contact.email edward@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://mi-frontend.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
	r.Use(gin.Recovery())
	r.Use(middleware2.RequestLogger())
	r.Use(middleware2.ErrorHandler())
	//r.Use(middleware.AuthMiddleware())

	/// versioning
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// Public
	authGroup := v1.Group("/auth")
	auth.SetupRoutes(authGroup, client)

	// Protected
	protected := v1.Group("/")
	protected.Use(middleware2.SetupAuthMiddleware(client))

	users.SetupRoutes(protected, client)

	// Simple test endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	r.Run(":8080")
}
