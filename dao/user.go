package dao

import (
	"context"
	"gin-skeleton/entity"
	"gin-skeleton/svc"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserDao struct {
	Dao
}

func NewUserDao(ctx context.Context, svcCtx *svc.ServiceContext) *UserDao {
	return &UserDao{
		Dao: *NewDao(ctx, svcCtx),
	}
}

func (that *UserDao) GetUserInfo(userId uint) (*entity.User, error) {
	var user entity.User
	err := that.db.First(&user, userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}
