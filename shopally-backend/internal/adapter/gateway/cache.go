package gateway

import (
	"github.com/redis/go-redis/v9"
	"github.com/shopally-ai/pkg/domain"
)

type CacheType string

const (
	CacheTypeRedis  CacheType = "redis"
	CacheTypeMemory CacheType = "memory"
)

type CacheConfig struct {
	Type     CacheType
	RedisURL string
	Prefix   string
}

func NewCache(config CacheConfig) (domain.ICachePort, error) {
	switch config.Type {
	case CacheTypeRedis:
		opts, err := redis.ParseURL(config.RedisURL)
		if err != nil {
			return nil, err
		}
		client := redis.NewClient(opts)
		return NewRedisCache(client, config.Prefix), nil

	case CacheTypeMemory:
		fallthrough
	default:
		return NewInMemoryCache(), nil
	}
}
