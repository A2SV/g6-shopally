package usecase

import (
	"context"
	"log"
	"sort"
	"sync"

	"fmt"
	"strings"
	"time"

	"github.com/shopally-ai/internal/contextkeys"
	"github.com/shopally-ai/pkg/domain"
	"golang.org/x/time/rate"
)

// SearchProductsUseCase contains the business logic for searching products.
// It orchestrates calls to external gateways (LLM, Alibaba, Cache).
type SearchProductsUseCase struct {
	alibabaGateway domain.AlibabaGateway
	llmGateway     domain.LLMGateway
	cacheGateway   domain.ICachePort
	rateLimiter    *rate.Limiter
}

// NewSearchProductsUseCase creates a new SearchProductsUseCase.
func NewSearchProductsUseCase(ag domain.AlibabaGateway, lg domain.LLMGateway, cg domain.ICachePort, requestPerSecond int) *SearchProductsUseCase {
	limiter := rate.NewLimiter(rate.Limit(requestPerSecond), 1)
	return &SearchProductsUseCase{
		alibabaGateway: ag,
		llmGateway:     lg,
		cacheGateway:   cg,
		rateLimiter:    limiter,
	}
}

// Search runs the mocked search pipeline: Parse -> Fetch (using intent as filters).
func (uc *SearchProductsUseCase) Search(ctx context.Context, query string) (interface{}, error) {

	// get context Header lang
	lang := ctx.Value(contextkeys.RespLang)

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
	var queryClass string
	var promptLang string

	if keywordsStr, ok := filters["keywords"].(string); ok && keywordsStr != "" {
		keywords = keywordsStr
	} else {
		keywords = query
	}
	if pl, ok := filters["language"].(string); ok {
		promptLang = pl
	} else {
		promptLang = lang.(string)
	}

	if qc, ok := filters["query_class"].(string); ok && qc != "" {
		// lowercase and trim
		// tolower, split, sort, join to string
		list := strings.Split(strings.ToLower(strings.TrimSpace(qc)), " ")
		// sorting
		sort.Strings(list)
		// join back to string
		normalizedQuery := strings.Join(list, " ")
		queryClass = normalizedQuery
	} else {
		list := strings.Split(strings.ToLower(strings.TrimSpace(keywords)), " ")
		// sorting
		sort.Strings(list)
		// join back to string
		normalizedQuery := strings.Join(list, " ")
		queryClass = normalizedQuery
	}

	log.Println("SearchProductsUseCase: final keywords for query:", query, "as", keywords)

	if promptLang == "am" {
		lang = promptLang
	}

	ctx = context.WithValue(ctx, contextkeys.RespLang, lang)

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
		sem := make(chan struct{}, 8)

		// ... [Inside your Summarization block] ...
		for i := range products {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				sem <- struct{}{}        // Acquire semaphore (concurrency control)
				defer func() { <-sem }() // Release semaphore

				if products[index] == nil {
					return
				}

				product := products[index]
				userPrompt := query

				// 1. First, check the cache
				cacheKey := fmt.Sprintf("product_summary:%s:%s:%s", product.ID, lang, queryClass)

				if err != nil {
					log.Printf("Cache error for product %s: %v", product.ID, err)
				}

				var cachedProduct domain.Product
				found, err := uc.cacheGateway.GetTypedObject(ctx, cacheKey, &cachedProduct)

				if err != nil {
					log.Printf("Cache error for product %s: %v", product.ID, err)
				}
				if found {
					log.Printf("SearchProductsUseCase: Cache HIT for product %s", product.ID)
					products[index] = &cachedProduct
					return // Exit early on cache hit
				}
				log.Printf("SearchProductsUseCase: Cache MISS for product %s", product.ID)
				// 2. CACHE MISS: Enforce rate limiting before calling the API

				// Use the request context so the wait can be cancelled if the client disconnects.
				err = uc.rateLimiter.Wait(ctx) // <-- WAIT HERE FOR RATE LIMITER
				if err != nil {
					// Context was cancelled (e.g., request timeout or client disconnect)
					log.Printf("Rate limiter wait cancelled for product %s: %v", product.ID, err)
					return // Abort processing this product
				}

				// 3. Now it's safe to call the LLM API
				enhancedProduct, err := uc.llmGateway.SummarizeProduct(ctx, product, userPrompt)

				// 4. Optional: Simple retry for 429s (though the rate limiter should prevent most of them)
				retries := 2
				for retries > 0 && err != nil && strings.Contains(err.Error(), "429") {
					log.Printf("SummarizeProduct hit 429 for product %s. Retries left: %d", product.ID, retries)
					backoff := time.Second * time.Duration(3-retries)
					time.Sleep(backoff)

					// Wait for the rate limiter again before retrying
					if limiterErr := uc.rateLimiter.Wait(ctx); limiterErr != nil {
						break
					}
					enhancedProduct, err = uc.llmGateway.SummarizeProduct(ctx, product, userPrompt)
					retries--
				}

				if err != nil {
					log.Printf("SummarizeProduct failed for product %s: %v", product.ID, err)
					return
				}
				if enhancedProduct != nil {
					// Cache the enhanced result for future requests
					err = uc.cacheGateway.SetObject(ctx, cacheKey, enhancedProduct, 24*time.Hour)
					if err != nil {
						log.Printf("Failed to cache enhanced product %s: %v", product.ID, err)
					}
					products[index] = enhancedProduct
				}
			}(i)
		}
		wg.Wait()
		// ... [Rest of your code] ...
	}

	// discarte nil products if any
	cleaned := make([]*domain.Product, 0, len(products))
	// uncleaned := []*domain.Product{}

	log.Println("****************************************************************")
	for prod := range products {
		log.Println("Product ID:", products[prod].ID, "RemoveProduct:", products[prod].RemoveProduct, "AIMatchPercentage:", products[prod].AIMatchPercentage)
	}
	for _, p := range products {
		if p != nil {
			if !p.RemoveProduct && p.AIMatchPercentage > 30 {
				cleaned = append(cleaned, p)
			}
			//  else {
			// 	uncleaned = append(uncleaned, p)
			// }
		}
	}

	// Sort product list by AI match percentage descending
	sort.SliceStable(cleaned, func(i, j int) bool {
		return cleaned[i].AIMatchPercentage > cleaned[j].AIMatchPercentage
	})

	// if len(cleaned) == 0 {
	// 	cleaned = uncleaned // fallback to uncleaned if all were removed
	// }

	products = cleaned

	log.Println("SearchProductsUseCase: cleaned products for query:", query, "final count:", len(products))

	log.Println("SearchProductsUseCase: summarized products for query:", query)

	// Return the envelope-compatible data payload
	return map[string]interface{}{"products": products}, nil
}

func defaultScore(p *domain.Product) float64 {
	// 0..5 rating scaled to 0..100, seller score is already 0..100
	// Weighted blend: 0.6 rating + 0.4 seller
	return 0.6*(p.ProductRating/5.0*100.0) + 0.4*float64(p.SellerScore)
}
