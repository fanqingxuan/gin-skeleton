package dao

import (
	"gin-demo/dto"
	"gin-demo/entity"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context, user *entity.User) error {
	return DB(ctx).Create(&user).Error
}

func ListUser(ctx *gin.Context) ([]dto.UserDTO, error) {
	var users []dto.UserDTO

	result := DB(ctx).Table("users").Select("Name", "Age").Find(&users)
	return users, result.Error
}

func FindUserById(ctx *gin.Context, userId uint, user *entity.User) (affectRows int64, err error) {
	result := DB(ctx).Where(entity.User{UserId: userId}, "UserId").Find(user)
	return result.RowsAffected, result.Error
}
