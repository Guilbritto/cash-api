package handlers

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) RegisterTransactionEndpoint(api fiber.Router) {
	api.Post("/transactions", h.CreateTransaction)
	api.Get("/transactions", h.GetTransactions)
}

// CreateTransaction godoc
// @Summary      Cria uma nova transação
// @Description  Cria uma transação de despesa ou receita para o usuário autenticado
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateTransactionRequest  true  "Dados da transação"
// @Success      201   {object}  models.Transaction
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      401
// @Router       /transactions [post]
func (h *Handlers) CreateTransaction(c *fiber.Ctx) error {
	transaction := new(dto.CreateTransactionRequest)

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid transaction",
		})
	}

	if err := errors.ValidateStruct(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	createdTransaction, err := h.UseCases.TransactionUseCase.Create(transaction, c.Locals("userId").(string))
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
// @Success      200   {object}  models.Transaction
// @Failure      401   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /transactions [get]
func (h *Handlers) GetTransactions(c *fiber.Ctx) error {
	transactions, err := h.UseCases.TransactionUseCase.GetAll(c.Locals("userId").(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get transactions",
		})
	}

	return c.JSON(transactions)

}
