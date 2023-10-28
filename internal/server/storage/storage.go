package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/arseniy96/GophKeeper/src/logger"
)

const (
	TimeOut = 3 * time.Second
)

var ErrConflict = errors.New(`already exists`)
var ErrNowRows = errors.New(`missing data`)

type Database struct {
	DB *pgx.Conn
}

func NewStorage(dsn string, l *logger.Logger) (*Database, error) {
	if err := runMigrations(dsn); err != nil {
		return nil, fmt.Errorf("migrations failed with error: %w", err)
	}

	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	database := &Database{
		DB: db,
	}
	l.Log.Info("Database connection was created")

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

	return db.DB.Ping(ctx)
}

func (db *Database) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), TimeOut)
	defer cancel()

	return db.DB.Close(ctx)
}

func (db *Database) CreateUser(ctx context.Context, login, password string) error {
	tx, err := db.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var pgErr *pgconn.PgError
	_, err = tx.Exec(ctx,
		`INSERT INTO users(login, password) VALUES($1, $2)`,
		login, password)
	if err != nil {
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return ErrConflict
		}
		return err
	}

	return tx.Commit(ctx)
}

func (db *Database) UpdateUserToken(ctx context.Context, login, token string) error {
	tx, err := db.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	_, err = tx.Exec(ctx,
		`UPDATE users SET token=$1 WHERE login=$2`,
		token, login)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (db *Database) FindUserByLogin(ctx context.Context, login string) (*User, error) {
	tx, err := db.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var u User
	err = tx.QueryRow(ctx,
		`SELECT id, login, password, token FROM users WHERE login=$1 LIMIT(1)`,
		login).Scan(&u.ID, &u.Login, &u.Password, &u.Token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNowRows
		}
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (db *Database) FindUserByToken(ctx context.Context, token string) (*User, error) {
	tx, err := db.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var u User
	err = tx.QueryRow(ctx,
		`SELECT id, login, password, token FROM users WHERE token=$1 LIMIT(1)`,
		token).Scan(&u.ID, &u.Login, &u.Password, &u.Token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNowRows
		}
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (db *Database) SaveUserData(ctx context.Context, userID int64, name, dataType string, data []byte) error {
	tx, err := db.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var pgErr *pgconn.PgError

	_, err = tx.Exec(ctx,
		`INSERT INTO user_records(name, data, data_type, user_id) VALUES($1, $2, $3, $4)`,
		name, data, dataType, userID)

	if err != nil {
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return ErrConflict
		}
		return err
	}

	return tx.Commit(ctx)
}

func (db *Database) GetUserData(ctx context.Context, userID int64) ([]ShortRecord, error) {
	tx, err := db.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	rows, err := tx.Query(ctx,
		`SELECT id, name, data_type, version from user_records where user_id=$1`,
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []ShortRecord
	for rows.Next() {
		var rec ShortRecord
		err = rows.Scan(&rec.ID, &rec.Name, &rec.DataType, &rec.Version)
		if err != nil {
			return nil, err
		}
		records = append(records, rec)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (db *Database) FindUserRecord(ctx context.Context, id, userID int64) (*Record, error) {
	tx, err := db.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var rec Record
	err = tx.QueryRow(ctx,
		`SELECT id, name, data, data_type, version FROM user_records where id=$1 AND user_id=$2`,
		id, userID).Scan(&rec.ID, &rec.Name, &rec.Data, &rec.DataType, &rec.Version)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNowRows
		}
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &rec, nil
}
