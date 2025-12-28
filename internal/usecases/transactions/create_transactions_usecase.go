package transactions

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/models"
)

func (s *Service) Create(transaction *dto.CreateTransactionRequest, userId string) (models.Transaction, error) {

	category, err := s.CategoryRepository.GetById(transaction.CategoryId)
	if err != nil {
		return models.Transaction{}, err
	}

	newTransaction, err := models.NewTransaction(
		transaction.Description,
		transaction.Amount,
		userId,
		models.TransactionType(transaction.Type),
		transaction.Date,
		category.Id,
	)

	if err != nil {
		return models.Transaction{}, err
	}

	createdTransaction, err := s.TransactionRepository.Save(newTransaction)
	if err != nil {
		return models.Transaction{}, err
	}

	return createdTransaction, nil
}
