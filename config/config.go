package config

type Config struct {
	App    `yaml:"App"`
	Redis  `yaml:"Redis"`
	Mysql  `yaml:"Mysql"`
	Log    `yaml:"Log"`
	CacheX `yaml:"CacheX"`
}
