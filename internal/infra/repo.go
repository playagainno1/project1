package infra

import (
	"context"
	. "taylor-ai-server/internal/domain"
)

type UserRepo interface {
	Find(ctx context.Context, id string) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Create(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
}
