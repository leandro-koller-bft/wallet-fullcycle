package entity_test

import (
	"testing"

	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestCreateNewClient(t *testing.T) {
	client, err := entity.NewClient("client", "client@aol.com")

	require.Nil(t, err)
	require.NotNil(t, client)
	require.Equal(t, "client", client.Name)
	require.Equal(t, "client@aol.com", client.Email)
}

func TestCreateNewClientFail(t *testing.T) {
	client, err := entity.NewClient("", "")

	require.NotNil(t, err)
	require.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := entity.NewClient("client", "client@aol.com")
	err := client.Update("client up", "email@up.com")

	require.Nil(t, err)
	require.NotNil(t, client)
	require.Equal(t, "client up", client.Name)
	require.Equal(t, "email@up.com", client.Email)
}

func TestUpdateClientFail(t *testing.T) {
	client, _ := entity.NewClient("client", "client@aol.com")
	err := client.Update("", "")

	require.NotNil(t, err)
	require.NotNil(t, client)
	require.Equal(t, "client", client.Name)
	require.Equal(t, "client@aol.com", client.Email)
}

func TestClientReceivingAccount(t *testing.T) {
	client, _ := entity.NewClient("client", "client@aol.com")
	account, _ := entity.NewAccount(client)
	err := client.AddAccount(account)

	require.Nil(t, err)
	require.Equal(t, 1, len(client.Accounts))
}
