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
	ID                string   `json:"id"`
	Title             string   `json:"title"`
	ImageURL          string   `json:"imageUrl"`
	AIMatchPercentage int      `json:"aiMatchPercentage"`
	Price             Price    `json:"price"`
	ProductRating     float64  `json:"productRating"`
	SellerScore       int      `json:"sellerScore"`
	DeliveryEstimate  string   `json:"deliveryEstimate"`
	SummaryBullets    []string `json:"summaryBullets"`
	DeeplinkURL       string   `json:"deeplinkUrl"`
	TaxRate           float64  `json:"taxRate"`
	Discount          float64  `json:"discount"`
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

// ComparisonResult holds multiple product comparisons (for side-by-side results).
type ComparisonResult struct {
	Products []ProductComparison `json:"products"`
}
