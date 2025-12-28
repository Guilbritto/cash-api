package transactions

import "github.com/Guilbritto/cash-api/internal/models"

func (s *Service) GetAll(userId string) ([]models.Transaction, error) {

	return s.TransactionRepository.GetAll(userId)
}
