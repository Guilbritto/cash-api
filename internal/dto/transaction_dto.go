package dto

import (
	"time"

	"github.com/Guilbritto/cash-api/internal/models"
	"github.com/google/uuid"
)

type CreateTransactionRequest struct {
	Amount      float64                `json:"amount" validate:"required,gt=0"`
	Type        models.TransactionType `json:"type" validate:"required"`
	Date        time.Time              `json:"date" validate:"required"` // parse depois
	Description string                 `json:"description" validate:"required"`
	CategoryId  string                 `json:"category_id" validate:"required,uuid4"`
}

type TransactionResponse struct {
	Id          uuid.UUID              `json:"id" validate:"required"`
	UserId      string                 `json:"-" validate:"required"`
	Amount      float64                `json:"amount" validate:"required"`
	Type        models.TransactionType `json:"type" validate:"required"`
	Date        time.Time              `json:"date" validate:"required"`
	Description string                 `json:"description"`
	Category    *CategoryResponse      `json:"category,omitempty" validate:"required"`
}

type GetTransactionResponse struct {
	transactions []TransactionResponse
}
