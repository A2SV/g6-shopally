package util

import (
	"context"
	"errors"
	"testing"

	"github.com/shopally-ai/internal/mocks"
	"github.com/stretchr/testify/suite"
)

type FXUtilSuite struct {
	suite.Suite
	ctx   context.Context
	cache *mocks.ICachePort
}

func (s *FXUtilSuite) SetupTest() {
	s.ctx = context.Background()
	s.cache = mocks.NewICachePort(s.T())
}

func (s *FXUtilSuite) TestUSDToETB_Success() {
	s.cache.On("Get", s.ctx, FXKeyUSDToETB).Return("56.500000", true, nil).Once()
	etb, rate, err := USDToETB(s.ctx, 10.0, s.cache)
	s.Require().NoError(err)
	s.InDelta(56.5, rate, 1e-9)
	s.InDelta(565.0, etb, 1e-9)
}

func (s *FXUtilSuite) TestUSDToETB_MissingKey() {
	s.cache.On("Get", s.ctx, FXKeyUSDToETB).Return("", false, nil).Once()
	etb, rate, err := USDToETB(s.ctx, 1.0, s.cache)
	s.Error(err)
	s.Equal(0.0, rate)
	s.Equal(0.0, etb)
}

func (s *FXUtilSuite) TestUSDToETB_ParseError() {
	s.cache.On("Get", s.ctx, FXKeyUSDToETB).Return("not-a-number", true, nil).Once()
	_, _, err := USDToETB(s.ctx, 1.0, s.cache)
	s.Error(err)
}

func (s *FXUtilSuite) TestUSDToETB_CacheError() {
	s.cache.On("Get", s.ctx, FXKeyUSDToETB).Return("", false, errors.New("redis down")).Once()
	_, _, err := USDToETB(s.ctx, 1.0, s.cache)
	s.Error(err)
}

func TestFXUtilSuite(t *testing.T) { suite.Run(t, new(FXUtilSuite)) }
