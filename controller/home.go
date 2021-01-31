package controller

import (
	"gin-demo/constant"
	"gin-demo/core"
	"gin-demo/core/logger"
	"gin-demo/entity"
	"gin-demo/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	users, _ := service.ListUser(ctx)
	Success(ctx, users)
}

func Create(ctx *gin.Context) {

	user := entity.User{
		Name: "测试",
	}
	service.CreateUser(ctx, &user)
	SuccessWithMessage(ctx, "成功了", user)
}

func TestRedis(ctx *gin.Context) {
	logger.Info(ctx, "成", "redis")
	val, err := core.Redis.Get(ctx, "hello").Result()
	Success(ctx, map[string]interface{}{"val": val, "err": err})
}

func TestLog(ctx *gin.Context) {
	logger.Error(ctx, "成功的表现", "成功呢了的")
	service.ListUser(ctx)
	Success(ctx, "成功")
}

func FindUser(ctx *gin.Context) {
	var user entity.User
	userId, _ := strconv.Atoi(ctx.DefaultQuery("id", "0"))
	exists, _ := service.FindUserById(ctx, uint(userId), &user)
	logger.Warn(ctx, "warn keyworkds", user)
	if exists <= 0 {
		ErrorWithMessage(ctx, constant.CODE_500, "用户不存在")
		return
	}
	Success(ctx, user)
}
