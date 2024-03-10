package database

import (
	"assignment-02/config"
	"database/sql"
	_ "embed"
	"fmt"

	_ "github.com/lib/pq"
)

func New(conf config.DB) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", conf.Username, conf.Password, conf.Host, conf.Port, conf.Name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

//go:embed db.sql
var q string

func createTables(db *sql.DB) error {
	_, err := db.Exec(q)
	return err
}
