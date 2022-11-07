package usecases

import (
	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	"github.com/leandro-koller-bft/wallet-ms/internal/gateway"
)

type CreateTransactionInputDTO struct {
	FromAccountId string
	ToAccountId   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	Id string
}

type CreateTransactionUsecase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUsecase(
	transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
) (
	uc *CreateTransactionUsecase,
) {
	uc = &CreateTransactionUsecase{
		transactionGateway,
		accountGateway,
	}
	return
}

func (uc *CreateTransactionUsecase) Apply(
	input *CreateTransactionInputDTO,
) (
	output *CreateTransactionOutputDTO,
	err error,
) {
	from, err := uc.AccountGateway.Get(input.FromAccountId)
	if err != nil {
		return
	}
	to, err := uc.AccountGateway.Get(input.ToAccountId)
	if err != nil {
		return
	}
	transaction, err := entity.NewTransaction(from, to, input.Amount)
	if err != nil {
		return
	}
	err = uc.TransactionGateway.Save(transaction)
	if err != nil {
		return nil, err
	}
	output = &CreateTransactionOutputDTO{transaction.Id}
	return
}
