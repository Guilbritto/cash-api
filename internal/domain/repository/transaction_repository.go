package repository

import "github.com/Guilbritto/cash-api/internal/domain/entities"

type TransactionRepository interface {
	Save(transaction *entities.Transaction) (entities.Transaction, error)
	GetAll(userId string) ([]entities.Transaction, error)
}
