package router

import (
	"net/http"

	"taylor-ai-server/internal/api"
	"taylor-ai-server/internal/config"
	"taylor-ai-server/internal/domain"
	"taylor-ai-server/internal/router/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Router struct {
	h *api.HTTP
}

func NewRouter(h *api.HTTP) *Router {
	return &Router{h: h}
}

func (ro *Router) Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(middlewares.RequestLogger)

	r.Use(gin.CustomRecoveryWithWriter(panicWriter{}, func(c *gin.Context, e interface{}) {
		var err error
		if v, ok := e.(error); ok {
			err = v
		} else {
			err = errors.Errorf("%v", e)
		}
		responseError(c, err, domain.ErrInternalError)
	}))

	r.Use(middlewares.Cookie(config.Config.Cookie))
	r.Use(middlewares.User)

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/api/user/profile", WrapHandler(ro.h.Profile))

	r.Use(middlewares.Authorize)

	r.GET("/api/ranks", WrapHandler(ro.h.Ranks))

	return r
}
