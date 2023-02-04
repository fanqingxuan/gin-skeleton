package sqlx

import (
	"context"
	"database/sql"
	"gin-skeleton/svc/logx"
)

// ErrNotFound is an alias of sql.ErrNoRows
var ErrNotFound = sql.ErrNoRows

type (
	// Session stands for raw connections or transaction sessions
	Session interface {
		Exec(query string, args ...any) (sql.Result, error)
		ExecCtx(ctx context.Context, query string, args ...any) (sql.Result, error)
		Prepare(query string) (StmtSession, error)
		PrepareCtx(ctx context.Context, query string) (StmtSession, error)
		QueryRow(v any, query string, args ...any) error
		QueryRowCtx(ctx context.Context, v any, query string, args ...any) error
		QueryRowPartial(v any, query string, args ...any) error
		QueryRowPartialCtx(ctx context.Context, v any, query string, args ...any) error
		QueryRows(v any, query string, args ...any) error
		QueryRowsCtx(ctx context.Context, v any, query string, args ...any) error
		QueryRowsPartial(v any, query string, args ...any) error
		QueryRowsPartialCtx(ctx context.Context, v any, query string, args ...any) error
	}

	// SqlConn only stands for raw connections, so Transact method can be called.
	SqlConn interface {
		Session
		// RawDB is for other ORM to operate with, use it with caution.
		// Notice: don't close it.
		RawDB() (*sql.DB, error)
		Transact(fn func(Session) error) error
		TransactCtx(ctx context.Context, fn func(context.Context, Session) error) error
	}

	// SqlOption defines the method to customize a sql connection.
	SqlOption func(*commonSqlConn)

	// StmtSession interface represents a session that can be used to execute statements.
	StmtSession interface {
		Close() error
		Exec(args ...any) (sql.Result, error)
		ExecCtx(ctx context.Context, args ...any) (sql.Result, error)
		QueryRow(v any, args ...any) error
		QueryRowCtx(ctx context.Context, v any, args ...any) error
		QueryRowPartial(v any, args ...any) error
		QueryRowPartialCtx(ctx context.Context, v any, args ...any) error
		QueryRows(v any, args ...any) error
		QueryRowsCtx(ctx context.Context, v any, args ...any) error
		QueryRowsPartial(v any, args ...any) error
		QueryRowsPartialCtx(ctx context.Context, v any, args ...any) error
	}

	// thread-safe
	// Because CORBA doesn't support PREPARE, so we need to combine the
	// query arguments into one string and do underlying query without arguments
	commonSqlConn struct {
		connProv connProvider
		onError  func(error)
		beginTx  beginnable
		accept   func(error) bool
	}

	connProvider func() (*sql.DB, error)

	sessionConn interface {
		Exec(query string, args ...any) (sql.Result, error)
		ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
		Query(query string, args ...any) (*sql.Rows, error)
		QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	}

	statement struct {
		query string
		stmt  *sql.Stmt
	}

	stmtConn interface {
		Exec(args ...any) (sql.Result, error)
		ExecContext(ctx context.Context, args ...any) (sql.Result, error)
		Query(args ...any) (*sql.Rows, error)
		QueryContext(ctx context.Context, args ...any) (*sql.Rows, error)
	}
)

// NewSqlConn returns a SqlConn with given driver name and datasource.
func NewSqlConn(driverName, datasource string, opts ...SqlOption) SqlConn {
	conn := &commonSqlConn{
		connProv: func() (*sql.DB, error) {
			return getSqlConn(driverName, datasource)
		},
		onError: func(err error) {
			logInstanceError(datasource, err)
		},
		beginTx: begin,
	}
	for _, opt := range opts {
		opt(conn)
	}

	return conn
}

// NewSqlConnFromDB returns a SqlConn with the given sql.DB.
// Use it with caution, it's provided for other ORM to interact with.
func NewSqlConnFromDB(db *sql.DB, opts ...SqlOption) SqlConn {
	conn := &commonSqlConn{
		connProv: func() (*sql.DB, error) {
			return db, nil
		},
		onError: func(err error) {
			logx.Errorf("Error on getting sql instance: %v", err)
		},
		beginTx: begin,
	}
	for _, opt := range opts {
		opt(conn)
	}

	return conn
}

func (db *commonSqlConn) Exec(q string, args ...any) (result sql.Result, err error) {
	return db.ExecCtx(context.Background(), q, args...)
}

func (db *commonSqlConn) ExecCtx(ctx context.Context, q string, args ...any) (
	result sql.Result, err error) {

	var conn *sql.DB
	conn, err = db.connProv()
	if err != nil {
		db.onError(err)
		return nil, err
	}

	result, err = exec(ctx, conn, q, args...)
	return

}

func (db *commonSqlConn) Prepare(query string) (stmt StmtSession, err error) {
	return db.PrepareCtx(context.Background(), query)
}

