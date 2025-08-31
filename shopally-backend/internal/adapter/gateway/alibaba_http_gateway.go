package gateway

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/shopally-ai/internal/config"
	"github.com/shopally-ai/pkg/domain"
)

// MapAliExpressResponseToProducts transforms the raw AliExpress API response JSON
// into a slice of internal `domain.Product` pointers. It is resilient to missing
// fields and uses sensible defaults/placeholders where mapping data is not
// available from the upstream response.
func MapAliExpressResponseToProducts(data []byte) ([]*domain.Product, error) {
	// Define minimal structs for the parts of the AliExpress response we care about.
	type aliProduct struct {
		AppSalePrice        string `json:"app_sale_price"`
		OriginalPrice       string `json:"original_price"`
		ProductDetailURL    string `json:"product_detail_url"`
		Discount            string `json:"discount"`
		ProductMainImageURL string `json:"product_main_image_url"`
		TaxRate             string `json:"tax_rate"`
		ProductID           string `json:"product_id"`
		ShipToDays          string `json:"ship_to_days"`
		EvaluateRate        string `json:"evaluate_rate"`
		SalePrice           string `json:"sale_price"`
		ProductTitle        string `json:"product_title"`
	}

	type resultBlock struct {
		Products []aliProduct `json:"products"`
	}

	type respResult struct {
		Result resultBlock `json:"result"`
	}

	type aliResp struct {
		RespResult respResult `json:"resp_result"`
	}

	// SG sync endpoint returns a different envelope: aliexpress_affiliate_product_query_response.resp_result.result.products.product
	type sgResp struct {
		AliexpressResp struct {
			RespResult struct {
				Result struct {
					Products struct {
						Product []aliProduct `json:"product"`
					} `json:"products"`
				} `json:"result"`
			} `json:"resp_result"`
		} `json:"aliexpress_affiliate_product_query_response"`
	}

	// Try SG shape first
	var sg sgResp
	if err := json.Unmarshal(data, &sg); err == nil {
		if len(sg.AliexpressResp.RespResult.Result.Products.Product) > 0 {
			out := make([]*domain.Product, 0, len(sg.AliexpressResp.RespResult.Result.Products.Product))
			for _, p := range sg.AliexpressResp.RespResult.Result.Products.Product {
				usd := parseFloatOrZero(p.SalePrice)
				if usd == 0 {
					usd = parseFloatOrZero(p.AppSalePrice)
				}

				tax := parseFloatOrZero(p.TaxRate)
				discount := parsePercentOrZero(p.Discount)
				rating := parsePercentOrZero(p.EvaluateRate)

				prod := &domain.Product{
					ID:                strings.TrimSpace(p.ProductID),
					Title:             strings.TrimSpace(p.ProductTitle),
					ImageURL:          strings.TrimSpace(p.ProductMainImageURL),
					AIMatchPercentage: 0,
					Price: domain.Price{
						ETB:         0,
						USD:         usd,
						FXTimestamp: time.Now().UTC(),
					},
					ProductRating:    rating,
					SellerScore:      0,
					DeliveryEstimate: strings.TrimSpace(p.ShipToDays),
					SummaryBullets:   []string{},
					DeeplinkURL:      strings.TrimSpace(p.ProductDetailURL),
					TaxRate:          tax,
					Discount:         discount,
				}

				out = append(out, prod)
			}
			return out, nil
		}
	}

	// Fallback to legacy shape
	var r aliResp
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}

	out := make([]*domain.Product, 0, len(r.RespResult.Result.Products))
	for _, p := range r.RespResult.Result.Products {
		// sale price fallback to app_sale_price
		usd := parseFloatOrZero(p.SalePrice)
		if usd == 0 {
			usd = parseFloatOrZero(p.AppSalePrice)
		}

		tax := parseFloatOrZero(p.TaxRate)
		discount := parsePercentOrZero(p.Discount)
		rating := parsePercentOrZero(p.EvaluateRate)

		prod := &domain.Product{
			ID:                strings.TrimSpace(p.ProductID),
			Title:             strings.TrimSpace(p.ProductTitle),
			ImageURL:          strings.TrimSpace(p.ProductMainImageURL),
			AIMatchPercentage: 0, // placeholder
			Price: domain.Price{
				ETB:         0, // not converted yet
				USD:         usd,
				FXTimestamp: time.Now().UTC(),
			},
			ProductRating:    rating,
			SellerScore:      0, // placeholder
			DeliveryEstimate: strings.TrimSpace(p.ShipToDays),
			SummaryBullets:   []string{},
			DeeplinkURL:      strings.TrimSpace(p.ProductDetailURL),
			TaxRate:          tax,
			Discount:         discount,
		}

		out = append(out, prod)
	}

	return out, nil
}

func parseFloatOrZero(s string) float64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	// remove commas
	s = strings.ReplaceAll(s, ",", "")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func parsePercentOrZero(s string) float64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	s = strings.TrimSuffix(s, "%")
	return parseFloatOrZero(s)
}

// AlibabaHTTPGateway is a development implementation of usecase.AlibabaGateway
// that returns mapped products from a mock AliExpress JSON response. Later
// this should be replaced with a real HTTP client that calls the AliExpress
// affiliate API and maps the real response.
type AlibabaHTTPGateway struct {
	client *http.Client
	cfg    *config.Config
}

// NewAlibabaHTTPGateway returns an implementation of usecase.AlibabaGateway
// suitable for development (returns mock mapped products).
func NewAlibabaHTTPGateway(cfg *config.Config) domain.AlibabaGateway {
	return &AlibabaHTTPGateway{
		client: &http.Client{Timeout: 10 * time.Second},
		cfg:    cfg,
	}
}

