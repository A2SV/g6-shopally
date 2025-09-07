package gateway

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/shopally-ai/pkg/domain"
)

type InMemoryCache struct {
	store map[string]memoryCacheItem
	mu    sync.RWMutex
}

type memoryCacheItem struct {
	value      string
	expiration time.Time
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		store: make(map[string]memoryCacheItem),
	}
}

// Get returns value, found, error
func (c *InMemoryCache) Get(ctx context.Context, key string) (string, bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.store[key]
	if !exists {
		return "", false, nil
	}

	if !item.expiration.IsZero() && time.Now().After(item.expiration) {
		return "", false, nil // Expired
	}

	return item.value, true, nil
}

// Set stores a value with TTL
func (c *InMemoryCache) Set(ctx context.Context, key, val string, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var expiration time.Time
	if ttl > 0 {
		expiration = time.Now().Add(ttl)
	}

	c.store[key] = memoryCacheItem{
		value:      val,
		expiration: expiration,
	}
	return nil
}

// Delete removes a key from cache
func (c *InMemoryCache) Delete(ctx context.Context, key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.store, key)
	return nil
}

// Exists checks if a key exists in cache
func (c *InMemoryCache) Exists(ctx context.Context, key string) (bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.store[key]
	if !exists {
		return false, nil
	}

	if !item.expiration.IsZero() && time.Now().After(item.expiration) {
		return false, nil // Expired
	}

	return true, nil
}

// GetObject retrieves and returns the unmarshaled object as interface{}
func (c *InMemoryCache) GetObject(ctx context.Context, key string) (interface{}, bool, error) {
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
func (c *InMemoryCache) GetTypedObject(ctx context.Context, key string, dest interface{}) (bool, error) {
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
func (c *InMemoryCache) SetObject(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(ctx, key, string(jsonData), ttl)
}

// Expire sets expiration time for a key
func (c *InMemoryCache) Expire(ctx context.Context, key string, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, exists := c.store[key]
	if !exists {
		return nil // Key doesn't exist
	}

	item.expiration = time.Now().Add(ttl)
	c.store[key] = item
	return nil
}

// TTL gets the time-to-live for a key
func (c *InMemoryCache) TTL(ctx context.Context, key string) (time.Duration, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.store[key]
	if !exists {
		return -2 * time.Second, nil // Key doesn't exist (mimics Redis)
	}

	if item.expiration.IsZero() {
		return -1 * time.Second, nil // No expiration set (mimics Redis)
	}

	remaining := time.Until(item.expiration)
	if remaining < 0 {
		return -2 * time.Second, nil // Expired
	}

	return remaining, nil
}

// Close clears the cache and releases resources
func (c *InMemoryCache) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store = make(map[string]memoryCacheItem)
	return nil
}

// Ensure interface compliance
var _ domain.ICachePort = (*InMemoryCache)(nil)
