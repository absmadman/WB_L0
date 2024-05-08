package db

import (
	"WB_L0/pkg/config"
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	connection *sql.DB
}

func NewDatabase(cfg *config.Config) *Database {
	return &Database{
		connection: NewDBConnection(cfg),
	}
}

func NewDBConnection(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.DatabaseLogin, cfg.DatabasePassword, cfg.DatabaseName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
