package controllers

import "github.com/Guilbritto/cash-api/internal/domain/transaction"

type ControllerBase struct {
	TransactionService transaction.TransactionService
}
