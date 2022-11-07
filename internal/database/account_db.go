package database

import (
	"database/sql"

	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{
		DB: db,
	}
}

func (a *AccountDB) Get(id string) (account *entity.Account, err error) {
	client := &entity.Client{}
	account = &entity.Account{Client: client}
	stmt, err := a.DB.Prepare(`
		SELECT a.id, a.balance, a.created_at, c.id, c.name, c.email, c.created_at 
		FROM accounts a 
		INNER JOIN clients c 
		ON a.client_id = c.id 
		WHERE a.id = ?
	`)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&account.Id,
		&account.Balance,
		&account.CreatedAt,
		&client.Id,
		&client.Name,
		&client.Email,
		&client.CreatedAt,
	)
	return
}

func (a *AccountDB) Save(account *entity.Account) (err error) {
	stmt, err := a.DB.Prepare("INSERT INTO accounts (id, client_id, balance, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		account.Id,
		account.Client.Id,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)
	return
}
