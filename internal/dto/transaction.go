package dto

import (
	"time"

	"github.com/Guilbritto/cash-api/internal/models"
)

type CreateTransactionRequest struct {
	Amount      float64                `json:"amount"`
	Type        models.TransactionType `json:"type"`
	Date        time.Time              `json:"date"`
	Description string                 `json:"description"`
	CategoryId  string                 `json:"category_id"`
}
type TransactionResponse struct {
	Id          string                 `json:"id" validate:"required"`
	UserId      string                 `json:"-" validate:"required"`
	Amount      float64                `json:"amount" validate:"required"`
	Type        models.TransactionType `json:"type" validate:"required"`
	Date        time.Time              `json:"date" validate:"required"`
	Description string                 `json:"description"`
	CategoryId  string                 `json:"category_id,omitempty" validate:"required"`
}

type GetTransactionResponse struct {
	transactions []TransactionResponse
}
