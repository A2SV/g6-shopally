package usecase

import (
	"context"
	"log"
	"sort"
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
		log.Println("SearchProductsUseCase: LLM intent parsing failed for query:", query, "error:", err)
		intent = map[string]interface{}{}
	}

	log.Println("SearchProductsUseCase: parsed intent for query:", query, "as", intent)

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
	log.Println("SearchProductsUseCase: using filters for query:", query, "as", filters)

	var keywords string

	if keywordsStr, ok := filters["keywords"].(string); ok && keywordsStr != "" {
		keywords = keywordsStr
	} else {
		keywords = query
	}

	// Fetch products from the gateway
	products, err := uc.alibabaGateway.FetchProducts(ctx, keywords, filters)

	log.Println("SearchProductsUseCase: fetched", len(products), "products for query:", query, "with filters:", filters)
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

	log.Println("SearchProductsUseCase: ranked products for query:", query)

	// Parallel summarization: each product summary is independent.
	// Parallel summarization: each product summary is independent.
	if uc.llmGateway != nil {
		var wg sync.WaitGroup
		wg.Add(len(products))

		for i := range products {
			go func(index int) {
				defer wg.Done()
				if products[index] == nil {
					return
				}

				// Get the original product and user prompt from context if available
				userPrompt := query

				// Get enhanced product with all details
				enhancedProduct, err := uc.llmGateway.SummarizeProduct(ctx, products[index], userPrompt)
				if err == nil && enhancedProduct != nil {
					// Replace the entire product with enhanced version
					products[index] = enhancedProduct
				}
			}(i)
		}
		wg.Wait()
	}

	// Return the envelope-compatible data payload
	return map[string]interface{}{"products": products}, nil
}

func defaultScore(p *domain.Product) float64 {
	// 0..5 rating scaled to 0..100, seller score is already 0..100
	// Weighted blend: 0.6 rating + 0.4 seller
	return 0.6*(p.ProductRating/5.0*100.0) + 0.4*float64(p.SellerScore)
}
