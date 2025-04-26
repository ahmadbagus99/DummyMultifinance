package models

type Limits struct {
	ID         int     `json:"id"`
	ConsumerID int     `json:"consumer_id"`
	TenorID    int     `json:"tenor_id"`
	Amount     float64 `json:"amount"`
}
