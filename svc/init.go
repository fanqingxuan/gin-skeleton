package svc

import (
	"gin-skeleton/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	Response *Response
	Redis    *redis.Client
	DB       *gorm.DB
	Log      *Log
}

func NewServiceContext(c config.Config) *ServiceContext {
	log := NewLog("app/", c.Log.Level, BusinessLogType)

	var Redis = redis.NewClient(&redis.Options{
		Network:  c.Redis.Network,
		Addr:     c.Redis.Addr,
		Username: c.Redis.Username,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})
	return &ServiceContext{
		Config:   c,
		Response: NewResponse(),
		Redis:    Redis,
		Log:      log,
		DB:       NewDB(c, log),
	}
}
