package impl

import (
	"gym-api/internal/modules/users/service/impl"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetUsersPaginated(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	GetMe(ctx *gin.Context)
}

type UserHandler struct {
	service impl.Service
}

func NewUserHandler(service impl.Service) Handler {
	return &UserHandler{service: service}
}
