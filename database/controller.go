package database

import (
	"github.com/go-pg/pg"
	"github.com/voltento/walletManager/walletErrors"
)

type Transaction interface {
	Commit() error
	Rollback() error
}

type WalletManager interface {
	StartTransaction() (Transaction, error)
	AddAccount(ac *Account) error
	GetAllAccounts() ([]Account, error)
	GetPayments() ([]Payment, error)
	Close() error
	GetAccount(id string) (*Account, error)
	UpdateAccount(id string, acc *Account) error
	AddPayment(p *Payment) error
}

type Closer func()

type WalletMgrCluster interface {
	GetWalletMgr() (WalletManager, Closer)
	Close() error
}

type walletMgrCluster struct {
	mgrPool    chan WalletManager
	mgrStorage []WalletManager
}

func (c walletMgrCluster) Close() error {
	var er error
	for _, v := range c.mgrStorage {
		newEr := v.Close()
		if er == nil && newEr != nil {
			er = newEr
		}
	}

	return er
}

func (c walletMgrCluster) GetWalletMgr() (WalletManager, Closer) {
	mgr := <-c.mgrPool

	return mgr, func() { c.mgrPool <- mgr }
}

func CreateWalletMgrCluster(user string, pswrd string, dbName string, addr string, sz int) (WalletMgrCluster, error) {
	cluster := walletMgrCluster{mgrPool: make(chan WalletManager, sz), mgrStorage: make([]WalletManager, sz)}
	var er error
	var newMgr WalletManager
	for i := 0; i < sz; i += 1 {
		newMgr, er = createPsqlWalletMgr(user, pswrd, dbName, addr)
		if er != nil {
			return nil, er
		}

		cluster.mgrPool <- newMgr
		cluster.mgrStorage[i] = newMgr
	}
	return cluster, nil
}

type psqlManager struct {
	db *pg.DB

	insertStmt        *pg.Stmt
	getAccountsStmt   *pg.Stmt
	getAccountStmt    *pg.Stmt
	updateAccountStmt *pg.Stmt
	addPaymentStmt    *pg.Stmt
	getPaymentsStmt   *pg.Stmt
}

func (m psqlManager) Close() error {
	return m.db.Close()
}

func createPsqlWalletMgr(user string, pswrd string, dbName string, addr string) (WalletManager, error) {
	db := pg.Connect(&pg.Options{
		User:     user,
		Database: dbName,
		Password: pswrd,
		Addr:     addr,
	})

	var err error

	var insertStmt *pg.Stmt
	insertStmt, err = db.Prepare("insert into account (id, currency, amount) values ($1, $2, $3);")
	if err != nil {
		return nil, err
	}

	var getPaymentsStmt *pg.Stmt
	getPaymentsStmt, err = db.Prepare("select id, currency, amount from account;")
	if err != nil {
		return nil, err
	}

	var getAccountStmt *pg.Stmt
	getAccountStmt, err = db.Prepare("select id, currency, amount from account where id=$1;")
	if err != nil {
		return nil, err
	}

	var updateAccountStmt *pg.Stmt
	updateAccountStmt, err = db.Prepare("update account set id=$1, currency=$2, amount=$3  where id=$4;")
	if err != nil {
		return nil, err
	}

	var getAccountsStmt *pg.Stmt
	getAccountsStmt, err = db.Prepare("select id, from_account, to_account, amount from payment;")
	if err != nil {
		return nil, err
	}

	var addPaymentStmt *pg.Stmt
	addPaymentStmt, err = db.Prepare("insert into payment (from_account, to_account, amount) values ($1, $2, $3);")
	if err != nil {
		return nil, err
	}

	mgr := psqlManager{
		db:                db,
		insertStmt:        insertStmt,
		getAccountsStmt:   getAccountsStmt,
		getAccountStmt:    getAccountStmt,
		updateAccountStmt: updateAccountStmt,
		addPaymentStmt:    addPaymentStmt,
		getPaymentsStmt:   getPaymentsStmt}
	return mgr, nil
}

func (m psqlManager) StartTransaction() (Transaction, error) {
	return m.db.Begin()
}

func (m psqlManager) AddAccount(ac *Account) error {
	_, er := m.insertStmt.Exec(ac.Id, ac.Currency, ac.Amount)
	return er
}

func (m psqlManager) GetAllAccounts() ([]Account, error) {
	var acc []Account
	_, er := m.getAccountsStmt.Query(&acc)
	if er != nil {
		return nil, er
	}

	return acc, nil
}

func (m psqlManager) GetPayments() ([]Payment, error) {
	var ps []Payment
	_, er := m.getAccountsStmt.Query(&ps)
	if er != nil {
		return nil, er
	}

	return ps, nil
}

func (m psqlManager) UpdateAccount(id string, acc *Account) error {
	r, er := m.updateAccountStmt.Exec(acc.Id, acc.Currency, acc.Amount, id)
	if er != nil {
		return er
	}

	if r.RowsAffected() == 0 {
		return walletErrors.BuildFindAccountError(id)
	}

	return nil
}

func (m psqlManager) GetAccount(id string) (*Account, error) {
	acc := new(Account)
	result, er := m.getAccountStmt.Query(acc, id)

	if er != nil {
		return nil, er
	}

	if result.RowsReturned() == 0 {
		return nil, walletErrors.BuildFindAccountError(id)
	}

	return acc, nil
}

func (m psqlManager) AddPayment(p *Payment) error {
	_, er := m.addPaymentStmt.Exec(p.From_account, p.To_account, p.Amount)
	return er
}
