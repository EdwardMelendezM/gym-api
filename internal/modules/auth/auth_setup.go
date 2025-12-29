package auth

import (
	"gym-api/internal/ent"
	"gym-api/internal/modules/sessions"
	"gym-api/internal/modules/users"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, client *ent.Client) {
	userRepository := users.NewUserEntRepository(client)
	sessionRepository := sessions.NewSessionEntRepository(client)
	authService := NewAuthService(userRepository, sessionRepository)
	authHandler := NewAuthHandler(authService)

	AuthRoutes(router, authHandler)
}
