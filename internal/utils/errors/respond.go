package errors

import (
	"net/http"

	"gym-api/internal/utils/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Respond(c *gin.Context, err error) {
	if err == nil {
		return
	}

	log := logger.FromContext(c.Request.Context())

	if appErr, ok := err.(*AppError); ok {
		if appErr.Err != nil {
			log.Error(
				appErr.Message,
				zap.String("layer", appErr.Layer),
				zap.String("function", appErr.Function),
				zap.Error(appErr.Err),
			)
		}

		c.JSON(appErr.Status, gin.H{
			"message": appErr.Message,
			"code":    appErr.Code,
		})
		return
	}

	log.Error("unhandled error", zap.Error(err))

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "internal server error",
		"code":    "INTERNAL_SERVER_ERROR",
	})
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
