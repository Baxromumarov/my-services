package config

import (
	"os"

    "github.com/spf13/cast"
)

// Config ...
type Config struct {
    Environment       string // develop, staging, production
    PostgresHost      string
    PostgresPort      int
    PostgresDatabase  string
    PostgresUser      string
    PostgresPassword  string
    LogLevel          string
    RPCPort           string
    ReviewServiceHost string
    ReviewServicePort int

    PostServicePort    int
    PostServiceHost    string
    KafkaHost   string
    KafkaPort   int
}

// Load loads environment vars and inflates Config
func Load() Config {
    c := Config{}

    c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

    c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
    c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
    c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "testdb"))
    c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
    c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1234"))

    c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "localhost"))
    c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 7070))
    c.KafkaHost = cast.ToString(getOrReturnDefault("KAFKA_HOST", "localhost"))
    c.KafkaPort = cast.ToInt(getOrReturnDefault("KAFKA_PORT", 9092))

    c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

    c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":8000"))

    return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
    _, exists := os.LookupEnv(key)
    if exists {
        return os.Getenv(key)
    }

    return defaultValue
}
