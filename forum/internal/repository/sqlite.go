package repository

import (
	"database/sql"
)

const (
	postTable = "post"
	categoryTable = "category"
	postCategoryTable = "post_category"
)

type Config struct {
	DSN string
}

func NewSQLiteDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.DSN)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
