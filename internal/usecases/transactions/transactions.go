package transactions

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/models"
	"github.com/Guilbritto/cash-api/internal/repositories"
)

type UseCase interface {
	Create(transaction *dto.CreateTransactionRequest, userId string) (models.Transaction, error)
	GetAll(userId string) ([]models.Transaction, error)
}

type Service struct {
	TransactionRepository repositories.TransactionRepository
	CategoryRepository    repositories.CategoryRepository
}
