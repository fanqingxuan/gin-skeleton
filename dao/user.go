package dao

import (
	"context"
	"gin-skeleton/entity"
	"gin-skeleton/svc"
)

type UserDao struct {
	Dao
}

func NewUserDao(ctx context.Context, svcCtx *svc.ServiceContext) *UserDao {
	return &UserDao{
		Dao: *NewDao(ctx, svcCtx),
	}
}

func (that *UserDao) GetUserInfo(userId uint) *entity.User {
	var user entity.User
	that.db.First(&user, userId)
	return &user
}
