package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) (entity *Account, err error) {
	entity = &Account{
		Id:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = entity.Validate()
	if err != nil {
		entity = nil
		return
	}
	client.AddAccount(entity)
	return
}

func (a *Account) Validate() (err error) {
	if a.Client == nil {
		err = errors.New("client is required")
	}
	return
}

func (a *Account) Credit(amount float64) (err error) {
	if amount < 0 {
		err = errors.New("invalid negative amount")
		return
	}
	a.Balance += amount
	a.UpdatedAt = time.Now()
	return
}

func (a *Account) Debit(amount float64) (err error) {
	if amount < 0 {
		err = errors.New("invalid negative amount")
		return
	}
	a.Balance -= amount
	a.UpdatedAt = time.Now()
	return
}
