package models

import (
	defaultError "errors"
	"time"

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
	Id          uuid.UUID       `gorm:"type:uuid;primaryKey" db:"id" json:"id"`
	UserId      string          `gorm:"column:user_id;type:varchar(64);not null" db:"user_id" json:"-"`
	Amount      float64         `gorm:"type:numeric(12,2);not null" db:"amount" json:"amount"`
	Type        TransactionType `gorm:"type:smallint;not null" db:"type" json:"type"`
	Date        time.Time       `gorm:"type:date;not null" db:"date" json:"date"`
	Description string          `gorm:"type:text" db:"description" json:"description"`

	CategoryId uuid.UUID `gorm:"column:category_id;type:uuid;not null" db:"category_id" json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryId;references:Id" json:"category"`

	CreatedAt time.Time `gorm:"not null" db:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" db:"updated_at" json:"updated_at"`
}

func NewTransaction(description string, amount float64, userId string, transactionType TransactionType, date time.Time, categoryId uuid.UUID) (*Transaction, error) {
	tx := &Transaction{
		Id:          uuid.New(),
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

	return tx, nil
}
