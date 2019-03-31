package database

import (
	"github.com/go-pg/pg"
)

type Transaction interface {
	Commit() error
	Rollback() error
}

type WalletManager interface {
	StartTransaction() (Transaction, error)
	CreateAccount(ac *Account) error
	GetAllAccounts() ([]*Account, error)
	Close() error
}

type psqlManager struct {
	db *pg.DB

	insertStmt      *pg.Stmt
	getAccountsStmt *pg.Stmt
}

func (m psqlManager) Close() error {
	return m.db.Close()
}

func CreatePsqlWalletMgr() (WalletManager, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "wallets",
		Password: "123",
	})

	var err error

	var insertStmt *pg.Stmt
	insertStmt, err = db.Prepare("insert into account (id, currency, amount) values ($1, $2, $3);")
	if err != nil {
		return nil, err
	}

	var getAccountsStmt *pg.Stmt
	getAccountsStmt, err = db.Prepare("select id, currency, amount from account;")

	mgr := psqlManager{
		db:              db,
		insertStmt:      insertStmt,
		getAccountsStmt: getAccountsStmt}
	return mgr, nil
}

func (m psqlManager) StartTransaction() (Transaction, error) {
	return m.db.Begin()
}

func (m psqlManager) CreateAccount(ac *Account) error {
	_, err := m.insertStmt.Exec(ac.Id, ac.Currency, ac.Amount)
	return err
}

func (m psqlManager) GetAllAccounts() ([]*Account, error) {
	var acc []*Account
	_, err := m.getAccountsStmt.Query(&acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
