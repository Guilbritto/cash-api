package repositories

import "github.com/Guilbritto/cash-api/internal/models"

type TransactionRepository interface {
	Save(transaction *models.Transaction) (models.Transaction, error)
	GetAll(userId string) ([]models.Transaction, error)
	GetByID(transactionId string) (models.Transaction, error)
}
