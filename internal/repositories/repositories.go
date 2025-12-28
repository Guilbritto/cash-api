package repositories

import "github.com/Guilbritto/cash-api/internal/infraestructure/database"

type Repositories struct {
	Transactions TransactionRepository
	Categories   CategoryRepository
}

func New() *Repositories {
	db := database.NewDb()

	return &Repositories{
		Transactions: &database.TransactionRepository{Db: db},
		Categories:   &database.CategoryRepository{Db: db},
	}
}
