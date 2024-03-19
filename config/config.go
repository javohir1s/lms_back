package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresPassword string
	PostgresUser     string
	PostgresDatabase string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println("error!!!", err)
	}
	cfg := Config{}

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "lms"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "javohir"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "12345"))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {

	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
