package errors

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.Status, gin.H{
			"error": appErr.Message,
		})
		return
	}

	// fallback (nunca debería pasar)
	c.JSON(500, gin.H{
		"error": "internal server error",
	})
}
