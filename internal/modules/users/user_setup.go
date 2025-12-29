package users

import (
	"gym-api/internal/ent"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, client *ent.Client) {
	userRepository := NewUserEntRepository(client)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	UserRoutes(router, userHandler)
}
