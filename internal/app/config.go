package app

import (
	"errors"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServiceAddress string `env:"SERVICE_ADDRESS" env-default:":8080"`
	DBHost         string `env:"DATA_PROTECTION_DB_ADDRESS" env-default:"localhost"`
	DBPort         int    `env:"DATA_PROTECTION_DB_PORT" env-default:"5432"`
	DBUser         string `env:"DATA_PROTECTION_DB_USER" env-default:"postgres"`
	DBPassword     string `env:"DATA_PROTECTION_DB_PASSWORD" env-default:"Tunaiku2018"`
	DBName         string `env:"DATA_PROTECTION_DB_NAME" env-default:"go_machine"`
	DBPoolSize     int    `env:"DATA_PROTECTION_DB_POOL_SIZE" env-default:"10"`
	DBMaxRetries   int    `env:"DATA_PROTECTION_DB_MAX_RETRIES" env-default:"3"`
	DBRetryDelay   int    `env:"DATA_PROTECTION_DB_RETRY_DELAY" env-default:"3"`
	DBIdleTimeout  int    `env:"DATA_PROTECTION_DB_IDLE_TIMEOUT" env-default:"30"`
	DBWriteTimeout int    `env:"DATA_PROTECTION_DB_WRITE_TIMEOUT" env-default:"30"`
	DBPoolTimeout  int    `env:"DATA_PROTECTION_DB_POOL_TIMEOUT" env-default:"30"`
}

func loadConfig() (Config, error) {
	var config Config
	err := cleanenv.ReadConfig(".env", &config)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return config, err
	}

	if errors.Is(err, os.ErrNotExist) {
		err := cleanenv.ReadEnv(&config)
		if err != nil {
			return config, err
		}
	}

	return config, nil
}
