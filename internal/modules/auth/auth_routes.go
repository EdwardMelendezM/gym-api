package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.RouterGroup, handler *Handler) {
	r.GET("/login", handler.Login)
	r.POST("/register", handler.Register)
}
