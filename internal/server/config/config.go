package config

import (
	"encoding/json"
	"os"

	"github.com/caarlos0/env"
)

type Config struct {
	DatabaseDSN string `json:"database_dsn" env:"DATABASE_DSN"`
	Host        string `json:"host" env:"HOST"`
	LogLevel    string `json:"log_level" env:"LOG_LEVEL"`
}

var Settings = &Config{}

func Initialize() error {
	configFile, err := os.Open("./config/server/settings/production.json")
	if err != nil {
		return err
	}
	err = json.NewDecoder(configFile).Decode(Settings)
	if err != nil {
		return err
	}
	err = env.Parse(Settings)
	return err
}
