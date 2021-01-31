package service

import (
	"gin-demo/core/logger"
	"gin-demo/dao"
	"gin-demo/dto"
	"gin-demo/entity"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context, user *entity.User) error {
	return dao.CreateUser(ctx, user)
}

func ListUser(ctx *gin.Context) ([]dto.UserDTO, error) {

	dto, err := dao.ListUser(ctx)
	logger.Debug(ctx, "测试下", dto)
	return dto, err
}

func FindUserById(ctx *gin.Context, userId uint, user *entity.User) (int64, error) {
	return dao.FindUserById(ctx, userId, user)
}
