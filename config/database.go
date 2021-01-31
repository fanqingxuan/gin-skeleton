package config

import "time"

type db struct {
	DATABASE_URL    string        `ini:"database_url"`
	MaxIdleConns    int           `ini:"max_idle_conns"`
	MaxOpenConns    int           `ini:"max_open_conns"`
	ConnMaxLifeTime int           `ini:"conn_max_lifetime"`
	ShowSql         bool          `ini:"print_sql"`
	ENGINE          string        `ini:"engine"`
	CHARSET         string        `ini:"charset"`
	PREFIX          string        `ini:"prefix"`
	SlowThreshold   time.Duration `ini:"slow_threshold"`
}
