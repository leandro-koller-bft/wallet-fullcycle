package entity_test

import (
	"testing"

	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestCreateTransaction(t *testing.T) {
	client, _ := entity.NewClient("client", "client@aol.com")
	accountFrom, _ := entity.NewAccount(client)
	accountTo, _ := entity.NewAccount(client)
	// no funds
	transaction, err := entity.NewTransaction(accountFrom, accountTo, 10)

	require.Error(t, err)
	require.Nil(t, transaction)

	// no amount
	transaction, err = entity.NewTransaction(accountFrom, accountTo, 0)

	require.Error(t, err)
	require.Nil(t, transaction)

	// no account to
	transaction, err = entity.NewTransaction(accountFrom, nil, 10)

	require.Error(t, err)
	require.Nil(t, transaction)

	// no account from
	transaction, err = entity.NewTransaction(nil, accountTo, 10)

	require.Error(t, err)
	require.Nil(t, transaction)

	// successful test
	accountFrom.Credit(10)

	transaction, err = entity.NewTransaction(accountFrom, accountTo, 10)

	require.Nil(t, err)
	require.NotNil(t, transaction)
	require.Equal(t, float64(0), accountFrom.Balance)
	require.Equal(t, float64(10), accountTo.Balance)
}
