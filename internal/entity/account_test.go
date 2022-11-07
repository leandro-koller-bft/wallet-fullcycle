package entity_test

import (
	"testing"

	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	client, _ := entity.NewClient("client", "client@aol.com")
	account, err := entity.NewAccount(client)

	require.Nil(t, err)
	require.NotNil(t, account)
	require.Equal(t, client, account.Client)
	require.Equal(t, float64(0), account.Balance)
}

func TestCreateAccountFail(t *testing.T) {
	account, err := entity.NewAccount(nil)

	require.NotNil(t, err)
	require.Nil(t, account)
}

func TestAccountBalance(t *testing.T) {
	client, _ := entity.NewClient("client", "client@aol.com")
	account, _ := entity.NewAccount(client)

	err := account.Credit(10)

	require.Nil(t, err)
	require.Equal(t, float64(10), account.Balance)

	err = account.Debit(1)

	require.Nil(t, err)
	require.Equal(t, float64(9), account.Balance)

	err = account.Credit(-10)
	require.Error(t, err)

	err = account.Debit(-1)
	require.Error(t, err)
}
