package util

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/shopally-ai/pkg/domain"
)

// PriceService provides reusable utilities around product prices.
// It reuses the existing AlibabaGateway to fetch product data from AliExpress.
type PriceService struct {
	ag domain.AlibabaGateway
}

// New creates a new PriceService.
func New(ag domain.AlibabaGateway) *PriceService {
	return &PriceService{ag: ag}
}

// UpdatePriceIfChanged fetches the product identified by productID from AliExpress
// and returns the numeric USD price from the mapped product. If the fetched price
// differs from currentUSD (beyond a tiny epsilon) the function returns the
// updated price and changed=true. If the product does not exist or the upstream
// call fails, an error is returned.
//
// Returned values: (updatedPrice, changed, error)
// Note: this function intentionally does not persist or cache results — it
// performs a lookup and returns the numeric USD price plus a boolean that
// indicates whether the price differs from `currentUSD`.
func (s *PriceService) UpdatePriceIfChanged(ctx context.Context, productID string, currentUSD float64) (float64, bool, error) {
	productID = strings.TrimSpace(productID)
	if productID == "" {
		return 0, false, fmt.Errorf("product id is empty")
	}

	// Request a single product by id
	filters := map[string]interface{}{
		"product_id": productID,
		"page_no":    1,
		"page_size":  1,
	}

	products, err := s.ag.FetchProducts(ctx, "", filters)
	if err != nil {
		return 0, false, fmt.Errorf("aliexpress fetch error: %w", err)
	}
	if len(products) == 0 {
		return 0, false, fmt.Errorf("product %s not found on AliExpress", productID)
	}

	p := products[0]
	updated := p.Price.USD

	// Small epsilon to avoid floating point noise
	const eps = 1e-6
	if math.Abs(updated-currentUSD) > eps {
		return updated, true, nil
	}
	return updated, false, nil
}
