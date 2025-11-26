package dto

type TransactionRequest struct {
	Description string  `validate:"required"`
	Amount      float64 `validate:"required"`
}

type TransactionResponse struct {
	Id          string
	Description string
	Amount      float64
}

type GetTransactionResponse struct {
	transactions []TransactionResponse
}
