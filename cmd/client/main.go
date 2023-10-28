package main

import (
	"log"

	"github.com/arseniy96/GophKeeper/internal/client/application"
	"github.com/arseniy96/GophKeeper/internal/client/commands"
	"github.com/arseniy96/GophKeeper/internal/client/config"
	"github.com/arseniy96/GophKeeper/src/logger"
)

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

	client := application.NewClient(l, settings)

	return commands.Start(client)
}
