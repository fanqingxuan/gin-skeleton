package model

import (
	"context"
	"gin-skeleton/svc/sqlx"
)

type User struct {
	Uid      int64  `db:"uid"`
	Username string `db:"username"`
	Age      int64  `db:"age"`
}

type UserModel struct {
	db  sqlx.SqlConn
	ctx context.Context
}

func NewUserModel(ctx context.Context, db sqlx.SqlConn) UserModel {
	return UserModel{
		db:  db,
		ctx: ctx,
	}
}

func (that *UserModel) FindOne(pk uint) (*User, error) {
	var resp User
	err := that.db.QueryRowCtx(that.ctx, &resp, "select * from users where uid=?", pk)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}
}

func (that *UserModel) List(age int) ([]User, error) {
	var resp []User
	err := that.db.QueryRowsCtx(that.ctx, &resp, "select * from users where age>?", age)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}
}

func (that *UserModel) Insert(user *User) (int64, error) {
	const insertsql = `insert into users (username,age) values(?, ?)`
	// insert op
	res, err := that.db.Exec(insertsql, user.Username, user.Age)
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
	const deletesql = `delete from users where uid=?`
	// insert op
	_, err := that.db.Exec(deletesql, pk)
	return err
}
