package svc

import (
	"fmt"
	"gin-skeleton/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config   config.Config
	Response *Response
	Redis    *redis.Client
	DB       *gorm.DB
	Log      *Log
}

func NewServiceContext(c config.Config) *ServiceContext {

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.DbName, c.DB.Charset)
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.DB.TablePrefix,
			SingularTable: c.DB.SingularTable,
		},
	})
	if err != nil {
		fmt.Println(DSN)
		panic(err)
	}

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
		DB:       DB,
		Log:      NewLog("app/", c.Log.Level),
	}
}
