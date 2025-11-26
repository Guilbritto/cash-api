package transaction

type TransactionRepository interface {
	Save(transaction *Transaction) (Transaction, error)
	GetAll() ([]Transaction, error)
}
