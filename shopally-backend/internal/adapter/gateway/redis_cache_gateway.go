package gateway

import (
	"context"
	"encoding/json"
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
		prefix = "sa:" // default namespace
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

// Set stores a value with TTL
func (c *RedisCache) Set(ctx context.Context, key, val string, ttl time.Duration) error {
	if ttl < 0 {
		ttl = 0
	}
	return c.client.Set(ctx, c.key(key), val, ttl).Err()
}

// Delete removes a key from cache
func (c *RedisCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, c.key(key)).Err()
}

// Exists checks if a key exists in cache
func (c *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	result, err := c.client.Exists(ctx, c.key(key)).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

// GetObject retrieves and returns the unmarshaled object as interface{}
func (c *RedisCache) GetObject(ctx context.Context, key string) (interface{}, bool, error) {
	val, found, err := c.Get(ctx, key)
	if err != nil || !found {
		return nil, found, err
	}

	var result interface{}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, true, err
	}
	return result, true, nil
}

// GetTypedObject unmarshals into the provided destination interface
func (c *RedisCache) GetTypedObject(ctx context.Context, key string, dest interface{}) (bool, error) {
	val, found, err := c.Get(ctx, key)
	if err != nil {
		return false, err
	}
	if !found {
		return false, nil
	}

	err = json.Unmarshal([]byte(val), dest)
	if err != nil {
		return true, err
	}
	return true, nil
}

// SetObject marshals an object to JSON and stores it in cache
func (c *RedisCache) SetObject(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(ctx, key, string(jsonData), ttl)
}

// Expire sets expiration time for a key
func (c *RedisCache) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return c.client.Expire(ctx, c.key(key), ttl).Err()
}

// TTL gets the time-to-live for a key
func (c *RedisCache) TTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := c.client.TTL(ctx, c.key(key)).Result()
	if err != nil {
		return 0, err
	}
	return ttl, nil
}

func (c *RedisCache) ClearAllWithPrefix(ctx context.Context) error {
	// Use SCAN to find all keys with the prefix and delete them in batches
	iter := c.client.Scan(ctx, 0, c.prefix+"*", 0).Iterator()

	var keys []string
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
		// Delete in batches of 1000 to avoid memory issues
		if len(keys) >= 1000 {
			if err := c.client.Del(ctx, keys...).Err(); err != nil {
				return err
			}
			keys = []string{}
		}
	}

	if err := iter.Err(); err != nil {
		return err
	}

	// Delete any remaining keys
	if len(keys) > 0 {
		return c.client.Del(ctx, keys...).Err()
	}

	return nil
}

// Close closes the Redis connection
func (c *RedisCache) Close() error {
	return c.client.Close()
}

// Ensure interface compliance
var _ domain.ICachePort = (*RedisCache)(nil)
