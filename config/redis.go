package config

type Redis struct {
	// The network type, either tcp or unix.
	// Default is tcp.
	Network string `yaml:"Network"`
	// host:port address.
	Addr string `yaml:"Addr"`

	// Use the specified Username to authenticate the current connection
	// with one of the connections defined in the ACL list when connecting
	// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
	Username string `yaml:"Username"`
	// Optional password. Must match the password specified in the
	// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
	// or the User Password when connecting to a Redis 6.0 instance, or greater,
	// that is using the Redis ACL system.
	Password string `yaml:"Password"`

	// Database to be selected after connecting to the server.
	DB int `yaml:"DB"`
}
