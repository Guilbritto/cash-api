package transaction

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

const (
	description = "TestDescription"
	amount      = 30.0
)

func TestNewTransaction(t *testing.T) {
	assert := assert.New(t)

	transaction, _ := NewTransaction(description, amount)

	assert.Equal(transaction.Description, description)
	assert.Equal(transaction.Amount, amount)
}

func TestTransactionShouldNotHaveA(t *testing.T) {
	assert := assert.New(t)

	_, error := NewTransaction("", amount)

	assert.NotNil(error)
	assert.Equal("description cannot be empty", error.Error())
}
