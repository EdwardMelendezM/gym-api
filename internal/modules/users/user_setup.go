package users

import (
	"github.com/gin-gonic/gin"

	"gym-api/internal/ent"
	"gym-api/internal/modules/users/handler"
	"gym-api/internal/modules/users/handler/impl"
	impl3 "gym-api/internal/modules/users/repository/impl"
	impl2 "gym-api/internal/modules/users/service/impl"
)

func SetupRoutes(router *gin.RouterGroup, client *ent.Client) {
	userRepository := impl3.NewUserEntRepository(client)
	userService := impl2.NewUserService(userRepository)
	userHandler := impl.NewUserHandler(userService)

	handler.UserRoutes(router, userHandler)
}
