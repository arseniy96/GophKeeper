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
	TimeOut        = 3 * time.Second
	DefaultVersion = 1
)

type Database struct {
	DB *pgx.Conn
	l  *logger.Logger
}

func NewStorage(dsn string, l *logger.Logger) (*Database, error) {
	if err := runMigrations(dsn); err != nil {
		return nil, fmt.Errorf("%w: Init storage error: %v", ErrMigrationsFailed, err)
	}

	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("%w: Connect storage error: %v", ErrConnectionRefused, err)
	}
	database := &Database{
		DB: db,
		l:  l,
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

func (db *Database) CreateUser(ctx context.Context, login, password string) (int64, error) {
	var pgErr *pgconn.PgError
	var id int64
	err := db.DB.QueryRow(ctx,
		`INSERT INTO users(login, password) VALUES($1, $2) RETURNING id`,
		login, password).Scan(&id)
	if err != nil {
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return 0, ErrConflict
		}
		return 0, fmt.Errorf("%w: Create user error: %v", ErrCreateUser, err)
	}

	return id, nil
}

func (db *Database) FindUserByLogin(ctx context.Context, login string) (*User, error) {
	var u User
	err := db.DB.QueryRow(ctx,
		`SELECT id, login, password FROM users WHERE login=$1 LIMIT(1)`,
		login).Scan(&u.ID, &u.Login, &u.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNowRows
		}

		return nil, fmt.Errorf("%w: Find user error: %v", ErrFindUser, err)
	}

	return &u, nil
}

func (db *Database) SaveUserData(ctx context.Context, userID int64, name, dataType string, data []byte) error {
	var pgErr *pgconn.PgError

	_, err := db.DB.Exec(ctx,
		`INSERT INTO user_records(name, data, data_type, user_id) VALUES($1, $2, $3, $4)`,
		name, data, dataType, userID)

	if err != nil {
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return ErrConflict
		}
		return fmt.Errorf("%w: Save user data error: %v", ErrSaveUserData, err)
	}

	return nil
}

func (db *Database) GetUserData(ctx context.Context, userID int64) ([]ShortRecord, error) {
	rows, err := db.DB.Query(ctx,
		`SELECT id, name, data_type, version from user_records where user_id=$1`,
		userID)
	if err != nil {
		return nil, fmt.Errorf("%w: Request error: %v", ErrGetUserData, err)
	}
	defer rows.Close()

	var records []ShortRecord
	for rows.Next() {
		var rec ShortRecord
		err = rows.Scan(&rec.ID, &rec.Name, &rec.DataType, &rec.Version)
		if err != nil {
			return nil, fmt.Errorf("%w: Scan error: %v", ErrGetUserData, err)
		}
		records = append(records, rec)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%w: Internal error: %v", ErrGetUserData, err)
	}

	return records, nil
}

func (db *Database) FindUserRecord(ctx context.Context, id, userID int64) (*Record, error) {
	var rec Record
	err := db.DB.QueryRow(ctx,
		`SELECT id, name, data, data_type, version FROM user_records where id=$1 AND user_id=$2`,
		id, userID).Scan(&rec.ID, &rec.Name, &rec.Data, &rec.DataType, &rec.Version)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNowRows
		}
		return nil, fmt.Errorf("%w: Find user record error: %v", ErrFindUserRecord, err)
	}

	return &rec, nil
}

func (db *Database) UpdateUserRecord(ctx context.Context, rec *Record) error {
	_, err := db.DB.Exec(ctx, `UPDATE user_records SET data=$1, version=version+1 WHERE id=$2`, rec.Data, rec.ID)
	if err != nil {
		return fmt.Errorf("%w: Update user record error: %v", ErrUpdateUserRecord, err)
	}

	return nil
}
