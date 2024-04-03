package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const DEFAULT_CONFIG_ID = 0

type Database struct {
	_db *sql.DB
	Path string
}

func (db *Database) Open() error {
	sqlDb, err := sql.Open("sqlite3", db.Path)
	if err != nil {
		return err
	}

	db._db = sqlDb
	return nil
}

func (db *Database) Close() error {
	return db._db.Close()
}
