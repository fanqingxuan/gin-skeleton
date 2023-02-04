package sqlx

import (
	"context"
	"database/sql"
	"fmt"
)

type (
	beginnable func(*sql.DB) (trans, error)

	trans interface {
		Session
		Commit() error
		Rollback() error
	}

	txSession struct {
		*sql.Tx
	}
)

// NewSessionFromTx returns a Session with the given sql.Tx.
// Use it with caution, it's provided for other ORM to interact with.
func NewSessionFromTx(tx *sql.Tx) Session {
	return txSession{Tx: tx}
}

func (t txSession) Exec(q string, args ...any) (sql.Result, error) {
	return t.ExecCtx(context.Background(), q, args...)
}

func (t txSession) ExecCtx(ctx context.Context, q string, args ...any) (result sql.Result, err error) {

	result, err = exec(ctx, t.Tx, q, args...)

	return
}

func (t txSession) Prepare(q string) (StmtSession, error) {
	return t.PrepareCtx(context.Background(), q)
}

func (t txSession) PrepareCtx(ctx context.Context, q string) (stmtSession StmtSession, err error) {

	stmt, err := t.Tx.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}

	return statement{
		query: q,
		stmt:  stmt,
	}, nil
}

func (t txSession) QueryRow(v any, q string, args ...any) error {
	return t.QueryRowCtx(context.Background(), v, q, args...)
}

func (t txSession) QueryRowCtx(ctx context.Context, v any, q string, args ...any) (err error) {

	return query(ctx, t.Tx, func(rows *sql.Rows) error {
		return unmarshalRow(v, rows, true)
	}, q, args...)
}

func (t txSession) QueryRowPartial(v any, q string, args ...any) error {
	return t.QueryRowPartialCtx(context.Background(), v, q, args...)
}

func (t txSession) QueryRowPartialCtx(ctx context.Context, v any, q string,
	args ...any) (err error) {

	return query(ctx, t.Tx, func(rows *sql.Rows) error {
		return unmarshalRow(v, rows, false)
	}, q, args...)
}

func (t txSession) QueryRows(v any, q string, args ...any) error {
	return t.QueryRowsCtx(context.Background(), v, q, args...)
}

func (t txSession) QueryRowsCtx(ctx context.Context, v any, q string, args ...any) (err error) {

	return query(ctx, t.Tx, func(rows *sql.Rows) error {
		return unmarshalRows(v, rows, true)
	}, q, args...)
}

func (t txSession) QueryRowsPartial(v any, q string, args ...any) error {
	return t.QueryRowsPartialCtx(context.Background(), v, q, args...)
}

func (t txSession) QueryRowsPartialCtx(ctx context.Context, v any, q string,
	args ...any) (err error) {

	return query(ctx, t.Tx, func(rows *sql.Rows) error {
		return unmarshalRows(v, rows, false)
	}, q, args...)
}

func begin(db *sql.DB) (trans, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	return txSession{
		Tx: tx,
	}, nil
}

func transact(ctx context.Context, db *commonSqlConn, b beginnable,
	fn func(context.Context, Session) error) (err error) {
	conn, err := db.connProv()
	if err != nil {
		db.onError(err)
		return err
	}

	return transactOnConn(ctx, conn, b, fn)
}

func transactOnConn(ctx context.Context, conn *sql.DB, b beginnable,
	fn func(context.Context, Session) error) (err error) {
	var tx trans
	tx, err = b(conn)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("recover from %#v, rollback failed: %w", p, e)
			} else {
				err = fmt.Errorf("recoveer from %#v", p)
			}
		} else if err != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("transaction failed: %s, rollback failed: %w", err, e)
			}
		} else {
			err = tx.Commit()
		}
	}()

	return fn(ctx, tx)
}
