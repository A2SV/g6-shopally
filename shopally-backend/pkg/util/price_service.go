package util

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/shopally-ai/pkg/domain"
)

// PriceService provides reusable utilities around product prices.
// It reuses the existing AlibabaGateway to fetch product data from AliExpress.
type PriceService struct {
	ag  domain.AlibabaGateway
	poc *priceOnlyClient // optional lightweight AliExpress price-only client
}

// New creates a new PriceService.
func New(ag domain.AlibabaGateway) *PriceService {
	return &PriceService{ag: ag}
}

// NewWithAli creates a PriceService that uses a lightweight AliExpress client directly,
// bypassing the heavier gateway mapping. Provide app credentials and optional baseURL.
func NewWithAli(appKey, appSecret, baseURL string) *PriceService {
	if strings.TrimSpace(baseURL) == "" {
		baseURL = "https://api-sg.aliexpress.com/sync"
	}
	return &PriceService{poc: &priceOnlyClient{
		appKey:    appKey,
		appSecret: appSecret,
		baseURL:   baseURL,
		http:      &http.Client{Timeout: 10 * time.Second},
	}}
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

	var updated float64
	if s.poc != nil {
		m, err := s.poc.fetchPrices(ctx, []string{productID})
		if err != nil {
			return 0, false, fmt.Errorf("aliexpress fetch error: %w", err)
		}
		v, ok := m[productID]
		if !ok {
			return 0, false, fmt.Errorf("product %s not found on AliExpress", productID)
		}
		updated = v.USD
	} else {
		// Request a single product by id via gateway
		filters := map[string]interface{}{
			"product_id": productID,
			"page_no":    1,
			"page_size":  1,
			// optimize upstream detail.get to return only price fields
			"price_only": true,
		}
		products, err := s.ag.FetchProducts(ctx, "", filters)
		if err != nil {
			return 0, false, fmt.Errorf("aliexpress fetch error: %w", err)
		}
		if len(products) == 0 {
			return 0, false, fmt.Errorf("product %s not found on AliExpress", productID)
		}
		updated = products[0].Price.USD
	}

	// Small epsilon to avoid floating point noise
	const eps = 1e-6
	if math.Abs(updated-currentUSD) > eps {
		return updated, true, nil
	}
	return updated, false, nil
}

// PriceChange represents the current price and whether it changed compared to the provided baseline.
type PriceChange struct {
	Price float64
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
	for i := 0; i < len(uniq); i += chunkSize {
		end := i + chunkSize
		if end > len(uniq) {
			end = len(uniq)
		}
		chunk := uniq[i:end]
		if s.poc != nil {
			prices, err := s.poc.fetchPrices(ctx, chunk)
			if err != nil {
				return nil, fmt.Errorf("aliexpress batch fetch error: %w", err)
			}
			for id, pr := range prices {
				updated := pr.USD
				out[id] = PriceChange{Price: updated}
			}
		} else {
			filters := map[string]interface{}{ // let gateway translate this appropriately
				"product_ids": strings.Join(chunk, ","),
				"page_no":     1,
				"page_size":   len(chunk),
				// optimize upstream detail.get to return only price fields
				"price_only": true,
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
				out[id] = PriceChange{Price: updated}
			}
		}
	}
	return out, nil
}

// GetCurrentPriceUSDAndETB returns the current price for a product in USD and ETB.
// Uses the lightweight price-only client if configured; otherwise falls back to the gateway.
func (s *PriceService) GetCurrentPriceUSDAndETB(ctx context.Context, productID string) (float64, float64, error) {
	productID = strings.TrimSpace(productID)
	if productID == "" {
		return 0, 0, fmt.Errorf("product id is empty")
	}
	if s.poc != nil {
		m, err := s.poc.fetchPrices(ctx, []string{productID})
		if err != nil {
			return 0, 0, err
		}
		pr, ok := m[productID]
		if !ok {
			return 0, 0, fmt.Errorf("product %s not found on AliExpress", productID)
		}
		return pr.USD, pr.ETB, nil
	}
	// Fallback to gateway
	filters := map[string]interface{}{
		"product_id": productID,
		"page_no":    1,
		"page_size":  1,
		"price_only": true,
	}
	products, err := s.ag.FetchProducts(ctx, "", filters)
	if err != nil {
		return 0, 0, fmt.Errorf("aliexpress fetch error: %w", err)
	}
	if len(products) == 0 {
		return 0, 0, fmt.Errorf("product %s not found on AliExpress", productID)
	}
	usd := products[0].Price.USD
	etb, _, convErr := USDToETB(usd)
	if convErr != nil {
		etb = 0
	}
	return usd, etb, nil
}

