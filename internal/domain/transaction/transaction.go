package transaction

import (
	"time"

	"github.com/Guilbritto/cash-api/internal/errors"
	"github.com/google/uuid"
)

type Transaction struct {
	Id          string    `validate:"required"`
	Description string    `validate:"required"`
	Amount      float64   `validate:"required"`
	CreatedAt   time.Time `validate:"required"`
}

func NewTransaction(description string, amount float64) (*Transaction, error) {

	transaction := &Transaction{
		Id:          uuid.New().String(),
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	err := errors.ValidateStruct(transaction)

	if err == nil {
		return transaction, nil
	}

	return nil, err
}
