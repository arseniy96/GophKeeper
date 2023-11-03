package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caarlos0/env"
)

type Config struct {
	DatabaseDSN       string `json:"database_dsn" env:"DATABASE_DSN"`
	Host              string `json:"host" env:"HOST"`
	LogLevel          string `json:"log_level" env:"LOG_LEVEL"`
	SecretKey         string `json:"secret_key"`
	ConnectionTimeout int64  `json:"connection_timeout" env:"CONNECTION_TIMEOUT"`
}

func Initialize() (*Config, error) {
	configFile, err := os.Open("./config/server/settings/production.json")
	if err != nil {
		return nil, fmt.Errorf("open file error: %w", err)
	}
	var c = &Config{}
	err = json.NewDecoder(configFile).Decode(c)
	if err != nil {
		return nil, fmt.Errorf("parse JSON error: %w", err)
	}
	err = env.Parse(c)
	return c, err
}
