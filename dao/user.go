package dao

import (
	"fmt"
	"gin-skeleton/entity"
	"gin-skeleton/svc"
)

type UserDao struct {
	Dao
}

func NewUserDao(svcCtx *svc.ServiceContext) *UserDao {
	return &UserDao{
		Dao: *NewDao(svcCtx),
	}
}

func (that *UserDao) GetUserInfo(userId uint) *entity.User {
	var user entity.User
	that.db.First(&user, userId)
	fmt.Println(that.db)
	return &user
}
