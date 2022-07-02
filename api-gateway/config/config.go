package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	UserServiceHost string
	UserServicePort int

	PostgresHost string
	PostgresPort int
	PostgresUser string
	PostgresPassword string
	PostgresDatabase string

	//context timeout in seconds
	CtxTimeout int
	RedisHost  string
	RedisPort  int

	LogLevel string
	HTTPPort string


	SigningKey string

	CasbinConfigPath string
	CasbinConfigPathCsv string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

    c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
    c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
    c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "casbindb"))
    c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
    c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1234"))
	c.CasbinConfigPath = cast.ToString(getOrReturnDefault("CASBIN_CONFIG_PATH", "config/rbac_model.conf"))
	c.CasbinConfigPathCsv = cast.ToString(getOrReturnDefault("CASBIN_CONFIG_PATH_CSV", "config/rbac_model.csv"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8010"))
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "127.0.0.1"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 8000))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))

	c.SigningKey = cast.ToString(getOrReturnDefault("SIGNING_KEY","hbzfpetosnncegaoigsouhlfwbhhwwgzjixcrmiayjmeaoclmtjqtbwvgysvkdudjtfeaoopdstljhqxwhsobwootanhihawjhyqfbarnbpuvirmxfoqepmhkxqtlktdzzneqnxqxvputkpetolylqhxneiauopj"))

	
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
