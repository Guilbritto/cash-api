package transaction

import (
	"github.com/Guilbritto/cash-api/internal/domain/entities"
	"github.com/Guilbritto/cash-api/internal/domain/repository"
	"github.com/Guilbritto/cash-api/internal/dto"
)

type TransactionService struct {
	TransactionRepository repository.TransactionRepository
	CategoryRepository    repository.CategoryRepository
}

func (s *TransactionService) Create(transaction *dto.CreateTransactionRequest, userId string) (entities.Transaction, error) {

	category, err := s.CategoryRepository.GetById(transaction.CategoryId)
	if err != nil {
		return entities.Transaction{}, err
	}

	newTransaction, err := entities.NewTransaction(
		transaction.Description,
		transaction.Amount,
		userId,
		entities.TransactionType(transaction.Type),
		transaction.Date,
		category.Id,
	)

	if err != nil {
		return entities.Transaction{}, err
	}

	createdTransaction, err := s.TransactionRepository.Save(newTransaction)
	if err != nil {
		return entities.Transaction{}, err
	}

	return createdTransaction, nil
}

func (s *TransactionService) GetAll(userId string) ([]entities.Transaction, error) {

	return s.TransactionRepository.GetAll(userId)
}
