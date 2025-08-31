package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopally-ai/internal/mocks"
	"github.com/stretchr/testify/suite"
)

type FXHandlerSuite struct {
	suite.Suite
	mockFX *mocks.IFXClient
	h      *FXHandler
}

func (s *FXHandlerSuite) SetupTest() {
	s.mockFX = mocks.NewIFXClient(s.T())
	s.h = NewFXHandler(s.mockFX)
}

func (s *FXHandlerSuite) TestGetFX_HappyPath() {
	req := httptest.NewRequest(http.MethodGet, "/fx?from=usd&to=etb&amount=2.5", nil)
	rr := httptest.NewRecorder()

	s.mockFX.On("GetRate", req.Context(), "USD", "ETB").Return(56.0, nil).Once()

	s.h.GetFX(rr, req)
	s.Equal(http.StatusOK, rr.Result().StatusCode)
	s.Equal("application/json", rr.Header().Get("Content-Type"))

	var resp fxResponse
	s.Require().NoError(json.Unmarshal(rr.Body.Bytes(), &resp))
	s.Equal("USD", resp.From)
	s.Equal("ETB", resp.To)
	s.InDelta(56.0, resp.Rate, 1e-9)
	s.InDelta(2.5, resp.Amount, 1e-9)
	s.InDelta(140.0, resp.Converted, 1e-9)
}

func (s *FXHandlerSuite) TestGetFX_DefaultsAndRounding() {
	req := httptest.NewRequest(http.MethodGet, "/fx", nil)
	rr := httptest.NewRecorder()

	s.mockFX.On("GetRate", req.Context(), "USD", "ETB").Return(1.23, nil).Once()

	s.h.GetFX(rr, req)
	s.Equal(http.StatusOK, rr.Result().StatusCode)

	var resp fxResponse
	s.Require().NoError(json.Unmarshal(rr.Body.Bytes(), &resp))
	s.Equal("USD", resp.From)
	s.Equal("ETB", resp.To)
	s.InDelta(1.23, resp.Rate, 1e-9)
	s.InDelta(1.0, resp.Amount, 1e-9)
	s.InDelta(1.23, resp.Converted, 1e-9)
}

func (s *FXHandlerSuite) TestGetFX_InvalidAmount() {
	req := httptest.NewRequest(http.MethodGet, "/fx?amount=abc", nil)
	rr := httptest.NewRecorder()

	s.h.GetFX(rr, req)
	s.Equal(http.StatusBadRequest, rr.Result().StatusCode)
}

func (s *FXHandlerSuite) TestGetFX_UpstreamError() {
	req := httptest.NewRequest(http.MethodGet, "/fx?from=USD&to=ETB", nil)
	rr := httptest.NewRecorder()

	s.mockFX.On("GetRate", req.Context(), "USD", "ETB").Return(0.0, assertAnError("oops")).Once()

	s.h.GetFX(rr, req)
	s.Equal(http.StatusBadGateway, rr.Result().StatusCode)
}

// helper provides an error without importing extra packages here
type assertAnError string

func (e assertAnError) Error() string { return string(e) }

func TestFXHandlerSuite(t *testing.T) { suite.Run(t, new(FXHandlerSuite)) }
