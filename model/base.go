package model

import (
	"context"
	"database/sql"
	"gin-skeleton/svc/sqlx"
	"reflect"
)

type Model struct {
	db  sqlx.SqlConn
	ctx context.Context
}

func NewModel(ctx context.Context, db sqlx.SqlConn) Model {
	return Model{
		db:  db,
		ctx: ctx,
	}
}

func (that *Model) Query(v interface{}, query string, args ...interface{}) error {
	if reflect.TypeOf(v).Elem().Kind() == reflect.Slice {
		return that.db.QueryRowsPartialCtx(that.ctx, v, query, args...)
	} else {
		return that.db.QueryRowPartialCtx(that.ctx, v, query, args...)
	}
}

func (that *Model) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := that.db.ExecCtx(that.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (that *Model) Transaction(fn func(context.Context, sqlx.Session) error) error {
	return that.db.TransactCtx(that.ctx, fn)
}
