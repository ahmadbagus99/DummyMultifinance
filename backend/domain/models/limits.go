package models

type Limits struct {
	ID         int     `json:"id"`
	ConsumerID int     `json:"consumer_id"` // Mengacu ke User
	Limit1     float64 `json:"limit_1"`
	Limit2     float64 `json:"limit_2"`
	Limit3     float64 `json:"limit_3"`
	Limit6     float64 `json:"limit_6"`
}
