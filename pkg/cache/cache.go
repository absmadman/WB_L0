package cache

import (
	"WB_L0/internal/entities"
	"WB_L0/pkg/config"
	lru "github.com/hashicorp/golang-lru/v2"
)

type Cache struct {
	Cache *lru.Cache[int, *entities.Item]
}

func NewCache(cfg config.Config) *Cache {
	return &Cache{
		Cache: lru.New[int, string](cfg.CacheSize),
	}
}
