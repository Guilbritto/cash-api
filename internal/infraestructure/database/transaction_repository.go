package database

import (
	"github.com/Guilbritto/cash-api/internal/domain/entities"
	gorm "gorm.io/gorm"
)

type TransactionRepository struct {
	Db *gorm.DB
}

func (t *TransactionRepository) Save(transaction *entities.Transaction) (entities.Transaction, error) {
	tx := t.Db.Create(transaction)
	return *transaction, tx.Error
}

func (t *TransactionRepository) GetAll(userId string) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	tx := t.Db.Find(&transactions).Where("user_id = ?", userId)

	return transactions, tx.Error
}

func (t *TransactionRepository) GetByID(id string) (entities.Transaction, error) {
	var transactions entities.Transaction
	tx := t.Db.First(&transactions, id)

	return transactions, tx.Error
}
