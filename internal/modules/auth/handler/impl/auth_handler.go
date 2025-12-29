package impl

import (
	"gym-api/internal/modules/auth/service/impl"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
}

type AuthHandler struct {
	service impl.Service
}

func NewAuthHandler(service impl.Service) Handler {
	return &AuthHandler{service: service}
}
