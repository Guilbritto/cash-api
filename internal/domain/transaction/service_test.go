package transaction

import (
	"testing"

	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(transaction *Transaction) (Transaction, error) {
	args := r.Called(transaction)
	return Transaction{}, args.Error(0)
}

func (r *repositoryMock) GetAll() ([]Transaction, error) {
	args := r.Called()
	return []Transaction{}, args.Error(0)
}

func TestService_Create(t *testing.T) {
	assert := assert.New(t)

	service := TransactionService{}
	newTransaction := dto.TransactionRequest{
		Amount:      100,
		Description: "test Description",
	}

	transaction, err := service.Create(newTransaction)

	assert.NotNil(transaction)
	assert.Nil(err)
}

func TestService_SaveTransaction(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	service := TransactionService{Repository: repositoryMock}

	newTransaction := dto.TransactionRequest{
		Amount:      100,
		Description: "test Description",
	}

	_, err := service.Create(newTransaction)

	assert.Nil(err)
}
