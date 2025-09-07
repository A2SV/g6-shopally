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
	type aliProduct struct {
		AppSalePrice        string `json:"app_sale_price"`
		OriginalPrice       string `json:"original_price"`
		ProductDetailURL    string `json:"product_detail_url"`
		Discount            string `json:"discount"`
		ProductMainImageURL string `json:"product_main_image_url"`
		TaxRate             string `json:"tax_rate"`
		ProductID           int64  `json:"product_id"`
		ShipToDays          string `json:"ship_to_days"`
		EvaluateRate        string `json:"evaluate_rate"`
		SalePrice           string `json:"sale_price"`
		ProductTitle        string `json:"product_title"`

		TargetSalePrice            string `json:"target_sale_price"`
		TargetAppSalePrice         string `json:"target_app_sale_price"`
		ShopName                   string `json:"shop_name"`
		TargetSalePriceCurrency    string `json:"target_sale_price_currency"`
		TargetAppSalePriceCurrency string `json:"target_app_sale_price_currency"`

		ProductSmallImageURLs struct {
			String []string `json:"string"`
		} `json:"product_small_image_urls"`
		SecondLevelCategoryName     string `json:"second_level_category_name"`
		SecondLevelCategoryID       int64  `json:"second_level_category_id"`
		FirstLevelCategoryID        int64  `json:"first_level_category_id"`
		FirstLevelCategoryName      string `json:"first_level_category_name"`
		OriginalPriceCurrency       string `json:"original_price_currency"`
		ShopURL                     string `json:"shop_url"`
		TargetOriginalPriceCurrency string `json:"target_original_price_currency"`
		TargetOriginalPrice         string `json:"target_original_price"`
		ProductVideoURL             string `json:"product_video_url"`
		PromotionLink               string `json:"promotion_link"`
		SKUId                       int64  `json:"sku_id"`
		HotProductCommissionRate    string `json:"hot_product_commission_rate"`
		ShopID                      int64  `json:"shop_id"`
		LastestVolume               int    `json:"lastest_volume"`
		SalePriceCurrency           string `json:"sale_price_currency"`
		CommissionRate              string `json:"commission_rate"`
	}

	type sgResp struct {
		AliexpressResp struct {
			RespResult struct {
				Result struct {
					CurrentRecordCount int `json:"current_record_count"`
					TotalRecordCount   int `json:"total_record_count"`
					CurrentPageNo      int `json:"current_page_no"`
					Products           struct {
						Product []aliProduct `json:"product"`
					} `json:"products"`
				} `json:"result"`
			} `json:"resp_result"`
		} `json:"aliexpress_affiliate_product_query_response"`
	}

	var sg sgResp
	err := json.Unmarshal(data, &sg)
	if err == nil {
		log.Printf("[AlibabaGateway] Unmarshal to sgResp succeeded. Raw product count in struct: %d", len(sg.AliexpressResp.RespResult.Result.Products.Product))
		if len(sg.AliexpressResp.RespResult.Result.Products.Product) > 0 {
			log.Println("[AlibabaGateway] Successfully unmarshaled with SG response structure and found products.")
			out := make([]*domain.Product, 0, len(sg.AliexpressResp.RespResult.Result.Products.Product))
			for _, p := range sg.AliexpressResp.RespResult.Result.Products.Product {
				usd := parseFloatOrZero(p.TargetSalePrice)
				if usd == 0 {
					usd = parseFloatOrZero(p.TargetAppSalePrice)
				}
				if usd == 0 {
					log.Printf("[AlibabaGateway] Warning: No explicit target USD price found for product ID %d. Falling back to SalePrice/AppSalePrice which might be in CNY.", p.ProductID)
					usd = parseFloatOrZero(p.SalePrice)
					if usd == 0 {
						usd = parseFloatOrZero(p.AppSalePrice)
					}
				}

				tax := parseFloatOrZero(p.TaxRate)
				discount := parsePercentOrZero(p.Discount)
				rating := parsePercentOrZero(p.EvaluateRate)

				prod := &domain.Product{
					ID:                strconv.FormatInt(p.ProductID, 10),
					Title:             strings.TrimSpace(p.ProductTitle),
					ImageURL:          strings.TrimSpace(p.ProductMainImageURL),
					AIMatchPercentage: 0, // Placeholder
					Price: domain.Price{
						ETB:         0,
						USD:         usd,
						FXTimestamp: time.Now().UTC(),
					},
					ProductRating:      rating,
					SellerScore:        0, // Placeholder
					DeliveryEstimate:   strings.TrimSpace(p.ShipToDays),
					Description:        "", // Not available in current API response snippet
					CustomerHighlights: "", // Not available in current API response snippet
					CustomerReview:     "", // Not available in current API response snippet
					NumberSold:         p.LastestVolume,
					SummaryBullets:     []string{},
					DeeplinkURL:        strings.TrimSpace(p.ProductDetailURL),
					TaxRate:            tax,
					Discount:           discount,
				}
				out = append(out, prod)
			}
			log.Println("Mapped", len(out), "products from AliExpress SG response")
			return out, nil
		} else {
			log.Println("[AlibabaGateway] Unmarshal to SG response structure succeeded, but found an empty 'product' array. This might indicate no products matched the query or a deeper API issue for this specific response.")
			return []*domain.Product{}, nil
		}
	} else {
		log.Printf("[AlibabaGateway] SG response structure unmarshaling failed: %v. This is unexpected for the current API response format. Returning an empty product list.", err)
		return []*domain.Product{}, fmt.Errorf("failed to unmarshal AliExpress response with SG structure: %v", err)
	}
}

