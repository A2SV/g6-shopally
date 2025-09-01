package domain

type Alert struct {
	ID           string  `json:"alertId"`
	DeviceID     string  `json:"deviceId"`
	ProductID    string  `json:"productId"`
	CurrentPrice float64 `json:"currentPrice"`
	IsActive     bool    `json:"isActive"`
}
