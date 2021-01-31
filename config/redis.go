package config

import "time"

type redis struct {
	// The network type, either tcp or unix.
	// Default is tcp.
	Network string `int:"network"`
	// host:port address.
	Addr string `ini:"addr"`

	// Use the specified Username to authenticate the current connection
	// with one of the connections defined in the ACL list when connecting
	// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
	Username string `ini:"username"`
	// Optional password. Must match the password specified in the
	// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
	// or the User Password when connecting to a Redis 6.0 instance, or greater,
	// that is using the Redis ACL system.
	Password string `ini:"password"`

	// Database to be selected after connecting to the server.
	DB int `ini:"db"`

	// Maximum number of retries before giving up.
	// Default is to not retry failed commands.
	MaxRetries int `ini:"max_retries"`
	// Minimum backoff between each retry.
	// Default is 8 milliseconds; -1 disables backoff.
	MinRetryBackoff time.Duration `ini:"min_retry_backoff"`
	// Maximum backoff between each retry.
	// Default is 512 milliseconds; -1 disables backoff.
	MaxRetryBackoff time.Duration `ini:"max_retry_backoff"`

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration `ini:"dial_timeout"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout time.Duration `ini:"read_timeout"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout time.Duration `ini:"write_timeout"`

	// Maximum number of socket connections.
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	PoolSize int `ini:"pool_size"`
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int `ini:"min_idle_conns"`
	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	MaxConnAge time.Duration `ini:"max_conn_age"`
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout time.Duration `ini:"pool_timeout"`
	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout time.Duration `ini:"idle_timeout"`
	// Frequency of idle checks made by idle connections reaper.
	// Default is 1 minute. -1 disables idle connections reaper,
	// but idle connections are still discarded by the client
	// if IdleTimeout is set.
	IdleCheckFrequency time.Duration `ini:"idle_check_frequency"`

	// Enables read only queries on slave nodes.
	readOnly bool `ini:"readonly"`
}
