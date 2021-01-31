package config

type app struct {
	AppName string `ini:"appname"`
	Port    string `ini:"port"`
	Debug   bool   `ini:"debug"`
}
