package main

import (
	"WB_L0/internal/server"
	"WB_L0/pkg/cache"
	"WB_L0/pkg/config"
	"WB_L0/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.NewConfig()
	db := db.NewDatabase(cfg)
	cache := cache.NewCache(cfg)
	handler := server.NewHandler(db, gin.Default(), cache, cfg)
	handler.InitCache(cfg)
	err := db.CreateDatabase()
	if err != nil {
		log.Fatal("error creating database")
	}
	handler.Serve()
}
