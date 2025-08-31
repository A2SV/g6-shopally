package gateway

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/shopally-ai/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CachedFXClientSuite struct {
	suite.Suite
	ctx   context.Context
	fx    *mocks.IFXClient
	cache *mocks.ICachePort
	c     *CachedFXClient
}

func (s *CachedFXClientSuite) SetupTest() {
	s.ctx = context.Background()
	s.fx = mocks.NewIFXClient(s.T())
	s.cache = mocks.NewICachePort(s.T())
	s.c = &CachedFXClient{Inner: s.fx, Cache: s.cache, TTL: time.Minute, Prefix: "fx:"}
}

func (s *CachedFXClientSuite) TestMissThenHit() {
	key := "fx:USD:ETB"
	s.cache.On("Get", s.ctx, key).Return("", false, nil).Once()
	s.fx.On("GetRate", s.ctx, "USD", "ETB").Return(56.123456, nil).Once()
	s.cache.On("Set", s.ctx, key, "56.123456", time.Minute).Return(nil).Once()

	rate1, err1 := s.c.GetRate(s.ctx, "usd", "etb")
	s.Require().NoError(err1)
	s.InDelta(56.123456, rate1, 1e-6)

	s.cache.On("Get", s.ctx, key).Return("56.123456", true, nil).Once()
	rate2, err2 := s.c.GetRate(s.ctx, "USD", "ETB")
	s.Require().NoError(err2)
	s.InDelta(56.123456, rate2, 1e-6)
}

func (s *CachedFXClientSuite) TestCacheHitParseErrorFallsThrough() {
	key := "fx:USD:ETB"
	// bad cached value -> fall through to provider
	s.cache.On("Get", s.ctx, key).Return("not-a-number", true, nil).Once()
	s.fx.On("GetRate", s.ctx, "USD", "ETB").Return(57.5, nil).Once()
	s.cache.On("Set", s.ctx, key, "57.500000", time.Minute).Return(nil).Once()

	rate, err := s.c.GetRate(s.ctx, "USD", "ETB")
	s.Require().NoError(err)
	s.InDelta(57.5, rate, 1e-6)
}

func (s *CachedFXClientSuite) TestCacheErrorFallsThrough() {
	key := "fx:USD:ETB"
	// cache error -> treat as miss and continue
	s.cache.On("Get", s.ctx, key).Return("", false, errors.New("boom")).Once()
	s.fx.On("GetRate", s.ctx, "USD", "ETB").Return(60.25, nil).Once()
	s.cache.On("Set", s.ctx, key, "60.250000", time.Minute).Return(nil).Once()

	rate, err := s.c.GetRate(s.ctx, "USD", "ETB")
	s.Require().NoError(err)
	s.InDelta(60.25, rate, 1e-6)
}

func (s *CachedFXClientSuite) TestProviderError() {
	key := "fx:USD:ETB"
	s.cache.On("Get", s.ctx, key).Return("", false, nil).Once()
	s.fx.On("GetRate", s.ctx, "USD", "ETB").Return(0.0, errors.New("provider down")).Once()

	rate, err := s.c.GetRate(s.ctx, "USD", "ETB")
	s.Error(err)
	s.Equal(0.0, rate)
}

func TestCachedFXClientSuite(t *testing.T) { suite.Run(t, new(CachedFXClientSuite)) }

// quick unit for formatFloat
func TestFormatFloat(t *testing.T) {
	assert.Equal(t, "1.230000", formatFloat(1.23))
}
