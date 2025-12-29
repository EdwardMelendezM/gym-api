package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()

		// log simple (luego se puede cambiar por zap / zerolog)
		println(method, path, status, latency.String())
	}
}
