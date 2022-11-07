package usecases_test

import (
	"testing"

	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	"github.com/leandro-koller-bft/wallet-ms/internal/usecases"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type TransactionGatewayMock struct {
	mock.Mock
}

// func (m *TransactionGatewayMock) Get(id string) (*entity.Transaction, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*entity.Transaction), args.Error(1)
// }

func (m *TransactionGatewayMock) Save(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUsecase(t *testing.T) {
	client, _ := entity.NewClient("client", "client@aol.com")
	accountFrom, _ := entity.NewAccount(client)
	accountTo, _ := entity.NewAccount(client)
	accountFrom.Credit(10)
	accountTo.Credit(10)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("Get", accountFrom.Id).Return(accountFrom, nil)
	mockAccount.On("Get", accountTo.Id).Return(accountTo, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Save", mock.Anything).Return(nil)

	uc := usecases.NewCreateTransactionUsecase(mockTransaction, mockAccount)
	output, err := uc.Apply(&usecases.CreateTransactionInputDTO{
		accountFrom.Id,
		accountTo.Id,
		10,
	})

	require.Nil(t, err)
	require.NotNil(t, output)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "Get", 2)
	mockTransaction.AssertNumberOfCalls(t, "Save", 1)
}
