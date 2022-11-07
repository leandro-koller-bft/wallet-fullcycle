package usecases_test

import (
	"testing"

	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	"github.com/leandro-koller-bft/wallet-ms/internal/usecases"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateClientUsecase(t *testing.T) {
	m := &ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := usecases.NewCreateClientUsecase(m)

	output, err := uc.Apply(&usecases.CreateClientInputDTO{
		Name:  "client",
		Email: "client@email.com",
	})

	require.Nil(t, err)
	require.NotNil(t, output)
	require.Equal(t, "client", output.Name)
	require.Equal(t, "client@email.com", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
