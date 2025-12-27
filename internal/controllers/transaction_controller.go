package controllers

import (
	"github.com/Guilbritto/cash-api/internal/domain/entities"
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// CreateTransaction godoc
// @Summary      Cria uma nova transação
// @Description  Cria uma transação de despesa ou receita para o usuário autenticado
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateTransactionRequest  true  "Dados da transação"
// @Success      201   {object}  transaction.Transaction
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      401   {object}  dto.ErrorResponse
// @Router       /transactions [post]
func (h *ControllerBase) CreateTransaction(c *fiber.Ctx) error {
	transaction := new(dto.CreateTransactionRequest)

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid transaction",
		})
	}

	createdTransaction, err := h.TransactionService.Create(transaction, c.Locals("userId").(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(createdTransaction)
}

// GetTransactions godoc
// @Summary      Obtém todas as transações do usuário autenticado
// @Description  Retorna todas as transações do usuário autenticado no sistema
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Success      200   {object}  transaction.Transaction
// @Failure      401   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /transactions [get]
func (h *ControllerBase) GetTransactions(c *fiber.Ctx) error {
	transactions, err := h.TransactionService.GetAll(c.Locals("userId").(string))
	transactionResponse := make(map[string][]entities.Transaction)
	transactionResponse["transactions"] = transactions

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get transactions",
		})
	}

	return c.JSON(transactionResponse)

}
