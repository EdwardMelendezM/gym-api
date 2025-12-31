package auth

import (
	"gym-api/internal/modules/sessions/models"

	"github.com/gin-gonic/gin"
)

const sessionKey = "session"

func GetSession(c *gin.Context) (*models.Session, bool) {
	s, exists := c.Get(sessionKey)
	if !exists {
		return nil, false
	}

	session, ok := s.(*models.Session)
	return session, ok
}
