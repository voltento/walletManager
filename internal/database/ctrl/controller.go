package ctrl

import (
	"github.com/go-pg/pg"
	"github.com/voltento/walletManager/internal/database/errorcheck"
	"github.com/voltento/walletManager/internal/database/model"
	"github.com/voltento/walletManager/internal/database/stmt-mw"
	"github.com/voltento/walletManager/internal/utils"
)

type Account = model.Account
type Payment = model.Payment

type WalletManager interface {
	// Run some logic in the transaction.
	// Be careful, it uses read commitment transaction isolation in basic implementation
	RunInTransaction(func() error) error

	// Add an account
	AddAccount(ac Account) error

	// Add an payment
	AddPayment(p Payment) error

	// Get all accounts
	GetAllAccounts() ([]Account, error)

	// Get the account by id
	GetAccount(id string) (*Account, error)

	// Get all payments
	GetPayments() ([]Payment, error)

	// Update the account
	UpdateAccount(id string, acc Account) error

	// Change the account balance
	ChangeAccountBalance(id string, changeAmount float64) error

	Close() error
}

type Closer func()

type WalletMgrCluster interface {
	// Get wallet manager from the pool
	GetWalletMgr() (WalletManager, Closer)

	Close() error
}

// Wallet manager pull
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

// Take wallet manager from the pool, it will be returned when the closer will be closed
func (c walletMgrCluster) GetWalletMgr() (WalletManager, Closer) {
	mgr := <-c.mgrPool

	return mgr, func() { c.mgrPool <- mgr }
}

// Create cluster for wallet managers
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

type stmt = stmt_mw.Decorator

// Implementation of WalletManager
type psqlManager struct {
	db *pg.DB

	addAccStmt        stmt
	getAccountsStmt   stmt
	getAccountStmt    stmt
	updateAccountStmt stmt
	addPaymentStmt    stmt
	getPaymentsStmt   stmt
	incAccBalanceStmt stmt
}

func (m psqlManager) Close() error {
	return m.db.Close()
}

// Create one wallet manager
func createPsqlWalletMgr(user string, pswrd string, dbName string, addr string) (WalletManager, error) {
	db := pg.Connect(&pg.Options{
		User:     user,
		Database: dbName,
		Password: pswrd,
		Addr:     addr,
	})

	var err error

	var addAccStmt stmt
	addAccStmt, err = db.Prepare("insert into account (id, currency, amount) values ($1, $2, $3);")
	if err != nil {
		return nil, err
	}
	addAccStmt = stmt_mw.LoseConWithDb(addAccStmt)
	addAccStmt = stmt_mw.NotEmptyResp(addAccStmt, "account id")
	addAccStmt = stmt_mw.UniqViolation(addAccStmt, "account id")
	addAccStmt = stmt_mw.NotEmptyRowEffected(addAccStmt, "account")

	var getPaymentsStmt stmt
	getPaymentsStmt, err = db.Prepare("select id, from_account, to_account, amount from payment;")
	if err != nil {
		return nil, err
	}
	getPaymentsStmt = stmt_mw.LoseConWithDb(getPaymentsStmt)

	var getAccountStmt stmt
	getAccountStmt, err = db.Prepare("select id, currency, amount from account where id=$1;")
	if err != nil {
		return nil, err
	}
	getAccountStmt = stmt_mw.LoseConWithDb(getAccountStmt)
	getAccountStmt = stmt_mw.NotEmptyResp(getAccountStmt, "account")

	var updateAccountStmt stmt
	updateAccountStmt, err = db.Prepare("update account set id=$1, currency=$2, amount=$3  where id=$4;")
	if err != nil {
		return nil, err
	}
	updateAccountStmt = stmt_mw.LoseConWithDb(updateAccountStmt)
	updateAccountStmt = stmt_mw.NotEmptyRowEffected(updateAccountStmt, "account")

	var getAccountsStmt stmt
	getAccountsStmt, err = db.Prepare("select id, currency, amount from account;")
	if err != nil {
		return nil, err
	}
	getAccountsStmt = stmt_mw.LoseConWithDb(getAccountsStmt)
	getAccountsStmt = stmt_mw.NotEmptyResp(getAccountsStmt, "account")

	var addPaymentStmt stmt
	addPaymentStmt, err = db.Prepare("insert into payment (from_account, to_account, amount) values ($1, $2, $3);")
	if err != nil {
		return nil, err
	}
	addPaymentStmt = stmt_mw.LoseConWithDb(addPaymentStmt)
	addPaymentStmt = stmt_mw.NotEmptyResp(addPaymentStmt, "account")
	updateAccountStmt = stmt_mw.NotEmptyRowEffected(updateAccountStmt, "account")

	var incAccBalanceStmt stmt
	incAccBalanceStmt, err = db.Prepare("update account set amount=amount+$1 where id=$2;")
	if err != nil {
		return nil, err
	}
	incAccBalanceStmt = stmt_mw.LoseConWithDb(incAccBalanceStmt)
	incAccBalanceStmt = stmt_mw.NotEmptyRowEffected(incAccBalanceStmt, "account")

	mgr := psqlManager{
		db:                db,
		addAccStmt:        addAccStmt,
		getAccountsStmt:   getAccountsStmt,
		getAccountStmt:    getAccountStmt,
		updateAccountStmt: updateAccountStmt,
		addPaymentStmt:    addPaymentStmt,
		getPaymentsStmt:   getPaymentsStmt,
		incAccBalanceStmt: incAccBalanceStmt}
	return mgr, nil
}

func (m psqlManager) RunInTransaction(fn func() error) error {
	fnWrp := func(tx *pg.Tx) error {
		return fn()
	}
	return m.db.RunInTransaction(fnWrp)
}

func (m psqlManager) AddAccount(ac Account) error {
	_, er := m.addAccStmt.Exec(ac.Id, ac.Currency, ac.Amount)
	return er
}

func (m psqlManager) GetAllAccounts() ([]Account, error) {
	var acc []Account
	_, er := m.getAccountsStmt.Query(&acc)
	return acc, er
}

func (m psqlManager) GetPayments() ([]Payment, error) {
	var ps []Payment
	_, er := m.getPaymentsStmt.Query(&ps)
	return ps, er
}

func (m psqlManager) UpdateAccount(id string, acc Account) error {
	_, er := m.updateAccountStmt.Exec(acc.Id, acc.Currency, acc.Amount, id)
	return er
}

func (m psqlManager) GetAccount(id string) (*Account, error) {
	acc := new(Account)
	_, er := m.getAccountStmt.Query(acc, id)
	return acc, er
}

func (m psqlManager) AddPayment(p Payment) error {
	_, er := m.addPaymentStmt.Exec(p.From_account, p.To_account, p.Amount)
	return er
}

func (m psqlManager) ChangeAccountBalance(id string, changeAmount float64) error {
	_, er := m.incAccBalanceStmt.Exec(changeAmount, id)
	if er != nil {
		if errorcheck.IsConstraintViolationError(er) {
			er = utils.BuildFewBalanceError(id)
		}
	}

	return er
}
