package config

import (
	"encoding/json"
	"os"

	"github.com/caarlos0/env"
)

type Config struct {
	ConnectionTimout int64  `json:"connection_timeout" env:"CONNECTION_TIMEOUT"`
	Host             string `json:"host" env:"HOST"`
	LogLevel         string `json:"log_level" env:"LOG_LEVEL"`
}

func Initialize() (*Config, error) {
	configFile, err := os.Open("./config/client/settings/production.json")
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