const mockAliExpressResponse = `{
	"code": "0",
	"resp_result": {
		"result": {
			"products": [
				{
					"app_sale_price": "362",
					"original_price": "100",
					"product_detail_url": "https://www.aliexpress.com/item/33006951782.html",
					"discount": "50%",
					"product_main_image_url": "https://example.com/img.jpg",
					"tax_rate": "0.1",
					"product_id": "33006951782",
					"ship_to_days": "ship to RU in 7 days",
					"evaluate_rate": "92.1%",
					"sale_price": "15.9",
					"product_title": "Spring Autumn mother daughter dress matching outfits"
				}
			]
		},
		"resp_code": "200",
		"resp_msg": "success"
	},
	"request_id": "0ba2887315178178017221014"
}`

// FetchProducts implements usecase.AlibabaGateway. For now it uses a mock
// JSON response and the mapper `MapAliExpressResponseToProducts` to return
// []*domain.Product.
func (a *AlibabaHTTPGateway) FetchProducts(ctx context.Context, query string, filters map[string]interface{}) ([]*domain.Product, error) {
	// Build form-encoded params according to AliExpress docs and compute per-request signature.
	ts := time.Now().UTC().UnixNano() / 1e6
	tsStr := strconv.FormatInt(ts, 10)

	params := map[string]string{
		"method":      "aliexpress.affiliate.product.query",
		"app_key":     a.cfg.Aliexpress.AppKey,
		"timestamp":   tsStr,
		"sign_method": "sha256",
		"keywords":    query,
		// defaults; can be overridden by filters mapping later
		"page_no":   "1",
		"page_size": "20",
	}

	// optional overrides from filters
	if v, ok := filters["page_no"]; ok {
		switch t := v.(type) {
		case int:
			params["page_no"] = strconv.Itoa(t)
		case float64:
			params["page_no"] = strconv.Itoa(int(t))
		case string:
			if t != "" {
				params["page_no"] = t
			}
		}
	}
	if v, ok := filters["page_size"]; ok {
		switch t := v.(type) {
		case int:
			params["page_size"] = strconv.Itoa(t)
		case float64:
			params["page_size"] = strconv.Itoa(int(t))
		case string:
			if t != "" {
				params["page_size"] = t
			}
		}
	}

	sign := computeAliSign(params, a.cfg.Aliexpress.AppSecret)
	params["sign"] = sign

	// Build GET URL using configured base (use SG sync endpoint by default)
	base := a.cfg.Aliexpress.BaseURL
	if strings.TrimSpace(base) == "" {
		base = "https://api-sg.aliexpress.com/sync"
	}

	u, err := url.Parse(base)
	if err != nil {
		log.Printf("[AlibabaGateway] invalid base url %s: %v", base, err)
		return nil, err
	}

	qv := url.Values{}
	for k, v := range params {
		qv.Set(k, v)
	}
	u.RawQuery = qv.Encode()

	log.Printf("[AlibabaGateway] GET URL preview: %s", preview([]byte(u.String()), 800))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		log.Printf("[AlibabaGateway] new request error: %v", err)
		return nil, err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		log.Printf("[AlibabaGateway] http request error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	respBody := new(bytes.Buffer)
	_, _ = respBody.ReadFrom(resp.Body)
	// log status and a short body preview for diagnostics
	log.Printf("[AlibabaGateway] response status=%d body_preview=%s", resp.StatusCode, preview(respBody.Bytes(), 800))

	// Explicitly detect redirects (3xx) and surface the Location header so callers
	// see why the API returned HTML/maintenance pages instead of JSON.
	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		loc := resp.Header.Get("Location")
		log.Printf("[AlibabaGateway] redirect detected: status=%d location=%s", resp.StatusCode, loc)
		return nil, fmt.Errorf("aliexpress API redirected: status=%d location=%s", resp.StatusCode, loc)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("[AlibabaGateway] non-200 response: %d", resp.StatusCode)
		return nil, fmt.Errorf("aliexpress API returned status %d: %s", resp.StatusCode, preview(respBody.Bytes(), 1000))
	}

	// Map the real API response to domain products; if mapping fails, log and try mock
	prods, err := MapAliExpressResponseToProducts(respBody.Bytes())
	if err != nil {
		log.Printf("[AlibabaGateway] mapping error: %v", err)
		return nil, err
	}
	return prods, nil
}

// computeAliSign computes the signature expected by the AliExpress affiliate API.
// Algorithm: sort keys, concatenate key+value (skip empty), signBase = appSecret + concatenated + appSecret,
// SHA256 and return uppercase hex.
func computeAliSign(params map[string]string, appSecret string) string {
	// Build concatenated key+value string (sorted by key) — do NOT include appSecret here
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

	unsigned := b.String()
	// Debug: log a preview of the unsigned concatenation used to compute the signature
	if len(unsigned) > 200 {
		log.Printf("[AlibabaGateway] sign unsigned preview: %s...", unsigned[:200])
	} else {
		log.Printf("[AlibabaGateway] sign unsigned preview: %s", unsigned)
	}

	// Compute HMAC-SHA256 using appSecret as the key (common requirement for API signatures)
	mac := hmac.New(sha256.New, []byte(appSecret))
	_, _ = mac.Write([]byte(unsigned))
	signature := strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
	log.Printf("[AlibabaGateway] computed sign (HMAC-SHA256) preview: %s", signature)
	return signature
}

// preview returns a safe string preview of bytes up to n chars
func preview(b []byte, n int) string {
	s := string(b)
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
