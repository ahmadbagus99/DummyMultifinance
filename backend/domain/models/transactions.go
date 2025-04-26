package models

import "time"

type Transactions struct {
	ID              int       `json:"id"`
	ContractNumber  string    `json:"contract_number"`
	ConsumerID      int       `json:"consumer_id"`
	OTR             float64   `json:"otr"`
	AdminFee        float64   `json:"admin_fee"`
	Installment     float64   `json:"installment"`
	Interest        float64   `json:"interest"`
	AssetName       string    `json:"asset_name"`
	Approved        bool      `json:"approved"`
	TransactionDate time.Time `json:"transaction_date"`
}
