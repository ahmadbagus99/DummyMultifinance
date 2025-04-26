package models

import "time"

type Consumers struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"` // Mengacu ke User
	NIK         string    `json:"nik"`
	FullName    string    `json:"full_name"`
	LegalName   string    `json:"legal_name"`
	BirthPlace  string    `json:"birth_place"`
	BirthDate   string    `json:"birth_date"`
	Salary      float64   `json:"salary"`
	KTPPhoto    string    `json:"ktp_photo"`
	SelfiePhoto string    `json:"selfie_photo"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ConsumersLimit struct {
	ConsumerID int     `json:"consumer_id"`
	TenorID    int     `json:"tenor_id"`
	Tenor      string  `json:"tenor"`
	Amount     float64 `json:"amount"`
}
