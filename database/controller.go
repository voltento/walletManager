package database

import (
	"errors"
	"fmt"
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
	GetAccount(id string) (*Account, error)
	UpdateAccount(id string, acc *Account) error
}

type psqlManager struct {
	db *pg.DB

	insertStmt        *pg.Stmt
	getAccountsStmt   *pg.Stmt
	getAccountStmt    *pg.Stmt
	updateAccountStmt *pg.Stmt
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

	var getAccountStmt *pg.Stmt
	getAccountStmt, err = db.Prepare("select id, currency, amount from account where id=$1;")

	var updateAccountStmt *pg.Stmt
	updateAccountStmt, err = db.Prepare("update account set id=$1, currency=$2, amount=$3  where id=$4;")

	mgr := psqlManager{
		db:                db,
		insertStmt:        insertStmt,
		getAccountsStmt:   getAccountsStmt,
		getAccountStmt:    getAccountStmt,
		updateAccountStmt: updateAccountStmt}
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

func (m psqlManager) UpdateAccount(id string, acc *Account) error {
	r, err := m.updateAccountStmt.Exec(acc.Id, acc.Currency, acc.Amount, id)
	if err != nil {
		return err
	}

	if r.RowsAffected() == 0 {
		return buildCantFindRecordError(id)
	}

	return nil
}

func (m psqlManager) GetAccount(id string) (*Account, error) {
	acc := new(Account)
	result, err := m.getAccountStmt.Query(acc, id)

	if err != nil {
		return nil, err
	}

	if result.RowsReturned() == 0 {
		return nil, buildCantFindRecordError(id)
	}

	print(result)

	return acc, nil
}

func buildCantFindRecordError(id string) error {
	return errors.New(fmt.Sprintf("Can't find the record for id '%v'", id))
}