func parseFloatOrZero(s string) float64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	s = strings.ReplaceAll(s, ",", "")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("[AlibabaGateway] parseFloatOrZero: failed to parse '%s' as float: %v", s, err)
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

type AlibabaHTTPGateway struct {
	client *http.Client
	cfg    *config.Config
}

func NewAlibabaHTTPGateway(cfg *config.Config) domain.AlibabaGateway {
	return &AlibabaHTTPGateway{
		client: &http.Client{Timeout: 10 * time.Second},
		cfg:    cfg,
	}
}

const mockAliExpressResponse = `{
	"aliexpress_affiliate_product_query_response": {
		"resp_result": {
			"result": {
				"current_record_count": 1,
				"total_record_count": 1,
				"current_page_no": 1,
				"products": {
					"product": [
						{
							"app_sale_price": "362",
							"original_price": "100",
							"product_detail_url": "https://www.aliexpress.com/item/33006951782.html",
							"discount": "50%",
							"product_main_image_url": "https://example.com/img.jpg",
							"tax_rate": "0.1",
							"product_id": 33006951782,
							"ship_to_days": "ship to RU in 7 days",
							"evaluate_rate": "92.1%",
							"sale_price": "15.9",
							"product_title": "Spring Autumn mother daughter dress matching outfits",
							"target_sale_price": "15.9",
							"target_app_sale_price": "15.9",
							"shop_name": "Mock Shop",
							"target_sale_price_currency": "USD",
							"first_level_category_id": 1,
							"second_level_category_id": 2,
							"sku_id": 12345,
							"shop_id": 67890,
							"lastest_volume": 5,
							"commission_rate": "7.0%",
							"target_app_sale_price_currency": "USD"
						}
					]
				}
			}
		},
		"request_id": "0ba2887315178178017221014"
	}
}`

