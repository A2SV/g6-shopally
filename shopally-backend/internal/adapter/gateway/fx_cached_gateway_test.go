// internal/adapter/gateway/fx_cached_gateway_test.go
package gateway

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// MockIFXClient implements domain.IFXClient for testing
type MockIFXClient struct {
	GetRateFunc func(ctx context.Context, from, to string) (float64, error)
}

func (m *MockIFXClient) GetRate(ctx context.Context, from, to string) (float64, error) {
	if m.GetRateFunc != nil {
		return m.GetRateFunc(ctx, from, to)
	}
	return 55.0, nil // Default rate
}

// MockICachePort implements domain.ICachePort for testing
type MockICachePort struct {
	store map[string]cacheItem
}

type cacheItem struct {
	value      string
	expiration time.Time
}

func NewMockICachePort() *MockICachePort {
	return &MockICachePort{
		store: make(map[string]cacheItem),
	}
}

func (m *MockICachePort) Get(ctx context.Context, key string) (string, bool, error) {
	item, exists := m.store[key]
	if !exists {
		return "", false, nil
	}
	if time.Now().After(item.expiration) {
		return "", false, nil
	}
	return item.value, true, nil
}

func (m *MockICachePort) Set(ctx context.Context, key, val string, ttl time.Duration) error {
	m.store[key] = cacheItem{
		value:      val,
		expiration: time.Now().Add(ttl),
	}
	return nil
}

func (m *MockICachePort) Delete(ctx context.Context, key string) error {
	delete(m.store, key)
	return nil
}

func (m *MockICachePort) Exists(ctx context.Context, key string) (bool, error) {
	_, exists := m.store[key]
	return exists, nil
}

func (m *MockICachePort) GetObject(ctx context.Context, key string) (interface{}, bool, error) {
	val, found, err := m.Get(ctx, key)
	if err != nil || !found {
		return nil, found, err
	}
	return val, true, nil
}

func (m *MockICachePort) GetTypedObject(ctx context.Context, key string, dest interface{}) (bool, error) {
	val, found, err := m.Get(ctx, key)
	if err != nil {
		return false, err
	}
	if !found {
		return false, nil
	}
	// Simple implementation for testing
	if strPtr, ok := dest.(*string); ok {
		*strPtr = val
		return true, nil
	}
	return true, nil
}

func (m *MockICachePort) SetObject(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if strVal, ok := value.(string); ok {
		return m.Set(ctx, key, strVal, ttl)
	}
	return m.Set(ctx, key, "test-object", ttl)
}

func (m *MockICachePort) Expire(ctx context.Context, key string, ttl time.Duration) error {
	if item, exists := m.store[key]; exists {
		item.expiration = time.Now().Add(ttl)
		m.store[key] = item
	}
	return nil
}

func (m *MockICachePort) TTL(ctx context.Context, key string) (time.Duration, error) {
	if item, exists := m.store[key]; exists {
		return time.Until(item.expiration), nil
	}
	return -1, nil
}

func (m *MockICachePort) Close() error {
	m.store = make(map[string]cacheItem)
	return nil
}

// Test suite
type CachedFXClientSuite struct {
	suite.Suite
	client *CachedFXClient
	inner  *MockIFXClient
	cache  *MockICachePort
}

func (s *CachedFXClientSuite) SetupTest() {
	s.inner = &MockIFXClient{}
	s.cache = NewMockICachePort()
	s.client = NewCachedFXClient(s.inner, s.cache, time.Hour)
}

func (s *CachedFXClientSuite) TearDownTest() {
	s.cache.Close()
}

func (s *CachedFXClientSuite) TestCacheMissThenHit() {
	// Set up the inner client to return a specific rate
	s.inner.GetRateFunc = func(ctx context.Context, from, to string) (float64, error) {
		return 55.0, nil
	}

	// First call - should miss cache and call inner client
	rate1, err := s.client.GetRate(context.Background(), "USD", "ETB")
	s.NoError(err)
	s.Equal(55.0, rate1)

	// Change the inner client to return a different rate to prove cache is working
	s.inner.GetRateFunc = func(ctx context.Context, from, to string) (float64, error) {
		return 99.0, nil // Different rate
	}

	// Second call - should hit cache and return the original rate
	rate2, err := s.client.GetRate(context.Background(), "USD", "ETB")
	s.NoError(err)
	s.Equal(55.0, rate2) // Should be cached value, not 99.0
}

func (s *CachedFXClientSuite) TestCacheWithInvalidData() {
	// Put invalid data in cache
	err := s.cache.Set(context.Background(), "fx:USD:ETB", "not-a-float", time.Hour)
	s.NoError(err)

	// Set up inner client to return valid rate
	s.inner.GetRateFunc = func(ctx context.Context, from, to string) (float64, error) {
		return 55.0, nil
	}

	// Should ignore invalid cache data and fall back to inner client
	rate, err := s.client.GetRate(context.Background(), "USD", "ETB")
	s.NoError(err)
	s.Equal(55.0, rate)
}

