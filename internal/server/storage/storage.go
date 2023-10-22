package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/arseniy96/GophKeeper/internal/server/logger"
)

const (
	TimeOut = 3 * time.Second
)

var ErrConflict = errors.New(`already exists`)
var ErrNowRows = errors.New(`missing data`)

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

func (db *Database) Close() error {
	return db.DB.Close()
}

func (db *Database) CreateUser(ctx context.Context, login, password string) error {
	var pgErr *pgconn.PgError

	_, err := db.DB.ExecContext(ctx,
		`INSERT INTO users(login, password) VALUES($1, $2)`,
		login, password)
	if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		return ErrConflict
	}

	return err
}

func (db *Database) UpdateUserToken(ctx context.Context, login, token string) error {
	_, err := db.DB.ExecContext(ctx,
		`UPDATE users SET token=$1 WHERE login=$2`,
		token, login)
	return err
}

func (db *Database) FindUserByLogin(ctx context.Context, login string) (*User, error) {
	var u User
	err := db.DB.QueryRowContext(ctx,
		`SELECT id, login, password, token FROM users WHERE login=$1 LIMIT(1)`,
		login).Scan(&u.ID, &u.Login, &u.Password, &u.Token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNowRows
		}
		return nil, err
	}

	return &u, nil
}

func (db *Database) FindUserByToken(ctx context.Context, token string) (*User, error) {
	var u User
	err := db.DB.QueryRowContext(ctx,
		`SELECT id, login, password, token FROM users WHERE token=$1 LIMIT(1)`,
		token).Scan(&u.ID, &u.Login, &u.Password, &u.Token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNowRows
		}
		return nil, err
	}

	return &u, nil
}

func (db *Database) SaveUserData(ctx context.Context, userID int, name, dataType string, data []byte) error {
	var pgErr *pgconn.PgError

	_, err := db.DB.ExecContext(ctx,
		`INSERT INTO user_records(name, data, data_type, user_id) VALUES($1, $2, $3, $4)`,
		name, data, dataType, userID)

	if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		return ErrConflict
	}

	return err
}
