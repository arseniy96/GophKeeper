package main

import (
	"log"

	"github.com/arseniy96/GophKeeper/internal/server/config"
	"github.com/arseniy96/GophKeeper/internal/server/handlers"
	"github.com/arseniy96/GophKeeper/internal/server/storage"
	"github.com/arseniy96/GophKeeper/src/logger"
)

// @Title GophKeeper Server.
// @Description Сервер для хранения приватных данных, таких, как пароли, карточные данные или файлы.
// @Version 1.0.
// @Contact.email arsenzhar@yandex.ru.
// @Host localhost:3200.

const (
	configPath = "./config/server/settings/production.json"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	settings, err := config.Initialize(configPath)
	if err != nil {
		return err
	}
	l, err := logger.Initialize(settings.LogLevel)
	if err != nil {
		return err
	}

	rep, err := storage.NewStorage(settings.DatabaseDSN, l)
	if err != nil {
		l.Log.Errorf("database error: %v", err)
		return err
	}
	defer func() {
		if err = rep.Close(); err != nil {
			l.Log.Errorf("database close error: %v", err)
		}
	}()

	serverGRPC := handlers.NewServer(rep, settings, l)

	return serverGRPC.Start()
}
