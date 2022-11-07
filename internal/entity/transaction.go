package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id           string
	FromAccount  *Account
	ToAccount    *Account
	Amount       float64
	TransactedAt time.Time
}

func NewTransaction(from, to *Account, amount float64) (entity *Transaction, err error) {
	entity = &Transaction{
		Id:          uuid.New().String(),
		FromAccount: from,
		ToAccount:   to,
		Amount:      amount,
	}
	err = entity.Validate()
	if err != nil {
		entity = nil
		return
	}
	entity.Commit()
	return
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.FromAccount == nil {
		return errors.New("invalid nil 'from' account")
	}
	if t.ToAccount == nil {
		return errors.New("invalid nil 'to' account")
	}
	if t.FromAccount.Balance < t.Amount {
		return errors.New("insufficient funds")
	}
	return nil
}

func (t *Transaction) Commit() {
	t.FromAccount.Debit(t.Amount)
	t.ToAccount.Credit(t.Amount)
	t.TransactedAt = time.Now()
}
