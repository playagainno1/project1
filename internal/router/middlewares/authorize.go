package middlewares

import (
	"net/http"

	"taylor-ai-server/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const currentUserKey = "currentUser"

func GetCurrentUser(c *gin.Context) string {
	if v, ok := c.Get(currentUserKey); ok {
		return v.(string)
	}
	return ""
}

func loadCurrentUser(c *gin.Context) string {
	s := GetSession(c)
	if s == nil {
		return ""
	}
	user := s.Get(currentUserKey)
	if user == nil {
		return ""
	}
	return user.(string)
}

func SaveCurrentUser(c *gin.Context, user string) {
	s := GetSession(c)
	if s == nil {
		logrus.Error("save current user to session: session not found")
		return
	}
	s.Clear()
	s.Set(currentUserKey, user)
	if err := s.Save(); err != nil {
		logrus.WithError(err).WithField("user", user).Error("save current user to session")
	}
	c.Set(currentUserKey, user)
}

func RemoveCurrentUser(c *gin.Context) {
	s := GetSession(c)
	if s == nil {
		logrus.Error("remove current user from session: session not found")
		return
	}
	s.Delete(currentUserKey)
	if err := s.Save(); err != nil {
		logrus.WithError(err).Error("remove current user from session")
	}
}

func User(c *gin.Context) {
	user := loadCurrentUser(c)
	if user != "" {
		c.Set(currentUserKey, user)
	}

	c.Next()
}

func Authorize(c *gin.Context) {
	user := GetCurrentUser(c)
	if user == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrUnauthorized)
		return
	}

	c.Next()
}
