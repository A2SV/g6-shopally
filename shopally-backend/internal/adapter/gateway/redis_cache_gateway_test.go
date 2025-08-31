package gateway

import (
	"context"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type RedisCacheSuite struct {
	suite.Suite
	ctx    context.Context
	mr     *miniredis.Miniredis
	client *redis.Client
	cache  *RedisCache
}

func (s *RedisCacheSuite) SetupTest() {
	s.ctx = context.Background()
	mr, err := miniredis.Run()
	s.Require().NoError(err)
	s.mr = mr
	s.client = redis.NewClient(&redis.Options{Addr: mr.Addr()})
	s.cache = NewRedisCache(s.client, "sa:")
}

func (s *RedisCacheSuite) TearDownTest() {
	_ = s.client.Close()
	s.mr.Close()
}

func (s *RedisCacheSuite) TestGetMissThenSetThenGetHit() {
	val, ok, err := s.cache.Get(s.ctx, "foo")
	s.NoError(err)
	s.False(ok)
	s.Equal("", val)

	s.Require().NoError(s.cache.Set(s.ctx, "foo", "bar", time.Minute))

	val, ok, err = s.cache.Get(s.ctx, "foo")
	s.NoError(err)
	s.True(ok)
	s.Equal("bar", val)
}

func (s *RedisCacheSuite) TestTTLExpiry() {
	s.Require().NoError(s.cache.Set(s.ctx, "temp", "1", 10*time.Second))
	s.True(s.mr.TTL("sa:temp") > 0)
}

func TestRedisCacheSuite(t *testing.T) { suite.Run(t, new(RedisCacheSuite)) }
