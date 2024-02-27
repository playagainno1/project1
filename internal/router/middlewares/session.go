package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetSession(c *gin.Context) sessions.Session {
	session := sessions.Default(c)
	return session
}

type CookieConfig struct {
	Name     string `yaml:"name" validate:"required"`
	Secret   string `yaml:"secret" validate:"required"`
	Domain   string `yaml:"domain" validate:"required"`
	MaxAge   int    `yaml:"maxAge" validate:"required"`
	Path     string `yaml:"path" validate:"required"`
	Secure   bool   `yaml:"secure" validate:"-"`
	HTTPOnly bool   `yaml:"httpOnly" validate:"-"`
	SameSite int    `yaml:"sameSite" validate:"required"`
}

func Cookie(c CookieConfig) gin.HandlerFunc {
	store := cookie.NewStore([]byte(c.Secret), []byte(c.Secret))
	options := sessions.Options{
		Path:     c.Path,
		Domain:   c.Domain,
		MaxAge:   c.MaxAge,
		Secure:   c.Secure,
		HttpOnly: c.HTTPOnly,
		SameSite: http.SameSite(c.SameSite),
	}
	store.Options(options)
	return sessions.Sessions(c.Name, store)
}
