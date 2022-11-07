package gateway

import "github.com/leandro-koller-bft/wallet-ms/internal/entity"

type AccountGateway interface {
	Get(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
