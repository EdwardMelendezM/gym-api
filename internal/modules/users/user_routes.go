package users

import (
	"github.com/gin-gonic/gin"

	paginationMiddleware "gym-api/internal/middleware"
)

func UserRoutes(r *gin.RouterGroup, handler *Handler) {
	r.GET("/users", paginationMiddleware.PaginationMiddleware(), handler.List)
	r.POST("/users", handler.Create)
	r.GET("/users/me", handler.GetMe)
	r.GET("/users/:id", handler.GetById)

}
