package models

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	ContractNumber  string    `json:"contract_number"`
	CustomerID      int       `json:"customer_id"`
	OTR             float64   `json:"otr"`
	AdminFee        float64   `json:"admin_fee"`
	Installment     float64   `json:"installment"`
	Interest        float64   `json:"interest"`
	AssetName       string    `json:"asset_name"`
	TransactionDate time.Time `json:"transaction_date"`
}
