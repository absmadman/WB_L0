package main

import (
	"WB_L0/internal/server"
	"WB_L0/pkg/cache"
	"WB_L0/pkg/config"
	"WB_L0/pkg/db"
	"WB_L0/pkg/nats"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.NewConfig()

	db := db.NewDatabase(cfg)
	cache := cache.NewCache(cfg)

	streaming, err := nats.NewStreaming(cfg, db, cache)
	if err != nil {
		log.Fatalln("Nats: error connect nats streaming")
	}

	handler := server.NewHandler(db, gin.Default(), cache, cfg)
	handler.InitCache(cfg)

	err = db.CreateDatabase()
	if err != nil {
		log.Fatalln("Database: error creating table")
	}

	streaming.Sub()

	handler.Serve()
}
