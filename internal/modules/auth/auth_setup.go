package auth

import (
	"gym-api/internal/ent"
	"gym-api/internal/modules/auth/handler"
	"gym-api/internal/modules/auth/handler/impl"
	impl2 "gym-api/internal/modules/auth/service/impl"
	"gym-api/internal/modules/sessions"
	"gym-api/internal/modules/users"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, client *ent.Client) {
	userRepository := users.NewUserEntRepository(client)
	sessionRepository := sessions.NewSessionEntRepository(client)
	authService := impl2.NewAuthService(userRepository, sessionRepository)
	authHandler := impl.NewAuthHandler(authService)

	handler.AuthRoutes(router, authHandler)
}
