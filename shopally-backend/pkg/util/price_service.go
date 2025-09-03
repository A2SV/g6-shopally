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

// PriceChange represents the current price and whether it changed compared to the provided baseline.
type PriceChange struct {
	Price   float64
	Changed bool
}

// UpdatePricesIfChangedBatch fetches current USD prices for multiple productIDs at once (when supported by the gateway)
// and compares them against the provided current map. It returns a map[productID]PriceChange for all product IDs that
// were found upstream. Missing products won't appear in the returned map. Errors are returned only for upstream failures.
// Note: This uses the AlibabaGateway's FetchProducts with a "product_ids" filter when available; the underlying gateway
// should translate this to the appropriate AliExpress API call.
func (s *PriceService) UpdatePricesIfChangedBatch(ctx context.Context, productIDs []string, current map[string]float64) (map[string]PriceChange, error) {
	out := make(map[string]PriceChange, len(productIDs))
	// Sanitize and dedupe IDs
	uniq := make([]string, 0, len(productIDs))
	seen := make(map[string]struct{}, len(productIDs))
	for _, id := range productIDs {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		uniq = append(uniq, id)
	}

	if len(uniq) == 0 {
		return out, nil
	}

	// Chunk to reasonable size; AliExpress typically allows dozens per call. We'll use 20 as a safe default.
	const chunkSize = 20
	const eps = 1e-6
	for i := 0; i < len(uniq); i += chunkSize {
		end := i + chunkSize
		if end > len(uniq) {
			end = len(uniq)
		}
		chunk := uniq[i:end]
		filters := map[string]interface{}{ // let gateway translate this appropriately
			"product_ids": strings.Join(chunk, ","),
			"page_no":     1,
			"page_size":   len(chunk),
		}
		products, err := s.ag.FetchProducts(ctx, "", filters)
		if err != nil {
			return nil, fmt.Errorf("aliexpress batch fetch error: %w", err)
		}
		for _, p := range products {
			id := strings.TrimSpace(p.ID)
			if id == "" {
				continue
			}
			updated := p.Price.USD
			base := current[id]
			changed := math.Abs(updated-base) > eps
			out[id] = PriceChange{Price: updated, Changed: changed}
		}
	}
	return out, nil
}
