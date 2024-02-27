package connection

import (
	"context"
	"time"

	"taylor-ai-server/internal/config"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var rd *redis.Client

func Redis() *redis.Client {
	return rd
}

func InitRedis() {
	logrus.Info("Connecting to redis...")
	c, err := newRedis(config.Config.Redis)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to redis")
	}
	rd = c
}

func newRedis(opt *redis.Options) (*redis.Client, error) {
	c := redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err := c.Ping(ctx).Result()
	return c, errors.WithStack(err)
}
