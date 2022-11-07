package database

import (
	"database/sql"

	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
)

type ClientDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{
		DB: db,
	}
}

func (c *ClientDB) Get(id string) (client *entity.Client, err error) {
	client = &entity.Client{}
	stmt, err := c.DB.Prepare("SELECT id, name, email, created_at FROM clients WHERE id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&client.Id,
		&client.Name,
		&client.Email,
		&client.CreatedAt,
	)
	return
}

func (c *ClientDB) Save(client *entity.Client) (err error) {
	stmt, err := c.DB.Prepare("INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.Id, client.Name, client.Email, client.CreatedAt, client.UpdatedAt)
	return
}
