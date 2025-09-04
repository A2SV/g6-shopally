package domain

import "time"

// Price represents the price of a product in different currencies.
type Price struct {
	ETB         float64   `json:"etb"`
	USD         float64   `json:"usd"`
	FXTimestamp time.Time `json:"fxTimestamp"`
}

// Product represents a product found on an e-commerce platform.
type Product struct {
	ID                    string   `json:"id"`
	Title                 string   `json:"title"`
	ImageURL              string   `json:"imageUrl"`
	AIMatchPercentage     int      `json:"aiMatchPercentage"`
	Price                 Price    `json:"price"`
	ProductRating         float64  `json:"productRating"`
	SellerScore           int      `json:"-"`
	DeliveryEstimate      string   `json:"deliveryEstimate"`
	Description           string   `json:"description"`
	ProductSmallImageURLs []string `json:"productSmallImageUrls"`
	NumberSold            int      `json:"numberSold"`
	SummaryBullets        []string `json:"summaryBullets"`
	DeeplinkURL           string   `json:"deeplinkUrl"`
	TaxRate               float64  `json:"taxRate"`
	Discount              float64  `json:"discount"`
	RemoveProduct         bool     `json:"removeProduct,omitempty"` // Flag to indicate if product should be removed from results
}

// Synthesis captures comparison insights for a product.
type Synthesis struct {
	Pros        []string          `json:"pros"`
	Cons        []string          `json:"cons"`
	IsBestValue bool              `json:"isBestValue"`
	Features    map[string]string `json:"features"`
}

// ProductComparison wraps a product and its synthesis insights.
type ProductComparison struct {
	Product   Product   `json:"product"`
	Synthesis Synthesis `json:"synthesis"`
}

// OverallComparison holds summary-level comparison insights.
type OverallComparison struct {
	BestValueProduct string   `json:"bestValueProduct"`
	BestValueLink    string   `json:"bestValueLink"`  // Deeplink for the best product
	BestValuePrice   Price    `json:"bestValuePrice"` // Price object of the best product
	KeyHighlights    []string `json:"keyHighlights"`
	Summary          string   `json:"summary"`
}

// ComparisonResult holds multiple product comparisons plus overall insights.
type ComparisonResult struct {
	Products          []ProductComparison `json:"products"`
	OverallComparison OverallComparison   `json:"overallComparison"`
}
