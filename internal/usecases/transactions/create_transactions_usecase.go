package transactions

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/mappers"
	"github.com/Guilbritto/cash-api/internal/models"
)

func (s *Service) Create(transaction *dto.CreateTransactionRequest, userId string) (*dto.TransactionResponse, error) {

	category, err := s.CategoryRepository.GetById(transaction.CategoryId)
	if err != nil {
		return &dto.TransactionResponse{}, err
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
		return &dto.TransactionResponse{}, err
	}

	createdTransaction, err := s.TransactionRepository.Save(newTransaction)
	if err != nil {
		return &dto.TransactionResponse{}, err
	}

	return mappers.TransactionToResponse(createdTransaction), nil
}
