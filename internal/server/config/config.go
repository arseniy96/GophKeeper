package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caarlos0/env"
)

// Config – объект конфигурации сервера.
type Config struct {
	// DatabaseDSN – dsn для подключения к БД.
	DatabaseDSN string `json:"database_dsn" env:"DATABASE_DSN"`
	// Host – адрес сервера.
	Host string `json:"host" env:"HOST"`
	// LogLevel – уровень логгирования.
	LogLevel string `json:"log_level" env:"LOG_LEVEL"`
	// SecretKey – ключ шифрования.
	SecretKey string `json:"secret_key"`
}

// Initialize – функция инициализации конфига.
func Initialize(configPath string) (*Config, error) {
	configFile, err := os.Open(configPath)
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
