package database

import (
	"github.com/Guilbritto/cash-api/internal/models"
	gorm "gorm.io/gorm"
)

type TransactionRepository struct {
	Db *gorm.DB
}

func (t *TransactionRepository) Save(transaction *models.Transaction) (models.Transaction, error) {
	tx := t.Db.Create(transaction)
	return *transaction, tx.Error
}

func (t *TransactionRepository) GetAll(userId string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	tx := t.Db.Preload("Category").Where("user_id = ?", userId).Find(&transactions)

	return transactions, tx.Error
}

func (t *TransactionRepository) GetByID(id string) (models.Transaction, error) {
	var transactions models.Transaction
	tx := t.Db.First(&transactions, id)

	return transactions, tx.Error
}
