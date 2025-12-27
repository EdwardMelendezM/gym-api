package users

import (
	"net/http"

	"gym-api/errors"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, service Service) {
	r.GET("/users", func(c *gin.Context) {
		users, err := service.ListUsers()
		if err != nil {
			errors.Respond(c, err)
			return
		}

		response := make([]UserResponse, 0, len(users))
		for _, u := range users {
			response = append(response, UserResponse{
				ID:    u.ID,
				Name:  u.Name,
				Email: u.Email,
			})
		}

		c.JSON(http.StatusOK, response)
	})

	r.POST("/users", func(c *gin.Context) {
		var req CreateUserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			errors.Respond(c, errors.BadRequest("invalid request body"))
			return
		}

		user, err := service.CreateUser(req)
		if err != nil {
			errors.Respond(c, err)
			return
		}

		c.JSON(http.StatusCreated, UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	})
}
