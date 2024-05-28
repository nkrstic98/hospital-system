package config

import "time"

var DefaultConfig = Config{
	Web: WebConfig{
		Host:    "localhost",
		Port:    "8081",
		Timeout: 0 * time.Second,
	},
	Database: DatabaseConfig{
		Host: "localhost",
		Port: "5432",
		DB:   "authz",
		User: "hospital_role",
		Pwd:  "5tBsPvvXUBDw25zt",
	},
	Registry: Registry{
		RegistrarHost: "localhost",
		RegistrarPort: "8500",
		ServiceName:   "authorization",
	},
}
