package database_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/leandro-koller-bft/wallet-ms/internal/database"
	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *database.ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	s.clientDB = database.NewClientDB(db)

}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("John", "j@j.com")
	s.clientDB.Save(client)

	clientDB, err := s.clientDB.Get(client.Id)
	s.Nil(err)
	s.Equal(client.Id, clientDB.Id)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}

func (s *ClientDBTestSuite) TestSave() {
	client := &entity.Client{
		Id:        "1",
		Name:      "Test",
		Email:     "j@j.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.clientDB.Save(client)
	s.Nil(err)
}
