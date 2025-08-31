package gateway

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

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
	// future fields: http client, cfg, logger
}

// NewAlibabaHTTPGateway returns an implementation of usecase.AlibabaGateway
// suitable for development (returns mock mapped products).
func NewAlibabaHTTPGateway() domain.AlibabaGateway {
	return &AlibabaHTTPGateway{}
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
	// In the future, perform an HTTP request here using `query` and `filters`.
	// For dev/testing we parse the embedded mock JSON and map it.
	return MapAliExpressResponseToProducts([]byte(mockAliExpressResponse))
}
