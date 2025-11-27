package database

import (
	"github.com/Guilbritto/cash-api/internal/domain/transaction"
	gorm "gorm.io/gorm"
)

type TransactionRepository struct {
	Db *gorm.DB
}

func (t *TransactionRepository) Save(transaction *transaction.Transaction) (transaction.Transaction, error) {
	tx := t.Db.Create(transaction)
	return *transaction, tx.Error
}

func (t *TransactionRepository) GetAll() ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction
	tx := t.Db.Find(&transactions)

	return transactions, tx.Error
}

func (t *TransactionRepository) GetByID(id string) (transaction.Transaction, error) {
	var transactions transaction.Transaction
	tx := t.Db.First(&transactions, id)

	return transactions, tx.Error
}
