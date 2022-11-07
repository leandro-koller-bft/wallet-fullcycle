package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Accounts  []*Account
}

func NewClient(name, email string) (entity *Client, err error) {
	entity = &Client{
		Id:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = entity.Validate()
	if err != nil {
		entity = nil
	}
	return
}

func (c *Client) Validate() (err error) {
	if c.Name == "" {
		err = errors.New("name is required")
	}
	if c.Email == "" {
		err = errors.New("email is required")
	}
	return
}

func (c *Client) Update(name, email string) (err error) {
	backup := *c

	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()

	err = c.Validate()
	if err != nil {
		*c = backup
	}
	return
}

func (c *Client) AddAccount(account *Account) error {
	if account == nil {
		return errors.New("invalid nil account")
	}
	if account.Client == nil {
		return errors.New("invalid nil account's client")
	}
	if account.Client.Id != c.Id {
		return errors.New("invalid account owned by other client")
	}
	for _, acc := range c.Accounts {
		if acc.Id == account.Id {
			return nil
		}
	}
	c.Accounts = append(c.Accounts, account)

	return nil
}
