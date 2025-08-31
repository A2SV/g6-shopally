package domain

import (
	"context"
	"time"
)

// AlibabaGateway defines the contract for fetching products from an external source.
type AlibabaGateway interface {
	FetchProducts(ctx context.Context, query string, filters map[string]interface{}) ([]*Product, error)
}

// LLMGateway defines the contract for a Large Language Model service
// to parse user intent from a search query.
type LLMGateway interface {
	ParseIntent(ctx context.Context, query string) (map[string]interface{}, error)
}

// CacheGateway defines the contract for a caching service.
type CacheGateway interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type IFXClient interface {
	GetRate(ctx context.Context, from, to string) (float64, error)
}

type ICachePort interface {
	// Get returns the value, whether it was found, and any error.
	Get(ctx context.Context, key string) (string, bool, error)
	// Set stores the value with a TTL; use 0 for no expiration.
	Set(ctx context.Context, key, val string, ttl time.Duration) error
}

type AlertRepository interface {
	CreateAlert(alert *Alert) error
	GetAlert(alertID string) (*Alert, error)
	DeleteAlert(alertID string) error
}
