package config

type Config struct {
	App   `yaml:"App"`
	Redis `yaml:"Redis"`
	DB    `yaml:"DB"`
	Log   `yaml:"Log"`
}
