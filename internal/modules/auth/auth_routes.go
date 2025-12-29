package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.RouterGroup, handler *Handler) {
	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
	r.POST("/refresh-token", handler.RefreshToken)
}
