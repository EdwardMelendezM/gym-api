package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, err error) {
	if err == nil {
		return
	}

	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.Status, gin.H{
			"message":  appErr.Message,
			"layer":    appErr.Layer,
			"function": appErr.Function,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "internal server error",
	})
}
