package auth

import (
	"gym-api/internal/ent"
	"gym-api/internal/modules/auth/handler"
	"gym-api/internal/modules/auth/handler/impl"
	impl2 "gym-api/internal/modules/auth/service/impl"
	impl3 "gym-api/internal/modules/sessions/repository/impl"
	"gym-api/internal/modules/users"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, client *ent.Client) {
	userRepository := users.NewUserEntRepository(client)
	sessionRepository := impl3.NewSessionEntRepository(client)
	authService := impl2.NewAuthService(userRepository, sessionRepository)
	authHandler := impl.NewAuthHandler(authService)

	handler.AuthRoutes(router, authHandler)
}
