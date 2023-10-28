package config

import (
	"encoding/json"
	"os"

	"github.com/caarlos0/env"
)

type Config struct {
	ConnectionTimeout int64  `json:"connection_timeout" env:"CONNECTION_TIMEOUT"`
	DatabaseDSN       string `json:"database_dsn" env:"DATABASE_DSN"`
	Host              string `json:"host" env:"HOST"`
	LogLevel          string `json:"log_level" env:"LOG_LEVEL"`
}

func Initialize() (*Config, error) {
	configFile, err := os.Open("./config/server/settings/production.json")
	if err != nil {
		return nil, err
	}
	var c = &Config{}
	err = json.NewDecoder(configFile).Decode(c)
	if err != nil {
		return nil, err
	}
	err = env.Parse(c)
	return c, err
}
