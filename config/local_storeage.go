package config

type LocalStorage struct {
	DefaultExpiration int    `yaml:"DefaultExpiration"`
	CleanupInterval   int    `yaml:"CleanupInterval"`
	KeyPrefix         string `yaml:"KeyPrefix"`
}
