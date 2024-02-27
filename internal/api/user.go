package api

import (
	"context"
	. "taylor-ai-server/internal/domain"
	"taylor-ai-server/internal/infra"
	"taylor-ai-server/internal/infra/repo"
	"taylor-ai-server/internal/router/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type UserFetcher struct {
	UserRepo infra.UserRepo
}

func NewUserFetcher() *UserFetcher {
	return &UserFetcher{
		UserRepo: repo.NewUserRepo(),
	}
}

func (s *UserFetcher) MustUser(ctx context.Context) User {
	u, err := s.User(ctx)
	if err != nil {
		e := errors.Wrap(ErrUnauthorized, err.Error())
		panic(e)
	}
	return u
}

func (s *UserFetcher) User(ctx context.Context) (User, error) {
	id := middlewares.GetCurrentUser(ctx.(*gin.Context))
	if id == "" {
		return User{}, ErrNotFound
	}
	u, err := s.UserRepo.Find(ctx, id)
	return u, err
}
