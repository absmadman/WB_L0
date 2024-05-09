package config

import (
	"os"
)

type Config struct {
	DatabaseName     string
	DatabaseLogin    string
	DatabasePassword string
	Client           string
	Server           string
}

func NewConfig() *Config {
	return &Config{
		/*
			DatabaseName:     os.Getenv("DB_NAME"),
			DatabaseLogin:    os.Getenv("DB_LOGIN"),
			DatabasePassword: os.Getenv("DB_PASSWORD"),
		*/
		DatabaseName:     "postgres",
		DatabaseLogin:    "postgres",
		DatabasePassword: "postgres",
		Server:           os.Getenv("NATS_SERVER"),
		Client:           os.Getenv("NATS_CLIENT"),
	}
}
