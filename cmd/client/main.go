package main

import (
	"log"

	"github.com/arseniy96/GophKeeper/internal/client/application"
	"github.com/arseniy96/GophKeeper/internal/client/config"
	"github.com/arseniy96/GophKeeper/src/logger"
)

// @Title GophKeeper Client.
// @Description Клиент для работы с приватными данными. Поддерживает сохранение, получение и редактирование.
// @Version 1.0.
// @Contact.email arsenzhar@yandex.ru.

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	settings, err := config.Initialize()
	if err != nil {
		return err
	}
	l, err := logger.Initialize(settings.LogLevel)
	if err != nil {
		return err
	}

	client, err := application.NewClient(l, settings)
	if err != nil {
		return err
	}

	return client.Start()
}
