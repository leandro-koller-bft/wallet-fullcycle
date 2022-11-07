package database_test

import (
	"database/sql"
	"testing"

	"github.com/leandro-koller-bft/wallet-ms/internal/database"
	"github.com/leandro-koller-bft/wallet-ms/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	fromAcc       *entity.Account
	toAcc         *entity.Account
	transactionDB *database.TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date, updated_at date)")
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), from_account_id varchar(255), to_account_id varchar(255), amount float, transacted_at date)")
	s.transactionDB = database.NewTransactionDB(db)

	s.client, _ = entity.NewClient("John", "j@j.com")
	s.fromAcc, _ = entity.NewAccount(s.client)
	s.toAcc, _ = entity.NewAccount(s.client)

	s.fromAcc.Balance = 10
	s.toAcc.Balance = 10
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestSave() {
	transaction, _ := entity.NewTransaction(s.fromAcc, s.toAcc, 10)
	err := s.transactionDB.Save(transaction)

	s.Nil(err)
	s.Equal(float64(0), s.fromAcc.Balance)
	s.Equal(float64(20), s.toAcc.Balance)
}
