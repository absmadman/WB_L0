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
		log.Fatalln(fmt.Sprintf("Database: %s", err))
	}
	return db
}

func (db *Database) CreateDatabase() error {
	_, err := db.connection.Exec(createTable)
	if err != nil {
		log.Println(fmt.Sprintf("Database: %s", err))
	}
	return err
}

func (db *Database) InsertDatabase(message []byte) (int, error) {
	id := 0
	err := db.connection.QueryRow(fmt.Sprintf(insertTable, message)).Scan(&id)
	return int(id), err
}

func (db *Database) GetTotal() int {
	count := 0
	db.connection.QueryRow(getTotal).Scan(&count)
	return count
}

func (db *Database) SelectById(id int) (entities.Item, error) {
	var item entities.Item
	err := db.connection.QueryRow(selectById, id).Scan(&item.Id, &item.Message)
	if err != nil {
		return item, err
	}
	return item, nil
}
