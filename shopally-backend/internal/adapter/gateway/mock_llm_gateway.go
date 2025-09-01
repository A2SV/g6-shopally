package gateway

import (
	"context"

	"github.com/shopally-ai/pkg/domain"
)

// MockLLMGateway implements domain.LLMGateway and returns a hardcoded parsed intent.
type MockLLMGateway struct{}

// CompareProducts implements domain.LLMGateway.
func (m *MockLLMGateway) CompareProducts(ctx context.Context, productDetails []*domain.Product) (map[string]interface{}, error) {
	panic("unimplemented")
}

func NewMockLLMGateway() domain.LLMGateway {
	return &MockLLMGateway{}
}

func (m *MockLLMGateway) ParseIntent(ctx context.Context, query string) (map[string]interface{}, error) {
	// Very simple mocked intent
	return map[string]interface{}{
		"category":      "smartphone",
		"price_max_ETB": 5000,
	}, nil
}

// SummarizeProduct returns a mocked summarized product for testing
func (m *MockLLMGateway) SummarizeProduct(ctx context.Context, p *domain.Product, summaryType string) (*domain.Product, error) {
	// Return a copy of the product with a mocked summary field if needed
	mockedProduct := *p
	// You can add a mocked summary to a field if domain.Product has one, or just return the product as is
	return &mockedProduct, nil
}
