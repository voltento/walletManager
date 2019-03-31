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
	Close() error
}

type psqlManager struct {
	db *pg.DB
}

func (m psqlManager) Close() error {
	return m.db.Close()
}

func CreatePsqlWalletMgr() WalletManager {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "wallets",
		Password: "123",
	})
	return psqlManager{db}
}

func (m psqlManager) StartTransaction() (Transaction, error) {
	return m.db.Begin()
}

func (m psqlManager) CreateAccount(ac *Account) error {
	_, err := m.db.Exec("insert into account (id, currency, amount) values ('1', 'USD', 1.0);")
	return err
}
