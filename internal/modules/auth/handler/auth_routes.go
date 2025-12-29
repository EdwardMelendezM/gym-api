package handler

import (
	"gym-api/internal/modules/auth/handler/impl"

	"github.com/gin-gonic/gin"
)

// AuthRoutes registers auth endpoints
// @Summary Auth endpoints
// @Description All auth related endpoints
// @Tags Auth
func AuthRoutes(r *gin.RouterGroup, handler impl.Handler) {
	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
	r.POST("/refresh-token", handler.RefreshToken)
}
