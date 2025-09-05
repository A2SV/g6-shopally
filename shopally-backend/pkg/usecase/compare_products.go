package usecase

import (
	"context"
	"time"

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

	// if the comapring products fail retry three times with exponential backoff

	var err error

	for try := 0; try < 3; try++ {

		result, err := uc.llmGateway.CompareProducts(ctx, products)
		if err == nil {
			return result, nil
		}
		if try < 2 {
			// wait before retrying
			time.Sleep(time.Duration(1<<try) * time.Second)
		}
	}

	return nil, err
}

// NewCompareProductsUseCase creates a new use case instance.
func NewCompareProductsUseCase(lg domain.LLMGateway) *CompareProductsUseCase {
	return &CompareProductsUseCase{
		llmGateway: lg,
	}
}
