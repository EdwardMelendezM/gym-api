package users

import (
	paginationMiddleware "gym-api/internal/modules/shared/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, handler *Handler) {
	r.GET("/users", paginationMiddleware.PaginationMiddleware(), handler.List)
	r.POST("/users", handler.Create)
	r.GET("/users/me", handler.GetMe)
	r.GET("/users/:id", handler.GetById)

}
