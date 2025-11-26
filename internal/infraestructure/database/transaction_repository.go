package database

import "cashflow/internal/domain/transaction"

type TransactionRepository struct {
	transactions []transaction.Transaction
}

func (t *TransactionRepository) Save(transaction *transaction.Transaction) (transaction.Transaction, error) {
	t.transactions = append(t.transactions, *transaction)
	return *transaction, nil
}

func (t *TransactionRepository) GetAll() ([]transaction.Transaction, error) {
	return t.transactions, nil
}
