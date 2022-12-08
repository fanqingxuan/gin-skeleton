package svc

import (
	"context"
	"gin-skeleton/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Ctx          context.Context
	Config       config.Config
	Response     *Response
	Redis        *AWRedis
	DB           *gorm.DB
	Log          *Log
	LocalStorage *LocalStorage
}

func (that *ServiceContext) WithContext(ctx context.Context) *ServiceContext {
	return &ServiceContext{
		Ctx:          ctx,
		Config:       that.Config,
		Response:     that.Response.WithContext(ctx),
		Redis:        NewRedis(ctx, that.Redis.client),
		DB:           that.DB.WithContext(ctx),
		Log:          that.Log.WithContext(ctx),
		LocalStorage: that.LocalStorage,
	}
}

func (that *ServiceContext) WithLog(log *Log) *ServiceContext {
	return &ServiceContext{
		Ctx:          that.Ctx,
		Config:       that.Config,
		Response:     that.Response,
		Redis:        that.Redis,
		DB:           NewDB(that.Config, log),
		Log:          log,
		LocalStorage: that.LocalStorage,
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
		Config:       c,
		Response:     NewResponse(),
		Redis:        NewRedis(context.Background(), Redis),
		Log:          log,
		DB:           NewDB(c, log),
		LocalStorage: NewLocalStorage(c),
	}
}
