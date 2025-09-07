package usecase

import (
	"context"

	"github.com/shopally-ai/pkg/domain"
	"github.com/stretchr/testify/mock"
)

// MockCompareProductsUseCase is a testify-based mock for testing.
type MockCompareProductsUseCase struct {
	mock.Mock
}

var _ CompareProductsExecutor = (*MockCompareProductsUseCase)(nil)

func (m *MockCompareProductsUseCase) Execute(ctx context.Context, products []*domain.Product) (interface{}, error) {
	args := m.Called(ctx, products)
	return args.Get(0), args.Error(1)
}
