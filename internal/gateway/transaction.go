package gateway

import "github.com/leandro-koller-bft/wallet-ms/internal/entity"

type TransactionGateway interface {
	Save(transaction *entity.Transaction) error
}
