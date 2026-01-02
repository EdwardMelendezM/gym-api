package handler

import (
	"github.com/gin-gonic/gin"

	"gym-api/internal/modules/users/handler/impl"
	paginationMiddleware "gym-api/internal/utils/middleware"
)

func UserRoutes(r *gin.RouterGroup, handler impl.Handler) {
	r.GET("/users", paginationMiddleware.PaginationMiddleware(), handler.GetUsersPaginated)
	r.POST("/users", handler.CreateUser)
	r.GET("/users/me", handler.GetMe)
	r.GET("/users/:id", handler.GetUserById)
}
