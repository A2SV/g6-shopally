package gateway

import (
	"context"
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

	products := []*domain.Product{
		{
			ID:                "MOCK-123",
			Title:             "Mock Smartphone - High Quality",
			ImageURL:          "https://via.placeholder.com/150",
			AIMatchPercentage: 92,
			Price:             domain.Price{ETB: 4999.00, USD: 45.45, FXTimestamp: fxTs},
			ProductRating:     4.6,
			SellerScore:       95,
			DeliveryEstimate:  "15-30 days",
			SummaryBullets:    []string{"This is a mock summary bullet."},
			DeeplinkURL:       "#",
		},
		{
			ID:                "MOCK-124",
			Title:             "Mock Budget Phone",
			ImageURL:          "https://via.placeholder.com/150",
			AIMatchPercentage: 88,
			Price:             domain.Price{ETB: 3999.00, USD: 36.36, FXTimestamp: fxTs},
			ProductRating:     4.4,
			SellerScore:       90,
			DeliveryEstimate:  "12-25 days",
			SummaryBullets:    []string{"Good battery life"},
			DeeplinkURL:       "#",
		},
		{
			ID:                "MOCK-125",
			Title:             "Mock Midrange Phone",
			ImageURL:          "https://via.placeholder.com/150",
			AIMatchPercentage: 90,
			Price:             domain.Price{ETB: 5499.00, USD: 50.00, FXTimestamp: fxTs},
			ProductRating:     4.7,
			SellerScore:       93,
			DeliveryEstimate:  "10-20 days",
			SummaryBullets:    []string{"Fast charging"},
			DeeplinkURL:       "#",
		},
		{
			ID:                "MOCK-126",
			Title:             "Mock Premium Phone",
			ImageURL:          "https://via.placeholder.com/150",
			AIMatchPercentage: 94,
			Price:             domain.Price{ETB: 9999.00, USD: 90.90, FXTimestamp: fxTs},
			ProductRating:     4.9,
			SellerScore:       98,
			DeliveryEstimate:  "7-15 days",
			SummaryBullets:    []string{"High refresh rate display"},
			DeeplinkURL:       "#",
		},
		{
			ID:                "MOCK-127",
			Title:             "Mock Accessory Bundle",
			ImageURL:          "https://via.placeholder.com/150",
			AIMatchPercentage: 80,
			Price:             domain.Price{ETB: 799.00, USD: 7.27, FXTimestamp: fxTs},
			ProductRating:     4.2,
			SellerScore:       85,
			DeliveryEstimate:  "10-18 days",
			SummaryBullets:    []string{"Budget friendly"},
			DeeplinkURL:       "#",
		},
	}

	return products, nil
}
