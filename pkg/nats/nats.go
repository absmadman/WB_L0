package nats

import (
	"WB_L0/pkg/config"
	"fmt"
	"github.com/nats-io/stan.go"
)

type Streaming struct {
	connection stan.Conn
}

func NewStreaming(cfg *config.Config) (*Streaming, error) {
	//conn, err := stan.Connect(cfg.Server, cfg.Client)
	conn, err := stan.Connect("server", "server")
	//conn, err := stan.Connect("server", "1")
	if err != nil {
		return nil, err
	}

	return &Streaming{
		connection: conn,
	}, nil
}

func (s *Streaming) Sub() {
	s.connection.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}, stan.DurableName("my-durable"))
}

func (s *Streaming) Pub() {
	s.connection.Publish("foo", []byte("Hello world"))
}
