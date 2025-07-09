package app

import (
	"errors"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServiceAddress string `env:"SERVICE_ADDRESS" env-default:":8080"`
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
