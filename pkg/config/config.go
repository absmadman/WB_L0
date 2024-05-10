package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

// структура для конфигурации различных параметров сервиса
type Config struct {
	DatabaseName     string `yaml:"DATABASE_NAME"`
	DatabaseLogin    string `yaml:"DATABASE_LOGIN"`
	DatabasePassword string `yaml:"DATABASE_PASS"`
	ClientId         string `yaml:"CLIENT_ID"`
	ClusterId        string `yaml:"CLUSTER_ID"`
	CacheSize        int    `yaml:"CACHE_SIZE"`
}

// конструктор для структуры Config, конфиг читается из файла config.yml
func NewConfig() (*Config, error) {
	var cfg Config
	file, err := os.ReadFile("../config/config.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
