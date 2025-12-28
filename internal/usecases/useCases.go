package usecases

import (
	"github.com/Guilbritto/cash-api/internal/repositories"
	"github.com/Guilbritto/cash-api/internal/usecases/categories"
	"github.com/Guilbritto/cash-api/internal/usecases/transactions"
)

type UseCases struct {
	CategoriesUseCase  categories.Service
	TransactionUseCase transactions.Service
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		CategoriesUseCase: categories.Service{
			CategoryRepository: repos.Categories,
		},
		TransactionUseCase: transactions.Service{
			TransactionRepository: repos.Transactions,
			CategoryRepository:    repos.Categories,
		},
	}
}
