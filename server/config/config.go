package config

type Config struct {
	Web       WebConfig
	Database  DatabaseConfig
	Redis     RedisConfig
	AuthToken AuthTokenConfig
	Registry  Registry
}

type WebConfig struct {
	Host         string
	Port         string
	ClientAppUrl string
}

type DatabaseConfig struct {
	Host string
	Port string
	DB   string
	User string
	Pwd  string
}

type RedisConfig struct {
	Host          string
	Port          string
	Password      string
	DatabaseIndex int
}

type AuthTokenConfig struct {
	Key               string
	ExpirationSeconds int
}

type Registry struct {
	RegistrarHost string
	RegistrarPort string
	ServiceName   string
}
