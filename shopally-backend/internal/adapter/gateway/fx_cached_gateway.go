package gateway

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/shopally-ai/pkg/domain"
)

type CachedFXClient struct {
	Inner  domain.IFXClient
	Cache  domain.ICachePort
	TTL    time.Duration
	Prefix string // optional key prefix, e.g., "fx:"
}

func NewCachedFXClient(inner domain.IFXClient, cache domain.ICachePort, ttl time.Duration) *CachedFXClient {
	return &CachedFXClient{
		Inner:  inner,
		Cache:  cache,
		TTL:    ttl,
		Prefix: "fx:",
	}
}

func (c *CachedFXClient) key(from, to string) string {
	prefix := c.Prefix
	if prefix == "" {
		prefix = "fx:"
	}

	t := strings.ToUpper(strings.TrimSpace(to))
	f := strings.ToUpper(strings.TrimSpace(from))
	return prefix + f + ":" + t
}

func (c *CachedFXClient) GetRate(ctx context.Context, from, to string) (float64, error) {
	f := strings.ToUpper(strings.TrimSpace(from))
	t := strings.ToUpper(strings.TrimSpace(to))
	key := c.key(f, t)

	// 1) Try cache
	if c.Cache != nil {
		if val, ok, err := c.Cache.Get(ctx, key); err == nil && ok {
			if rate, perr := strconv.ParseFloat(val, 64); perr == nil {
				return rate, nil
			}
			// fall through on parse error
		}
	}

	// 2) Cache miss -> fetch from provider
	rate, err := c.Inner.GetRate(ctx, f, t)
	if err != nil {
		return 0, err
	}

	// 3) Write-through
	if c.Cache != nil {
		_ = c.Cache.Set(ctx, key, formatFloat(rate), c.TTL)
	}

	return rate, nil
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 6, 64)
}

var _ domain.IFXClient = (*CachedFXClient)(nil)
