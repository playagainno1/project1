package api

import (
	"context"
	"taylor-ai-server/internal/router/middlewares"

	"github.com/gin-gonic/gin"
)

type HTTP struct {
	userFetcher *UserFetcher
	profile     *Profile
	ranks       *Ranks
}

func NewHTTP() *HTTP {
	return &HTTP{
		userFetcher: NewUserFetcher(),
		profile:     newProfile(),
		ranks:       newRanks(),
	}
}

func (h *HTTP) Profile(ctx context.Context, r *ProfileRequest) (*ProfileResponse, error) {
	user, _ := h.userFetcher.User(ctx)
	res, err := h.profile.Handle(ctx, user)
	if err != nil {
		return nil, err
	}
	middlewares.SaveCurrentUser(ctx.(*gin.Context), res.ID)
	return res, nil
}

func (h *HTTP) Ranks(ctx context.Context, r *EmptyRequest) (*RanksResponse, error) {
	user, _ := h.userFetcher.User(ctx)
	res, err := h.ranks.Handle(ctx, user)
	return res, err
}
