package svc

import (
	"context"
	"gin-skeleton/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Redis  *AWRedis
	DB     *gorm.DB
	Log    *Log
	CacheX *CacheX
}

func (that *ServiceContext) WithLog(log *Log) *ServiceContext {
	return &ServiceContext{
		Config: that.Config,
		Redis:  that.Redis,
		DB:     NewDB(that.Config, log),
		Log:    log,
		CacheX: that.CacheX,
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
		Config: c,
		Redis:  NewRedis(context.Background(), Redis),
		Log:    log,
		DB:     NewDB(c, log),
		CacheX: NewCacheX(c),
	}
}
