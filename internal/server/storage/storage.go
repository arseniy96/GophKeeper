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

func NewStorage(dsn string, l *logger.Logger) (*Database, error) {
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

	return db.DB.PingContext(ctx)
}

func (db *Database) Close() error {
	return db.DB.Close()
}

func (db *Database) CreateUser(ctx context.Context, login, password string) error {
	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var pgErr *pgconn.PgError
	_, err = tx.ExecContext(ctx,
		`INSERT INTO users(login, password) VALUES($1, $2)`,
		login, password)
	if err != nil {
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return ErrConflict
		}
		return err
	}

	return tx.Commit()
}

func (db *Database) UpdateUserToken(ctx context.Context, login, token string) error {
	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		`UPDATE users SET token=$1 WHERE login=$2`,
		token, login)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (db *Database) FindUserByLogin(ctx context.Context, login string) (*User, error) {
	tx, err := db.DB.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var u User
	err = tx.QueryRowContext(ctx,
		`SELECT id, login, password, token FROM users WHERE login=$1 LIMIT(1)`,
		login).Scan(&u.ID, &u.Login, &u.Password, &u.Token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNowRows
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (db *Database) FindUserByToken(ctx context.Context, token string) (*User, error) {
	tx, err := db.DB.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var u User
	err = tx.QueryRowContext(ctx,
		`SELECT id, login, password, token FROM users WHERE token=$1 LIMIT(1)`,
		token).Scan(&u.ID, &u.Login, &u.Password, &u.Token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNowRows
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (db *Database) SaveUserData(ctx context.Context, userID int64, name, dataType string, data []byte) error {
	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var pgErr *pgconn.PgError

	_, err = tx.ExecContext(ctx,
		`INSERT INTO user_records(name, data, data_type, user_id) VALUES($1, $2, $3, $4)`,
		name, data, dataType, userID)

	if err != nil {
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return ErrConflict
		}
		return err
	}

	return tx.Commit()
}

func (db *Database) GetUserData(ctx context.Context, userID int64) ([]ShortRecord, error) {
	tx, err := db.DB.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx,
		`SELECT id, name, data_type, version, created_at from user_records where user_id=$1`,
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []ShortRecord
	for rows.Next() {
		var rec ShortRecord
		err = rows.Scan(&rec.ID, &rec.Name, &rec.DataType, &rec.Version, &rec.CreatedAt)
		if err != nil {
			return nil, err
		}
		records = append(records, rec)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (db *Database) FindUserRecord(ctx context.Context, id, userID int64) (*Record, error) {
	tx, err := db.DB.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var rec Record
	err = tx.QueryRowContext(ctx,
		`SELECT id, name, data, data_type, version, created_at FROM user_records where id=$1 AND user_id=$2`,
		id, userID).Scan(&rec.ID, &rec.Name, &rec.Data, &rec.DataType, &rec.Version, &rec.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNowRows
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &rec, nil
}
