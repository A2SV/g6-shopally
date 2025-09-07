package gateway

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shopally-ai/pkg/domain"
)

type RedisCache struct {
	client *redis.Client
	prefix string
}

func NewRedisCache(client *redis.Client, prefix string) *RedisCache {
	if prefix == "" {
		prefix = "sa:" //default namespace
	}
	return &RedisCache{
		client: client,
		prefix: prefix,
	}
}

func (c *RedisCache) key(k string) string {
	return c.prefix + k
}

// Get returns value, found, error
func (c *RedisCache) Get(ctx context.Context, key string) (string, bool, error) {
	val, err := c.client.Get(ctx, c.key(key)).Result()
	if err == redis.Nil {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}

	return val, true, nil
}

func (c *RedisCache) Set(ctx context.Context, key, val string, ttl time.Duration) error {
	if ttl < 0 {
		ttl = 0
	}
	return c.client.Set(ctx, c.key(key), val, ttl).Err()
}

// Ensure interface compliance at compile time
var _ domain.ICachePort = (*RedisCache)(nil)
