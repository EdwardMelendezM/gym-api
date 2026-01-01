package main

import (
	_ "gym-api/docs"
	"gym-api/internal/modules/auth"
	"gym-api/internal/modules/users"
	"gym-api/internal/utils/config"
	"net/http"

	"gym-api/internal/utils/database"
	middleware2 "gym-api/internal/utils/middleware"

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

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Create Gin engine (modern default: logger + recovery)
	r := gin.Default()

	// Load configuration and initialize database client
	cfg := config.LoadConfig()
	client := database.NewEntClient(cfg.DatabaseURL)

	// middlewares base
	r.Use(middleware2.CorsMiddleware(cfg.AllowOrigins))
	r.Use(gin.Recovery())
	r.Use(middleware2.RequestLogger())
	r.Use(middleware2.ErrorMiddleware())
	r.Use(middleware2.RequestIDMiddleware())

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
