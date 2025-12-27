package entities

import (
	defaultError "errors"
	"time"

	"github.com/Guilbritto/cash-api/internal/errors"
	"github.com/google/uuid"
)

type TransactionType int

const (
	Expense TransactionType = iota
	Income
)

func (t TransactionType) IsValid() bool {
	return t == Expense || t == Income
}

type Transaction struct {
	Id          string          `db:"id" json:"id" validate:"required"`
	UserId      string          `db:"user_id" json:"-" validate:"required"`
	Amount      float64         `db:"amount" json:"amount" validate:"required"`
	Type        TransactionType `db:"type" json:"type" validate:"required"`
	Date        time.Time       `db:"date" json:"date" validate:"required"`
	Description string          `db:"description" json:"description"`
	CategoryId  string          `db:"category_id" json:"category_id,omitempty" validate:"required"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at" validate:"required"`
	UpdatedAt   time.Time       `db:"updated_at" json:"updated_at"`
}

func NewTransaction(description string, amount float64, userId string, transactionType TransactionType, date time.Time, categoryId string) (*Transaction, error) {
	tx := &Transaction{
		Id:          uuid.New().String(),
		UserId:      userId,
		Description: description,
		Amount:      amount,
		Type:        transactionType,
		Date:        date,
		CategoryId:  categoryId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if !tx.Type.IsValid() {
		return nil, defaultError.New("Invalid transaction type")
	}

	if err := errors.ValidateStruct(tx); err != nil {
		return nil, err
	}

	return tx, nil
}
