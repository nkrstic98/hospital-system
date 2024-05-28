package config

var DefaultConfig = Config{
	Web: WebConfig{
		Host:         "localhost",
		Port:         "8080",
		ClientAppUrl: "http://localhost:3000",
	},
	Database: DatabaseConfig{
		Host: "localhost",
		Port: "5432",
		DB:   "hospital",
		User: "hospital_role",
		Pwd:  "5tBsPvvXUBDw25zt",
	},
	Redis: RedisConfig{
		Host:          "localhost",
		Port:          "6379",
		Password:      "cWk5U29GSmw0KzI3",
		DatabaseIndex: 0,
	},
	AuthToken: AuthTokenConfig{
		Key:               "4b/7D7BnGZ+MuTEmLsfI6CwZ6kz3tI3EjD0InHb2c04=",
		ExpirationSeconds: 900,
	},
	Registry: Registry{
		RegistrarHost: "localhost",
		RegistrarPort: "8500",
	},
}
