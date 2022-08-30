package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
 * Configuration
 */
const (
	_SESSION_COOKIE_NAME string = "session"
	_SESSION_MAX_AGE     int    = 60 * 60 * 24 * 2 // 2 days (in seconds)
)

//////////

func GetSessionCookie(c *gin.Context) (string, error) {
	return c.Cookie(_SESSION_COOKIE_NAME)
}

func SetSessionCookie(c *gin.Context, session_id string) {
	c.SetCookie(_SESSION_COOKIE_NAME, session_id, _SESSION_MAX_AGE, "/", "localhost", false, true)
}

func GenSessionCookie(c *gin.Context) string {
	id := uuid.New().String()
	SetSessionCookie(c, id)
	return id
}

func ClearSessionCookie(c *gin.Context) {
	session_id, err := GetSessionCookie(c)
	if err != nil {
		// Cookie does not exist !
		return
	}
	// Set the lifetime of the cookie to 0.
	c.SetCookie(_SESSION_COOKIE_NAME, session_id, 0, "/", "localhost", false, true)
}
