package users

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.RouterGroup, handler *Handler) {
	r.GET("/users", handler.List)
	r.POST("/users", handler.Create)
}
