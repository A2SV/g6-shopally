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

// SummarizeProduct returns hardcoded bullets for mocks
func (m *MockLLMGateway) SummarizeProduct(ctx context.Context, p *domain.Product) ([]string, error) {
	return []string{"This is a mock summary bullet."}, nil
}
