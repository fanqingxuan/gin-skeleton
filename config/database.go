package config

type DB struct {
	Username      string `yaml:"Username"`
	Password      string `yaml:"Password"`
	Host          string `yaml:"Host"`
	Port          int    `yaml:"Port"`
	DbName        string `yaml:"DbName"`
	Charset       string `yaml:"Charset"`
	TablePrefix   string `yaml:"TablePrefix"`
	SingularTable bool   `yaml:"SingularTable"`
}
