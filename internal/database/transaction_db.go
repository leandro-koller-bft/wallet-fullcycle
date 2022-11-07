package database

import (
	"database/sql"

	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{
		DB: db,
	}
}

func (a *TransactionDB) Save(transaction *entity.Transaction) (err error) {
	stmt, err := a.DB.Prepare("INSERT INTO transactions (id, from_account_id, to_account_id, amount, transacted_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		transaction.Id,
		transaction.FromAccount.Id,
		transaction.ToAccount.Id,
		transaction.Amount,
		transaction.TransactedAt,
	)
	return
}