// GetCurrentPricesUSDAndETBBatch returns current USD and ETB prices for multiple product IDs.
func (s *PriceService) GetCurrentPricesUSDAndETBBatch(ctx context.Context, productIDs []string) (map[string]struct{ USD, ETB float64 }, error) {
	// sanitize and dedupe
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
	out := make(map[string]struct{ USD, ETB float64 }, len(uniq))
	if len(uniq) == 0 {
		return out, nil
	}

	const chunkSize = 20
	for i := 0; i < len(uniq); i += chunkSize {
		end := i + chunkSize
		if end > len(uniq) {
			end = len(uniq)
		}
		chunk := uniq[i:end]
		if s.poc != nil {
			mp, err := s.poc.fetchPrices(ctx, chunk)
			if err != nil {
				return nil, err
			}
			for id, pr := range mp {
				out[id] = struct{ USD, ETB float64 }{USD: pr.USD, ETB: pr.ETB}
			}
		} else {
			filters := map[string]interface{}{
				"product_ids": strings.Join(chunk, ","),
				"page_no":     1,
				"page_size":   len(chunk),
				"price_only":  true,
			}
			prods, err := s.ag.FetchProducts(ctx, "", filters)
			if err != nil {
				return nil, err
			}
			for _, p := range prods {
				usd := p.Price.USD
				etb, _, convErr := USDToETB(usd)
				if convErr != nil {
					etb = 0
				}
				out[strings.TrimSpace(p.ID)] = struct{ USD, ETB float64 }{USD: usd, ETB: etb}
			}
		}
	}
	return out, nil
}

// priceOnlyClient performs minimal AliExpress calls to get current prices.
type priceOnlyClient struct {
	appKey    string
	appSecret string
	baseURL   string
	http      *http.Client
}

type pricePair struct{ USD, ETB float64 }

func (c *priceOnlyClient) fetchPrices(ctx context.Context, productIDs []string) (map[string]pricePair, error) {
	// Build params
	ts := time.Now().UTC().UnixNano() / 1e6
	params := map[string]string{
		"method":          "aliexpress.affiliate.productdetail.get",
		"app_key":         c.appKey,
		"timestamp":       fmt.Sprintf("%d", ts),
		"sign_method":     "sha256",
		"target_currency": "USD",
		"target_language": "en",
		"product_ids":     strings.Join(productIDs, ","),
		// request only what we need
		"fields": "product_id,target_sale_price",
	}
	// Sign
	params["sign"] = computeAliSignLite(params, c.appSecret)

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("aliexpress status %d", resp.StatusCode)
	}
	// Minimal response struct
	var dr struct {
		Detail struct {
			RespResult struct {
				Result struct {
					Products struct {
						Product []struct {
							ProductID       int64  `json:"product_id"`
							TargetSalePrice string `json:"target_sale_price"`
						} `json:"product"`
					} `json:"products"`
				} `json:"result"`
			} `json:"resp_result"`
		} `json:"aliexpress_affiliate_productdetail_get_response"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&dr); err != nil {
		return nil, err
	}
	out := make(map[string]pricePair, len(productIDs))
	for _, p := range dr.Detail.RespResult.Result.Products.Product {
		id := fmt.Sprintf("%d", p.ProductID)
		usd := parseFloatPriceLite(p.TargetSalePrice)
		if usd <= 0 {
			continue
		}
		etb, _, err := USDToETB(usd)
		if err != nil {
			etb = 0 // leave ETB as 0 if rate not available
		}
		out[id] = pricePair{USD: usd, ETB: etb}
	}
	return out, nil
}

func computeAliSignLite(params map[string]string, appSecret string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		b.WriteString(k)
		b.WriteString(v)
	}
	mac := hmac.New(sha256.New, []byte(appSecret))
	_, _ = mac.Write([]byte(b.String()))
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
}

func parseFloatPriceLite(s string) float64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	s = strings.ReplaceAll(s, ",", "")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}
