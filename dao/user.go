package dao

import (
	"gin-skeleton/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserDao struct {
	*Dao
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		Dao: NewDao(db),
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
