package usecase

import (
	"context"

	"github.com/shopally-ai/pkg/domain"
)

// CompareProductsExecutor defines the contract for comparing products.
type CompareProductsExecutor interface {
	Execute(ctx context.Context, products []*domain.Product) (interface{}, error)
}

// CompareProductsUseCase is the real implementation that calls the LLM gateway.
type CompareProductsUseCase struct {
	llmGateway domain.LLMGateway
}

var _ CompareProductsExecutor = (*CompareProductsUseCase)(nil)

// Execute delegates to the LLMGateway to compare products.
func (uc *CompareProductsUseCase) Execute(ctx context.Context, products []*domain.Product) (interface{}, error) {
	result, err := uc.llmGateway.CompareProducts(ctx, products)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// NewCompareProductsUseCase creates a new use case instance.
func NewCompareProductsUseCase(lg domain.LLMGateway) *CompareProductsUseCase {
	return &CompareProductsUseCase{
		llmGateway: lg,
	}
}
