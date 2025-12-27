package controllers

import (
	"github.com/Guilbritto/cash-api/internal/application/category"
	"github.com/Guilbritto/cash-api/internal/application/transaction"
)

type ControllerBase struct {
	TransactionService transaction.TransactionService
	CategoryService    category.CategoryService
}
