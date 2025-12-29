package utils

import (
	"gym-api/internal/modules/sessions"

	"github.com/gin-gonic/gin"
)

const sessionKey = "session"

func GetSession(c *gin.Context) (*sessions.Session, bool) {
	s, exists := c.Get(sessionKey)
	if !exists {
		return nil, false
	}

	session, ok := s.(*sessions.Session)
	return session, ok
}
