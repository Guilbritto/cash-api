package dto

import "time"

type CreateTransactionRequest struct {
	Amount      float64   `json:"amount"`
	Type        int       `json:"type"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	CategoryId  string    `json:"category_id"`
}
type TransactionResponse struct {
	Id          string
	Description string
	Amount      float64
}

type GetTransactionResponse struct {
	transactions []TransactionResponse
}
