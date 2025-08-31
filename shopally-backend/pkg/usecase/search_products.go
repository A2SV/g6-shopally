package usecase

import (
	"context"
	"sort"
	"strings"
	"sync"

	"github.com/shopally-ai/pkg/domain"
)

// SearchProductsUseCase contains the business logic for searching products.
// It orchestrates calls to external gateways (LLM, Alibaba, Cache).
type SearchProductsUseCase struct {
	alibabaGateway domain.AlibabaGateway
	llmGateway     domain.LLMGateway
	cacheGateway   domain.CacheGateway
}

// NewSearchProductsUseCase creates a new SearchProductsUseCase.
func NewSearchProductsUseCase(ag domain.AlibabaGateway, lg domain.LLMGateway, cg domain.CacheGateway) *SearchProductsUseCase {
	return &SearchProductsUseCase{
		alibabaGateway: ag,
		llmGateway:     lg,
		cacheGateway:   cg,
	}
}

// Search runs the mocked search pipeline: Parse -> Fetch (using intent as filters).
func (uc *SearchProductsUseCase) Search(ctx context.Context, query string) (interface{}, error) {
	// Parse intent via LLM
	intent, err := uc.llmGateway.ParseIntent(ctx, query)
	if err != nil {
		// For V1 mock, fail soft by using empty filters
		intent = map[string]interface{}{}
	}

	// Infer category if missing using lightweight keyword mapping
	if _, ok := intent["category"]; !ok || intent["category"] == "" {
		if cat := inferCategory(query); cat != "" {
			intent["category"] = cat
		}
	}

	// Prune empty filters (nil/empty string) before passing to gateway
	filters := make(map[string]interface{})
	for k, v := range intent {
		switch vv := v.(type) {
		case nil:
			continue
		case string:
			if vv == "" {
				continue
			}
		}
		filters[k] = v
	}

	// Fetch products from the gateway
	products, err := uc.alibabaGateway.FetchProducts(ctx, query, filters)
	if err != nil {
		return nil, err
	}

	// Default ranking if filters are sparse (no price or delivery constraints)
	if _, ok1 := filters["min_price"]; !ok1 {
		if _, ok2 := filters["max_price"]; !ok2 {
			if _, ok3 := filters["delivery_days_max"]; !ok3 {
				sort.SliceStable(products, func(i, j int) bool {
					si := defaultScore(products[i])
					sj := defaultScore(products[j])
					return si > sj
				})
			}
		}
	}

	// Parallel summarization: each product summary is independent.
	if uc.llmGateway != nil {
		var wg sync.WaitGroup
		wg.Add(len(products))
		for _, p := range products {
			prod := p
			go func() {
				defer wg.Done()
				if prod == nil {
					return
				}
				bullets, err := uc.llmGateway.SummarizeProduct(ctx, prod)
				if err == nil && len(bullets) > 0 {
					prod.SummaryBullets = bullets
				}
			}()
		}
		wg.Wait()
	}

	// Return the envelope-compatible data payload
	return map[string]interface{}{"products": products}, nil
}

// inferCategory does a minimal keyword-based category guess.
func inferCategory(q string) string {
	l := strings.ToLower(q)
	switch {
	case strings.Contains(l, "phone") || strings.Contains(l, "smartphone") || strings.Contains(l, "iphone") || strings.Contains(l, "galaxy"):
		return "smartphone"
	case strings.Contains(l, "laptop") || strings.Contains(l, "notebook") || strings.Contains(l, "macbook"):
		return "laptop"
	case strings.Contains(l, "earbud") || strings.Contains(l, "headphone") || strings.Contains(l, "airpods"):
		return "headphone"
	case strings.Contains(l, "watch") || strings.Contains(l, "smartwatch"):
		return "watch"
	default:
		return ""
	}
}

func defaultScore(p *domain.Product) float64 {
	// 0..5 rating scaled to 0..100, seller score is already 0..100
	// Weighted blend: 0.6 rating + 0.4 seller
	return 0.6*(p.ProductRating/5.0*100.0) + 0.4*float64(p.SellerScore)
}
