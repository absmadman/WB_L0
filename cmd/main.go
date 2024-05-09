package main

import (
	"WB_L0/internal/server"
	"WB_L0/pkg/config"
	"WB_L0/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.NewConfig()
	db := db.NewDatabase(cfg)
	handler := server.NewHandler(db, gin.Default(), cfg)
	err := db.CreateDatabase()
	if err != nil {
		log.Fatal("error creating database")
	}
	handler.Serve()
}
