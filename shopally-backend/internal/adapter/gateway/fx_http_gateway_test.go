package gateway

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FXHTTPGatewaySuite struct {
	suite.Suite
	ctx context.Context
}

func (s *FXHTTPGatewaySuite) SetupTest() {
	s.ctx = context.Background()
}

func (s *FXHTTPGatewaySuite) newGatewayWithServer(handler http.HandlerFunc) (*FXHTTPGateway, *httptest.Server) {
	srv := httptest.NewServer(handler)
	g := NewFXHTTPGateway(srv.URL, "apikey123", srv.Client())
	return g, srv
}

func (s *FXHTTPGatewaySuite) TestGetRate_XHostResult() {
	g, srv := s.newGatewayWithServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"result": 56.78}`))
	})
	defer srv.Close()

	rate, err := g.GetRate(s.ctx, "usd", "etb")
	s.Require().NoError(err)
	s.InDelta(56.78, rate, 1e-9)
}

func (s *FXHTTPGatewaySuite) TestGetRate_CurrencyFreaksStringRate() {
	g, srv := s.newGatewayWithServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"rates": {"ETB": "56.78"}}`))
	})
	defer srv.Close()

	rate, err := g.GetRate(s.ctx, "USD", "ETB")
	s.Require().NoError(err)
	s.InDelta(56.78, rate, 1e-9)
}

func (s *FXHTTPGatewaySuite) TestGetRate_OpenERAPI_FloatRate() {
	g, srv := s.newGatewayWithServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"rates": {"ETB": 56.78}}`))
	})
	defer srv.Close()

	rate, err := g.GetRate(s.ctx, "USD", "ETB")
	s.Require().NoError(err)
	s.InDelta(56.78, rate, 1e-9)
}

func (s *FXHTTPGatewaySuite) TestGetRate_ERA_V6_ConversionRates() {
	g, srv := s.newGatewayWithServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"conversion_rates": {"ETB": 56.78}}`))
	})
	defer srv.Close()

	rate, err := g.GetRate(s.ctx, "USD", "ETB")
	s.Require().NoError(err)
	s.InDelta(56.78, rate, 1e-9)
}

func (s *FXHTTPGatewaySuite) TestGetRate_BadStatus() {
	g, srv := s.newGatewayWithServer(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusBadGateway)
	})
	defer srv.Close()

	_, err := g.GetRate(s.ctx, "USD", "ETB")
	s.Error(err)
}

func (s *FXHTTPGatewaySuite) TestGetRate_InvalidArgs() {
	g := NewFXHTTPGateway("", "", nil)
	_, err := g.GetRate(s.ctx, "", "ETB")
	s.Error(err)
	_, err = g.GetRate(s.ctx, "USD", "")
	s.Error(err)
}

func (s *FXHTTPGatewaySuite) TestBuildRequestURLVariants() {
	g := NewFXHTTPGateway("", "k123", nil)
	// default falls back to exchangerate.host template when empty
	u := g.buildRequestURL("USD", "ETB")
	s.Contains(u, "exchangerate.host/convert")

	g.APIURL = "https://api.currencyfreaks.com/latest?apikey={APIKEY}&symbols={TO}"
	u = g.buildRequestURL("USD", "ETB")
	s.Contains(u, "currencyfreaks.com")
	s.Contains(u, "apikey=k123")
	s.Contains(u, "symbols=ETB")

	g.APIURL = "https://open.er-api.com/v6/latest/{FROM}"
	u = g.buildRequestURL("USD", "ETB")
	s.Contains(u, "/latest/USD")

	g.APIURL = "https://v6.exchangerate-api.com/v6/{APIKEY}/latest/{FROM}"
	u = g.buildRequestURL("USD", "ETB")
	s.Contains(u, "/v6/k123/latest/USD")
}

func TestFXHTTPGatewaySuite(t *testing.T) { suite.Run(t, new(FXHTTPGatewaySuite)) }
