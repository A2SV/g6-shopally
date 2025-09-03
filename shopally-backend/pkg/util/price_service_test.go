package util

import (
	"context"
	"errors"
	"testing"
	"time"

	imocks "github.com/shopally-ai/internal/mocks"
	"github.com/shopally-ai/pkg/domain"
	"github.com/stretchr/testify/mock"
)

// mockGateway is a tiny test double implementing domain.AlibabaGateway for unit tests.
type mockGateway struct {
	products []*domain.Product
	err      error
}

func (m *mockGateway) FetchProducts(ctx context.Context, query string, filters map[string]interface{}) ([]*domain.Product, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.products, nil
}

func TestUpdatePriceIfChanged_Changed(t *testing.T) {
	mg := &mockGateway{
		products: []*domain.Product{{
			ID:    "p1",
			Price: domain.Price{USD: 20.0, FXTimestamp: time.Now()},
		}},
	}

	svc := New(mg)
	updated, changed, err := svc.UpdatePriceIfChanged(context.Background(), "p1", 10.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updated != 20.0 {
		t.Fatalf("expected updated 20.0, got %v", updated)
	}
	if !changed {
		t.Fatalf("expected changed=true")
	}
}

func TestUpdatePriceIfChanged_Unchanged(t *testing.T) {
	mg := &mockGateway{
		products: []*domain.Product{{
			ID:    "p2",
			Price: domain.Price{USD: 15.5, FXTimestamp: time.Now()},
		}},
	}

	svc := New(mg)
	updated, changed, err := svc.UpdatePriceIfChanged(context.Background(), "p2", 15.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updated != 15.5 {
		t.Fatalf("expected updated 15.5, got %v", updated)
	}
	if changed {
		t.Fatalf("expected changed=false")
	}
}

func TestUpdatePriceIfChanged_ProductNotFound(t *testing.T) {
	mg := &mockGateway{products: []*domain.Product{}}
	svc := New(mg)
	_, _, err := svc.UpdatePriceIfChanged(context.Background(), "missing", 0)
	if err == nil {
		t.Fatalf("expected error for missing product, got nil")
	}
}

func TestUpdatePriceIfChanged_GatewayError(t *testing.T) {
	mg := &mockGateway{err: errors.New("upstream failure")}
	svc := New(mg)
	_, _, err := svc.UpdatePriceIfChanged(context.Background(), "p3", 0)
	if err == nil {
		t.Fatalf("expected error when gateway returns error, got nil")
	}
}

func TestUpdatePricesIfChangedBatch_ReturnsFoundPrices(t *testing.T) {
	mg := &mockGateway{
		products: []*domain.Product{
			{ID: "x", Price: domain.Price{USD: 2}},
			{ID: "y", Price: domain.Price{USD: 4}},
		},
	}
	svc := New(mg)
	current := map[string]float64{"x": 2, "y": 3}
	res, err := svc.UpdatePricesIfChangedBatch(context.Background(), []string{"x", "y", "y", "z", ""}, current)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res) != 2 {
		t.Fatalf("expected 2 results for found products, got %d", len(res))
	}
	if res["x"].Price != 2 {
		t.Fatalf("expected x=2, got %v", res["x"].Price)
	}
	if res["y"].Price != 4 {
		t.Fatalf("expected y=4, got %v", res["y"].Price)
	}
	if _, ok := res["z"]; ok {
		t.Fatalf("did not expect missing product 'z' to be present")
	}
}

func TestGetCurrentPriceUSDAndETB_SuccessWithRate(t *testing.T) {
	// Set mock FX cache with rate 60
	c := imocks.NewICachePort(t)
	c.On("Get", mock.Anything, FXKeyUSDToETB).Return("60", true, nil)
	SetFXCache(c)

	mg := &mockGateway{products: []*domain.Product{{ID: "p1", Price: domain.Price{USD: 10}}}}
	svc := New(mg)
	usd, etb, err := svc.GetCurrentPriceUSDAndETB(context.Background(), "p1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if usd != 10 {
		t.Fatalf("expected usd=10, got %v", usd)
	}
	if etb != 600 { // 10 * 60
		t.Fatalf("expected etb=600, got %v", etb)
	}
}

func TestGetCurrentPriceUSDAndETB_RateMissingEtbZero(t *testing.T) {
	// Use mocked cache that misses so conversion fails -> etb should be 0, but no error returned
	c := imocks.NewICachePort(t)
	c.On("Get", mock.Anything, FXKeyUSDToETB).Return("", false, nil)
	SetFXCache(c)

	mg := &mockGateway{products: []*domain.Product{{ID: "p2", Price: domain.Price{USD: 7.5}}}}
	svc := New(mg)
	usd, etb, err := svc.GetCurrentPriceUSDAndETB(context.Background(), "p2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if usd != 7.5 {
		t.Fatalf("expected usd=7.5, got %v", usd)
	}
	if etb != 0 {
		t.Fatalf("expected etb=0 when rate missing, got %v", etb)
	}
}

func TestGetCurrentPricesUSDAndETBBatch_Success(t *testing.T) {
	// Set mock FX cache with rate 50
	c := imocks.NewICachePort(t)
	c.On("Get", mock.Anything, FXKeyUSDToETB).Return("50", true, nil)
	SetFXCache(c)

	mg := &mockGateway{products: []*domain.Product{
		{ID: "a", Price: domain.Price{USD: 3}},
		{ID: "b", Price: domain.Price{USD: 4}},
	}}
	svc := New(mg)
	out, err := svc.GetCurrentPricesUSDAndETBBatch(context.Background(), []string{"a", "b", "c", "a", ""})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(out))
	}
	if out["a"].USD != 3 || out["a"].ETB != 150 {
		t.Fatalf("expected a={3,150}, got %+v", out["a"])
	}
	if out["b"].USD != 4 || out["b"].ETB != 200 {
		t.Fatalf("expected b={4,200}, got %+v", out["b"])
	}
	if _, ok := out["c"]; ok {
		t.Fatalf("did not expect missing product 'c' to be present")
	}
}
