package model

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID       uint `gorm:"primaryKey;column:uuid"`
	Username string
	Age      int
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt time.Time
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (that *UserModel) FindOne(pk uint) (*User, error) {
	var user *User
	err := that.db.Limit(1).Find(&user, pk).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
