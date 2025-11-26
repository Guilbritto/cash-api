package controllers

import (
	"cashflow/internal/domain/transaction"
)

type ControllerBase struct {
	TransactionService transaction.TransactionService
}
