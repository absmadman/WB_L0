package db

import (
	"WB_L0/internal/entities"
	"WB_L0/pkg/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
	//connStr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.DatabaseLogin, cfg.DatabasePassword, cfg.DatabaseName)
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", cfg.DatabaseLogin, cfg.DatabaseName, cfg.DatabasePassword)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (db *Database) CreateDatabase() error {
	_, err := db.connection.Exec(createTable)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (db *Database) InsertDatabase(message string) error {
	_, err := db.connection.Exec(insertTable, message)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (db *Database) SelectById(id int) (*entities.Item, error) {
	var item entities.Item
	err := db.connection.QueryRow(selectById, id).Scan(&item.Id, &item.Message)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &item, nil
}
