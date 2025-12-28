package transactions

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/mappers"
)

func (s *Service) GetAll(userId string) (*[]dto.TransactionResponse, error) {
	transactions, err := s.TransactionRepository.GetAll(userId)
	if err != nil {
		return &[]dto.TransactionResponse{}, err
	}

	return mappers.TransactionsToResponse(transactions), nil
}