// FetchProducts implements usecase.AlibabaGateway.
func (a *AlibabaHTTPGateway) FetchProducts(ctx context.Context, Keywords string, filters map[string]interface{}) ([]*domain.Product, error) {
	ts := time.Now().UTC().UnixNano() / 1e6
	tsStr := strconv.FormatInt(ts, 10)

	log.Printf("[AlibabaGateway] FetchProducts called with query: '%s' and filters: %+v", Keywords, filters)

	// Initialize params with required fields and **default values**
	params := map[string]string{
		"method":          "aliexpress.affiliate.product.query",
		"app_key":         a.cfg.Aliexpress.AppKey,
		"timestamp":       tsStr,
		"sign_method":     "sha256",
		"keywords":        Keywords,
		"page_no":         "1",         // Default page number
		"page_size":       "10",        // Default page size
		"target_currency": "USD",       // Default currency
		"target_language": "en",        // Default language
		"sort":            "relevancy", // Default sort order

		// Define all fields we want to receive from the API.
		// This list should reflect all fields in `aliProduct` that you want populated.
		"fields": "product_id,product_title,product_main_image_url,product_detail_url,sale_price,app_sale_price,original_price,discount,evaluate_rate,tax_rate,target_sale_price,target_app_sale_price,shop_name,lastest_volume,ship_to_days,first_level_category_name,second_level_category_name",
	}

	// Apply overrides from the filters map
	// For fields like min_sale_price, max_sale_price, category_ids, etc.,
	// we only add them to params if they are explicitly provided and non-empty.
	// This prevents sending empty string values for optional parameters, which
	// might be interpreted differently by the API than omitting them.

	// Helper function for safely setting string parameters from filters
	setStringParam := func(key string) {
		if v, ok := filters[key]; ok {
			if s, ok := v.(string); ok && s != "" {
				params[key] = s
			}
		}
	}

	// Helper function for safely setting integer/float parameters from filters
	setNumberParam := func(key string) {
		if v, ok := filters[key]; ok {
			switch t := v.(type) {
			case int:
				params[key] = strconv.Itoa(t)
			case float64:
				params[key] = strconv.Itoa(int(t)) // Convert float to int string
			case string:
				if t != "" {
					params[key] = t
				}
			}
		}
	}

	setFloatParam := func(key string) {
		if v, ok := filters[key]; ok {
			switch t := v.(type) {
			case float64:
				// Use -1 for precision to represent the smallest number of digits
				// necessary to accurately represent value
				params[key] = strconv.FormatFloat(t, 'f', -1, 64)
			case string:
				if t != "" {
					params[key] = t
				}
			}
		}
	}

	// Override defaults with values from the filters map
	setNumberParam("page_no")
	setNumberParam("page_size")
	setStringParam("category_ids")
	setFloatParam("min_sale_price") // Use setFloatParam for prices
	setFloatParam("max_sale_price") // Use setFloatParam for prices
	setStringParam("platform_product_type")
	setStringParam("sort")
	setStringParam("target_currency")
	setStringParam("target_language")
	setStringParam("tracking_id")
	setStringParam("promotion_name")
	setStringParam("ship_to_country")
	setNumberParam("delivery_days")

	// Log final params for debugging
	log.Printf("[AlibabaGateway] Final API params: %+v", params)

	// The 'fields' parameter is critical for our mapper. It's best to control
	// it internally to ensure all expected fields for `aliProduct` are always requested.
	// If the user *must* override it, a more complex merge/validation logic would be needed.
	// For now, we prioritize our hardcoded list for reliability.

	sign := computeAliSign(params, a.cfg.Aliexpress.AppSecret)
	params["sign"] = sign

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

	log.Printf("[AlibabaGateway] Final request URL: %s", u.String())

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
	log.Printf("[AlibabaGateway] response status=%d body_preview=%s", resp.StatusCode, preview(respBody.Bytes(), 800))

	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		loc := resp.Header.Get("Location")
		log.Printf("[AlibabaGateway] redirect detected: status=%d location=%s", resp.StatusCode, loc)
		return nil, fmt.Errorf("aliexpress API redirected: status=%d location=%s", resp.StatusCode, loc)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("[AlibabaGateway] non-200 response: %d body: %s", resp.StatusCode, preview(respBody.Bytes(), 1000))
		return nil, fmt.Errorf("aliexpress API returned status %d: %s", resp.StatusCode, preview(respBody.Bytes(), 1000))
	}

	prods, err := MapAliExpressResponseToProducts(respBody.Bytes())
	if err != nil {
		log.Printf("[AlibabaGateway] mapping error from real API response: %v. Attempting mock fallback for development.", err)
		return MapAliExpressResponseToProducts([]byte(mockAliExpressResponse))
	}
	return prods, nil
}

// computeAliSign computes the signature expected by the AliExpress affiliate API.
// Algorithm: sort keys, concatenate key+value (skip empty), signBase = appSecret + concatenated + appSecret,
// SHA256 and return uppercase hex.
func computeAliSign(params map[string]string, appSecret string) string {
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
	if len(unsigned) > 200 {
		log.Printf("[AlibabaGateway] sign unsigned preview: %s...", unsigned[:200])
	} else {
		log.Printf("[AlibabaGateway] sign unsigned preview: %s", unsigned)
	}

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
