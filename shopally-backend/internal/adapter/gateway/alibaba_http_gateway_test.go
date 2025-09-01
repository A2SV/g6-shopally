package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockAliExpressResponseValid = `{
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

const mockAliExpressResponseEmpty = `{
    "aliexpress_affiliate_product_query_response": {
        "resp_result": {
            "result": {
                "current_record_count": 0,
                "total_record_count": 0,
                "current_page_no": 1,
                "products": {
                    "product": []
                }
            }
        }
    }
}`

func TestMapAliExpressResponseToProducts(t *testing.T) {
	t.Run("valid response with 1 product", func(t *testing.T) {
		products, err := MapAliExpressResponseToProducts([]byte(mockAliExpressResponseValid))
		require.NoError(t, err)
		require.Len(t, products, 1)

		p := products[0]
		assert.Equal(t, "33006951782", p.ID)
		assert.Equal(t, "Spring Autumn mother daughter dress matching outfits", p.Title)
		assert.Equal(t, "Mock Shop", p.SellerName)

		assert.InDelta(t, 15.9, p.Price.USD, 0.0001)
		assert.InDelta(t, 0.1, p.TaxRate, 0.0001)
		assert.InDelta(t, 50.0, p.Discount, 0.0001)
		assert.InDelta(t, 92.1, p.ProductRating, 0.0001)
		assert.Equal(t, 5, p.NumberSold)
	})

	t.Run("valid response with empty product list", func(t *testing.T) {
		products, err := MapAliExpressResponseToProducts([]byte(mockAliExpressResponseEmpty))
		require.NoError(t, err)
		assert.Empty(t, products)
	})
}
