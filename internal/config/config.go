package config

import (
	"os"

	"taylor-ai-server/internal/router/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Env      string    `validate:"required"`
	Addr     string    `validate:"required"`
	Cookie   Cookie    `validate:"required"`
	Database *Database `validate:"required"`
	Redis    *Redis    `validate:"required"`
}

var Config = Configuration{}

type (
	Cookie = middlewares.CookieConfig
	Redis  = redis.Options
)

type Database struct {
	Driver     string            `validate:"required"`
	DSN        string            `validate:"required"`
	Parameters map[string]string `validate:"-"`
}

func LoadConfig(name string) error {
	data, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}
	return validator.New().Struct(Config)
}
