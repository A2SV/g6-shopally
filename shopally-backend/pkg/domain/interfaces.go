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
	// SummarizeProduct generates short bullet points for a product based on provided fields.
	SummarizeProduct(context.Context, *Product, string) (*Product, error)
	CompareProducts(ctx context.Context, productDetails []*Product) (*ComparisonResult, error)
}

type IFXClient interface {
	GetRate(ctx context.Context, from, to string) (float64, error)
}

type ICachePort interface {
	// Basic operations
	Get(ctx context.Context, key string) (string, bool, error)
	Set(ctx context.Context, key, val string, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)

	// Object operations
	GetObject(ctx context.Context, key string) (interface{}, bool, error)
	GetTypedObject(ctx context.Context, key string, dest interface{}) (bool, error)
	SetObject(ctx context.Context, key string, value interface{}, ttl time.Duration) error

	// Utility operations
	Expire(ctx context.Context, key string, ttl time.Duration) error
	TTL(ctx context.Context, key string) (time.Duration, error)
	Close() error
}

type AlertRepository interface {
	CreateAlert(alert *Alert) error
	GetAlert(alertID string) (*Alert, error)
	DeleteAlert(alertID string) error
}

type IPushNotificationGateway interface {
	Send(ctx context.Context, token, title, body string, data map[string]string) (string, error)
}
