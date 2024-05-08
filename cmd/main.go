package main

import (
	"WB_L0/internal/handler"
	"WB_L0/pkg/config"
	"WB_L0/pkg/nats"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.NewConfig()
	stream, err := nats.NewStreaming(cfg)
	if err != nil {
		log.Println("ERROr")
	}
	stream.Sub()
	stream.Pub()
	server := handler.NewHandler(gin.Default())
	server.Serve()
}
