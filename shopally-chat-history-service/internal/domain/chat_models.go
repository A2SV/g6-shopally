package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Price struct {
	ETB         float64   `json:"etb" bson:"etb"`
	USD         float64   `json:"usd" bson:"usd"`
	FXTimestamp time.Time `json:"fxTimestamp" bson:"fxTimestamp"`
}

type Product struct {
	ID                    string   `json:"id" bson:"id"`
	Title                 string   `json:"title" bson:"title"`
	ImageURL              string   `json:"imageUrl" bson:"imageUrl"`
	AIMatchPercentage     int      `json:"aiMatchPercentage" bson:"aiMatchPercentage"`
	Price                 Price    `json:"price" bson:"price"` // Nested struct, ensure Price fields also have bson tags
	ProductRating         float64  `json:"productRating" bson:"productRating"`
	SellerScore           int      `json:"-" bson:"sellerScore,omitempty"` // json:"-" ignores for JSON, bson:"sellerScore,omitempty" stores in DB but omits if 0
	DeliveryEstimate      string   `json:"deliveryEstimate" bson:"deliveryEstimate"`
	Description           string   `json:"description" bson:"description"`
	ProductSmallImageURLs []string `json:"productSmallImageUrls" bson:"productSmallImageUrls"`
	NumberSold            int      `json:"numberSold" bson:"numberSold"`
	SummaryBullets        []string `json:"summaryBullets" bson:"summaryBullets"`
	DeeplinkURL           string   `json:"deeplinkUrl" bson:"deeplinkUrl"`
	TaxRate               float64  `json:"taxRate" bson:"taxRate"`
	Discount              float64  `json:"discount" bson:"discount"`
	RemoveProduct         bool     `json:"removeProduct,omitempty" bson:"removeProduct,omitempty"`
}

type Message struct {
	UserPrompt string             `json:"user_prompt" bson:"user_prompt"`
	CreatedAt  primitive.DateTime `json:"created_at" bson:"created_at"`
	Products   []Product          `json:"products" bson:"products"`
}

type ChatSession struct {
	ChatID      string             `json:"chat_id" bson:"chat_id"`
	ChatTitle   string             `json:"chat_title" bson:"chat_title"`
	StartTime   primitive.DateTime `json:"start_time" bson:"start_time"`
	LastUpdated primitive.DateTime `json:"last_updated" bson:"last_updated"`
	Messages    []Message          `json:"messages" bson:"messages"`
}

type ChatHistory struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserEmail    string             `json:"user_email" bson:"user_email"`
	LastActivity primitive.DateTime `json:"last_activity" bson:"last_activity"`
	ChatSessions []ChatSession      `json:"chat_sessions" bson:"chat_sessions"`
}

type AddMessageRequest struct {
	UserPrompt string    `json:"user_prompt" bson:"user_prompt"`
	Products   []Product `json:"products" bson:"products"`
}

type CreateChatRequest struct {
	ChatTitle string `json:"chat_title"`
}

type APIError struct {
	Code    string `json:"code"`    // e.g., "INVALID_INPUT", "NOT_FOUND", "INTERNAL_SERVER_ERROR"
	Message string `json:"message"` // Human-readable error description
}

