package gateway

import (
	"testing"
)

const mockAliExpressResponseTest = `
{
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
}
`

func TestMapAliExpressResponseToProducts(t *testing.T) {
	products, err := MapAliExpressResponseToProducts([]byte(mockAliExpressResponseTest))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(products) != 1 {
		t.Fatalf("expected 1 product, got %d", len(products))
	}
	p := products[0]
	if p.ID != "33006951782" {
		t.Errorf("ID mismatch: got %s", p.ID)
	}
	if p.Title == "" {
		t.Errorf("expected Title to be set")
	}
	if p.Price.USD == 0 {
		t.Errorf("expected USD price parsed, got 0")
	}
	if p.TaxRate == 0 {
		t.Errorf("expected TaxRate parsed, got 0")
	}
	if p.Discount == 0 {
		t.Errorf("expected Discount parsed, got 0")
	}
	if p.ProductRating == 0 {
		t.Errorf("expected ProductRating parsed, got 0")
	}
}
