package svc

import (
	"context"
	"gin-skeleton/config"
	"gin-skeleton/svc/logx"
	"gin-skeleton/svc/sqlx"

	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	Config config.Config
	Redis  *AWRedis
	Mysql  sqlx.SqlConn
	CacheX *CacheX
}

func NewServiceContext(c config.Config) *ServiceContext {

	var Redis = redis.NewClient(&redis.Options{
		Network:  c.Redis.Network,
		Addr:     c.Redis.Addr,
		Username: c.Redis.Username,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})
	logx.NewLog(c.Log.Level)
	return &ServiceContext{
		Config: c,
		Redis:  NewRedis(context.Background(), Redis),
		Mysql:  sqlx.NewMysql(c.Mysql.DataSource),
		CacheX: NewCacheX(c),
	}
}
