package mappers

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/models"
)

func TransactionsToResponse(transactions []models.Transaction) *[]dto.TransactionResponse {
	response := make([]dto.TransactionResponse, 0, len(transactions))
	for _, transaction := range transactions {
		response = append(response, dto.TransactionResponse{
			Id:          transaction.Id,
			Amount:      transaction.Amount,
			Type:        transaction.Type,
			Date:        transaction.Date,
			Description: transaction.Description,
			CategoryId:  transaction.CategoryId,
		})
	}
	return &response
}

func TransactionToResponse(transaction models.Transaction) *dto.TransactionResponse {
	response := dto.TransactionResponse{
		Id:          transaction.Id,
		Amount:      transaction.Amount,
		Type:        transaction.Type,
		Date:        transaction.Date,
		Description: transaction.Description,
		CategoryId:  transaction.CategoryId,
	}
	return &response
}
