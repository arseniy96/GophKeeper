package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/arseniy96/GophKeeper/internal/server/logger"
)

const (
	TimeOut = 3 * time.Second
)

type Database struct {
	DB *sqlx.DB
}

func NewStorage(dsn string) (*Database, error) {
	if err := runMigrations(dsn); err != nil {
		return nil, fmt.Errorf("migrations failed with error: %w", err)
	}
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	database := &Database{
		DB: db,
	}
	logger.Log.Info("Database connection was created")

	return database, nil
}

func runMigrations(dsn string) error {
	const migrationsPath = "db/migrations"
	m, err := migrate.New(fmt.Sprintf("file://%s", migrationsPath), dsn)
	if err != nil {
		return fmt.Errorf("failed to get a new migrate instance: %w", err)
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed to apply migrations: %w", err)
		}
	}

	return nil
}

func (db *Database) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), TimeOut)
	defer cancel()

	return db.DB.PingContext(ctx)
}
