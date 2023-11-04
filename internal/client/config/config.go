package config

import (
	"encoding/json"
	"os"

	"github.com/caarlos0/env"
)

// Config – объект конфигурации клиента.
type Config struct {
	// Host – адрес сервера.
	Host string `json:"host" env:"HOST"`
	// LogLevel – уровень логгирования.
	LogLevel string `json:"log_level" env:"LOG_LEVEL"`
	// ConnectionTimeout – таймаут запроса к серверу.
	ConnectionTimeout int64 `json:"connection_timeout" env:"CONNECTION_TIMEOUT"`
}

// Initialize – функция инициализации конфига.
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
