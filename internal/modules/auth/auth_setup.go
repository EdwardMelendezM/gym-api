package auth

import (
	"github.com/gin-gonic/gin"

	"gym-api/internal/ent"
	"gym-api/internal/modules/auth/handler"
	"gym-api/internal/modules/auth/handler/impl"
	impl2 "gym-api/internal/modules/auth/service/impl"
	impl3 "gym-api/internal/modules/sessions/repository/impl"
	impl4 "gym-api/internal/modules/users/repository/impl"
)

func SetupRoutes(router *gin.RouterGroup, client *ent.Client) {
	userRepository := impl4.NewUserEntRepository(client)
	sessionRepository := impl3.NewSessionEntRepository(client)
	authService := impl2.NewAuthService(userRepository, sessionRepository)
	authHandler := impl.NewAuthHandler(authService)

	handler.AuthRoutes(router, authHandler)
}
