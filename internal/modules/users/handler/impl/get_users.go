package impl

import (
	"gym-api/internal/utils/errors"
	"gym-api/internal/utils/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) List(c *gin.Context) {
	ctx := c.Request.Context()
	p := c.MustGet(pagination.Key).(pagination.Params)

	result, err := h.service.ListUsers(ctx, p)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