func (s *CachedFXClientSuite) TestInnerClientError() {
	// Set up inner client to return error
	s.inner.GetRateFunc = func(ctx context.Context, from, to string) (float64, error) {
		return 0, errors.New("network error")
	}

	// Should propagate the error
	_, err := s.client.GetRate(context.Background(), "USD", "ETB")
	s.Error(err)
	s.Contains(err.Error(), "network error")
}

func (s *CachedFXClientSuite) TestCacheKeyFormat() {
	s.inner.GetRateFunc = func(ctx context.Context, from, to string) (float64, error) {
		return 55.0, nil
	}

	// Call with different currencies
	_, err := s.client.GetRate(context.Background(), "USD", "ETB")
	s.NoError(err)

	// Check that cache key was formatted correctly
	val, found, err := s.cache.Get(context.Background(), "fx:USD:ETB")
	s.NoError(err)
	s.True(found)
	s.Equal("55.000000", val)
}

func (s *CachedFXClientSuite) TestCaseInsensitiveCurrencyHandling() {
	s.inner.GetRateFunc = func(ctx context.Context, from, to string) (float64, error) {
		return 55.0, nil
	}

	// Test with different case combinations
	_, err := s.client.GetRate(context.Background(), "usd", "etb")
	s.NoError(err)

	_, err = s.client.GetRate(context.Background(), "USD", "etb")
	s.NoError(err)

	_, err = s.client.GetRate(context.Background(), "usd", "ETB")
	s.NoError(err)

	// All should use the same cache key
	val, found, err := s.cache.Get(context.Background(), "fx:USD:ETB")
	s.NoError(err)
	s.True(found)
	s.Equal("55.000000", val)
}

func (s *CachedFXClientSuite) TestDifferentCurrencyPairs() {
	s.inner.GetRateFunc = func(ctx context.Context, from, to string) (float64, error) {
		if from == "USD" && to == "ETB" {
			return 55.0, nil
		}
		if from == "EUR" && to == "USD" {
			return 1.1, nil
		}
		return 0, errors.New("unknown currency pair")
	}

	// Test USD to ETB
	rate1, err := s.client.GetRate(context.Background(), "USD", "ETB")
	s.NoError(err)
	s.Equal(55.0, rate1)

	// Test EUR to USD
	rate2, err := s.client.GetRate(context.Background(), "EUR", "USD")
	s.NoError(err)
	s.Equal(1.1, rate2)

	// Verify different cache keys
	val1, found1, _ := s.cache.Get(context.Background(), "fx:USD:ETB")
	val2, found2, _ := s.cache.Get(context.Background(), "fx:EUR:USD")

	s.True(found1)
	s.True(found2)
	s.Equal("55.000000", val1)
	s.Equal("1.100000", val2)
}

func (s *CachedFXClientSuite) TestCacheExpiration() {
	s.inner.GetRateFunc = func(ctx context.Context, from, to string) (float64, error) {
		return 55.0, nil
	}

	// Create client with very short TTL
	shortTTLClient := NewCachedFXClient(s.inner, s.cache, time.Millisecond*100)

	// First call
	rate1, err := shortTTLClient.GetRate(context.Background(), "USD", "ETB")
	s.NoError(err)
	s.Equal(55.0, rate1)

	// Wait for cache to expire
	time.Sleep(time.Millisecond * 150)

	// Change inner client to return different rate
	s.inner.GetRateFunc = func(ctx context.Context, from, to string) (float64, error) {
		return 66.0, nil
	}

	// Should call inner client again due to expiration
	rate2, err := shortTTLClient.GetRate(context.Background(), "USD", "ETB")
	s.NoError(err)
	s.Equal(66.0, rate2)
}

func (s *CachedFXClientSuite) TestNilCache() {
	// Create client with nil cache (should still work, just no caching)
	noCacheClient := NewCachedFXClient(s.inner, nil, time.Hour)

	s.inner.GetRateFunc = func(ctx context.Context, from, to string) (float64, error) {
		return 55.0, nil
	}

	// First call
	rate1, err := noCacheClient.GetRate(context.Background(), "USD", "ETB")
	s.NoError(err)
	s.Equal(55.0, rate1)

	// Change inner client to return different rate
	s.inner.GetRateFunc = func(ctx context.Context, from, to string) (float64, error) {
		return 66.0, nil
	}

	// Second call - should call inner client again (no caching)
	rate2, err := noCacheClient.GetRate(context.Background(), "USD", "ETB")
	s.NoError(err)
	s.Equal(66.0, rate2)
}

func TestCachedFXClientSuite(t *testing.T) {
	suite.Run(t, new(CachedFXClientSuite))
}
