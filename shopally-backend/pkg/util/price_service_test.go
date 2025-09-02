package util

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/shopally-ai/pkg/domain"
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
