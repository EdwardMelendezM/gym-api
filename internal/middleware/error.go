package middleware

import (
	"net/http"

	"gym-api/internal/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		last := c.Errors.Last()
		if last == nil {
			return
		}

		if appErr, ok := last.Err.(*errors.AppError); ok {
			c.JSON(appErr.Status, gin.H{
				"message":  appErr.Message,
				"layer":    appErr.Layer,
				"function": appErr.Function,
			})
			return
		}

		// fallback
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}
}
