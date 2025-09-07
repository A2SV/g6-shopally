package gateway

import (
	"context"

	"github.com/shopally-ai/internal/contextkeys"
	"github.com/shopally-ai/pkg/domain"
)

// MockLLMGateway implements domain.LLMGateway and returns a hardcoded parsed intent.
type MockLLMGateway struct{}

// CompareProducts implements domain.LLMGateway.
func (m *MockLLMGateway) CompareProducts(ctx context.Context, productDetails []*domain.Product) (map[string]interface{}, error) {
	// Simple mock: best value = lowest USD price
	bestIdx := 0
	if len(productDetails) > 0 {
		best := productDetails[0].Price.USD
		for i := 1; i < len(productDetails); i++ {
			if productDetails[i].Price.USD < best {
				best = productDetails[i].Price.USD
				bestIdx = i
			}
		}
	}

	lang, _ := ctx.Value(contextkeys.RespLang).(string)
	prosLabelEn := []string{"Good price", "Decent rating"}
	consLabelEn := []string{"May lack accessories"}
	prosLabelAm := []string{"መልካም ዋጋ", "ጥሩ እውቅና"}
	consLabelAm := []string{"አንዳንድ ንብረቶች ሊጎዱ ይችላሉ"}

	comparisons := make([]map[string]interface{}, 0, len(productDetails))
	for i, p := range productDetails {
		var pros, cons []string
		if lang == "am" {
			pros = append([]string{}, prosLabelAm...)
			cons = append([]string{}, consLabelAm...)
		} else {
			pros = append([]string{}, prosLabelEn...)
			cons = append([]string{}, consLabelEn...)
		}

		comparisons = append(comparisons, map[string]interface{}{
			"product": p,
			"synthesis": map[string]interface{}{
				"pros":        pros,
				"cons":        cons,
				"isBestValue": i == bestIdx,
				"features": map[string]string{
					"Screen Type": "Unknown",
					"Processor":   "Unknown",
				},
			},
		})
	}

	return map[string]interface{}{
		"comparison": comparisons,
	}, nil
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
