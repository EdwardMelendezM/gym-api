package middleware

import (
	"gym-api/internal/utils/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID, _ := c.Get(RequestIDKey)

		ctx := logger.WithContext(
			c.Request.Context(),
			zap.String("request_id", reqID.(string)),
			zap.String("path", c.FullPath()),
			zap.String("method", c.Request.Method),
		)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
