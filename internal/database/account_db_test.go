package database_test

import (
	"database/sql"
	"testing"

	"github.com/leandro-koller-bft/wallet-ms/internal/database"
	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	clientDB  *database.ClientDB
	accountDB *database.AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date, updated_at date)")
	s.accountDB = database.NewAccountDB(db)
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	s.clientDB = database.NewClientDB(db)

	s.client, _ = entity.NewClient("John", "j@j.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account, _ := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestGet() {
	s.clientDB.Save(s.client)
	account, _ := entity.NewAccount(s.client)
	s.accountDB.Save(account)

	accountDB, err := s.accountDB.Get(account.Id)

	s.Nil(err)
	s.Equal(account.Id, accountDB.Id)
	s.Equal(account.Client.Id, accountDB.Client.Id)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.Client.Name, accountDB.Client.Name)
	s.Equal(account.Client.Email, accountDB.Client.Email)
}
