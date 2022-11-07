package usecases

import (
	"time"

	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	"github.com/leandro-koller-bft/wallet-ms/internal/gateway"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	Id        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUsecase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUsecase(gateway gateway.ClientGateway) *CreateClientUsecase {
	return &CreateClientUsecase{
		ClientGateway: gateway,
	}
}

func (uc *CreateClientUsecase) Apply(input *CreateClientInputDTO) (
	output *CreateClientOutputDTO,
	err error,
) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return
	}
	err = uc.ClientGateway.Save(client)
	if err != nil {
		return
	}
	output = &CreateClientOutputDTO{
		Id:        client.Id,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}
	return
}
