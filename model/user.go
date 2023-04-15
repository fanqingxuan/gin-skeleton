package model

import (
	"context"
	"gin-skeleton/svc/sqlx"

	"github.com/golang-module/carbon/v2"
)

type User struct {
	Uid      int64            `db:"uid"`
	Username string           `db:"username"`
	Age      int64            `db:"age"`
	Ctime    carbon.Timestamp `db:"create_time"`
}

type UserModel struct {
	Model
}

func NewUserModel(ctx context.Context, db sqlx.SqlConn) UserModel {
	return UserModel{
		Model: NewModel(ctx, db),
	}
}

func (that *UserModel) FindOne(pk uint) (*User, error) {
	var user User
	err := that.QueryOne(&user, "select * from users where uid=?", pk)
	switch err {
	case nil:
		return &user, nil
	case sqlx.ErrNotFound:
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}
}

func (that *UserModel) List(age int) ([]User, error) {
	var users []User
	err := that.QueryAll(&users, "select age,username from users where age>?", age)
	switch err {
	case nil:
		return users, nil
	case sqlx.ErrNotFound:
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}
}

func (that *UserModel) Insert(user *User) (int64, error) {
	const sql = `insert into users (username,age) values(?, ?)`
	// insert op
	res, err := that.Execute(sql, user.Username, user.Age)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (that *UserModel) Delete(pk uint64) error {
	const sql = `delete from users where uid=?`
	// insert op
	_, err := that.Execute(sql, pk)
	return err
}
