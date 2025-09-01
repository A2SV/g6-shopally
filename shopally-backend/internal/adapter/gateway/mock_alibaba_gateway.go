package gateway

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/shopally-ai/pkg/domain"
)

// MockAlibabaGateway implements usecase.AlibabaGateway and returns hardcoded products.
type MockAlibabaGateway struct{}

func NewMockAlibabaGateway() domain.AlibabaGateway {
	return &MockAlibabaGateway{}
}

func (m *MockAlibabaGateway) FetchProducts(ctx context.Context, query string, filters map[string]interface{}) ([]*domain.Product, error) {
	fxTs, _ := time.Parse(time.RFC3339, "2025-08-22T10:00:00Z")

	// If caller requested a specific product_id, return a simple inline mock
	// product for easier local testing (no external file required).
	if filters != nil {
		if v, ok := filters["product_id"]; ok {
			switch t := v.(type) {
			case string:
				if strings.TrimSpace(t) == "33006951782" {
					return []*domain.Product{
						{
							ID:                "33006951782",
							Title:             "Mock Sample Phone",
							ImageURL:          "https://via.placeholder.com/300",
							AIMatchPercentage: 90,
							Price:             domain.Price{ETB: 0, USD: 15.90, FXTimestamp: fxTs},
							ProductRating:     4.5,
							SellerScore:       90,
							DeliveryEstimate:  "7-15 days",
							Description:       "Inline mock sample product used for price tests.",
							NumberSold:        1234,
							DeeplinkURL:       "https://www.aliexpress.com/item/33006951782.html",
						},
					}, nil
				}
			case int:
				if strconv.Itoa(t) == "33006951782" {
					return []*domain.Product{
						{
							ID:                "33006951782",
							Title:             "Mock Sample Phone",
							ImageURL:          "https://via.placeholder.com/300",
							AIMatchPercentage: 90,
							Price:             domain.Price{ETB: 0, USD: 15.90, FXTimestamp: fxTs},
							ProductRating:     4.5,
							SellerScore:       90,
							DeliveryEstimate:  "7-15 days",
							Description:       "Inline mock sample product used for price tests.",
							NumberSold:        1234,
							DeeplinkURL:       "https://www.aliexpress.com/item/33006951782.html",
						},
					}, nil
				}
			}
		}
	}

	products := []*domain.Product{
		{
			ID:                 "MOCK-123",
			Title:              "Mock Smartphone - High Quality",
			ImageURL:           "https://via.placeholder.com/150",
			AIMatchPercentage:  92,
			Price:              domain.Price{ETB: 4999.00, USD: 45.45, FXTimestamp: fxTs},
			ProductRating:      4.6,
			SellerScore:        95,
			DeliveryEstimate:   "15-30 days",
			Description:        "High quality smartphone suitable for everyday use.",
			CustomerHighlights: "Good camera, solid battery life",
			CustomerReview:     "Customers praise its durability and battery.",
			NumberSold:         1200,
			DeeplinkURL:        "#",
		},
		{
			ID:                 "MOCK-124",
			Title:              "Mock Budget Phone",
			ImageURL:           "https://via.placeholder.com/150",
			AIMatchPercentage:  88,
			Price:              domain.Price{ETB: 3999.00, USD: 36.36, FXTimestamp: fxTs},
			ProductRating:      4.4,
			SellerScore:        90,
			DeliveryEstimate:   "12-25 days",
			Description:        "Affordable smartphone with essential features.",
			CustomerHighlights: "Long battery life",
			CustomerReview:     "Great value for the price.",
			NumberSold:         2450,
			DeeplinkURL:        "#",
		},
		{
			ID:                 "MOCK-125",
			Title:              "Mock Midrange Phone",
			ImageURL:           "https://via.placeholder.com/150",
			AIMatchPercentage:  90,
			Price:              domain.Price{ETB: 5499.00, USD: 50.00, FXTimestamp: fxTs},
			ProductRating:      4.7,
			SellerScore:        93,
			DeliveryEstimate:   "10-20 days",
			Description:        "Balanced performance and features for most users.",
			CustomerHighlights: "Fast charging",
			CustomerReview:     "Users like the smooth performance.",
			NumberSold:         1780,
			DeeplinkURL:        "#",
		},
		{
			ID:                 "MOCK-126",
			Title:              "Mock Premium Phone",
			ImageURL:           "https://via.placeholder.com/150",
			AIMatchPercentage:  94,
			Price:              domain.Price{ETB: 9999.00, USD: 90.90, FXTimestamp: fxTs},
			ProductRating:      4.9,
			SellerScore:        98,
			DeliveryEstimate:   "7-15 days",
			Description:        "Premium device with high-end features.",
			CustomerHighlights: "High refresh rate display",
			CustomerReview:     "Top-notch screen and performance.",
			NumberSold:         950,
			DeeplinkURL:        "#",
		},
		{
			ID:                 "MOCK-127",
			Title:              "Mock Accessory Bundle",
			ImageURL:           "https://via.placeholder.com/150",
			AIMatchPercentage:  80,
			Price:              domain.Price{ETB: 799.00, USD: 7.27, FXTimestamp: fxTs},
			ProductRating:      4.2,
			SellerScore:        85,
			DeliveryEstimate:   "10-18 days",
			Description:        "Budget-friendly accessory kit for phones.",
			CustomerHighlights: "Budget friendly",
			CustomerReview:     "Great for everyday needs.",
			NumberSold:         5200,
			DeeplinkURL:        "#",
		},
	}

	return products, nil
}
