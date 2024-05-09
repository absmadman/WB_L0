package cache

import (
	"WB_L0/pkg/config"
	lru "github.com/hashicorp/golang-lru/v2"
)

type Cache struct {
	Cache *lru.Cache[int, []byte]
}

func NewCache(cfg *config.Config) *Cache {
	tmp, _ := lru.New[int, []byte](cfg.CacheSize)
	return &Cache{
		Cache: tmp,
	}
}
