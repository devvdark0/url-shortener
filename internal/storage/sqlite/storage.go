package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/devvdark0/url-shortener/internal/storage"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(ctx context.Context, urlToSave, alias string) (int64, error) {
	const op = "storage.sqlite.SaveURL"

	stmt, err := s.db.Prepare(`INSERT INTO urls(alias, url) VALUES(?, ?)`)
	if err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	res, err := stmt.ExecContext(ctx, alias, urlToSave)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return 0, fmt.Errorf("%s:%w", op, storage.ErrUrlExists)
		}

		return 0, fmt.Errorf("%s:%w", op, err)
	}

	urlId, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	return urlId, nil
}

func (s *Storage) GetURL(ctx context.Context, alias string) (string, error) {
	const op = "storage.sqlite.GetURL"

	stmt, err := s.db.Prepare(`SELECT url FROM urls WHERE alias=?`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s:%w", op, storage.ErrUrlNotFound)
		}

		return "", fmt.Errorf("%s:%w", op, err)
	}
	row := stmt.QueryRowContext(ctx, alias)

	var url string
	if err := row.Scan(&url); err != nil {
		return "", fmt.Errorf("%s:%w", op, err)
	}

	return url, nil
}

func (s *Storage) DeleteURL(ctx context.Context, alias string) error {
	const op = "storage.sqlite.DeleteURL"

	stmt, err := s.db.Prepare(`DELETE FROM urls WHERE alias=?`)
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	if _, err = stmt.ExecContext(ctx, alias); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s:%w", op, storage.ErrUrlNotFound)
		}

		return fmt.Errorf("%s:%w", op, err)
	}

	return nil
}
