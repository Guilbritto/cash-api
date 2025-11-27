package transaction

import (
	"time"

	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/google/uuid"
)

type TransactionService struct {
	Repository TransactionRepository
}

func (s *TransactionService) Create(transaction dto.TransactionRequest) (Transaction, error) {

	newTransaction := Transaction{
		Amount:      transaction.Amount,
		Description: transaction.Description,
		CreatedAt:   time.Now(),
		Id:          uuid.New().String(),
	}

	createdTransaction, err := s.Repository.Save(&newTransaction)
	if err != nil {
		return Transaction{}, err
	}

	return createdTransaction, nil
}

func (s *TransactionService) GetAll() ([]Transaction, error) {

	return s.Repository.GetAll()
}
