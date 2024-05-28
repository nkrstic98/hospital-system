package config

import "time"

type Config struct {
	Web      WebConfig
	Database DatabaseConfig
	Registry Registry
}

type WebConfig struct {
	Host    string
	Port    string
	Timeout time.Duration
}

type DatabaseConfig struct {
	Host string
	Port string
	DB   string
	User string
	Pwd  string
}

type Registry struct {
	RegistrarHost string
	RegistrarPort string
	ServiceName   string
}