func (db *commonSqlConn) PrepareCtx(ctx context.Context, query string) (stmt StmtSession, err error) {

	var conn *sql.DB
	conn, err = db.connProv()
	if err != nil {
		db.onError(err)
		return nil, err
	}

	st, err := conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	stmt = statement{
		query: query,
		stmt:  st,
	}

	return
}

func (db *commonSqlConn) QueryRow(v any, q string, args ...any) error {
	return db.QueryRowCtx(context.Background(), v, q, args...)
}

func (db *commonSqlConn) QueryRowCtx(ctx context.Context, v any, q string,
	args ...any) (err error) {

	return db.queryRows(ctx, func(rows *sql.Rows) error {
		return unmarshalRow(v, rows, true)
	}, q, args...)
}

func (db *commonSqlConn) QueryRowPartial(v any, q string, args ...any) error {
	return db.QueryRowPartialCtx(context.Background(), v, q, args...)
}

func (db *commonSqlConn) QueryRowPartialCtx(ctx context.Context, v any,
	q string, args ...any) (err error) {

	return db.queryRows(ctx, func(rows *sql.Rows) error {
		return unmarshalRow(v, rows, false)
	}, q, args...)
}

func (db *commonSqlConn) QueryRows(v any, q string, args ...any) error {
	return db.QueryRowsCtx(context.Background(), v, q, args...)
}

func (db *commonSqlConn) QueryRowsCtx(ctx context.Context, v any, q string,
	args ...any) (err error) {

	return db.queryRows(ctx, func(rows *sql.Rows) error {
		return unmarshalRows(v, rows, true)
	}, q, args...)
}

func (db *commonSqlConn) QueryRowsPartial(v any, q string, args ...any) error {
	return db.QueryRowsPartialCtx(context.Background(), v, q, args...)
}

func (db *commonSqlConn) QueryRowsPartialCtx(ctx context.Context, v any,
	q string, args ...any) (err error) {

	return db.queryRows(ctx, func(rows *sql.Rows) error {
		return unmarshalRows(v, rows, false)
	}, q, args...)
}

func (db *commonSqlConn) RawDB() (*sql.DB, error) {
	return db.connProv()
}

func (db *commonSqlConn) Transact(fn func(Session) error) error {
	return db.TransactCtx(context.Background(), func(_ context.Context, session Session) error {
		return fn(session)
	})
}

func (db *commonSqlConn) TransactCtx(ctx context.Context, fn func(context.Context, Session) error) (err error) {

	err = transact(ctx, db, db.beginTx, fn)

	return
}

func (db *commonSqlConn) acceptable(err error) bool {
	ok := err == nil || err == sql.ErrNoRows || err == sql.ErrTxDone || err == context.Canceled
	if db.accept == nil {
		return ok
	}

	return ok || db.accept(err)
}

func (db *commonSqlConn) queryRows(ctx context.Context, scanner func(*sql.Rows) error,
	q string, args ...any) (err error) {
	var qerr error
	conn, err := db.connProv()
	if err != nil {
		db.onError(err)
		return err
	}

	return query(ctx, conn, func(rows *sql.Rows) error {
		qerr = scanner(rows)
		return qerr
	}, q, args...)

}

func (s statement) Close() error {
	return s.stmt.Close()
}

func (s statement) Exec(args ...any) (sql.Result, error) {
	return s.ExecCtx(context.Background(), args...)
}

func (s statement) ExecCtx(ctx context.Context, args ...any) (result sql.Result, err error) {

	return execStmt(ctx, s.stmt, s.query, args...)
}

func (s statement) QueryRow(v any, args ...any) error {
	return s.QueryRowCtx(context.Background(), v, args...)
}

func (s statement) QueryRowCtx(ctx context.Context, v any, args ...any) (err error) {

	return queryStmt(ctx, s.stmt, func(rows *sql.Rows) error {
		return unmarshalRow(v, rows, true)
	}, s.query, args...)
}

func (s statement) QueryRowPartial(v any, args ...any) error {
	return s.QueryRowPartialCtx(context.Background(), v, args...)
}

func (s statement) QueryRowPartialCtx(ctx context.Context, v any, args ...any) (err error) {

	return queryStmt(ctx, s.stmt, func(rows *sql.Rows) error {
		return unmarshalRow(v, rows, false)
	}, s.query, args...)
}

func (s statement) QueryRows(v any, args ...any) error {
	return s.QueryRowsCtx(context.Background(), v, args...)
}

func (s statement) QueryRowsCtx(ctx context.Context, v any, args ...any) (err error) {

	return queryStmt(ctx, s.stmt, func(rows *sql.Rows) error {
		return unmarshalRows(v, rows, true)
	}, s.query, args...)
}

func (s statement) QueryRowsPartial(v any, args ...any) error {
	return s.QueryRowsPartialCtx(context.Background(), v, args...)
}

func (s statement) QueryRowsPartialCtx(ctx context.Context, v any, args ...any) (err error) {

	return queryStmt(ctx, s.stmt, func(rows *sql.Rows) error {
		return unmarshalRows(v, rows, false)
	}, s.query, args...)
}
