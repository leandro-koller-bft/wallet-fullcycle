package usecases_test

import (
	"testing"

	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	"github.com/leandro-koller-bft/wallet-ms/internal/usecases"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Get(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestCreateAccountUsecase(t *testing.T) {
	client, _ := entity.NewClient("client", "client@email.com")
	clientMock := &ClientGatewayMock{}
	clientMock.On("Get", client.Id).Return(client, nil)

	accountMock := &AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := usecases.NewCreateAccountUsecase(accountMock, clientMock)

	output, err := uc.Apply(&usecases.CreateAccountInputDTO{client.Id})

	require.Nil(t, err)
	require.NotNil(t, output)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
