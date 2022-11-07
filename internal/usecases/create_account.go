package usecases

import (
	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	"github.com/leandro-koller-bft/wallet-ms/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientId string
}

type CreateAccountOutputDTO struct {
	Id string
}

type CreateAccountUsecase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUsecase(
	accountGateway gateway.AccountGateway,
	clientGateway gateway.ClientGateway,
) (
	uc *CreateAccountUsecase,
) {
	uc = &CreateAccountUsecase{
		accountGateway,
		clientGateway,
	}
	return
}

func (uc *CreateAccountUsecase) Apply(input *CreateAccountInputDTO) (
	output *CreateAccountOutputDTO,
	err error,
) {
	client, err := uc.ClientGateway.Get(input.ClientId)
	if err != nil {
		return
	}
	account, err := entity.NewAccount(client)
	if err != nil {
		return
	}
	err = uc.AccountGateway.Save(account)
	if err != nil {
		return
	}
	output = &CreateAccountOutputDTO{account.Id}
	return
}
