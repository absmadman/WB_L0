package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

type Streaming struct {
	connection stan.Conn
}

func NewStreaming() (*Streaming, error) {
	conn, err := stan.Connect("test-cluster", "publisher")
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

func (s *Streaming) PubValid() {
	data, _ := os.ReadFile("../data/model.json")
	s.connection.Publish("post", data)
}

func (s *Streaming) PubInvalid() {
	data, _ := os.ReadFile("../data/invalid_model.json")
	s.connection.Publish("post", data)
}

func main() {
	str, err := NewStreaming()
	if err != nil {
		log.Fatalln("Error while streaming connection")
	}
	str.PubValid()
	str.PubInvalid()
}
