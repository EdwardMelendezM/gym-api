package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lucsky/cuid"
)

const RequestIDKey = "requestID"
const RequestIDHeader = "X-Request-ID"

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := cuid.New()
		c.Set(RequestIDKey, id)
		c.Writer.Header().Set(RequestIDHeader, id)
		c.Next()
	}
}
