package controllers

import (
	"cashflow/internal/domain/transaction"
	"cashflow/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func (h *ControllerBase) CreateTransaction(c *fiber.Ctx) error {
	transaction := new(dto.TransactionRequest)

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid transaction",
		})
	}

	createdTransaction, err := h.TransactionService.Create(*transaction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create transaction",
		})
	}

	return c.JSON(createdTransaction)
}

func (h *ControllerBase) GetTransactions(c *fiber.Ctx) error {
	transactions, err := h.TransactionService.GetAll()
	transactionResponse := make(map[string][]transaction.Transaction)
	transactionResponse["transactions"] = transactions

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get transactions",
		})
	}

	return c.JSON(transactionResponse)

}
