package svc

import (
	"context"
	"gin-skeleton/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	Response *Response
	Redis    *AWRedis
	DB       *gorm.DB
	Log      *Log
	XCache   *XCache
}

func (that *ServiceContext) WithLog(log *Log) *ServiceContext {
	return &ServiceContext{
		Config:   that.Config,
		Response: that.Response,
		Redis:    that.Redis,
		DB:       NewDB(that.Config, log),
		Log:      log,
		XCache:   that.XCache,
	}
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
		Redis:    NewRedis(context.Background(), Redis),
		Log:      log,
		DB:       NewDB(c, log),
		XCache:   NewXCache(c),
	}
}
