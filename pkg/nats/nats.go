package nats

import (
	"WB_L0/internal/entities"
	"WB_L0/pkg/cache"
	"WB_L0/pkg/config"
	"WB_L0/pkg/db"
	"github.com/nats-io/stan.go"
	"log"
)

type Streaming struct {
	connection stan.Conn
	database   *db.Database
	cache      *cache.Cache
}

func NewStreaming(cfg *config.Config, db *db.Database, cache *cache.Cache) (*Streaming, error) {
	conn, err := stan.Connect(cfg.ClusterId, cfg.ClientId)
	if err != nil {
		return nil, err
	}
	return &Streaming{
		connection: conn,
		database:   db,
		cache:      cache,
	}, nil
}

func (s *Streaming) Sub() {
	s.connection.Subscribe("post", func(m *stan.Msg) {
		if !entities.IsValid(m.Data) {
			log.Println("Nats: invalid received data")
			return
		}
		id, err := s.database.InsertDatabase(m.Data)
		if err != nil {
			log.Println("Database: error indexing data")
			return
		}
		s.cache.Cache.Add(id, m.Data)
	}, stan.DurableName("my-durable"))
}
