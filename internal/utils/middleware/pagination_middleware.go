package middleware

import (
	"gym-api/internal/utils/pagination"

	"github.com/gin-gonic/gin"
)

func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var p pagination.Params
		_ = c.ShouldBindQuery(&p)
		p.Normalize()

		c.Set(pagination.Key, p)
		c.Next()
	}
}
